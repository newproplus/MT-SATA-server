package backend

import (
	"net"
)

type StrStrMap map[string]int
type StMtTimeframe int
type StMtAppliedPrice int
type StReachStd int

type StStringBoolMap = map[string]bool
type StTimeframeMap = map[string]StMtTimeframe
type StAppliedPriceMap = map[string]StMtAppliedPrice

type StIndParam struct {
	IndName      string           `json:"indName"`
	Timeframe    StMtTimeframe    `json:"timeframe"`
	Period       StMtTimeframe    `json:"period"`
	AppliedPrice StMtAppliedPrice `json:"appliedPrice"`
	UpLine       float32          `json:"upLine"`
	DownLine     float32          `json:"downLine"`
}

// Price limit range
type StPriceRange struct {
	EnablePriceLimit              bool `json:"enablePriceLimit"`
	AlertWhenPriceRangeIsExceeded bool `json:"alertWhenPriceRangeIsExceeded"`

	LongMin  float32 `json:"longMin"`  // long position minimum price
	LongMax  float32 `json:"longMax"`  // long position maximum price
	ShortMin float32 `json:"shortMin"` // short position minimum price
	ShortMax float32 `json:"shortMax"` // short position maximum price
}

type StSocketClientSetting struct {
	IndParam   StIndParam   `json:"indParam"`
	PriceRange StPriceRange `json:"priceRange"`
}

type StSocketClientData struct {
	Connected                bool         `json:"connected" toml:"-"` // connection status
	Conn                     *net.Conn    `json:"-" toml:"-"`         // connection
	LastCommunicateTimestamp int64        `json:"-"`                  // last Communicate time of connection
	IndParam                 StIndParam   `json:"indParam"`
	PriceRange               StPriceRange `json:"priceRange"`
}

type StWsBasicInfo struct {
	IndNameArr         []string          `json:"indNameArr"`
	TimeframeMap       StTimeframeMap    `json:"timeframeMap"`
	ClientStatusMap    StStringBoolMap   `json:"clientStatusMap"`
	AppliedPriceMap    StAppliedPriceMap `json:"appliedPriceMap"`
	WsPort             uint              `json:"wsPort"`
	WsHeartbeatSeconds uint              `json:"wsHeartbeatSeconds"`
	EmptytStrSign      string            `json:"emptytStrSign"`
}

type ConfigureWindow struct {
	Width  int
	Height int
}

type ConfigureWs struct {
	ServerPort       uint
	HeartbeatSeconds uint
}

type ConfigureSocket struct {
	ServerPort       uint
	HeartbeatSeconds uint
	ServerBufferSize uint
	DataTail         []byte
}

type ConfigureZap struct {
	Director       string
	FileNamePrefix string
	Prefix         string
	StacktraceKey  string
	ShowLine       bool
	ShowCaller     bool
	LogInConsole   bool
	Format         string
	EncodeLevel    string
}

type Configure struct {
	Window           ConfigureWindow
	Ws               ConfigureWs
	Socket           ConfigureSocket
	Zap              ConfigureZap
	ClientSettingMap map[string]StSocketClientData
}

type StCommunicateData struct {
	Action  int    `json:"action"`
	Symbol  string `json:"symbol"`
	DataStr string `json:"dataStr"`
	Time    string `json:"time"`
}

type StCommunicateDataActSvrRegisterRes struct {
	EmptytStrSign    string `json:"emptytStrSign"`
	HeartbeatSeconds uint   `json:"heartbeatSeconds"`
}

type StCommunicateDataReachStd struct {
	IndReachStd  int `json:"indReachStd"`
	FormReachStd int `json:"formReachStd"`
}
