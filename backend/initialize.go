package backend

import (
	"bytes"
	"io"
	"os"

	GUF "gitee.com/vhash/gutils/files"
	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
)

func Init() {
	Logger = initZap()

	InitConfFile()
}

func InitConfFile() error {
	_, err := toml.DecodeFile(CONIFG_FILE, &Conf)
	if err != nil {
		Logger.Error("InitConfFile err: ", zap.Error(err))

		if err.(*os.PathError).Op == "open" {
			SaveConfFile() // save the default configure
		}
	}

	return err
}

func SaveConfFile() (err error) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(Conf); err != nil {
		Logger.Error("", zap.Error(err))
	}
	confString := buf.String()

	var f *os.File
	defer f.Close()
	if GUF.FileExist(CONIFG_FILE) {
		f, _ = os.OpenFile(CONIFG_FILE, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	} else {
		f, _ = os.Create(CONIFG_FILE)
	}
	_, err = io.WriteString(f, confString)
	if err != nil {
		Logger.Error("SaveConfFile err: ", zap.Error(err))
	}

	return
}
