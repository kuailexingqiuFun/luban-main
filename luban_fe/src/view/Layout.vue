<template>
  <div>
    <el-container class="header-container">
      <!--头部-->
      <el-header style="display: flex; background-color: #fff;align-items: stretch; justify-content: space-between;height: 50px;">
        <div style="display: flex; align-items: center">
          <span style="font-size: large; margin-left: 30px">多集群管理平台</span>
        </div>
      </el-header>
      <el-container>
        <el-aside style="background-color: #f5f5f5">
          <el-col style="height: 100%">
            <el-menu
                active-text-color="#ffd04b"
                background-color="#fff"
                border-right="none"
                default-active="1"
                text-color="#555"
                router
                class="el-menu-vertical-no-collapse">
              <el-submenu index="1">
                <template slot="title">
                  <i class="el-icon-location"></i>
                  <span>容器管理</span>
                </template>
                <el-menu-item-group>
                  <template slot="title">kubernetes</template>
                  <el-menu-item index="/kubernetes/workloads">工作负载</el-menu-item>
                  <el-menu-item index="1-2">选项2</el-menu-item>
                </el-menu-item-group>
                <el-menu-item-group title="分组2">
                  <el-menu-item index="1-3">选项3</el-menu-item>
                </el-menu-item-group>
                <el-submenu index="1-4">
                  <template slot="title">选项4</template>
                  <el-menu-item index="1-4-1">选项1</el-menu-item>
                </el-submenu>
              </el-submenu>
              <el-menu-item index="2">
                <i class="el-icon-menu"></i>
                <span slot="title">导航二</span>
              </el-menu-item>
              <el-menu-item index="3" disabled>
                <i class="el-icon-document"></i>
                <span slot="title">导航三</span>
              </el-menu-item>
              <el-menu-item index="4">
                <i class="el-icon-setting"></i>
                <span slot="title">导航四</span>
              </el-menu-item>
            </el-menu>
          </el-col>
        </el-aside>
        <el-main>
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "Layout"
}
</script>

<style scoped>
.header-container {
  height: 100%;
  width: 100%;
  font-weight: 400;
  font-style: normal;
   position: fixed;
  z-index: 999;
  top: 0;
  left: 0;
}
.el-menu-vertical-no-collapse:not(.el-menu--collapse) {
  width: 200px;
  height: 100%;
}

.el-header {
  background: rgb(55, 61, 61);
  display: flex;
  justify-content: space-between;
  line-height: 60px;
  color: #555;
  font-size: 25px;
}
.el-aside {
  height: 100%;
  width: auto !important;
}
.el-menu {
  border-right: none;
  height: 100%;
}

.el-main {
  background: #e4ebf1;
}

.el-menu-class {
  height: 51px;
}
.header-input {
  margin-right: 30px;
  width: 200px;
  transition: width 0.5s;
}
.header-input:hover {
  width: 350px;
}
.el-menu-class {
  height: 51px;
}
.header-bottom {
  font-size: small;
  margin-right: 25px;
  color: #adb0bb;
}



</style>
