package backend

import (
	"encoding/json"
	"net"

	GUDT "gitee.com/vhash/gutils/datetime"
	"github.com/gen2brain/beeep"
	"go.uber.org/zap"
)

func GenActDataBytesBtStr(act int, dataStr string) []byte {
	var data StCommunicateData
	data.Action = act
	data.DataStr = dataStr
	data.Time = GUDT.NowStrYMDHMS()

	var bbb, _ = json.Marshal(data)
	return bbb
}

func GenActDataBytesWithStruct(act int, dataWithStruct any) []byte {
	allClientSymbolNameJsonByte, err := json.Marshal(dataWithStruct)
	if err != nil {
		Logger.Error("GenActDataBytesWithStruct err", zap.Error(err), zap.Any("input:", dataWithStruct))
	}

	return GenActDataBytesBtStr(act, string(allClientSymbolNameJsonByte))
}

func GenIndicatorDataStr(symbol string) []byte {
	indParam := GetSymbolSettings(symbol)
	var indParamsJson, _ = json.Marshal(indParam)
	return GenActDataBytesBtStr(ActSvrSetIndicator, string(indParamsJson))
}

func GenRegisterDataStr(symbol string, conn *net.Conn) []byte {
	var client StSocketClientData
	if v, ok := Conf.ClientSettingMap[symbol]; ok {
		client = v
	}
	client.Conn = conn
	client.Connected = true
	Conf.ClientSettingMap[symbol] = client

	// return the basic data
	var regRes StCommunicateDataActSvrRegisterRes
	regRes.EmptytStrSign = EmptytStrSign
	regRes.HeartbeatSeconds = Conf.Socket.HeartbeatSeconds
	return GenActDataBytesWithStruct(ActSvrRegisterRes, regRes)
}

func GetWsBasicInfo() StWsBasicInfo {
	return StWsBasicInfo{
		WsPort:             Conf.Ws.ServerPort,
		WsHeartbeatSeconds: Conf.Ws.HeartbeatSeconds,
		IndNameArr:         IndNameArr,
		TimeframeMap:       TimePeriodMap,
		PriceTypeMap:       priceTypeMap,
		ClientStatusMap:    GetClientStatusMap(),
		EmptytStrSign:      EmptytStrSign,
	}
}

func DeleteClient(symbol string) {
	if _, ok := Conf.ClientSettingMap[symbol]; ok {
		delete(Conf.ClientSettingMap, symbol)
		SaveConfFile()
		UpdateClientCh()
	}
}

func GetSymbolSettings(symbol string) StSocketClientSetting {
	var indParam StIndParam
	var priceRange StPriceRange
	if v, ok := Conf.ClientSettingMap[symbol]; ok {
		indParam = v.IndParam
		priceRange = v.PriceRange
	}

	return StSocketClientSetting{
		IndParam:   indParam,
		PriceRange: priceRange,
	}
}

func SaveSymbolSettings(symbol string, indParam StIndParam, priceRange StPriceRange) bool {
	if v, ok := Conf.ClientSettingMap[symbol]; ok {
		v.IndParam.IndName = indParam.IndName
		v.IndParam.Period = indParam.Period
		v.IndParam.PriceType = indParam.PriceType
		v.IndParam.Timeframe = indParam.Timeframe
		v.IndParam.UpLine = indParam.UpLine
		v.IndParam.DownLine = indParam.DownLine

		v.PriceRange.EnablePriceLimit = priceRange.EnablePriceLimit
		v.PriceRange.AlertWhenPriceRangeIsExceeded = priceRange.AlertWhenPriceRangeIsExceeded
		v.PriceRange.LongMin = priceRange.LongMin
		v.PriceRange.LongMax = priceRange.LongMax
		v.PriceRange.ShortMin = priceRange.ShortMin
		v.PriceRange.ShortMax = priceRange.ShortMax

		Conf.ClientSettingMap[symbol] = v

		SaveConfFile()
		return true
	}

	return false
}

func GetClientStatusMap() StStringBoolMap {
	var res = make(StStringBoolMap)
	for k, v := range Conf.ClientSettingMap {
		res[k] = v.Connected
	}
	return res
}

func UpdateClientCh() {
	csm := GetClientStatusMap()
	allClientSymbolNameJsonByte, err := json.Marshal(csm)
	if err != nil {
		Logger.Error("UpdateClientCh Unmarshal error:", zap.Error(err), zap.Any("input:", csm))
	}
	allClientSymbolNameJson := string(allClientSymbolNameJsonByte)
	SocketClientCh <- allClientSymbolNameJson
}

// Notify on Windows OS
func NotifyWindows(title, message, iconPath string) {
	err := beeep.Notify(title, message, iconPath)
	if err != nil {
		Logger.Error("NotifyWindows Unmarshal error:", zap.Error(err))
	}
}
