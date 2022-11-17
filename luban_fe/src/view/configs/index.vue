<template>
  <div>
    <el-card>
      <div class="container">
        <el-tabs style="width:100%;">
          <el-tab-pane label="配置项">
            <!--子组件-->
            <ConfigMap/>
          </el-tab-pane>
          <el-tab-pane label="保密字典" >
            <!--子组件-->
            <Secret/>
          </el-tab-pane>
        </el-tabs>
      </div>
    </el-card>
  </div>
</template>

<script>
import  ConfigMap from './configmap/index'
import  Secret    from './secret/index'
export default {
  name:"configuration",
  components: {
    ConfigMap,
    Secret,
  },
}
</script>
<style lang="scss">

</style>
