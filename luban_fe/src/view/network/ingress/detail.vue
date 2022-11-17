<template>
  <div class="grid-content bg-purple-light">
    <el-row :gutter="24" class="row-box">
      <el-col :span="24">
        <el-card class="el-card">
          <el-col :span="24">
            <table style="width: 90%" class="myTable">
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
                    <el-tag type="warning" size="small">
                      {{ key }} = {{ value }}
                    </el-tag>
                  </div>
                </td>
              </tr>
              <tr>
                <td>注解:</td>
                <td colspan="4">
                  <div v-for="(value,key,index) in value.metadata.annotations" v-bind:key="index" class="myTag">
                    <el-tag type="success" size="small" v-if="value.length < 5">
                      {{ key }} = {{ value }}
                    </el-tag>
                    <el-tooltip v-if="value.length > 50" :content="value" placement="top">
                      <el-tag type="info" size="small" v-if="value.length >= 5">
                        {{ key }} = {{ value.substring(0, 50) + "..." }}
                      </el-tag>
                    </el-tooltip>
                  </div>
                </td>
              </tr>
            </table>
          </el-col>
        </el-card>
      </el-col>
    </el-row>
    <el-tabs style="margin-top: 20px" tab-position="top" type="border-card">
      <el-tab-pane lazy label="规则">
        <div>
          <div>
            <el-alert v-if="value.length === 0" type="info">
              <div slot="title">
                <i class="el-icon-info"></i>
                <span>此处显示与 {0} 同名的 Ingress: {1}，但是您并未定义此 Ingress。</span>
              </div>
            </el-alert>
            <div>
              <el-table v-if="value.spec.rules.length > 0" :data="value.spec.rules">
                <el-table-column label="域名" prop="host" />
                <el-table-column
                    label="路径"
                    prop="http"
                    type="scope">
                  <template slot-scope="scope" >
                    <div v-for="(item, index) in scope.row.http.paths" :key="index">
                      <span>{{ item.path }}</span>
                    </div>
                  </template>
                </el-table-column>
                <el-table-column
                    label="服务名"
                    prop="http"
                    type="scope">
                  <template slot-scope="scope" >
                    <div v-for="(item, index) in scope.row.http.paths" :key="index">
                      <span>{{ item.backend.serviceName }}</span>
                    </div>
                  </template>
                </el-table-column>
                <el-table-column
                    label="端口"
                    prop="http"
                    type="scope">
                  <template slot-scope="scope" >
                    <div v-for="(item, index) in scope.row.http.paths" :key="index">
                      <span>{{ item.backend.servicePort }}</span>
                    </div>
                  </template>
                </el-table-column>
              </el-table>
              <br><br>
            </div>
          </div>
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
