package backend

const WindowTitle = "Semi-automatic trading assistant"
const CONIFG_FILE = "./conf.toml"
const EmptytStrSign = "NONE" // Emptyt string sign in Communication
const StrIndicator = "Indicator"
const StrForm = "Form"
const StrBullish = "bullish"
const StrBearish = "bearish"

const IndReachStdLong = 1    // The technical indicator value of position long reach standard.
const IndReachStdShort = -1  // The technical indicator value of position short reach standard.
const FormReachStdLong = 1   // The form of position long reach standard.
const FormReachStdShort = -1 // The form of position short reach standard.

// Actions
const (
	ActNone = 0

	// Client side
	ActClientRegister  = 401
	ActClientHeartbeat = 402
	ActClientReachStd  = 411

	// Server side
	ActSvrRegisterRes  = 501
	ActSvrUpdateClient = 502
	ActSvrReciveData   = 503
	ActSvrSendData     = 504
	ActSvrSetIndicator = 541
)

// Meta Trader timeframe
// Refer: https://www.mql5.com/en/docs/constants/chartconstants/enum_timeframes
const (
	MtPeriodNone    StMtTimeframe = -1
	MtPeriodCurrent StMtTimeframe = 0
	MtPeriodM1      StMtTimeframe = 1
	MtPeriodM2      StMtTimeframe = 2
	MtPeriodM3      StMtTimeframe = 3
	MtPeriodM4      StMtTimeframe = 4
	MtPeriodM5      StMtTimeframe = 5
	MtPeriodM6      StMtTimeframe = 6
	MtPeriodM10     StMtTimeframe = 10
	MtPeriodM12     StMtTimeframe = 12
	MtPeriodM15     StMtTimeframe = 15
	MtPeriodM20     StMtTimeframe = 20
	MtPeriodM30     StMtTimeframe = 30
	MtPeriodH1      StMtTimeframe = 60
	MtPeriodH2      StMtTimeframe = 120
	MtPeriodH3      StMtTimeframe = 180
	MtPeriodH4      StMtTimeframe = 240
	MtPeriodH5      StMtTimeframe = 300
	MtPeriodH6      StMtTimeframe = 360
	MtPeriodH8      StMtTimeframe = 480
	MtPeriodH12     StMtTimeframe = 720
	MtPeriodD1      StMtTimeframe = 1440
	MtPeriodW1      StMtTimeframe = 10080
	MtPeriodMN1     StMtTimeframe = 43200
)

// Meta Trader applied price
// Refer: https://www.mql5.com/en/docs/constants/indicatorconstants/prices#enum_applied_price_enum
const (
	MtPriceTypeNone     StMtAppliedPrice = -1 // None
	MtPriceTypeClose    StMtAppliedPrice = 0  // Close price
	MtPriceTypeOpen     StMtAppliedPrice = 1  // Open price
	MtPriceTypeHigh     StMtAppliedPrice = 2  // The maximum price for the period
	MtPriceTypeLow      StMtAppliedPrice = 3  // The minimum price for the period
	MtPriceTypeMedian   StMtAppliedPrice = 4  // Median price, (high + low)/2
	MtPriceTypeTypical  StMtAppliedPrice = 5  // Typical price, (high + low + close)/3
	MtPriceTypeWeighted StMtAppliedPrice = 6  // Weighted close price, (high + low + close + close)/4
)
