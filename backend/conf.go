package backend

var Conf = Configure{
	Window: ConfigureWindow{
		Width:  1000,
		Height: 640,
	},

	Ws: ConfigureWs{
		ServerPort:       60200,
		HeartbeatSeconds: 10,
	},

	Socket: ConfigureSocket{
		ServerPort:       60100,
		HeartbeatSeconds: 10,
		ServerBufferSize: 4096,
		DataTail:         []byte("\t\t\t"),
	},

	Zap: ConfigureZap{
		Director:       "./log",
		FileNamePrefix: "",
		Prefix:         "",
		StacktraceKey:  "",
		ShowLine:       true,
		ShowCaller:     true,
		LogInConsole:   true,
		Format:         "txt",
		EncodeLevel:    "LowercaseColorLevelEncoder",
	},

	ClientSettingMap: map[string]StSocketClientData{},
}
