<template>
  <div>
    <el-table
      :data="value"
      stripe
      style="width: 100%"
    >
      <el-table-column
        label="名称"
        prop="name"
        type="scope"
      >
        <template slot-scope="scope">
          <span class="operate-span" @click="handleDetail(scope.row)">{{ scope.row.metadata.name }} </span>
        </template>
      </el-table-column>
      <el-table-column
        label="命名空间"
        prop="namespace"
        type="scope"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.metadata.namespace }}</span>
        </template>
      </el-table-column>
      <el-table-column
          label="容器组数量"
          prop="status"
          type="scope">
        <template slot-scope="scope" >
          <el-tag type="success" v-if="scope.row.status.readyReplicas / scope.row.status.replicas === 1"> {{ scope.row.status.readyReplicas || 0 }} / {{ scope.row.status.replicas }}</el-tag>
          <el-tag type="danger" v-else> {{ scope.row.status.readyReplicas || 0 }} / {{ scope.row.status.replicas }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        label="存活时间"
        prop="creationTimestamp"
        type="scope"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.metadata.creationTimestamp |formatDate }}</span>
        </template>
      </el-table-column>
      <el-table-column
        fixed="right"
        label="操作"
        width="200"
      >
        <template slot-scope="scope">
          <el-button size="small" type="primary" @click="handleEdit(scope.row)">编辑</el-button>
          <el-button type="danger" size="mini" @click="handleDelete(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>
import { AgeFormat } from '@/utils/age'
export default {
  name: 'ListBlock',
  filters: {
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
  },
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
