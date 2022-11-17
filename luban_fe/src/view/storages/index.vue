<template>
  <div>
    <el-card>
      <div class="container">
        <el-tabs style="width:100%;">
          <el-tab-pane label="Persistent Volume Claims">
            <!--子组件-->
            <persistentVolumeClaim/>
          </el-tab-pane>
          <el-tab-pane label="Persistent Volumes" >
            <!--子组件-->
            <persistentvolumes/>
          </el-tab-pane>
          <el-tab-pane label="Storage Class" >
            <!--子组件-->
            <storageclasses/>
          </el-tab-pane>
        </el-tabs>
      </div>
    </el-card>
  </div>
</template>

<script>
import  persistentVolumeClaim from './persistentVolumeClaim/index'
import  persistentvolumes from './persistentvolumes/index'
import  storageclasses from './storageclasses/index'
export default {
  name:"Storage",
  components: {
    persistentVolumeClaim,
    persistentvolumes,
    storageclasses
  },
}
</script>
<style lang="scss">

</style>
