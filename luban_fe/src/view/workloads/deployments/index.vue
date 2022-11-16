<template>
  <div>
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
        <el-form-item label="命名空间">
          <el-select v-model="searchInfo.namespace" size="mini" filterable placeholder="请选择命名空间" @change="NamespaceChange">
            <el-option
                v-for="item in namespace_list"
                :key="item.id"
                :label="item.name"
                :value="item.name"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-input v-model="searchInfo.name" size="mini" placeholder="名称" @change="Searchapp" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search">搜索</el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-refresh" @click="RefreshStatus">刷新</el-button>
          <el-button size="mini" type="primary" icon="el-icon-circle-plus" @click="handleYAMLAdd">YAML</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!--子组件表格-->
    <ListBlock :value="tableData" @edit="handleEditYAML" />

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

<!--    <div v-if="dialogYamlVisible">-->
<!--      <el-dialog-->
<!--          :visible.sync="dialogYamlVisible"-->
<!--          :title="title"-->
<!--          width="70%"-->
<!--          @close="handleYamlAddCancel"-->
<!--      >-->
<!--        <div class="container">-->
<!--          <YamlFormBlock-->
<!--              ref="FormBlock"-->
<!--              style="width:100%;"-->
<!--              :title="title"-->
<!--              :form="currentValue"-->
<!--              @submit="handleSubmit"-->
<!--              @cancel="handleYamlAddCancel"-->
<!--          />-->
<!--        </div>-->
<!--      </el-dialog>-->
<!--    </div>-->
  </div>
</template>

<script>
import { DeploymentsGet, DeploymentsList, DeploymentsUpdate } from '@/api/kubernetes/workloads/deployments'
import { getClusterList } from '@/api/kubernetes/clusters'
import { NamespaceList } from '@/api/kubernetes/namespaces'
// import YamlFormBlock from '@/components/yaml/YamlBlock.vue'
import ListBlock from './table.vue'
export default {
  name: 'Deployment',
  components: {
    ListBlock,
    // YamlFormBlock
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
      dialogYamlVisible: false
    }
  },
  created() {
    this.getCluster_list()
  },
  methods: {
    async getTableData(page = this.page, pageSize = this.pageSize, cluster_id = this.cluster_id, namespace = this.namespace, searchInfo = this.searchInfo.name) {
      const res = await DeploymentsList(cluster_id, page, pageSize, namespace, searchInfo)
      if (res.code === 0) {
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
          await this.NamespaceChange(this.namespace_list[0].id)
        }
      }
    },
    async getnamespace_list(cluster_id) {
      const res = await NamespaceList(cluster_id, '', '', '')
      if (res.code === 0) {
        this.namespace_list = []
        for (const ns of res.data.items) {
          const item = {
            id: Math.random(),
            name: ns.metadata.name
          }
          this.namespace_list.push(item)
        }
        this.namespace_list.push(this.defaultNamespace)
        this.searchInfo['namespace'] = this.namespace_list[0].name
      }
    },
    handleCurrentChange(val) {
      console.log(val)
    },
    handleSizeChange(val) {
      console.log(val)
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
      this.information_id = value
      await this.getnamespace_list(value)
      await this.getTableData(1, 10)
    },
    // 切换命名空间
    async NamespaceChange(value) {
      this.namespace = value
      await this.getTableData()
    },

    // 集群信息获取
    async getCluster_list() {
      const res = await getClusterList()
      if (res.data.code === 0) {
        this.cluster_list = res.data.data
        this.searchInfo['cluster_id'] = this.cluster_list[0].id
        this.cluster_id = this.cluster_list[0].id
        await this.getnamespace_list(this.cluster_list[0].id)
        // await this.getTableData(1, 10)
        // console.log(this.cluster_list)

      }
    },
    // 添加
    handleYAMLAdd() {
      console.log('111')
    },
    async handleEditYAML(value) {
      const res = await DeploymentsGet(this.cluster_id, value.metadata.namespace, value.metadata.name)
      if (res.code === 0) {
        this.currentValue = res.data.items
      }
      this.title = '编辑 YAML'
      this.dialogYamlVisible = true
    },
    handleYamlAddCancel() {
      this.dialogYamlVisible = false
    },
    async handleSubmit(value) {
      this.dialogYamlVisible = false
      const res = await DeploymentsUpdate(this.cluster_id, value.metadata.namespace, value.metadata.name, value)
      if (res.data.items.code) {
        this.$message({
          type: 'error',
          message: '更新失败: ' + res.data.items.reason + ': ' + res.data.items.message,
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
    }
  }
}
</script>

<style scoped>

</style>
