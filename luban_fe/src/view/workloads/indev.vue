<template>
  <div>
    <div class="container">
      <el-tabs style="width:100%;">
        <el-tab-pane label="Deployments">
          <!--子组件-->
          <Deployments />
        </el-tab-pane>
<!--        <el-tab-pane label="DaemonSets">-->
<!--          &lt;!&ndash;子组件&ndash;&gt;-->
<!--          <Daemonsets />-->
<!--        </el-tab-pane>-->
<!--        <el-tab-pane label="StatefulSets">-->
<!--          &lt;!&ndash;子组件&ndash;&gt;-->
<!--          <StatefulSets />-->
<!--        </el-tab-pane>-->
<!--        <el-tab-pane label="Jobs">-->
<!--          &lt;!&ndash;子组件&ndash;&gt;-->
<!--          <Jobs />-->
<!--        </el-tab-pane>-->
<!--        <el-tab-pane label="CronJobs">-->
<!--          &lt;!&ndash;子组件&ndash;&gt;-->
<!--          <CronJobs />-->
<!--        </el-tab-pane>-->
<!--        <el-tab-pane label="Pods">-->
<!--          &lt;!&ndash;子组件&ndash;&gt;-->
<!--          <Pods />-->
<!--        </el-tab-pane>-->
      </el-tabs>
    </div>
    <router-view />
  </div>
</template>
<script>
// import Pods from './pods/index.vue'
import Deployments from './deployments/index.vue'
// import Daemonsets from './daemonsets/index.vue'
// import StatefulSets from './statefulsets/index.vue'
// import Jobs from './jobs/index.vue'
// import CronJobs from './cronjobs/index.vue'
export default {
  name: 'Workloads',
  components: {
    // Pods,
    Deployments,
    // Daemonsets,
    // Jobs,
    // CronJobs,
    // StatefulSets
  }
}
</script>
