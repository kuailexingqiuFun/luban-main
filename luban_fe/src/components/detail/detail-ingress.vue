<template>
  <div>
    <h3>路由</h3>
    <el-alert v-if="ingress.length === 0" type="info">
      <div slot="title">
        <i class="el-icon-info"></i>
        <span> 此处显示与 {0} 同名的 Ingress: {1}，但是您并未定义此 Ingress。</span>
      </div>
    </el-alert>
    <div v-for="(item, index) in ingress" :key="index">
      <span>域名: {{ item.host }}</span>
      <el-table v-if="ingress.length !== 0" :data="item.details">
        <el-table-column
                label="名称"
                type="scope">
          <template slot-scope="scope" >
            <span v-if="scope.row.backend.serviceName">{{scope.row.backend.serviceName }}</span>
            <span v-if="scope.row.backend.service">{{scope.row.backend.service.name }}</span>
          </template>
        </el-table-column>
        <el-table-column label="路径" prop="path" />
        <el-table-column
                label="端口"
                prop="port"
                type="scope">
          <template slot-scope="scope" >
            <span v-if="scope.row.backend.servicePort">{{scope.row.backend.servicePort }}</span>
            <span v-if="scope.row.backend.service">{{scope.row.backend.service.port.number }}</span>
          </template>
        </el-table-column>
        <el-table-column label="类型" prop="pathType" />
      </el-table>
      <br><br>
    </div>
  </div>
</template>

<script>
import {IngressesList} from "@/api/kubernetes/ingress";

export default {
  name: "DetailIngress",
  components: {  },
  props: {
    cluster_id: Number,
    namespace: String,
    name: String,
    resourceType: String,
  },
  data() {
    return {
      loading: false,
      ingress: [],
    }
  },
  methods: {
    search() {
      this.ingress = []
      this.loading= true
      IngressesList(this.cluster_id, "", "", this.namespace, this.name,  "", "").then((res) => {
        if (res.data.items === null) {
          return
        }
        for (const ing of res.data.items) {
          if ((ing.metadata.name === this.name)) {
            if (ing.spec.rules) {
              for (const i of ing.spec.rules) {
                var item = {
                  host: i.host,
                }
                if (i.http) {
                  item.details = i.http.paths
                }
                this.ingress.push(item)
              }
            }
          }
        }
      })
    },
  },
  created() {
    this.search()
  },
}
</script>

<style scoped>
</style>
