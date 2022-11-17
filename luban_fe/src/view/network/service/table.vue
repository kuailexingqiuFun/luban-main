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
          <span class="operate-span"  @click="handleDetail(scope.row)">{{ scope.row.metadata.name }} </span>
        </template>
      </el-table-column>
      <el-table-column
          label="命名空间"
          prop="namespace"
          width="178px"
          type="scope">
        <template slot-scope="scope">
          {{ scope.row.metadata.namespace }}
        </template>
      </el-table-column>
      <el-table-column
          label="标签"
          prop="labels"
          width="300px"
          type="scope">
        <template slot-scope="scope">
              <span v-for="(k, v, index) in scope.row.metadata.labels" :key="index">
                <span class="label-custom wd" type="info">{{ k }}: {{ v }}</span>
              </span>
        </template>
      </el-table-column>
      <el-table-column
          label="类型"
          prop="type"
          width="178px"
          type="scope">
        <template slot-scope="scope">
          {{ scope.row.spec.type }}
        </template>
      </el-table-column>
      <el-table-column
          label="创建时间"
          prop="creationTimestamp"
          width="178px"
          type="scope">
        <template slot-scope="scope" >
          <span>{{scope.row.metadata.creationTimestamp |formatDate }}</span>
        </template>
      </el-table-column>
      <el-table-column
          label="IP"
          prop="clusterIP"
          width="178px"
          type="scope">
        <template slot-scope="scope" >
          <span>{{scope.row.spec.clusterIP }}</span>
        </template>
      </el-table-column>
      <el-table-column
          label="内部端点"
          prop="endpoint"
          width="300px"
          type="scope">
        <template slot-scope="scope" >
          <div v-if="scope.row.spec.ports">
            <div :key="item.index" v-for="item in scope.row.spec.ports">
              <span v-if="item.port">{{ scope.row.metadata.name }}: {{ item.port}} {{ item.protocol}} </span><br/>
              <span v-if="item.nodePort">{{ scope.row.metadata.name }}: {{ item.nodePort}}  {{ item.protocol}}</span> <br/>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column
          label="外部端点"
          prop="endpoint"
          width="300px"
          type="scope">
        <template slot-scope="scope" >
          <div v-if="scope.row.status.loadBalancer.ingress">
            <div :key="ingress.index" v-for="ingress in scope.row.status.loadBalancer.ingress">
              <div :key="item.index" v-for="item in scope.row.spec.ports">
                <span v-if="item.port">{{ ingress.ip }}: {{ item.port}} </span><br/>
              </div>
            </div>
          </div>
          <div v-else>-</div>
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
