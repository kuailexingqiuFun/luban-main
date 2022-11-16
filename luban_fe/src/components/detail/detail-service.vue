<template>
  <div>
    <h3>服务</h3>
      <el-alert v-if="services.length === 0" type="info">
        <div slot="title">
          <i class="el-icon-info"></i>
          <span>此处显示与 {0} 同名的 Service: {1}，但是您并未定义此 Service。</span>
        </div>
      </el-alert>
    <el-table v-if="services.length !== 0" :data="services">
      <el-table-column label="名称" prop="name" />
      <el-table-column label="协议" prop="protocol" />
      <el-table-column label="端口" prop="port" />
      <el-table-column label="NodePort" prop="nodePort" />
      <el-table-column label="容器端口" prop="targetPort" />
    </el-table>
  </div>
</template>

<script>
// import { getServiceList } from "@/api/k8s/service"

export default {
  name: "DetailSvcs",
  props: {
    information_id: Number,
    namespace: String,
    name: String,
    resourceType: String,
  },
  data() {
    return {
      loading: false,
      services: [],
    }
  },
  methods: {
    async search() {
      this.services = []
      // const res = await getServiceList("", "", this.information_id, this.name, this.namespace, "")
      // if (res.data.items === null) {
      //   return
      // }
      //   for (const svc of res.data.items) {
      //     if (svc.metadata.name === this.name) {
      //       if (svc.spec.ports) {
      //         for (const p of svc.spec.ports) {
      //           this.services.push(p)
      //         }
      //       }
      //     }
      //   }
      //   this.loading = false
    },
  },
  created() {
    this.search()
  },
}
</script>

<style scoped>
</style>
