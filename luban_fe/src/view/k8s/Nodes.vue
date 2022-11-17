<template>
  <div>
    <el-card>
      <el-button type="primary" size="small">节点排水</el-button>
      <el-button type="danger" size="small">设置不可调度</el-button>
      <el-table
          :data="nodeData"
          style="width: 100%">
        <el-table-column
            prop="name"
            label="节点名称">
          <template v-slot="scope">
            <span>{{scope.row.name}}</span><br/>
            <span>{{scope.row.nodeIP}}</span>
          </template>
        </el-table-column>
        <el-table-column
            prop="ready"
            label="状态">
          <template v-slot="scope">
            <el-tag v-if="scope.row.ready" type="success">就绪</el-tag>
            <el-tag v-else type="danger">异常</el-tag>
          </template>
        </el-table-column>
        <el-table-column
            prop="role"
            label="角色">
        </el-table-column>
        <el-table-column
            prop="version"
            label="版本">
          <template v-slot="scope">
            <span>{{scope.row.version}}</span><br/>
            <span>{{scope.row.nodeInfo.containerRuntimeVersion}}</span><br/>
            <span>{{scope.row.nodeInfo.osImage}}</span>
          </template>
        </el-table-column>
        <el-table-column
            prop="createAt"
            label="创建时间">
        </el-table-column>
        <el-table-column
            fixed="right"
            label="操作"
            width="200"
        >
          <template v-slot="scope">
            <el-button size="small"  @click="GetNodeDetail(scope.row.name)">详情</el-button>
            <el-button type="danger" size="mini" @click="RemoveNode(scope.row.name)">移除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script>
import {ListNodes} from "@/api/k8s";
import {Message} from "element-ui";

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Nodes",
  created() {
    this.GetNodes("test")
  },
  methods: {
    GetNodes(clusterName) {
      ListNodes(clusterName).then((res) => {
        if (res.code === 0) {
          this.nodeData = res.data.nodes
        } else {
          Message.success(res.msg)
        }
      })
    }
  },
  data() {
    return {
      nodeData: [],
    }
  }
}
</script>

<style scoped>

</style>