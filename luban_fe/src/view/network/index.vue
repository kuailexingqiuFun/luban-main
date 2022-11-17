<template>
  <div>
    <el-card>
      <div class="container">
        <el-tabs style="width:100%;">
          <el-tab-pane label="路由">
            <!--子组件-->
            <Ingress/>
          </el-tab-pane>
          <el-tab-pane label="服务" >
            <!--子组件-->
            <Service/>
          </el-tab-pane>
        </el-tabs>
      </div>
    </el-card>
  </div>

</template>

<script>
import  Ingress from './ingress/index.vue'
import  Service from './service/index.vue'
export default {
  name:"Network",
  components: {
    Ingress,
    Service,
  },
}
</script>
<style lang="scss">

</style>
