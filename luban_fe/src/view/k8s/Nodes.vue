<template>
  <div>
    <el-card>
      <el-button type="primary" size="small" @click="setCordon">节点排水</el-button>
      <el-button type="danger" size="small" @click="setUnschedule">设置不可调度</el-button>

      <div class="search-term">
        <el-form :inline="true" class="demo-form-inline">
          <el-form-item label="集群">
            <el-select v-model="searchInfo.clusterName" size="mini" filterable placeholder="请选择集群" @change="clusterChange">
              <el-option
                  v-for="item in cluster_list"
                  :key="item.id"
                  :label="item.clusterName"
                  :value="item.clusterName"
              />
            </el-select>
          </el-form-item>

          <el-form-item>
            <el-input v-model="searchInfo.name" size="mini" placeholder="名称" @change="Search" />
          </el-form-item>
          <el-form-item>
            <el-button size="mini" type="primary" icon="el-icon-search">搜索</el-button>
          </el-form-item>
        </el-form>
      </div>

      <el-table :data="nodeData" style="width: 100%" ref="multipleTable" @selection-change="handleSelectionChange">
        <el-table-column
            type="selection"
            width="55">
        </el-table-column>

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
            <el-button size="small"  @click="getNodeDetail(scope.row.name)">详情</el-button>
            <el-button type="danger" size="mini" @click="removeNode(scope.row.name)">移除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!--节点详情-->
    <div>
      <el-dialog title="节点详情" :visible.sync="NodeVisible" width:="100%" :modal="false" :destroy-on-close="true" :close-on-click-modal="false">
        <el-descriptions :column="2" border>
          <!--基础信息-->
          <el-descriptions-item>
            <template slot="label">
              节点名称
            </template>
            {{this.nodeDetailData?.name}}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              创建时间
            </template>
            {{ this.nodeDetailData?.createAt }}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              UUID
            </template>
            {{this.nodeDetailData?.nodeInfo?.systemUUID}}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              容器子网
            </template>
            {{this.nodeDetailData?.podCIDR}}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              调度状态
            </template>
            <el-tag size="small">{{ this.nodeDetailData?.ready }}</el-tag>
          </el-descriptions-item>

          <el-descriptions-item>
            <template slot="label">
              IP地址
            </template>
            {{this.nodeDetailData?.nodeIP}}
          </el-descriptions-item>

          <el-descriptions-item>
            <template slot="label">
              系统镜像
            </template>
            {{ this.nodeDetailData?.nodeInfo?.osImage }}
          </el-descriptions-item>

          <el-descriptions-item>
            <template slot="label">
              内核版本
            </template>
            {{this.nodeDetailData?.nodeInfo?.kernelVersion}}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              Kubelet版本
            </template>
            {{this.nodeDetailData?.nodeInfo?.kubeletVersion}}
          </el-descriptions-item>

          <el-descriptions-item>
            <template slot="label">
              KubeProxy版本
            </template>
            {{this.nodeDetailData?.nodeInfo?.kubeProxyVersion}}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              操作系统
            </template>
            {{this.nodeDetailData?.nodeInfo?.operatingSystem}}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              架构
            </template>
            {{this.nodeDetailData?.nodeInfo?.architecture}}
          </el-descriptions-item>

          <!--状态-->
          <el-descriptions-item>
            <template slot="label">
              状态
            </template>
            <div>
              <el-table :data="nodeDetailData?.conditions" style="width: 100%" height="250">
                <el-table-column
                    fixed
                    prop="type"
                    label="类型">
                </el-table-column>
                <el-table-column
                    prop="status"
                    label="状态">
                </el-table-column>
                <el-table-column
                    prop="lastTransitionTime"
                    label="最近更改">
                </el-table-column>
                <el-table-column
                    prop="reason"
                    label="内容">
                </el-table-column>
                <el-table-column
                    prop="message"
                    label="信息">
                </el-table-column>
              </el-table>
            </div>
          </el-descriptions-item>
        </el-descriptions>

        <!--节点上Pod-->
        <div>
          <el-table :data="nodeDetailData?.podList?.items" style="width: 100%" height="250">
            <el-table-column
                fixed
                prop="metadata.name"
                label="名称">
            </el-table-column>
            <el-table-column
                prop="status.phase"
                label="状态">
              <template v-slot="scope">
                <el-button type="success" round size="mini" v-if="scope.row.status.phase==='Running'">{{scope.row.status.phase}}</el-button>
                <el-button type="danger" round size="mini" v-else>{{scope.row.status.phase}}</el-button>
              </template>
            </el-table-column>
            <el-table-column
                prop="status.containerStatuses[0].restartCount"
                label="重启次数">
            </el-table-column>
            <el-table-column
                prop="status.podIP"
                label="容器IP">
            </el-table-column>
            <el-table-column
                prop="status.hostIP"
                label="节点">
            </el-table-column>
            <el-table-column
                prop="status.startTime"
                label="创建时间">
            </el-table-column>
          </el-table>
        </div>
        <el-descriptions-item>
          <template slot="label">
            容器
          </template>

        </el-descriptions-item>
      </el-dialog>
    </div>

  </div>
</template>

<script>
import {GetNodeDetail, ListNodes, SetCordon} from "@/api/k8s";
import {Message} from "element-ui";
import {getClusterList} from "@/api/kubernetes/clusters";

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Nodes",
  inject: ['reload'],
  created() {
    this.getClusterList()
  },
  methods: {
    async getNodes(clusterId, params) {
      await ListNodes(clusterId, params).then((res) => {
        if (res.code === 0) {
          this.nodeData = res.data.nodes
        } else {
          Message.success(res.msg)
        }
      })
    },
    async getClusterList() {
      const res = await getClusterList()
      if (res.code === 0) {
        this.cluster_list = res.data
        await this.getNodes(res.data[0]?.clusterName)
      }
    },
    // 切换集群
    async clusterChange(value) {
      this.clusterName = value
      localStorage.setItem("clusterName", value)
      await this.getNodes(value)
    },
    // 获取节点详情
    getNodeDetail(nodeName) {
      this.NodeVisible = true
      GetNodeDetail(localStorage.getItem("clusterName"), nodeName).then((res) => {
        if (res.code === 0) {
          this.nodeDetailData = res.data
        }
      })
    },
    removeNode(nodeName) {
      console.log("移除的节点", nodeName)
    },
    async Search(value) {
      await this.getNodes(this.clusterName, {
        keyword: value,
      })
    },
    // 表格多选
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    // 设置节点不可调度
    setUnschedule() {
      console.log(this.multipleSelection)
    },
    // 节点排水
    setCordon() {
      let nodeName = []
      for(let i=0; i<this.multipleSelection.length; i++) {
        nodeName.push(this.multipleSelection[i].name)
      }
      SetCordon(localStorage.getItem("clusterName"), {"name":nodeName}).then((res) => {
        if (res.code === 0) {
          Message.success("节点排水成功")
        }
      })

    },
  },
  data() {
    return {
      nodeData: [],
      searchInfo: {
        'name': ''
      },
      cluster_list: [],
      clusterName: undefined,
      NodeVisible: false,
      nodeDetailData: undefined,
      multipleSelection: [],
    }
  }
}
</script>

<style scoped>
.search-term {
  padding-top: 10px;
}
</style>