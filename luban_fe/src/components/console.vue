<template>
  <div>
    <div id="terminal" width="100%" style="overflow: auto;" class="console" />
  </div>
</template>
<script>
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
export default {
  name: 'LogsConsole',
  props: {
    url: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      term: null,
      terminalSocket: null
    }
  },
  mounted() {
    this.initWebSocket()
  },
  beforeDestroy() {
    this.terminalSocket.close()
  },
  methods: {
    initWebSocket() {
      const terminalContainer = document.getElementById('terminal')
      this.term = new Terminal(
          {
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
          }
      )
      const fitAddon = new FitAddon()
      this.term.loadAddon(fitAddon)
      fitAddon.fit()
      this.term.open(terminalContainer)
      this.terminalSocket = new WebSocket(this.url)
      this.terminalSocket.onopen = this.runRealTerminal
      this.terminalSocket.onmessage = this.onConnectionMessage
      this.terminalSocket.onclose = this.closeRealTerminal
      this.terminalSocket.onerror = this.errorRealTerminal
      this.term._initialized = true
      console.log('websocket初始化')
    },
    onConnectionMessage(event) {
      this.term.write(event.data)
    },
    runRealTerminal() {
      console.log('websocket connect onopen')
      this.term.focus()
      console.log('websocket连接成功')
    },
    errorRealTerminal() {
      console.log('error')
    },
    closeRealTerminal() {
      console.log('终端关闭连接')
    }
  }
}
</script>
