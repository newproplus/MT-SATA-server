package backend

import (
	"fmt"
	"os"
	"time"

	GUF "gitee.com/vhash/gutils/files"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initZap() (logger *zap.Logger) {
	if ok, _ := GUF.PathExists(Conf.Zap.Director); !ok {
		_ = os.Mkdir(Conf.Zap.Director, os.ModePerm)
	}
	// DebugLevel
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// InfoLevel
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// WarnLevel
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// ErrorLevel
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/%sdebug.log", Conf.Zap.Director, Conf.Zap.FileNamePrefix), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/%sinfo.log", Conf.Zap.Director, Conf.Zap.FileNamePrefix), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/%swarn.log", Conf.Zap.Director, Conf.Zap.FileNamePrefix), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/%serror.log", Conf.Zap.Director, Conf.Zap.FileNamePrefix), errorPriority),
	}
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if Conf.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  Conf.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	// Show caller
	if Conf.Zap.ShowCaller {
		config.EncodeCaller = zapcore.FullCallerEncoder
	}
	switch {
	case Conf.Zap.EncodeLevel == "LowercaseLevelEncoder": // lowercase encoder(default)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case Conf.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // colored lowercase encoder
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case Conf.Zap.EncodeLevel == "CapitalLevelEncoder": // uppdercase encoder
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case Conf.Zap.EncodeLevel == "CapitalColorLevelEncoder": // colored uppdercase encoder
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

func getEncoder() zapcore.Encoder {
	if Conf.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := GetWriteSyncer(fileName) 
	return zapcore.NewCore(getEncoder(), writer, level)
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(Conf.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}

func GetWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, 
		MaxSize:    10,   // The size of log file before splitting.(MB)
		MaxBackups: 200,  // The maximum number of old files to keep.
		MaxAge:     30,   // Maximum number of days to keep old files.
		Compress:   true, // Enable compress.
	}

	if Conf.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
