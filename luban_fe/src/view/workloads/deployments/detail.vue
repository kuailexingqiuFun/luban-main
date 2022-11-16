<template>
    <div>
      <el-row :gutter="20" class="row-box">
        <el-col :span="24">
          <el-card class="el-card">
            <detail-basic :item="form" />
          </el-card>
        </el-col>
          <el-col :span="24">
              <el-tabs v-model="activeName" tab-position="top" type="border-card"
                       @tab-click="handleClick" ref=tabs>
                  <el-tab-pane name="pods" label="容器组">
                      <detail-pods :cluster_id="cluster_id" :namespace="form.metadata.namespace" :name="form.metadata.name" :selector="selectors" />
                  </el-tab-pane>
                  <el-tab-pane name="service" label="服务">
                      <detail-service :cluster_id="cluster_id" :namespace="form.metadata.namespace" :name="form.metadata.name" resourceType="Deployment" />
                  </el-tab-pane>
                  <el-tab-pane name="ingress" label="路由">
                      <detail-ingress :cluster_id="cluster_id" :namespace="form.metadata.namespace" :name="form.metadata.name" resourceType="Deployment" />
                  </el-tab-pane>
                  <el-tab-pane name="conditions" label="运行时信息">
                      <detail-conditions :conditions="form.status.conditions" />
                  </el-tab-pane>
                  <el-tab-pane name="strategy" label="部署策略">
                      <detail-update :cluster_id="cluster_id"  :formInfo="form" resourceType="deployments" />
                  </el-tab-pane>
                  <el-tab-pane name="replicas" label="副本数">
                      <detail-replicas :cluster_id="cluster_id"  :formInfo="form" resourceType="deployments" />
                  </el-tab-pane>
                  <el-tab-pane name="image" label="容器镜像">
                      <detail-image  :cluster_id="cluster_id"  :formInfo="form" resourceType="deployments" />
                  </el-tab-pane>
                  <el-tab-pane name="pause" label="容器暂停">
                      <detail-pause :cluster_id="cluster_id"  :formInfo="form" resourceType="deployments" />
                  </el-tab-pane>
                  <el-tab-pane name="history" label="历史版本">
                      <detail-replicasets :cluster_id="cluster_id" :namespace="form.metadata.namespace" :name="form.metadata.name" :selector="selectors" />
                  </el-tab-pane>
              </el-tabs>
          </el-col>
      </el-row>
    </div>
</template>

<script>
  import DetailConditions from "@/components/detail/detail-conditions"
  import DetailReplicas from "@/components/detail/detail-replicas"
  import DetailBasic from "@/components/detail/detail-basic"
  import DetailPods from "@/components/detail/detail-pods"
  import DetailReplicasets from "@/components/detail/detail-replicasets"
  import DetailService from "@/components/detail/detail-service"
  import DetailIngress from "@/components/detail/detail-ingress"
  import DetailPause from "@/components/detail/detail-pause"
  import DetailUpdate from "@/components/detail/detail-update"
  import DetailImage from "@/components/detail/detail-image"
  import { formatTimeToStr } from '@/utils/date'
    export default {
        name: 'DetailBlock',
      components: { DetailBasic, DetailConditions,DetailReplicas, DetailService,DetailPause, DetailIngress, DetailUpdate,DetailPods, DetailImage, DetailReplicasets},
      props: {
        cluster_id: Number,
        form: {
                type: Object,
                default: function() {
                    return {}
                }
            }
        },
        data() {
            return {
                name: '',
                activeName: "pods",
                ContainerTitle: '',
                dialogVisibleContainerlog: false,
                currentValuedetailurl: '',
                terminal: {
                  pid: 1,
                  name: 'terminal',
                  cols: 150,
                  rows: 23
                },
                selectors: "",
            }
        },
        filters: {
            formatDate: function(time) {
                if (time != null && time !== '') {
                    var date = new Date(time)
                    return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
                } else {
                    return ''
                }
            }
        },
        created() {
            this.getDetail()
        },
        methods: {
            getDetail() {
                    this.loading = true
                    if (this.form.spec.selector.matchLabels) {
                     this.selectors = ""
                         for (const key in this.form.spec.selector.matchLabels) {
                                if (Object.prototype.hasOwnProperty.call(this.form.spec.selector.matchLabels, key)) {
                                    this.selectors += key + "=" + this.form.spec.selector.matchLabels[key] + ","
                                }
                            }
                            this.selectors = this.selectors.length !== 0 ? this.selectors.substring(0, this.selectors.length - 1) : ""
                        }
                        this.loading = false
                },
            handleClick (tab) {
                this.activeName = tab.name
            },
        }
    }
</script>
<style scoped lang="scss">

</style>
