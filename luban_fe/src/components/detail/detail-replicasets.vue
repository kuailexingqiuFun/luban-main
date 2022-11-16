<template>
  <div>
    <el-table :data="pods" v-loading="loading" @search="search">
      <el-table-column label="版本" prop="name" min-width="80" show-overflow-tooltip>
        <template v-slot:default="{row}">
          <span>{{ row.metadata.annotations["deployment.kubernetes.io/revision"] }}</span>
        </template>
      </el-table-column>
      <el-table-column label="镜像" prop="image" min-width="80" show-overflow-tooltip>
        <template v-slot:default="{row}">
          <span v-for="(k, index) in row.spec.template.spec.containers" :key="index">
            <span class="label-custom wd" type="info">{{ k.image }}</span>
          </span>
        </template>
      </el-table-column>
      <el-table-column label="时间" prop="time" min-width="80" show-overflow-tooltip>
        <template v-slot:default="{row}">
          <span>{{ row.metadata.creationTimestamp }}</span>
        </template>
      </el-table-column>
      <table-operations :buttons="buttons" label="操作"></table-operations>
    </el-table>

    <el-dialog title="详细信息" width="70%" :close-on-click-modal="false" :visible.sync="dialogSpecificVisible" append-to-body>
      <yaml-editor :form="itemVerisonForm" :read-only="true" />
      <div slot="footer" class="dialog-footer">
        <el-button size="small" @click="dialogSpecificVisible = false">取消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import TableOperations from "@/components/table-operations"
// import { listNsReplicaSetsWorkload } from "@/api/k8s/replicasets"
// import { patchDeployment } from "@/api/k8s/deployment"
import YamlEditor from "@/components/yaml/YamlBlock.vue"

export default {
  name: "DetailReplicasets",
  components: { TableOperations, YamlEditor },
  props: {
    cluster_id: Number,
    namespace: String,
    selector: String,
    fieldSelector: String,
    name: String,
  },
  watch: {
    selector: {
      handler(newSelector) {
        if (newSelector) {
          this.search()
        }
      },
      immediate: true,
    },
    fieldSelector: {
      handler(newSelector) {
        if (newSelector) {
          this.search()
        }
      },
      immediate: true,
    },
  },
  data() {
    return {
      buttons: [
        {
          label: "回滚到该版本",
          icon: "el-icon-refresh-left",
          click: (row) => {
            this.OptionRollback(row)
          },
          disabled: (row) => {
            return !(row.status.readyReplicas === undefined)
          }
        },
        {
          label: "具体信息",
          icon: "el-icon-tickets",
          click: (row) => {
            this.SpecificInformation(row)
          },
        },
      ],
      loading: false,
      pods: [],
      podUsage: [],
      itemVerisonForm: {},
      dialogSpecificVisible: false,
    }
  },
  methods: {
    search() {
      this.loading = true
      // listNsReplicaSetsWorkload(this.cluster_id, this.namespace, this.selector, this.fieldSelector).then((res) => {
      //   this.loading = false
      //   res.data.items.sort((a, b) => b.metadata.annotations["deployment.kubernetes.io/revision"] - a.metadata.annotations["deployment.kubernetes.io/revision"])
      //   for (var i = 0; i < res.data.items.length; i++) {
      //     this.pods.push(res.data.items[i])
      //   }
      // })
    },
    OptionRollback(row) {
      this.$confirm("回滚到该版本", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(() => {
        // patchDeployment(this.cluster_id, row.metadata.namespace, this.name, { spec: { template: row.spec.template } })
        //   .then(() => {
        //     this.dialogModifyVersionVisible = false
        //     this.loading = true
        //     this.$message({
        //       type: "success",
        //       message: "操作成功",
        //     })
        //   })
        //   .finally(() => {
        //     this.loading = false
        //   })
      })
    },
    SpecificInformation(row) {
      this.itemVerisonForm = row
      this.dialogSpecificVisible = true
    },
  },
}
</script>

<style scoped>
.btnSize {
  width: 28px;
  height: 28px;
}
</style>
