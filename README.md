# MT-SATA-server

## About

Semi-automatic trading assistant server for metatrader.

![screenshot](./docs/screenshot.png)

### About client source code

So far, the project is an open source demo project of **wails**, no client code provided. If you want to create a client, please refer to some structure involved in **ProcessJson** function in `backend/server_socket.go`.

### Technology stack

* backend:
  * golang (v1.20.1)
  * wails: A GUI framework that enables you to write desktop apps using Go and web technologies.
* frontend:
  * Node.js (v18.14.0)
  * pnpm: Fast, disk space efficient Node.js package manager. [installation document link](https://pnpm.io/installation)
  * typeScript
  * vuejs (v3.2)
* communication:
  * data type: json string
  * use socket with **metatrader(MT4/MT5)** or other trading client
  * use **web socket** with frontend

## Develop

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

### Live development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

If you want to modify the golang code multiple times in a short time, it's better to stop running the command `wails dev`.

### Building

To build a redistributable, production mode package, use `wails build`.

## Dante

### Cryptocurrency

* XMR: 46df6rwnqcUCFaSummLobcH3J9sWgqYASF8Znq5HnhgrLeASh8u4TPJ2LaLnoQk3uV6t18CgNuFVCDfLUR9G94AZUj1TtGr
* SOL: BbrRkLArfTeAieAtDpvBHNE4KBKX9fmbjPb5JDmKHWE7
* ETH: 0xA59186a08424BE262FBacA922E87Ab82F3C5245B

### Online payment

![DanteQrCodeImg]([./docs/dante_qr_code.png](https://github.com/newproplus/newproplus/blob/main/images/dante_qr_code.png))

## Trouble shutting

### websocket Error "connection to 'ws://localhost:12345/' failed: Invalid frame header"
Refer: https://qiita.com/dbgso/items/4bbfa52d99cae6c547a4, create: `vue.config.ts`, add the following content:

```js
export const devServer = {
    proxy: {
        '/': {
            target: 'http://localhost:12345',
            ws: true,
            changeOrigin: true
        }
    }
};
```