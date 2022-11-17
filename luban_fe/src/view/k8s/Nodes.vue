<template>
  <div>
    <el-card>
      <el-table
          :data="tableData"
          style="width: 100%">
        <el-table-column
            prop="name"
            label="节点名称">
        </el-table-column>
        <el-table-column
            prop="name"
            label="状态">
        </el-table-column>
        <el-table-column
            prop="address"
            label="角色">
        </el-table-column>
        <el-table-column
            prop="address"
            label="版本">
        </el-table-column>
        <el-table-column
            prop="address"
            label="创建时间">
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script>
import {ListNodes} from "@/api/k8s";

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Nodes",
  created() {
    this.GetNodes("test")
  },
  methods: {
    GetNodes(clusterName) {
      ListNodes(clusterName).then((res) => {
        console.log(res, "xxx")
      })
    }
  }
}
</script>

<style scoped>

</style>