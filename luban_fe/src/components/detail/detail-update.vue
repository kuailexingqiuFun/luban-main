<template>
  <div>
    <h3 style="display: inline-block;">部署策略</h3>
    <el-button v-if="!enableEdit" style="margin-left: 10px;" type="text" icon="el-icon-edit" @click="enableEdit = true">编辑信息</el-button>

    <div v-loading="loading">
      <el-form label-position="top" ref="form" :model="form">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="策略" prop="spec.strategy.type" required>
              <span style="margin-left: 10px;" v-if="!enableEdit">{{getLabels(form.spec.strategy.type)}}</span>
              <el-select v-else style="width:100%" @change="typeChange"  v-model="form.spec.strategy.type">
                <el-option v-for="(item, index) in deployment_strategy_list" :key="index" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20" v-if="form.spec.strategy.type === 'RollingUpdate'">
          <el-col :span="12">
            <el-form-item label="最大Pod数量" prop="spec.strategy.rollingUpdate.maxSurge">
              <span style="margin-left: 10px;" v-if="!enableEdit">{{form.spec.strategy.rollingUpdate.maxSurge}}</span>
              <el-input v-else @input="onInput()" v-model="form.spec.strategy.rollingUpdate.maxSurge" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="最大不可用数量" prop="spec.strategy.rollingUpdate.maxUnavailable">
              <span style="margin-left: 10px;" v-if="!enableEdit">{{form.spec.strategy.rollingUpdate.maxUnavailable}}</span>
              <el-input v-else @input="onInput()" v-model="form.spec.strategy.rollingUpdate.maxUnavailable" />
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
    name: "DetailUpdate",
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
        deployment_strategy_list: [
          { label: "滚动升级", value: "RollingUpdate" },
          { label: "重新创建", value: "Recreate" },
        ],
        strategy_list: [
          { label: "滚动升级", value: "RollingUpdate" },
          { label: "删除", value: "OnDelete" },
        ],
        form: {
          spec: {
            strategy: {
              maxSurge: 0,
              maxUnavailable: 0,
            },
            type: "",
          },
        },
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
      typeChange(value){
        switch (value) {
          case "RollingUpdate":
            this.form.spec.strategy.type = "RollingUpdate"
            this.form.spec.strategy.rollingUpdate =  {
              maxSurge: 0,
              maxUnavailable: 0,
            }
            return
          case "Recreate":
            this.form.spec.strategy.type = "Recreate"
            delete this.form.spec.strategy.rollingUpdate
            return
        }
      },
      updateForm() {
        this.loading = true
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
      getLabels(value) {
        switch (value) {
          case "RollingUpdate":
            return "滚动升级"
          case "Recreate":
            return "重新创建"
          case "OnDelete":
            return "删除"
        }
      },
    },
    created() {
      this.initForm()
    },
  }
</script>

<style scoped>
</style>
