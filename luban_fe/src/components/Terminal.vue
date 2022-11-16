<template>
  <div>
    <div>
      <div id="terminal" width="100%" style="overflow: auto;" className="console" />
    </div>
  </div>
</template>
<script>
import { Terminal } from 'xterm'
import 'xterm/css/xterm.css'
import { FitAddon } from 'xterm-addon-fit'
export default {
  name: 'Terminal',
  props: {
    url: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      xterm: null,
      ws: null,
      winHeight: window.innerHeight,
      containername: '',
      containers: [],
    }
  },
  mounted() {
    // 初始化xterm
    this.initTerminal()
    window.addEventListener('resize', this.onWindowResize)
  },
  methods: {
    initTerminal() {
      // 初始化xterm
      const height = document.body.clientHeight
      this.xterm = new Terminal({
        cols: 120,
        rows: 23,
        rendererType: 'canvas', // 渲染类型
        convertEol: true, // 启动时，光标设置为下一行的开头
        disableStdin: false, // 是否禁用输入
        scrollback: 100, // 终端中的回滚量
        cursorBlink: true, // 光标闪烁
        cursorStyle: 'bar', // 光标样式
        fontSize: 14,
        windowsMode: true, // 根据窗口换行
        theme: {
          // 配置主题
          foreground: '#FFFFFF', // 字体
          background: '#000000', // 背景色
          cursor: 'help', // 设置光标
          lineHeight: 16,
        }
      })

      const fitAddon = new FitAddon()
      this.xterm.loadAddon(fitAddon)
      fitAddon.fit()
      const terminalContainer = document.getElementById('terminal')
      this.xterm.open(terminalContainer)
      this.xterm.writeln('欢迎使用运维系统终端')
      this.xterm.writeln('')

      // 链接websocket
      this.wsConnect()

      // data事件处理，支持输入与粘贴
      this.xterm.onData(this.onTerminalData)
      // resize事件处理
      // 窗口大小变化时触发xterm的resize方法，向后端发送行列数，格式由后端决定
      // 这里不使用seize默认参数，因为改变窗口大小会 改变size中的列数而不能改变行数
      this.xterm.onResize(this.onTerminalResize)
    },

    onTerminalData(data) {
      const msg = {
        type: 'read',
        data: data
      }
      this.ws.send(JSON.stringify(msg))
    },

    onWindowResize() {
      const msg = {
        type: 'resize',
        cols: this.xterm.cols,
        rows: this.xterm.rows,
      }
      this.ws.send(JSON.stringify(msg))
    },
    onTerminalResize() {
      const msg = {
        type: 'resize',
        cols: this.xterm.cols,
        rows: this.xterm.rows,
      }
      this.ws.send(JSON.stringify(msg))
    },

    wsConnect() {
      this.ws = new WebSocket(this.url)
      this.ws.onopen = this.onConnectionOpen
      this.ws.onmessage = this.onConnectionMessage
      this.ws.onclose = this.onConnectionClose
      this.ws.onerror = this.onConnectionError
    },

    onConnectionOpen(event) {
      console.log('websocket connect onopen', event)
      this.onTerminalResize()
      this.xterm.focus()
    },

    onConnectionMessage(event) {
      const msg = JSON.parse(event.data)
      if (msg.type === 'write') {
        this.xterm.write(msg.data)
      } else {
        console.log('invalid msg type:' + msg)
      }
    },
    onConnectionClose(event) {
      console.log('终端关闭连接', event)
    },
    onConnectionError(event) {
      console.log('error', event)
    },
  }
}
</script>
