<template>
  <div>
    <el-card>
    <div class="search-term">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item label="集群">
          <el-select v-model="searchInfo.cluster_id" size="mini" filterable placeholder="请选择集群" @change="ClusterChange">
            <el-option
                v-for="item in cluster_list"
                :key="item.id"
                :label="item.clusterName"
                :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-input v-model="searchInfo.name" size="mini" placeholder="名称" @change="Searchapp" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-refresh" @click="RefreshStatus">刷新</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!--子组件表格-->
    <ListBlock :value="tableData" @edit="handleEditYAML" @detail="handleDetail" @delete="handleDelete" />

    <!--分页-->
    <el-pagination
        class="pagination"
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :style="{float:'right',padding:'20px'}"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
    />

    <div v-if="dialogAddYamlVisible">
      <el-dialog
          :visible.sync="dialogAddYamlVisible"
          :title="title"
          width="70%"
          append-to-body
          @close="handleYamlAddCancel"
      >
        <div class="container">
          <YamlFormBlock
              ref="FormBlock"
              style="width:100%;"
              :title="title"
              :form="currentValue"
              @submit="handleSubmitAdd"
              @cancel="handleYamlAddCancel"
          />
        </div>
      </el-dialog>
    </div>

    <div v-if="dialogYamlVisible">
      <el-dialog
          :visible.sync="dialogYamlVisible"
          :title="title"
          width="70%"
          append-to-body
          @close="handleYamlAddCancel"
      >
        <div class="container">
          <YamlFormBlock
              ref="FormBlock"
              style="width:100%;"
              :title="title"
              :form="currentValue"
              @submit="handleSubmit"
              @cancel="handleYamlAddCancel"
          />
        </div>
      </el-dialog>
    </div>
    <div v-if="dialogDetailVisible">
      <el-dialog
          :visible.sync="dialogDetailVisible"
          :title="title"
          width="70%"
          append-to-body
          @close="handleDetailCancel">
        <div class="container">
          <el-tabs style="width:100%;">
            <DetailBlock
                ref="DetailBlock"
                :value="currentValue"
                @cancel="handleDetailCancel" />
          </el-tabs>
        </div>
      </el-dialog>
    </div>
    </el-card>
  </div>
</template>

<script>
import { NodesList, NodesDelete, NodesGet, NodesUpdate } from '@/api/kubernetes/nodes'
import { NodeMetricsList  } from '@/api/kubernetes/metrics'
import { cpuUnitConvert, memoryUnitConvert } from '@/utils/unitConvert'
import { getClusterList } from '@/api/kubernetes/clusters'
import YamlFormBlock from '@/components/yaml/YamlBlock.vue'
import ListBlock from './table.vue'
import DetailBlock from './detail.vue'
export default {
  name: 'Deployment',
  components: {
    ListBlock,
    YamlFormBlock,
    DetailBlock
  },
  data() {
    return {
      tableData: [],
      title: '',
      pageSize: 10,
      page: 1,
      total: 0,
      searchInfo: {
        'name': ''
      },
      defaultNamespace: {
        'id': 0,
        'name': 'All Namespaces'
      },
      cluster_id: 1,
      currentValue: {},
      cluster_list: [],
      namespace_list: [],
      namespace: '',
      dialogYamlVisible: false,
      dialogAddYamlVisible: false,
      dialogDetailVisible: false,
    }
  },
  created() {
    this.getCluster_list()
  },
  methods: {
    async getTableData(page = this.page, pageSize = this.pageSize, cluster_id = this.cluster_id, searchInfo = this.searchInfo.name) {
      const res = await NodesList(cluster_id, page, pageSize, searchInfo)
      if (res.code === 0) {
        const metrics = await NodeMetricsList(this.cluster_id, { ...this.searchInfo })
        if (metrics.code === 0) {
          for (const node of res.data.items) {
            for (const item of metrics.data.items) {
              if (node.metadata.name === item.metadata.name) {
                if (item.usage?.cpu) {
                  node.cpuUsage = cpuUnitConvert(item.usage.cpu) + 'm'
                  node.cpuUsagePersent = Math.round((cpuUnitConvert(item.usage.cpu) / cpuUnitConvert(node.status.allocatable.cpu)).toFixed(2) * 100)
                }
                if (item.usage?.memory) {
                  node.memoryUsage = memoryUnitConvert(item.usage.memory).toFixed(2) + 'Mi'
                  node.memoryUsagePersent = Math.round((memoryUnitConvert(item.usage.memory) / memoryUnitConvert(node.status.allocatable.memory)).toFixed(2) * 100)
                }
              }
            }
          }
        }
        this.tableData = res.data.items
        this.total = res.data.total
        this.page = res.data.page
        this.pageSize = res.data.pageSize
        if (this.cluster_id === 0) {
          this.searchInfo['cluster_id'] = this.cluster_list[0].id
          this.searchInfo['namespace'] = this.namespace_list[0].id
          this.cluster_id = this.cluster_list[0].id
          this.namespace = this.namespace_list[0].name
          await this.ClusterChange(this.cluster_list[0].id)
        }
      }
    },
    handleCurrentChange(val) {
      this.page = val
      this.getTableData()
    },
    handleSizeChange(val) {
      this.pageSize = val
      this.getTableData()
    },
    // 提交搜索
    onSubmit(){
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    async RefreshStatus() {
      if (this.cluster_id !== '') {
        await this.getTableData(1, 10)
      } else {
        this.$message({
          type: 'warning',
          message: '请选择集群!'
        })
      }
    },
    async Searchapp(value) {
      this.name = value
      if (this.cluster_id !== '') {
        await this.getTableData(1, 10)
      } else {
        this.$message({
          type: 'warning',
          message: '请选择集群!'
        })
      }
    },
    // 切换集群
    async ClusterChange(value) {
      this.cluster_id = value
      await this.getTableData(1, 10)
    },
    // 集群信息获取
    async getCluster_list() {
      const res = await getClusterList()
      if (res.code === 0) {
        this.cluster_list = res.data
        this.searchInfo['cluster_id'] = this.cluster_list[0].id
        this.cluster_id = this.cluster_list[0].id
        await this.getTableData(1, 10)
      }
    },
    async handleEditYAML(value) {
      const res = await NodesGet(this.cluster_id, value.metadata.name)
      if (res.code === 0) {
        this.currentValue = res.data.items
      }
      this.title = '编辑 YAML'
      this.dialogYamlVisible = true
    },
    handleYamlAddCancel() {
      this.dialogYamlVisible = false
      this.dialogAddYamlVisible = false
    },
    async handleSubmit(value) {
      this.dialogYamlVisible = false
      const res = await NodesUpdate(this.cluster_id, value.metadata.name, value)
      if (res.code !== 0) {
        this.$message({
          type: 'error',
          message: '更新失败: ' + res.items.reason + ': ' + res.items.message,
          showClose: true
        })
      } else {
        this.$message({
          type: 'success',
          message: '更新成功',
          showClose: true
        })
        this.handleYamlAddCancel()
        this.getTableData()
      }
    },
    async handleDelete(value){
      this.$confirm('此操作将永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        const res = await  NodesDelete(this.cluster_id, value.metadata.name, value)
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '删除成功!'
          })
          await this.getTableData()
        }
      })
          .catch(() => {
            this.$message({
              type: 'info',
              message: '已取消删除'
            })
          })
    },
    handleDetailCancel(){
      this.dialogDetailVisible = false
    },
    async handleDetail(value) {
      this.title = value.metadata.name
      const res = await NodesGet(this.cluster_id, value.metadata.name)
      if(res.code === 0){
        this.currentValue = res.data.items
        this.dialogDetailVisible = true
      }
    },
  }
}
</script>

<style scoped>

</style>
