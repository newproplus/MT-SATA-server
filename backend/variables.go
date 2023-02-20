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
	"M1":  MtPeriodM1,
	"M2":  MtPeriodM2,
	"M3":  MtPeriodM3,
	"M4":  MtPeriodM4,
	"M5":  MtPeriodM5,
	"M6":  MtPeriodM6,
	"M10": MtPeriodM10,
	"M12": MtPeriodM12,
	"M15": MtPeriodM15,
	"M20": MtPeriodM20,
	"M30": MtPeriodM30,
	"H1":  MtPeriodH1,
	"H2":  MtPeriodH2,
	"H3":  MtPeriodH3,
	"H4":  MtPeriodH4,
	"H5":  MtPeriodH5,
	"H6":  MtPeriodH6,
	"H8":  MtPeriodH8,
	"H12": MtPeriodH12,
	"D1":  MtPeriodD1,
	"W1":  MtPeriodW1,
	"MN1": MtPeriodMN1,
}

var priceTypeMap StPriceTypeMap = StPriceTypeMap{
	"Close price":                           MtPriceTypeClose,
	"Open price":                            MtPriceTypeOpen,
	"High price":                            MtPriceTypeHigh,
	"Low price":                             MtPriceTypeLow,
	"Median price, (high + low)/2":          MtPriceTypeMedian,
	"Typical price, (high + low + close)/3": MtPriceTypeTypical,
	"Weighted close price, (high + low + close + close)/4": MtPriceTypeWeighted,
}
