package backend

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

var SocketRecvCh = make(chan string, 1) // channel of socket server received data
var SocketSendCh = make(chan string, 1) // channel of socket server send data
var SocketClientCh = make(chan string)  // channel of socket server's client

var IndNameArr = []string{"ATR", "CCI", "EMA", "LWMA", "MACD", "MTM", "RSI", "RVI", "SMA", "SMMA", "STO"}

var TimePeriodMap StTimeframeMap = StTimeframeMap{
	"M1":  MtTimeframeM1,
	"M2":  MtTimeframeM2,
	"M3":  MtTimeframeM3,
	"M4":  MtTimeframeM4,
	"M5":  MtTimeframeM5,
	"M6":  MtTimeframeM6,
	"M10": MtTimeframeM10,
	"M12": MtTimeframeM12,
	"M15": MtTimeframeM15,
	"M20": MtTimeframeM20,
	"M30": MtTimeframeM30,
	"H1":  MtTimeframeH1,
	"H2":  MtTimeframeH2,
	"H3":  MtTimeframeH3,
	"H4":  MtTimeframeH4,
	"H5":  MtTimeframeH5,
	"H6":  MtTimeframeH6,
	"H8":  MtTimeframeH8,
	"H12": MtTimeframeH12,
	"D1":  MtTimeframeD1,
	"W1":  MtTimeframeW1,
	"MN1": MtTimeframeMN1,
}

var appliedPriceMap StAppliedPriceMap = StAppliedPriceMap{
	"Close price":                           MtAppliedPriceClose,
	"Open price":                            MtAppliedPriceOpen,
	"High price":                            MtAppliedPriceHigh,
	"Low price":                             MtAppliedPriceLow,
	"Median price, (high + low)/2":          MtAppliedPriceMedian,
	"Typical price, (high + low + close)/3": MtAppliedPriceTypical,
	"Weighted close price, (high + low + close + close)/4": MtAppliedPriceWeighted,
}
