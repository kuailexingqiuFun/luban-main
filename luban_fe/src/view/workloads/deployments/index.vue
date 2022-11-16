<template>
<div>

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
</div>
</template>

<script>
import ListBlock from './table.vue'
export default {
  name: "index.vue",
  components: {
    ListBlock,
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
  methods: {
    handleCurrentChange(val) {
      console.log(val)
    },
    handleSizeChange(val) {
      console.log(val)
    },
    async handleEditYAML(value) {
      console.log(value)
    },
  }
}
</script>

<style scoped>

</style>
