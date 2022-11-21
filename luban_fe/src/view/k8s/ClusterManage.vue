<template>
  <div>
    <el-card>
      <div>
        <el-button type="primary" size="medium" @click="RegisterCluster">注册集群</el-button>
        <el-table
            :data="clusterData"
            style="width: 100%">
          <el-table-column
              prop="id"
              label="ID">
          </el-table-column>

          <el-table-column
              prop="clusterName"
              label="集群名称"
              column-key="clusterName">
          </el-table-column>
          <el-table-column
              prop="nodeNumber"
              label="节点数量"

          >
            <template v-slot="scope">
              <el-tag>{{scope.row.nodeNumber}}</el-tag>
            </template>
          </el-table-column>
          <el-table-column
              prop="api_address"
              label="APIServer"
              sortable
          >
          </el-table-column>
          <el-table-column
              prop="prometheus_url"
              label="普罗米修斯地址">
          </el-table-column>
          <el-table-column
              fixed="right"
              label="操作"
              width="200"
          >
            <template v-slot="scope">
              <el-button size="small"  @click="editCluster(scope.row.id)">编辑</el-button>
              <el-button type="danger" size="mini" @click="deleteCluster(scope.row.id)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>

    </el-card>
    <div>
      <el-dialog title="添加集群" :visible.sync="addClusterVisible" width="600px" :modal="false" :destroy-on-close="true">
        <el-form :model="clusterForm" label-width="100px" :rules="clusterRules" ref="clusterForm">
          <el-form-item label="集群名称" prop="clusterName">
            <el-input v-model="clusterForm.clusterName" autocomplete="off" placeholder="请输入集群名称"></el-input>
          </el-form-item>
          <el-form-item label="集群凭证" prop="kubeConfig">
            <el-input
                type="textarea"
                :rows="10"
                placeholder="请输入内容"
                v-model="clusterForm.kubeConfig">
            </el-input>
          </el-form-item>
          <el-form-item label="APIServer" prop="apiAddress">
            <el-input v-model="clusterForm.api_address" autocomplete="off" placeholder="请输入apiServer地址"></el-input>
          </el-form-item>
          <el-form-item label="普罗米修斯" prop="prometheus_url">
            <el-input v-model="clusterForm.prometheus_url" autocomplete="off" placeholder="请输入普罗米修斯地址"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="addClusterVisible = false">取 消</el-button>
          <el-button type="primary" @click="AddCluster('clusterForm')" >确 定</el-button>
        </div>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import {createCluster, getClusterList} from "@/api/kubernetes/clusters";
import {Message} from "element-ui";
import {DeleteCluster, GetClusterDetail} from "@/api/k8s";

export default {
  name: "ClusterManage.vue",
  inject: ['reload'],
  created() {
    this.GetCluster()
  },
  mounted() {
    console.log(this.$route.path)
  },
  methods: {
    GetCluster() {
      getClusterList().then((res) => {
        this.clusterData = res.data
      })
    },
    RegisterCluster() {
      this.addClusterVisible = true
      this.clusterForm.clusterName = ''
      this.clusterForm.kubeConfig = ''
      this.clusterForm.api_address = ''
      this.clusterForm.prometheus_url = ''
    },
    AddCluster(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          console.log(this.clusterForm, "===")
          createCluster(this.clusterForm).then((res) => {
            if (res.code ===0) {
              Message.success("集群注册成功")
              this.addClusterVisible = false
              this.reload()
            } else {
              Message.error(res.msg)
            }
          })
        } else {
          return false
        }
      })
    },
    // 编辑集群凭证
    editCluster(clusterId) {
      GetClusterDetail(clusterId).then((res) => {
        if (res.code === 0) {
          this.addClusterVisible = true
          this.clusterForm = res.data
        }
      })
    },
    deleteCluster(clusterId) {
      DeleteCluster({
        id: clusterId
      }).then((res) => {
        if (res.code ===0){
          Message.success("删除成功")
          this.reload()
        }else {
          Message.error(res.msg)
        }
      })
    }
  },
  data() {
    return {
      clusterData: [],
      addClusterVisible: false,
      clusterForm: {
        clusterName: '',
        kubeConfig: '',
        apiAddress: '',
        prometheus_url: ''
      },
      clusterRules: {
        clusterName: [
          {required: true, message: '请输入集群名称', trigger: 'blur'}
        ],
        kubeConfig: [
          {required: true, message: '请输入集群凭证', trigger: 'blur'}
        ],
        api_address: [
          {required: true, message: '请输入apiserver地址', trigger: 'blur'}
        ],
        prometheus_url: [
          {required: true, message: '请输入普罗米修斯地址', trigger: 'blur'}
        ]
      }
    }
  }
}
</script>

<style scoped>

</style>
