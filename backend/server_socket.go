package backend

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	GUC "gitee.com/vhash/gutils/convert"
	"go.uber.org/zap"
)

type StSocketServer struct{}

func (self *StSocketServer) Start() {
	portStr := GUC.Uint2Str(Conf.Socket.ServerPort)
	server, err := net.Listen("tcp", ":"+portStr)
	if err != nil {
		Logger.Error("socketServer Listen error:", zap.Error(err))
	}
	fmt.Println("TCP server started, port:" + portStr)

	for {
		conn, err := server.Accept()
		if err != nil {
			Logger.Error("socketServer Accept error:", zap.Error(err))
			continue
		}
		go self.handler(&conn)
	}
}

func (self *StSocketServer) handler(conn *net.Conn) {
	if conn == nil {
		return
	}

	buf := make([]byte, int(Conf.Socket.ServerBufferSize))
	for {
		cnt, err := (*conn).Read(buf)
		if err != nil {
			return
		}
		if cnt == 0 {
			continue
		}

		inByte := buf[0:cnt]
		inStr := strings.TrimSpace(string(inByte))
		inStr = strings.TrimSpace(inStr)

		if inStr == "" {
			actData := GenActDataBytesBtStr(ActNone, "")
			self.WriteRes(conn, actData)
			continue
		}

		SocketRecvCh <- inStr
		self.ProcessJson(conn, inByte, inStr)
	}
}

func (self *StSocketServer) ProcessJson(conn *net.Conn, inByte []byte, inStr string) {
	var inData StCommunicateData
	if err := json.Unmarshal(inByte, &inData); err != nil {
		Logger.Error("ProcessJson Unmarshal inData error:", zap.Error(err), zap.String("input:", inStr))
		return
	}

	if inData.Action != ActClientHeartbeat {
		fmt.Printf("socket received(not heartbeat): `%v`\n", inStr)
	}

	var resBytes []byte
	symbol := inData.Symbol

	switch inData.Action {
	case ActClientHeartbeat:
		resBytes = GenIndicatorDataStr(symbol)
	case ActClientReachStd:
		var reachStd StCommunicateDataReachStd
		if err := json.Unmarshal([]byte(inData.DataStr), &reachStd); err != nil {
			Logger.Error("ProcessJson Unmarshal ActClientReachStd data error:", zap.Error(err), zap.String("input:", inStr))
			return
		}

		var resArr []string
		if reachStd.IndReachStd == IndReachStdLong {
			resArr = append(resArr, StrIndicator+" "+StrBullish)
		} else if reachStd.IndReachStd == IndReachStdShort {
			resArr = append(resArr, StrIndicator+" "+StrBearish)
		}

		if reachStd.IndReachStd == FormReachStdLong {
			resArr = append(resArr, StrForm+" "+StrBullish)
		} else if reachStd.IndReachStd == FormReachStdShort {
			resArr = append(resArr, StrForm+" "+StrBearish)
		}

		var notifyMessage = strings.Join(resArr, "\n")
		NotifyWindows(symbol, notifyMessage, "")

	case ActClientRegister:
		resBytes = GenRegisterDataStr(symbol, conn)
		UpdateClientCh()
	}

	if len(resBytes) == 0 {
		resBytes = GenActDataBytesBtStr(ActNone, "")
	}

	if v, ok := Conf.ClientSettingMap[symbol]; ok {
		v.LastCommunicateTimestamp = time.Now().Unix()
	}

	self.WriteRes(conn, resBytes)
}

func (self *StSocketServer) WriteRes(conn *net.Conn, cmd []byte) {
	if conn == nil {
		return
	}

	res := append(cmd, Conf.Socket.DataTail...)
	(*conn).Write(res)
	SocketSendCh <- string(res)
}
