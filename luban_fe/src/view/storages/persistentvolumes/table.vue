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
          <span class="operate-span"  @click="handleDetail(scope.row)">{{ scope.row.metadata.name }} </span>
        </template>
      </el-table-column>
      <el-table-column
          label="状态"
          prop="status"
          type="scope">
        <template slot-scope="scope" >
          <el-tag type="success" v-if="scope.row.status.phase ==='Bound' || scope.row.status.phase ==='Available'">Bound</el-tag>
          <el-tag type="danger" v-else>{{scope.row.status.phase}}</el-tag>
        </template>
      </el-table-column>
      <el-table-column
          label="访问模式"
          prop="accessModes"
          type="scope">
        <template slot-scope="scope" >
          <el-tag v-for="(name,index) in scope.row.spec.accessModes " :key="index" type="info" size="mini">{{ name }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column
          label="卷类型"
          prop="status"
          type="scope">
        <template slot-scope="scope" >
          {{scope.row.spec.volumeMode }}
        </template>
      </el-table-column>
      <el-table-column
          label="回收策略"
          prop="persistentVolumeReclaimPolicy"
          type="scope">
        <template slot-scope="scope" >
          <span>{{ scope.row.spec.persistentVolumeReclaimPolicy }} </span>
        </template>
      </el-table-column>
      <el-table-column
          label="容量Gi"
          prop="storage"
          type="scope">
        <template slot-scope="scope" >
          <span>{{ scope.row.spec.capacity.storage }} </span>
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
