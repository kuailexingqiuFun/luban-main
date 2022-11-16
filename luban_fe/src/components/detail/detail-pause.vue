<template>
  <div>
    <h3>Deployment Processing</h3>
    <div v-loading="loading">
      <el-form label-position="top" ref="form" :model="form">
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="状态" prop="spec.paused">
              <el-radio-group v-model="form.spec.paused" @change="changeStatus">
                <el-radio :label="true">暂停</el-radio>
                <el-radio :label="false">激活</el-radio>
              </el-radio-group>
              <div><span style="font-size: 12px;">暂停状态下修改 Deployment，将不会执行处理动作</span></div>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="最小就绪时间(秒)" prop="spec.minReadySeconds">
              <span>{{form.spec.minReadySeconds || 0}}</span>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="进程截止时间(秒)" prop="spec.progressDeadlineSeconds">
              <span>{{form.spec.progressDeadlineSeconds || 0}}</span>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>
  </div>
</template>

<script>
import { updateWorkLoad } from "@/api/workloads"

export default {
  name: "DetailPause",
  props: {
    information_id: Number,
    formInfo: Object,
    resourceType: String,
  },
  watch: {
    formInfo: {
      handler(newformInfo) {
        if (newformInfo) {
          if (newformInfo.spec) {
            this.form = newformInfo
            if (newformInfo.spec.paused === undefined) {
              this.form.spec.paused = false
            }
            this.loading = false
          }
        }
      },
      immediate: true,
      deep: true,
    },
  },
  data() {
    return {
      loading: true,
      statu_list: [
        { label: "暂停", value: true },
        { label: "激活", value: false },
      ],
      form: {
        spec: {
          paused: false,
          minReadySeconds: 0,
          progressDeadlineSeconds: 600,
        },
      },
    }
  },
  methods: {
    initForm() {
      if (this.formInfo.spec) {
        this.form = this.formInfo
        this.loading = false
      }
    },
    changeStatus() {
      if (this.form.spec.paused) {
        this.$confirm("暂停" + "  Deployments ?", "提示", {
          confirmButtonText: "确认",
          cancelButtonText: "取消",
          type: "warning",
        })
          .then(() => {
            this.updateForm()
          })
          .catch(() => {
            this.form.spec.paused = false
          })
      } else {
        this.updateForm()
      }
    },
    updateForm() {
      this.loading = true
      updateWorkLoad(this.information_id, this.resourceType, this.form.metadata.namespace, this.form.metadata.name, this.form)
        .then(() => {
          this.$message({
            type: "success",
            message: "更新成功",
          })
          this.$emit("refresh")
        })
        .finally(() => {
          this.loading = false
        })
    },
  },
  created() {
    this.initForm()
  },
}
</script>

<style scoped>
</style>
