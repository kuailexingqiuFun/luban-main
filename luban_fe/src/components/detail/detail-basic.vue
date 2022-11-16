<template>
  <div>
    <h3>基本信息</h3>
    <table style="width: 100%" class="myTable">
      <tr>
        <td>名称</td>
        <td colspan="4">{{ item.metadata.name }}</td>
      </tr>
      <tr v-if="item.metadata.namespace">
        <td>命名空间</td>
        <td colspan="4">{{ item.metadata.namespace }}</td>
      </tr>
      <tr v-if="item.metadata.finalizers">
        <td>Finalizers</td>
        <td colspan="4">
          <div v-for="value in item.metadata.finalizers" v-bind:key="value" class="myTag">
            <el-tag type="info" size="small">
              {{ value.length > 100 ? value.substring(0, 100) + "..." : value }}
            </el-tag>
          </div>
        </td>
      </tr>
      <tr v-if="hasPodContainers()">
        <td>镜像名称</td>
        <td colspan="4">
          <div v-for="(item,index) in containers" v-bind:key="index" class="myTag">
            <el-tag type="info" size="small">
              {{ item.image }}
            </el-tag>
          </div>
        </td>
      </tr>
      <tr v-if="item.kind === 'Namespace'">
        <td>Resource Quotas</td>
        <td colspan="4">
          <div v-for="(item,index) in resourceQuotas" v-bind:key="index" class="my-tag">
            <el-link>{{item.metadata.name}}</el-link>
          </div>
        </td>
      </tr>
      <tr v-if="item.kind === 'Namespace'">
        <td>Limit Ranges</td>
        <td colspan="4">
          <div v-for="(item,index) in limitRanges" v-bind:key="index" class="my-tag">
            <el-link>
              {{item.metadata.name}}
            </el-link>
          </div>
        </td>
      </tr>
      <tr>
        <td>存活时间</td>
        <td colspan="4">{{ item.metadata.creationTimestamp | formatDate }}</td>
      </tr>
      <tr>
        <td>标签</td>
        <td colspan="4">
          <detail-key-value v-if="Object.keys(item.metadata).length > 0" :valueObj="item.metadata.labels"></detail-key-value>
        </td>
      </tr>
      <tr>
        <td>注解</td>
        <td colspan="4">
          <detail-key-value v-if="Object.keys(item.metadata).length > 0" :valueObj="item.metadata.annotations"></detail-key-value>
        </td>
      </tr>
    </table>
  </div>
</template>

<script>
import DetailKeyValue from "@/components/detail/detail-key-value"
import { AgeFormat } from '@/utils/age'
export default {
  name: "DetailBasic",
  components: { DetailKeyValue },
  props: {
    item: Object,
    yamlShow: Boolean,
  },
  data() {
    return {
      show: false,
      containers: [],
      limitRanges: [],
      resourceQuotas: [],
      cluster: ""
    }
  },
  methods: {
    hasPodContainers() {
      if (this.item.spec?.template?.spec?.containers) {
        this.containers = this.item.spec.template.spec.containers
        return true
      } else if (this.item.spec?.jobTemplate?.spec?.template?.spec?.containers) {
        this.containers = this.item.spec.jobTemplate.spec.template.spec.containers
        return true
      } else {
        return false
      }
    },
  },
  created() {
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time !== '') {
        return AgeFormat(time)
      } else {
        return ''
      }
    }
  },
}
</script>

<style scoped>
</style>
