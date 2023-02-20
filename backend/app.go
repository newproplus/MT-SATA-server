package backend

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	// Start non-GUI functions.
	Init()	

	var SocketServer StSocketServer
	go SocketServer.Start()

	var WsServer StWsServer
	go WsServer.Start()
	WsServer.MonitorChToWs()
}

func (a *App) XGetWsBasicInfo() StWsBasicInfo {
	return GetWsBasicInfo()
}

func (a *App) XGetSymbolSettings(symbol string) StSocketClientSetting {
	return GetSymbolSettings(symbol)
}

func (a *App) XSaveSymbolSettings(symbol string, indParam StIndParam, priceRange StPriceRange) bool {
	return SaveSymbolSettings(symbol, indParam, priceRange)
}

func (a *App) XDeleteClient(symbol string) {
	DeleteClient(symbol)
}
