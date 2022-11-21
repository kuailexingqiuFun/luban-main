<template>
  <div>
    <el-table :data="pods" v-loading="loading" @search="search">
      <el-table-column label="名称" prop="name" min-width="80" show-overflow-tooltip>
        <template v-slot:default="{row}">
          <span class="span-link">{{ row.metadata.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="镜像" min-width="120" show-overflow-tooltip>
        <template v-slot:default="{row}">
          <div v-for="(item,index) in row.spec.containers" v-bind:key="index" class="myTag">
            <el-tag type="info" size="small" style="display:flex">
              {{ item.image }}
            </el-tag>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="状态" min-width="45">
        <template v-slot:default="{row}">
          <el-button v-if="row.status.phase === 'Running' || row.status.phase === 'Succeeded'" type="success"
                     size="mini" plain round>
            {{ row.status.phase }}
          </el-button>
          <el-button v-if="row.status.phase !== 'Running' && row.status.phase !== 'Succeeded'" type="warning"
                     size="mini" plain round>
            {{ row.status.phase }}
          </el-button>
        </template>
      </el-table-column>
      <el-table-column :label="'重启次数'" min-width="40">
        <template v-slot:default="{row}">
          {{ getPodRestartCounts(row.status.containerStatuses ) }}
        </template>
      </el-table-column>

      <el-table-column label="节点" min-width="40" prop="spec.nodeName" show-overflow-tooltip/>
      <el-table-column :label="'Cpu'" min-width="45">
        <template v-slot:default="{row}">
          {{ getPodUsage(row.metadata.name, "cpu") }}
        </template>
      </el-table-column>
      <el-table-column :label="'Memory'" min-width="45">
        <template v-slot:default="{row}">
          {{ getPodUsage(row.metadata.name, "memory") }}
        </template>
      </el-table-column>
      <el-table-column label="存活时间" min-width="40" prop="metadata.creationTimestamp" show-overflow-tooltip fix>
        <template v-slot:default="{row}">
          {{ row.metadata.creationTimestamp |formatDate }}
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template v-slot:default="{row}">
          <el-dropdown>
            <el-button link size="small">更多</el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <div v-for="(item,index) in row.spec.containers" :key="'info-'+ index">
                  <el-dropdown-item @click.native="openTerminal(row, item.name)">
                    {{ item.name }} 终端
                  </el-dropdown-item>
                </div>
                <div v-for="(item,index1) in row.spec.containers" :key="index1">
                  <el-dropdown-item @click.native="openTerminalLogs(row, item.name)">
                    {{ item.name }} 终端日志
                  </el-dropdown-item>
                </div>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
      </el-table-column>
    </el-table>

    <div v-if="dialogVisibleTerminal">
      <el-dialog
              :title="Title"
              :visible.sync="dialogVisibleTerminal"
              width="70%"
              @close="closeTerminalDialog"
              append-to-body>
        <div>
          <div class="container">
            <Application-terminal
                    :url="currenturl"/>
          </div>
        </div>
      </el-dialog>
      <div/>
    </div> <!--最内层终端模态框结束位置-->

    <!--最内层日志模态框-->
    <div v-if="dialogVisibleContainerlog">
      <el-dialog
              :title="Title"
              :visible.sync="dialogVisibleContainerlog"
              width="70%"
              @close="closeContainerlogDialog"
              append-to-body>
        <div class="container">
          <Application-LogsConsole
                  :terminal="terminal"
                  :url="currenturl"/>
        </div>
      </el-dialog>
      <div/>
    </div> <!--最内层日志模态框结束位置-->
  </div>
</template>

<script>
   import { PodsList } from "@/api/kubernetes/workloads/pods"
   import { PodsMetricsList } from "@/api/kubernetes/metrics"
   import { AgeFormat } from '@/utils/age'
   import Terminal from "@/components/Terminal"
   import LogsConsole from "@/components/console"
export default {
  name: "DetailPods",
  props: {
    cluster_id: Number,
    namespace: String,
    selector: String,
    fieldSelector: String,
    name: String,
  },
  components: {
            'Application-terminal': Terminal,
            'Application-LogsConsole': LogsConsole,
    },
  watch: {
    selector: {
      handler (newSelector) {
        if (newSelector) {
          this.search()
        }
      },
      immediate: true,
    },
    fieldSelector: {
      handler (newSelector) {
        if (newSelector) {
          this.search()
        }
      },
      immediate: true,
    },
  },
  data () {
    return {
      dialogVisibleTerminal: false,
      dialogVisibleContainerlog: false,
      Title: "",
      currenturl: "",
      path:  process.env.VUE_APP_WS_URL,
      terminal: {
        pid: 1,
        name: 'terminal',
        cols: 150,
        rows: 23
      },
      loading: false,
      pods: [],
      podUsage: [],
    }
  },
  methods: {
    openTerminal(row, name){
      console.log(row)
      this.Title = '容器终端 '+  ' 容器名称: ' + row.metadata.name
      this.dialogVisibleTerminal = true
      console.log(this.name)
      this.currenturl = this.path + "/api/v1/kubernetes/proxy/terminal?pod_name=" + row.metadata.name + '&namespace=' + row.metadata.namespace + '&cluster_id=' + this.cluster_id + '&name='+ name
    },
    openTerminalLogs(row, name){
      this.Title = '容器日志 '+  ' 容器名称: ' + row.metadata.name
      this.dialogVisibleContainerlog = true
      this.currenturl = this.path + "/api/v1/kubernetes/proxy/logs?pod_name=" + row.metadata.name + '&namespace=' + row.metadata.namespace + '&cluster_id=' + this.cluster_id + '&name='+ name
    },
    closeTerminalDialog(){
      this.dialogVisibleTerminal = false
    },
    closeContainerlogDialog(){
      this.dialogVisibleContainerlog = false
    },
    search () {
      this.loading = true
      PodsList(this.cluster_id, '', '', this.namespace,'', this.selector, '').then((res) => {
        this.pods = res.data.items
        this.loading = false
        PodsMetricsList(this.cluster_id, this.namespace, this.selector).then(res => {
          if (res.code === 0) {
            this.podUsage = res.data.items
          }
        })
      })
    },
    getPodRestartCounts(containerStatuses){
      let counts = 0
      if (containerStatuses.length >0 ) {
        for ( let stats of containerStatuses) {
          counts += stats.restartCount
        }
      }
      return counts
    },
    getPodUsage (name, type) {
      let result = "0 m"
      if (this.podUsage.length > 0) {
        for (let item of this.podUsage) {
          if (item.metadata.name === name) {
            let usage = 0
            for (let container of item.containers) {
              if (type === "cpu") {
                if (container.usage.cpu.indexOf("n") > -1) {
                  usage = usage + parseInt(container.usage.cpu)
                }
                if (container.usage.cpu.indexOf("m") > -1) {
                  usage = usage + parseInt(container.usage.cpu) * 1000 * 1000
                }
              }
              if (type === "memory") {
                if (container.usage.memory.indexOf("Ki") > -1) {
                  usage = usage + parseInt(container.usage.memory)
                }
                if (container.usage.memory.indexOf("Mi") > -1) {
                  usage = usage + parseInt(container.usage.memory) * 1000
                }
              }
            }
            const unit = type === "cpu" ? "m" : "Mi"
            if (type === "cpu") {
              result = (usage / 1000000).toFixed(2)
            } else {
              result = (usage / 1000).toFixed(2)
            }
            result = result + unit
          }
        }
      }
      return result
    }
  },
  filters:{
    formatDate: function(time) {
      if (time != null && time !== '') {
        return AgeFormat(time)
      } else {
        return ''
      }
    }
  }
}
</script>

<style scoped>
</style>
