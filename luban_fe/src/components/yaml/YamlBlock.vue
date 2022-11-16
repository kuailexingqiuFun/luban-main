<template>
  <div style="overflow: hidden;width: 100%;">
    <el-form ref="YamlBlock">
      <div id="app" style="height: 100vh">
        <div id="code-editor" ref="code-editor" style="height: 100%; width: 100%" />
      </div>
    </el-form>
    <div v-if="!readOnly" class="dialog-footer" align="center">
      <el-button size="small" @click="cancel">取 消</el-button>
      <el-button size="small" type="primary" @click="submitForm">确 定</el-button>
    </div>
  </div>
</template>

<script>
import * as monaco from 'monaco-editor'
import yaml from 'js-yaml'
export default {
  name: 'App',
  props: {
    form: {
      type: Object,
      default() {
        return {}
      }
    },
    readOnly: {
      type: Boolean,
      default() {
        return false
      }
    }
  },
  data() {
    return {
      // 编辑器实例
      monacoEditor: null,
      // 编辑器主题
      theme: 'vs-dark', // 默认是 "vs"
      dataYaml: this.form
    }
  },
  watch: {
    value: function() {
      this.form = yaml.load(this.dataYaml)
    }
  },

  mounted() {
    // 初始化编辑器
    this.initmonacoEditor()
  },

  beforeDestroy() {
    // 销毁之前把monaco的实例也销毁了，不然会多次注册
    if (this.monacoEditor) {
      this.monacoEditor.dispose()
    }
  },
  created() {
    this.JsonToYamlFrom()
  },
  methods: {
    initmonacoEditor() {
      this.monacoEditor = monaco.editor.create(document.getElementById('code-editor'), {
        value: this.dataYaml, // 初始文字
        language: 'yaml', // 语言
        readOnly: false, // 是否只读
        automaticLayout: true, // 自动布局
        foldingStrategy: 'indentation',
        theme: this.theme, // vs | hc-black | vs-dark
        minimap: {
          enabled: true// 小地图
        },
        tabSize: 2, // tab缩进长度
        fontSize: 16 // 文字大小
      })
    },
    submitForm() {
      this.$emit('submit', yaml.load(this.monacoEditor.getValue()))
    },
    JsonToYamlFrom() {
      if (JSON.stringify(this.dataYaml) !== '{}') {
        this.dataYaml = yaml.dump(this.dataYaml)
      } else {
        this.dataYaml = ''
      }
    },
    cancel() {
      this.$emit('cancel')
    }
  }
}
</script>

<style scoped>

</style>
