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
          width="278px"
          type="scope">
        <template slot-scope="scope" >
          <span class="operate-span" @click="handleDetail(scope.row)">{{ scope.row.metadata.name }} </span>
        </template>
      </el-table-column>
      <el-table-column
          label="状态"
          prop="status"
          width="100px"
          type="scope">
        <template slot-scope="scope" >
          <el-tag type="success" v-if="scope.row.status.phase ==='Bound' || scope.row.status.phase ==='Available'">Bound</el-tag>
          <el-tag type="danger" v-else>{{scope.row.status.phase}}</el-tag>
        </template>
      </el-table-column>
      <el-table-column
          label="命名空间"
          prop="namespace"
          type="scope">
        <template slot-scope="scope" >
          <span>{{ scope.row.metadata.namespace}} </span>
        </template>
      </el-table-column>
      <el-table-column
          label="Volume"
          prop="volumeName"
          type="scope">
        <template slot-scope="scope" >
          <span>{{ scope.row.spec.volumeName}} </span>
        </template>
      </el-table-column>
      <el-table-column
          label="容量Gi"
          prop="storage"
          type="scope">
        <template slot-scope="scope" >
          <span>{{ scope.row.spec.resources.requests.storage }} </span>
        </template>
      </el-table-column>
      <el-table-column
          label="存储类型"
          prop="storageClassName"
          type="scope">
        <template slot-scope="scope" >
          <span>{{ scope.row.spec.storageClassName }} </span>
        </template>
      </el-table-column>
      <el-table-column
          label="volumeMode"
          prop="volumeMode"
          type="scope">
        <template slot-scope="scope" >
          <span>{{ scope.row.spec.volumeMode }} </span>
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
          width="385">
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
