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
	MtTimeframeNone    StMtTimeframe = -1
	MtTimeframeCurrent StMtTimeframe = 0
	MtTimeframeM1      StMtTimeframe = 1
	MtTimeframeM2      StMtTimeframe = 2
	MtTimeframeM3      StMtTimeframe = 3
	MtTimeframeM4      StMtTimeframe = 4
	MtTimeframeM5      StMtTimeframe = 5
	MtTimeframeM6      StMtTimeframe = 6
	MtTimeframeM10     StMtTimeframe = 10
	MtTimeframeM12     StMtTimeframe = 12
	MtTimeframeM15     StMtTimeframe = 15
	MtTimeframeM20     StMtTimeframe = 20
	MtTimeframeM30     StMtTimeframe = 30
	MtTimeframeH1      StMtTimeframe = 60
	MtTimeframeH2      StMtTimeframe = 120
	MtTimeframeH3      StMtTimeframe = 180
	MtTimeframeH4      StMtTimeframe = 240
	MtTimeframeH5      StMtTimeframe = 300
	MtTimeframeH6      StMtTimeframe = 360
	MtTimeframeH8      StMtTimeframe = 480
	MtTimeframeH12     StMtTimeframe = 720
	MtTimeframeD1      StMtTimeframe = 1440
	MtTimeframeW1      StMtTimeframe = 10080
	MtTimeframeMN1     StMtTimeframe = 43200
)

// Meta Trader applied price
// Refer: https://www.mql5.com/en/docs/constants/indicatorconstants/prices#enum_applied_price_enum
const (
	MtAppliedPriceNone     StMtAppliedPrice = -1 // None
	MtAppliedPriceClose    StMtAppliedPrice = 0  // Close price
	MtAppliedPriceOpen     StMtAppliedPrice = 1  // Open price
	MtAppliedPriceHigh     StMtAppliedPrice = 2  // The maximum price for the period
	MtAppliedPriceLow      StMtAppliedPrice = 3  // The minimum price for the period
	MtAppliedPriceMedian   StMtAppliedPrice = 4  // Median price, (high + low)/2
	MtAppliedPriceTypical  StMtAppliedPrice = 5  // Typical price, (high + low + close)/3
	MtAppliedPriceWeighted StMtAppliedPrice = 6  // Weighted close price, (high + low + close + close)/4
)
