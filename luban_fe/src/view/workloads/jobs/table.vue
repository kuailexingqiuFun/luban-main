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
          <span class="operate-span" @click="handleDetail(scope.row)">{{ scope.row.metadata.name }} </span>
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
      <el-table-column label="状态" min-width="60">
        <template v-slot:default="{row}">
          <el-tag style="margin-left: 5px" v-if="row.status.active" type="info">active: *{{row.status.active}}</el-tag>
          <el-tag style="margin-left: 5px" v-if="row.status.succeeded" type="success">succeeded: {{row.status.succeeded}}</el-tag>
          <el-tag style="margin-left: 5px" v-if="row.status.failed" type="danger">failed: {{row.status.failed}}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="耗时" min-width="30">
        <template v-slot:default="{row}">
          {{ getDuration(row) }}
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
          <el-button @click="handleEdit(scope.row)" size="small" type="primary">编辑</el-button>
          <el-button type="danger" size="small" @click="handleDelete(scope.row)">删除</el-button>
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
    handleDetail(value) {
      this.$emit('detail', value)
    },
    handleDelete(value) {
      this.$emit('delete', value)
    },
    getDuration(row) {
      let startTime = new Date(row.status.startTime)
      if (!row.status.completionTime) {
        return "/"
      } else {
        let endTime = new Date(row.status.completionTime)
        let t = Math.floor((endTime - startTime) / 1000)
        if (t % 60 !== 0) {
          return (t % 60) + " mins"
        }
        if (t % 3600 !== 0) {
          return (t % 60) + " hours"
        }
        return Math.floor((endTime - startTime) / 1000) + "S"
      }
    },
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
