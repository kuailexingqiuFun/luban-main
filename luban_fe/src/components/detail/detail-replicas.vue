<template>
  <div>
    <h3 style="display: inline-block;">副本数</h3>
    <el-button v-if="!enableEdit" style="margin-left: 10px;" type="text" icon="el-icon-edit" @click="enableEdit = true">编辑信息</el-button>

    <div v-loading="loading">
      <el-form label-position="top" ref="form" :model="form">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="副本数" prop="spec.replicas">
              <span style="margin-left: 10px;" v-if="!enableEdit">{{form.spec.replicas}}  个副本</span>
              <el-input v-else @input="onInput()" v-model="form.spec.replicas" />
            </el-form-item>
          </el-col>
        </el-row>
        <div style="float:right" v-if="enableEdit">
          <el-button @click="updateForm">提交</el-button>
          <el-button @click="enableEdit = false">取消</el-button>
        </div>
        <br>
      </el-form>
    </div>
  </div>
</template>

<script>
  import { updateWorkLoad } from "@/api/workloads"
  import Rule from "@/utils/rules"

  export default {
    name: "DetailReplicas",
    props: {
      cluster_id: Number,
      formInfo: Object,
      resourceType: String,
    },
    watch: {
      formInfo: {
        handler(newformInfo) {
          if (newformInfo) {
            if (newformInfo.spec) {
              this.form = newformInfo
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
        form: {},
        enableEdit: false,
        numberPersentRule: Rule.NumberPersentRule,
      }
    },
    methods: {
      initForm() {
        if (this.formInfo.spec) {
          this.form = this.formInfo
          this.loading = false
        }
      },
      onInput() {
        this.$forceUpdate();
      },
      updateForm() {
        this.loading = true
        this.form.spec.replicas = parseInt(this.form.spec.replicas)
        updateWorkLoad(this.cluster_id, this.resourceType, this.form.metadata.namespace, this.form.metadata.name, this.form)
                .then(() => {
                  this.$message({
                    type: "success",
                    message: "修改成功",
                  })
                  this.enableEdit = false
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
