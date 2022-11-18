<template>
  <div class="grid-content bg-purple-light">
    <el-row :gutter="24" class="row-box">
      <el-col :span="24">
        <el-card class="el-card">
          <el-col :span="24">
            <table style="width: 100%" class="myTable">
              <th scope="col" width="100%" align="left">
                <h3>基本信息</h3>
              </th>
              <tr>
                <td>名称:</td>
                <td colspan="4">{{ value.metadata.name }}</td>
              </tr>
              <tr>
                <td>命名空间:</td>
                <td colspan="4">{{ value.metadata.namespace }}</td>
              </tr>
              <tr>
                <td>创建时间:</td>
                <td colspan="4">{{   value.metadata.creationTimestamp |formatDate    }}</td>
              </tr>
              <tr>
                <td>标签:</td>
                <td colspan="4">
                  <div v-for="(value,key,index) in value.metadata.labels" v-bind:key="index" class="myTag">
                    <el-tag type="success" size="small" v-if="value.length < 50">
                      {{ key }} = {{ value }}
                    </el-tag>
                    <el-tooltip v-if="value.length > 50" :content="value" placement="top">
                      <el-tag type="info" size="small" v-if="value.length >= 50">
                        {{ key }} = {{ value.substring(0, 50) + "..." }}
                      </el-tag>
                    </el-tooltip>
                  </div>
                </td>
              </tr>
              <tr>
                <td>注解:</td>
                <td colspan="4">
                  <div v-for="(value,key,index) in value.metadata.annotations" v-bind:key="index" class="myTag">
                    <el-tag type="success" size="small" v-if="value.length < 50">
                      {{ key }} = {{ value }}
                    </el-tag>
                    <el-tooltip v-if="value.length > 50" :content="value" placement="top">
                      <el-tag type="info" size="small" v-if="value.length >= 50">
                        {{ key }} = {{ value.substring(0, 50) + "..." }}
                      </el-tag>
                    </el-tooltip>
                  </div>
                </td>
              </tr>
              <tr>
                <td>类型:</td>
                <td colspan="4">
                  <el-tag type="info" size="small">
                    {{ value.spec.type }}
                  </el-tag>
                </td>
              </tr>
              <tr>
                <td>集群IP:</td>
                <td colspan="4">
                  <el-tag size="small">
                    {{ value.spec.clusterIP }}
                  </el-tag>
                </td>
              </tr>
              <tr>
                <td>Session Affinity:</td>
                <td colspan="4">
                  <el-tag type="warning" size="small">
                    {{ value.spec.sessionAffinity }}
                  </el-tag>
                </td>
              </tr>
              <tr>
                <td>选择器:</td>
                <td colspan="4">
                  <div v-for="(value,key,index) in value.spec.selector" v-bind:key="index" class="myTag">
                    <el-tag type="danger" size="small">
                      {{ key }} = {{ value }}
                    </el-tag>
                  </div>
                </td>
              </tr>
            </table>
          </el-col>
        </el-card>
      </el-col>
    </el-row>
    <el-tabs style="margin-top: 20px" tab-position="top" type="border-card">
      <el-tab-pane lazy label="端口">
        <div>
          <el-table v-if="value.spec.ports.length > 0" :data="value.spec.ports">
            <el-table-column
                label="名称"
                prop="name">
            </el-table-column>
            <el-table-column
                label="端口"
                prop="port">
            </el-table-column>
            <el-table-column
                label="协议"
                prop="protocol">
            </el-table-column>
            <el-table-column
                label="目标端口"
                prop="targetPort">
            </el-table-column>
            <el-table-column
                label="Node Port"
                prop="nodePort">
            </el-table-column>
          </el-table>
        </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import { formatTimeToStr } from '@/utils/date'
export default {
  name: "DetailBlock",
  props: {
    value: {
      type: Object,
      default: function() {
        return {}
      }
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
    }
  },
}
</script>

<style scoped>

</style>
