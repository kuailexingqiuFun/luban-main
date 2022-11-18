<template>
  <div>
    <div class="table-viewer-header clearfix" style="clear: both">
      <span class="table-viewer-topbar-title">基本信息</span>
    </div>
    <table class="table-default-viewer">
      <tbody>
      <tr>
        <td style="width: 50%">
          <span>名称</span>
          <span class="margin-right"> :</span>
          <span> {{value.metadata.name}} </span>
        </td>
        <td style="width: 50%">
          <span>存活时间</span>
          <span class="margin-right"> :</span>
          <span> {{value.metadata.creationTimestamp | AgeformatDate}} </span>
        </td>
      </tr>
      <tr>
        <td style="width: 50%">
          <span>UID</span>
          <span class="margin-right">:</span>
          <span>  {{ value.metadata.uid }}</span>
        </td>
      </tr>
      <tr>
        <td>
                <span>
                  <span>容器组</span> CIDR
                </span>
          <span class="margin-right">:</span>
          <span> {{ value.spec.podCIDR }}</span>
        </td>
        <td>
          <span>调度状态：</span>
          <span v-if="value.spec.unschedulable">不可调度</span>
          <span v-else>可调度</span>
        </td>
      </tr>
      <tr>
        <td>
          <span>IP 地址</span>
          <span class="margin-right">:</span>
          <span v-for="(value) in value.status.addresses " :key="value.index">
                  <span>  {{value.type}} : {{ value.address}}</span>
            </span>
        </td>
      </tr>
      <tr>
        <td colspan="2">
          <span>标签</span>
          <span class="margin-right"> :</span>
          <span v-for="(k, v, index) in value.metadata.labels" :key="index">
                          <span class="label-custom" type="info"> {{ k }}: {{ v }}</span>
                     </span>
        </td>
      </tr>
      <tr>
        <td colspan="2">
          <span>注解</span>
          <span> :</span>
          <span v-for="(k, v, index) in value.metadata.annotations" :key="index">
                          <span class="margin-right label-custom" type="info">  {{ k }}: {{ v }}</span>
                     </span>
        </td>
      </tr>
      <tr>
        <td>
          <span>系统镜像</span>
          <span class="margin-right">:</span>
          <span>  {{ value.status.nodeInfo.osImage }}</span>
        </td>
        <td>
          <span>内核版本</span>
          <span class="margin-right">:</span>
          <span>  {{ value.status.nodeInfo.kernelVersion }}</span>
        </td>
      </tr>
      <tr>
        <td>
          <span>Kubelet 版本</span>
          <span class="margin-right">:</span>
          <span>{{ value.status.nodeInfo.kubeletVersion }}</span>
        </td>
        <td>
          <span>Kube-Proxy 版本</span>
          <span class="margin-right">:</span>
          <span>{{ value.status.nodeInfo.kubeProxyVersion }}</span>
        </td>
      </tr>
      <tr>
        <td>
          <span>机器 ID</span>
          <span class="margin-right">:</span>
          <span>{{ value.status.nodeInfo.machineID }}</span>
        </td>
        <td>
          <span>系统 UUID</span>
          <span class="margin-right">:</span>
          <span>{{ value.status.nodeInfo.systemUUID }}</span>
        </td>
      </tr>
      <tr>
        <td>
          <span>启动 ID</span>
          <span class="margin-right">:</span>
          <span>{{ value.status.nodeInfo.bootID }}</span>
        </td>
        <td>
          <span>容器运行时版本</span>
          <span class="margin-right">:</span>
          <span>{{ value.status.nodeInfo.containerRuntimeVersion }}</span>
        </td>
      </tr>
      <tr>
        <td>
          <span>操作系统</span>
          <span class="margin-right">:</span>
          <span>{{ value.status.nodeInfo.operatingSystem }}</span>
        </td>
        <td>
          <span>架构</span>
          <span class="margin-right">:</span>
          <span>{{ value.status.nodeInfo.architecture}}</span>
        </td>
      </tr>
      <tr>
        <td colspan="2">
          <span>污点 （Taints）</span>
          <span class="margin-right">:</span>
          <el-tag v-for="(train, index) in value.spec.taints" style="background: #999999; color: white" :key="index">{{ train.key }}: {{ train.value }} Effect: {{ train.effect }}</el-tag>
        </td>
      </tr>
      </tbody>
    </table>
    <div class="table-viewer-header clearfix" style="clear: both">
      <span class="table-viewer-topbar-title">状态</span>
    </div>
    <p></p>
    <div>
      <el-table
          :data="value.status.conditions"
          stripe
          style="width: 100%">
        <el-table-column
            label="类型"
            prop="type"
            type="scope">
          <template slot-scope="scope" >
            <span>{{ scope.row.type }} </span>
          </template>
        </el-table-column>
        <el-table-column
            label="状态"
            prop="status"
            type="scope">
          <template slot-scope="scope" >
            <span>{{ scope.row.status }} </span>
          </template>
        </el-table-column>
        <el-table-column
            label="更新时间"
            prop="status"
            type="scope">
          <template slot-scope="scope" >
            <span>{{ scope.row.lastTransitionTime |formatDate }} </span>
          </template>
        </el-table-column>
        <el-table-column
            label="内容"
            prop="reason"
            type="scope">
          <template slot-scope="scope" >
            <span v-if="scope.row.reason">{{ scope.row.reason }} </span>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column
            label="消息"
            prop="message"
            type="scope">
          <template slot-scope="scope" >
            <span v-if="scope.row.message">{{scope.row.message}}</span>
            <span v-else>-</span>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="table-viewer-header clearfix" style="clear: both">
      <span class="table-viewer-topbar-title">容器组</span>
    </div>
    <p></p>
<!--    <div class="container">-->
<!--      <el-tabs style="width:100%;">-->
<!--        <PodListBlock v-if="pods_list" :value="pods_list" :pageData="pageData" @pageinfo="handlePageData" @edit="handleEdit" @monitor="handleMonitor" @detail="handlePodDetail" @delete="handleDelete"/>-->
<!--      </el-tabs>-->
<!--    </div>-->

    <div v-if="dialogMonitorVisible">
      <el-dialog
          :visible.sync="dialogMonitorVisible"
          :title="podtitle"
          width="70%"
          append-to-body
          @close="handleMonitorCancel">
        <div class="container">
          <MonitorBlock
              style="width:100%;"
              ref="MonitorBlock"
              title="监控"
              :form="monitorcurrentValue"
              @cancel="handleMonitorCancel" />
        </div>
      </el-dialog>
    </div>

    <div v-if="dialogPodDetailVisible">
      <el-dialog
          :visible.sync="dialogPodDetailVisible"
          :title="podtitle"
          width="70%"
          append-to-body
          @close="handlePodDetailCancel">
        <div class="container">
          <el-tabs style="width:100%;">
            <PodDetailBlock
                ref="PodDetailBlock"
                :value="poddetailcurrentValue"
                @cancel="handlePodDetailCancel" />
          </el-tabs>
        </div>
      </el-dialog>
    </div>

    <div v-if="dialogDetailYamlVisible">
      <el-dialog
          :visible.sync="dialogDetailYamlVisible"
          :title="podtitle"
          width="70%"
          append-to-body
          @close="handleDetailYamlCancel">
        <div class="container">
          <YamlFormBlock
              style="width:100%;"
              ref="FormBlock"
              title="编辑 YAML"
              @submit="handleSubmit"
              :form="detailYamlcurrentValue"
              @cancel="handleDetailYamlCancel" />
        </div>
      </el-dialog>
    </div>
  </div>

</template>

<script>
import { formatTimeToStr } from '@/utils/date'
import { AgeFormat } from '@/utils/age'
// import PodListBlock from './podstable'
// import MonitorBlock from '../../k8s/workloads/pod/monitor'
// import PodDetailBlock from '../../k8s/workloads/pod/detail'
import YamlFormBlock   from '@/components/yaml/YamlBlock.vue'
// import { PodsList, podsDetail, podsUpdate, podsDelete, podsEvents} from '@/api/k8s/pod'
export default {
  name:"DetailBlock",
  components: {
    // PodListBlock,
    // MonitorBlock,
    // PodDetailBlock,
    YamlFormBlock,
  },
  props: {
    value: {
      type: Object,
      default: function() {
        return {}
      }
    }
  },
  data() {
    return {
      dialogMonitorVisible: false,
      dialogPodDetailVisible: false,
      dialogDetailYamlVisible: false,
      nodedetail: this.value,
      pods_list: [],
      podtitle: '',
      monitorcurrentValue: {},
      poddetailcurrentValue: {},
      detailYamlcurrentValue: {},
      information_id: 0,
      metadata: {},
      pageData: {
        page: 1,
        pageSize: 10,
        total: 0,
      },
    }
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time !== '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    },
    AgeformatDate: function(time) {
      if (time != null && time !== '') {
        return AgeFormat(time)
      } else {
        return ''
      }
    }
  },
  methods: {
    async  handlenodetail() {
      // this.information_id = this.nodedetail.information_id
      // const ret = await PodsList(this.pageData.page, this.pageData.pageSize,
      //     this.nodedetail.information_id, "","", "spec.nodeName="+this.nodedetail.metadata.name)
      // if(ret.code === 0){
      //   this.pods_list = ret.data.items
      //   this.pageData.page = ret.data.page
      //   this.pageData.pageSize = ret.data.pageSize
      //   this.pageData.total = ret.data.total
      // }
    },
    handlePageData(value){
      this.pageData = value
      this.handlenodetail()
    },
    async handleMonitor(value){
      this.podtitle = "Pod 名称: " + value.metadata.name
      this.monitorcurrentValue = value
      this.monitorcurrentValue['cluster_id'] = this.information_id
      this.dialogMonitorVisible = true
    },
    async handlePodDetail(value){
      this.podtitle = value.metadata.name
      const res = await podsDetail(this.information_id, value.metadata.namespace, value.metadata.name)
      if(res.code === 0){
        this.poddetailcurrentValue = res.data.items
        const resEvents = await podsEvents(this.information_id, value.metadata.namespace, value.metadata.name)
        if (resEvents.code === 0) {
          this.poddetailcurrentValue["events"] = resEvents.data
        }
        this.dialogPodDetailVisible = true
      }
    },
    async handleEdit(value){
      const res = await podsDetail(this.information_id, value.metadata.namespace, value.metadata.name)
      if(res.code === 0){
        this.detailYamlcurrentValue = res.data.items
      }
      this.podtitle = "编辑 YAML"
      this.dialogDetailYamlVisible = true
    },
    async handleSubmit(value){
      const res = await  podsUpdate(this.information_id, value.metadata.namespace, value.metadata.name, value)
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '更新成功',
          showClose: true
        })
        this.dialogDetailYamlVisible = false
      }

    },
    async handleDelete(value){
      console.log(value)
      this.$confirm('此操作将永久删除Pod, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        const res = await  podsDelete(this.information_id, value.metadata.namespace, value.metadata.name, value)
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '删除成功!'
          })
        }
        this.pods_list.splice(this.pods_list.indexOf(value),1);
      })
          .catch(() => {
            this.$message({
              type: 'info',
              message: '已取消删除'
            })
          })
    },
    handleMonitorCancel(){
      this.dialogMonitorVisible = false
    },
    handlePodDetailCancel(){
      this.dialogPodDetailVisible = false
    },
    async handleDetailYamlCancel(){
      this.dialogDetailYamlVisible = false
    },
  },
  created() {
    this.handlenodetail()
  }
}
</script>
<style scoped>
.label-custom {
  margin-right: 8px;
  margin-bottom: 8px;
  display: inline-block;
  padding: 0 8px;
  cursor: pointer;
  height: 24px;
  line-height: 24px;
  border-radius: 12px;
  -webkit-border-radius: 12px;
  -moz-border-radius: 12px;
  -ms-border-radius: 12px;
  -o-border-radius: 12px;
  border: 1px solid #bcbdbf;
}
</style>
