<script setup lang="ts">

import { ref } from "vue";
import { strArr, strNumObj, strBoolObj } from "@/types"
import { sortObjectByValue } from "@/utils/object"

import { XGetSymbolSettings, XGetWsBasicInfo, XSaveSymbolSettings, XDeleteClient } from "../../wailsjs/go/backend/App";
import { backend } from "../../wailsjs/go/models";

let emptytStrSign = ref("")

// ---------- actions ----------
// same as backend
const ActNone = 0

// Client side
const ActClientRegister = 401
const ActClientHeartbeat = 402
const ActClientReachStd = 411

// Server side
const ActSvrRegisterRes = 501
const ActSvrUpdateClient = 502
const ActSvrReciveData = 503
const ActSvrSendData = 504
const ActSvrSetIndicator = 541
// ========== actions end ==========

const timeframeMap = ref<strNumObj>({})
const appliedPriceMap = ref<strNumObj>({})
const indNameArr = ref<strArr>([])
const clientStatusMap = ref<strBoolObj>({})

const currentSymbol = ref(emptytStrSign.value)
const socketReceiveData = ref("")
const socketSendData = ref("")
const socketReceiveTime = ref("")
const socketSendTime = ref("")

const indParam = ref<backend.StIndParam>({
  indName: "",
  timeframe: 0,
  period: 0,
  appliedPrice: 0,
  upLine: 0,
  downLine: 0,
})
const priceRange = ref<backend.StPriceRange>({
  enablePriceLimit: false,
  alertWhenPriceRangeIsExceeded: false,
  longMin: 0,
  longMax: 0,
  shortMin: 0,
  shortMax: 0,
})

let wsConn: WebSocket;
const wsConnOk = ref(false)
const wsPort = ref(0)
const wsHeartbeatTime = ref(3000)

// ---------- basic info ----------
const getBasicInfo = async () => {
  XGetWsBasicInfo().then(basic => {
    indNameArr.value = basic.indNameArr
    clientStatusMap.value = basic.clientStatusMap
    timeframeMap.value = basic.timeframeMap
    appliedPriceMap.value = basic.appliedPriceMap
    emptytStrSign.value = basic.emptytStrSign

    wsPort.value = basic.wsPort
    wsHeartbeatTime.value = basic.wsHeartbeatSeconds * 1000
    wsConnectServer()
  })
}

getBasicInfo()
// ========== basic info end ==========

// ---------- WS ----------
const wsConnectServer = () => {
  wsConn = new WebSocket(`ws://localhost:${wsPort.value}`)
  wsConn.onopen = (ev: Event) => {
    console.log("ws contected, event: ", ev)
    wsWriteMessageActionData(ActClientHeartbeat, "")
    wsConnOk.value = true
    wsStartHeartbeat()
  }
  wsConn.onclose = (ev: Event) => {
    wsConnOk.value = false
    console.log("ws closed, event: ", ev)
    wsConnectServer() // reconnect
  }
  wsConn.onerror = (ev: Event) => {
    wsConnOk.value = false
    console.log("ws error, event: ", ev)
  }
  wsConn.onmessage = wsOnMessage
}

const wsWriteMessageActionData = (action: number, data: any) => {
  if (currentSymbol.value === emptytStrSign.value) {
    return
  }

  let dataJson = ""
  if (data) {
    dataJson = JSON.stringify(data)
  }

  const json = JSON.stringify({
    action,
    symbol: currentSymbol.value,
    data: dataJson
  })

  wsConn.send(json);
}

const wsStartHeartbeat = () => {
  let timer = window.setInterval(() => {
    if (wsConnOk) {
      wsWriteMessageActionData(ActClientHeartbeat, "")
    } else {
      clearInterval(timer)
    }
  }, wsHeartbeatTime.value)
}

// Process WS messages
const wsOnMessage = (evt: MessageEvent) => {
  const str = evt.data
  if (str === emptytStrSign || str === "") {
    return
  }

  const resAction = ActNone
  const resData = ""

  const cmd = JSON.parse(str)
  const action = cmd.action


  switch (action) {
    case ActSvrUpdateClient:
      const statusMap = JSON.parse(cmd.dataStr)
      clientStatusMap.value = statusMap
      break;
    case ActSvrReciveData:
      socketReceiveData.value = cmd.dataStr
      socketReceiveTime.value = (new Date).toLocaleString()
      break;
    case ActSvrSendData:
      socketSendData.value = cmd.dataStr
      socketSendTime.value = (new Date).toLocaleString()
      break;
    default:
      break;
  }

  // ignore ActNone, avoid forming an endless loop
  if (resAction != ActNone) {
    wsWriteMessageActionData(resAction, resData)
  }
}
// ========== WS end ==========

// ---------- symbol ----------
const onChangeSymbol = (symbol: string) => {
  priceRange.value.longMax = 0
  priceRange.value.longMin = 0
  priceRange.value.shortMax = 0
  priceRange.value.shortMin = 0

  currentSymbol.value = symbol
  getSymbolSettings(symbol)
}

const getSymbolSettings = async (symbol: string) => {
  XGetSymbolSettings(symbol).then((setting) => {
    indParam.value = setting.indParam
    priceRange.value = setting.priceRange
  })
}

const onSaveSymbolSettings = () => {
  XSaveSymbolSettings(currentSymbol.value, indParam.value, priceRange.value)
}

const onDeleteSymbol = async (symbol: string) => {
  var msg = `Confirm to delete client"${symbol}"?`;
  if (confirm(msg)) {
    XDeleteClient(symbol)
  }
}

const terminalData = ref(false)
const onToggleTerminal = () => {
  terminalData.value = !terminalData.value
}
// ========== symbol end ==========
</script>

<template>
  <div class="container">
    <div class="left p-2">
      <h2>Symbols</h2>

      <template v-for=" v, i in clientStatusMap">
        <h4 class="symbol py-1" v-if="typeof i === 'string'">
          <span class="status-block" :class="v ? 'ok' : 'err'"></span>
          <span :class="i === currentSymbol ? 'current' : ''" @click="onChangeSymbol(i)"> {{ i }}</span>
          <span class="delete-symbol" @click="onDeleteSymbol(i)">❎</span>
        </h4>
      </template>
    </div>
    <div class="right p-2">
      <h2>Setting of {{ currentSymbol }}</h2>

      <div class="apply-block">
        <button class="px-2" @click="onSaveSymbolSettings()">Apply</button>
        <span class="gap-x2"></span>
        <button class="px-2" @click="onToggleTerminal()">Toggle terminal data</button>
      </div>

      <div>
        <div class="block-title py-1">Indicator params:</div>
        <p>
          <span class="disp-inline-block w-8rem">Indicator name:</span>
          <select v-model="indParam.indName" class="w-8rem">
            <option v-for="v, i in indNameArr" :value="v">{{ v }}</option>
          </select>
          <span class="px-2">Up line</span>
          <input type="number" placeholder="Up line" v-model="indParam.upLine" class="w-4rem" />
          <span class="px-2">Down line</span>
          <input type="number" placeholder="Down line" v-model="indParam.downLine" class="w-4rem" />
        </p>
        <p>
          <span class="disp-inline-block w-8rem">Applied price:</span>
          <select v-model="indParam.appliedPrice">
            <option v-for="v, k in appliedPriceMap" :value="v">{{ k }}</option>
          </select>
        </p>
        <p>
          <span class="disp-inline-block w-8rem">Timeframe:</span>
          <select v-model="indParam.timeframe" class="w-8rem">
            <option v-for="v, k in sortObjectByValue(timeframeMap)" :value="v[1]">{{ v[0] }}</option>
          </select>
          <span class="px-2">Period:</span>
          <input v-model.number="indParam.period" type="number" class="w-4rem" />
        </p>
      </div>

      <div>
        <div class="block-title py-1">
          <input type="checkbox" id="enablePriceLimit" v-model="priceRange.enablePriceLimit">
          <span class="gap-x2"></span>
          <label for="enablePriceLimit">Price limit:</label>
        </div>
        <p>
          <label for="alertWhenPriceRangeIsExceeded">Alert when price range is exceeded:</label>
          <span class="gap-x2"></span>
          <input type="checkbox" id="alertWhenPriceRangeIsExceeded" v-model="priceRange.alertWhenPriceRangeIsExceeded" />
        </p>
        <p>
          <span class="disp-inline-block w-4rem">Long:</span>
          <span class="px-2">&gt;</span>
          <input type="number" placeholder="minimum price" v-model="priceRange.longMin" style="width:7rem" />
          <span class="px-2">and &lt;</span>
          <input type="number" placeholder="maximum price" v-model="priceRange.longMax" style="width:7rem" />
        </p>
        <p>
          <span class="disp-inline-block w-4rem">Short:</span>
          <span class="px-2">&gt;</span>
          <input type="number" placeholder="minimum price" v-model="priceRange.shortMin" style="width:7rem" />
          <span class="px-2">and &lt;</span>
          <input type="number" placeholder="maximum price" v-model="priceRange.shortMax" style="width:7rem" />
        </p>
      </div>

      <div class="transfer-data p-2" v-show="terminalData">
        <p class="block-title py-1"> {{ socketReceiveTime ? `${socketReceiveTime}` : "" }} Received:</p>
        <div class="terminal-data p-2">
          {{ socketReceiveData }}
        </div>
        <p class="block-title py-1"> {{ socketSendTime ? `${socketSendTime}` : "" }} Sent:</p>
        <div class="terminal-data p-2">
          {{ socketSendData }}
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
$leftWidth: 240px;

.container {
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: row;

  .left {
    background-color: #3f42c4;
    width: $leftWidth;

    .current {
      color: #b776dd;
    }

    .symbol {
      position: relative;
      cursor: pointer;

      .delete-symbol {
        position: absolute;
        right: 0;
        color: #e29d1e;

        &:hover {
          color: #dd22a5;
        }
      }
    }
  }

  .right {
    background-color: #a2a3db;
    width: calc(100% - $leftWidth);
    position: relative;

    .apply-block {
      position: fixed;
      top: 0.5rem;
      right: 0.5rem;
      z-index: 9;
    }

    .block-title {
      position: relative;
      font-weight: 700;
      color: #3f42c4;
    }

    .transfer-data {
      position: fixed;
      left: 0;
      bottom: 0;
      width: 100%;
      background-color: #400c5e;

      .block-title {
        color: #a2a3db;
      }

      .terminal-data {
        background-color: #290885;
        height: 8rem;
        width: 100%;
        word-break: break-all;
      }
    }
  }
}

.status-block {
  width: 0.5rem;
  height: 1rem;
  display: inline-block;
  margin-right: 0.5rem;

  &.ok {
    background-color: #17eb4c;
  }

  &.err {
    background-color: #f01d1d;
  }
}
</style>