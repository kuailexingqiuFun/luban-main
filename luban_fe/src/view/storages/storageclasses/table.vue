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
          <span >{{ scope.row.metadata.namespace }} </span>
        </template>
      </el-table-column>
      <el-table-column
          label="创建时间"
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
