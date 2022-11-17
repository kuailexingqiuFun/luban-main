<template>
  <div>
    <el-table
        :data="value"
        :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
        border
        row-key="ID"
        stripe
        size="mini"
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
          label="提供商"
          prop="provisioner"
          type="scope">
        <template slot-scope="scope" >
          <span>{{ scope.row.provisioner }} </span>
        </template>
      </el-table-column>
      <el-table-column
          label="回收策略"
          prop="reclaimPolicy"
          type="scope">
        <template slot-scope="scope" >
          <span>{{ scope.row.reclaimPolicy }} </span>
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
          label="操作">
        <template slot-scope="scope">
          <el-button @click="handleEdit(scope.row)" size="small" type="primary">编辑</el-button>
          <el-button @click="handleDelete(scope.row)" size="small" type="danger">删除</el-button>
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
