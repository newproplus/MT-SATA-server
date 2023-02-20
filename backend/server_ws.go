package backend

import (
	"encoding/json"
	"fmt"
	"net"
	"runtime"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsflate"
	"go.uber.org/zap"
)

var wsConn net.Conn

type StWsServer struct{}

func (self *StWsServer) Start() {
	path := fmt.Sprintf("localhost:%v", Conf.Ws.ServerPort)

	ln, err := net.Listen("tcp", path)
	if err != nil {
		Logger.Error("ws Listen error: ", zap.Error(err))
	}
	e := wsflate.Extension{
		// We are using default parameters here since we use
		// wsflate.{Compress,Decompress}Frame helpers below in the code.
		// This assumes that we use standard compress/flate package as flate
		// implementation.
		Parameters: wsflate.DefaultParameters,
	}
	u := ws.Upgrader{
		Negotiate: e.Negotiate,
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		// Reset extension after previous upgrades.
		e.Reset()

		_, err = u.Upgrade(conn)
		if err != nil {
			Logger.Error("ws Upgrade error: ", zap.Error(err))
			continue
		}

		if _, ok := e.Accepted(); !ok {
			Logger.Error("ws Accepted error: ", zap.Error(err))
			continue
		}

		wsConn = conn

		go func() {
			defer conn.Close()
			for {
				frame, err := ws.ReadFrame(conn)
				if err != nil {
					// Logger.Error("ws ReadFrame error: ", zap.Error(err))
					// conn.Close()
					runtime.Goexit()
					// continue
				}

				frame = ws.UnmaskFrameInPlace(frame)
				isSet, err := wsflate.IsCompressed(frame.Header)
				if isSet && err == nil {
					// Note that even after successful negotiation of
					// compression extension, both sides are able to send
					// non-compressed messages.
					frame, err = wsflate.DecompressFrame(frame)
					if err != nil {
						Logger.Error("ws DecompressFrame error: ", zap.Error(err))
						return
					}
				}

				// Data processing
				inBytes := frame.Payload
				inStr := string(inBytes)
				var inData StCommunicateData
				json.Unmarshal(inBytes, &inData)

				if inData.Action != ActClientHeartbeat {
					fmt.Printf("ws received(not heartbeat): `%v`\n", inStr)
				}

				var resBytes []byte
				switch inData.Action {
				case ActClientHeartbeat:
					resBytes = GenActDataBytesBtStr(ActNone, "")
				}

				// ignore ActNone, avoid forming an endless loop
				if len(resBytes) == 0 && inData.Action != ActNone {
					resBytes = GenActDataBytesBtStr(ActNone, "")
				}
				self.WriteWsBytes(resBytes)
			}
		}()
	}
}

func (self *StWsServer) WriteWsResStr(outStr string) {
	if outStr != "" {
		outStr = EmptytStrSign
	}
	outByte := []byte(outStr)
	self.WriteWsBytes(outByte)
}

func (self *StWsServer) WriteWsBytes(outByte []byte) {
	ack := ws.NewTextFrame(outByte)

	// Compress response unconditionally.
	ack, err := wsflate.CompressFrame(ack)
	if err != nil {
		Logger.Error("ws CompressFrame error: ", zap.Error(err))
		return
	}
	if err = ws.WriteFrame(wsConn, ack); err != nil {
		Logger.Error("ws WriteFrame error: ", zap.Error(err))
		return
	}
}

func (self *StWsServer) WriteWsActionData(str string, action int) {
	if str == "" {
		str = EmptytStrSign
	}

	var res StCommunicateData
	res.Action = action
	res.DataStr = str
	jsonBytes, _ := json.Marshal(res)
	self.WriteWsBytes(jsonBytes)
}

func (self *StWsServer) MonitorChToWs() {
	// Do not use many "go func(){ /*source  code*/ }()"" code, it will be blocked.
	go self.MonitorSocketReveiveCh()
	go self.MonitorSocketSendCh()
	go self.MonitorSocketClientCh()
}

func (self *StWsServer) MonitorSocketReveiveCh() {
	for {
		recvStr := <-SocketRecvCh
		self.WriteWsActionData(recvStr, ActSvrReciveData)
	}
}

func (self *StWsServer) MonitorSocketSendCh() {
	for {
		sendStr := <-SocketSendCh
		self.WriteWsActionData(sendStr, ActSvrSendData)
	}
}

func (self *StWsServer) MonitorSocketClientCh() {
	for {
		sendStr := <-SocketClientCh
		self.WriteWsActionData(sendStr, ActSvrUpdateClient)
	}
}
