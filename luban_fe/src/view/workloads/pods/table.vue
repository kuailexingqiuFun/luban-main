<template>
  <div>
    <el-table
        :data="value"
        stripe
        style="width: 100%">
      <el-table-column
          label="名称"
          prop="name"
          type="scope">
        <template slot-scope="scope" >
          <img style="width:14px;margin-right:2px" src="//g.alicdn.com/aliyun/cos/1.38.27/images/icon_docker.png">
          <span class="operate-span" @click="handleDetail(scope.row)">{{ scope.row.metadata.name }} </span>
        </template>
      </el-table-column>
      <el-table-column
          label="容器"
          prop="container"
          type="scope">
        <template slot-scope="scope" >
          <div class="gridChart" v-for="(p, index) in scope.row.status.containerStatuses" :key="index">
            <div v-if="p.ready" class="success">
              <el-tooltip class="item" effect="dark"  placement="top">
                <div slot="content">
                  {{ p.name}}（running, ready）
                </div>
                <div aria-haspopup="true" aria-expanded="false"></div>
              </el-tooltip>
            </div>
            <div v-else-if="!p.ready && p.state.waiting" class="danger">
              <el-tooltip class="item" effect="dark"  placement="top">
                <div slot="content">
                  {{ p.name}} 报错信息: {{ p.state.waiting.reason}}
                </div>
                <div aria-haspopup="true" aria-expanded="false"></div>
              </el-tooltip>
            </div>
            <div v-else-if="!p.ready && p.state.terminated" class="info">
              <el-tooltip class="item" effect="dark"  placement="top">
                <div slot="content">
                  {{ p.name}} {{ p.state.terminated.reason}}
                </div>
                <div aria-haspopup="true" aria-expanded="false"></div>
              </el-tooltip>
            </div>
            <div v-else class="danger">
              <el-tooltip class="item" effect="dark"  placement="top">
                <div slot="content">
                  {{ p.name}}
                </div>
                <div aria-haspopup="true" aria-expanded="false"></div>
              </el-tooltip>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column
          label="状态"
          prop="status"
          type="scope">
        <template slot-scope="scope" >
          <el-tag type="success" v-if="scope.row.status.phase ==='Running'">Running</el-tag>
          <el-tag type="info" v-else-if="scope.row.status.phase ==='Succeeded'">Completed</el-tag>
          <el-tag type="danger" v-else>{{scope.row.status.phase}}</el-tag>
        </template>
      </el-table-column>
      <el-table-column
          label="重启次数"
          prop="restarts"
          type="scope">
        <template slot-scope="scope" >
          <span>{{scope.row.status.containerStatuses  | restartCounts}}</span>
        </template>
      </el-table-column>
      <el-table-column
          label="命名空间"
          prop="namespace"
          type="scope">
        <template slot-scope="scope" >
          <span>{{scope.row.metadata.namespace}}</span>
        </template>
      </el-table-column>
      <el-table-column
          label="Pod IP"
          prop="Pod_IP"
          type="scope">
        <template slot-scope="scope" >
          <span>{{scope.row.status.podIP}}</span>
        </template>
      </el-table-column>
      <el-table-column
          label="调度节点"
          prop="hostIP"
          type="scope">
        <template slot-scope="scope" >
          <span  class="operate-span">{{scope.row.status.hostIP}}</span>
        </template>
      </el-table-column>
      <el-table-column
          label="存活时间"
          prop="creationTimestamp"
          type="scope">
        <template slot-scope="scope" >
          <span>{{scope.row.metadata.creationTimestamp |formatDate }}</span>
        </template>
      </el-table-column>
      <el-table-column
          fixed="right"
          label="操作"
          width="200">
        <template slot-scope="scope">
          <el-form :inline="true"  class="demo-form-inline">
            <el-form-item>
              <el-dropdown>
                    <span class="operate-span">
                      更多<i class="el-icon-arrow-down"></i>
                    </span>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item @click.native="handleEdit(scope.row)">编辑</el-dropdown-item>
                  <div v-for="(item,index) in scope.row.spec.containers" :key="index">
                    <el-dropdown-item @click.native="handleTerminal(scope.row, item.name)">{{  item.name }}终端</el-dropdown-item>
                    <el-dropdown-item @click.native="handleLogs(scope.row, item.name)">{{  item.name }}日志</el-dropdown-item>
                  </div>
                  <el-dropdown-item @click.native="handleDelete(scope.row)">删除</el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>
import { AgeFormat } from '@/utils/age'
export default {
  name: 'ListBlock',
  props: {
    value: {
      type: Array,
      default: function() {
        return []
      }
    }
  },
  data() {
    return {
      name: ''
    }
  },
  methods: {
    handleEdit(value) {
      this.$emit('edit', value)
    },
    handleTerminal(value, name) {
      this.$emit('terminal', value, name)
    },
    handleLogs(value, name) {
      this.$emit('logs', value, name)
    },
    handleDetail(value) {
      this.$emit('detail', value)
    },
    handleDelete(value) {
      this.$emit('delete', value)
    }
  },
  filters:{
    restartCounts: function(restarts) {
      let restart = 0
      if (restarts) {
        for (let i = 0; i < restarts.length; i++) {
          restart += restarts[i].restartCount
        }
      }
      return restart
    },
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
.operate-span {
  color: #409EFF;
  cursor: pointer;
  font-weight:bold;
}
</style>
