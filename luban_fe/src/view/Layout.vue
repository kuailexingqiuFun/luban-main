<template>
  <div>
    <el-container class="header-container">
      <!--头部-->
      <el-header style="display: flex; background-color: #fff;align-items: stretch; justify-content: space-between;height: 50px;">
        <div style="display: flex; align-items: center">
          <span style="font-size: large; margin-left: 30px">多集群管理平台</span>
        </div>

        <div style="display: flex; align-items: center">
          <div style="vertical-align: middle;margin-top: 30px;margin-right: 28px;cursor: pointer">
            <el-dropdown>
              <span>
                <el-avatar :size="30" :src="avatarImg"></el-avatar>
              </span>

              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item @click.native="logout()">注销登陆</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </div>
      </el-header>
      <el-container>
        <el-aside style="background-color: #f5f5f5;height: calc(100vh - 70px);">
          <el-col style="height: 100%">
            <el-menu
                active-text-color="#ffd04b"
                background-color="#fff"
                border-right="none"
                :default-active="this.$route.path"
                :default-openeds="['1']"
                text-color="#555"
                router
                class="el-menu-vertical-no-collapse">
              <el-submenu index="1">
                <template slot="title">
                  <i class="el-icon-s-platform"></i>
                  <span>容器管理</span>
                </template>
                <el-menu-item-group>
                  <el-menu-item index="/kubernetes/cluster">集群管理</el-menu-item>
                  <el-menu-item index="/kubernetes/node">节点管理</el-menu-item>
                  <el-menu-item index="/kubernetes/workloads">工作负载</el-menu-item>
                  <el-menu-item index="/kubernetes/network">服务发现</el-menu-item>
                  <el-menu-item index="/kubernetes/configs">配置</el-menu-item>
                  <el-menu-item index="/kubernetes/storages">存储</el-menu-item>
                  <el-menu-item index="/kubernetes/accessControls">访问控制</el-menu-item>
                </el-menu-item-group>
              </el-submenu>

              <el-submenu index="2">
                <template slot="title">
                  <i class="el-icon-user-solid"></i>
                  <span>用户管理</span>
                </template>
                <el-menu-item-group>
                  <el-menu-item index="/user/userList">用户列表</el-menu-item>
                </el-menu-item-group>
              </el-submenu>
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
  name: "Layout",
  data() {
    return {
      avatarImg: 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'
    }
  },
  methods: {
    logout() {
      localStorage.removeItem("onLine")
      this.$router.push("/user/login")
    },
  }
}
</script>

<style scoped>
.header-container {
  height: 100%;
  width: 100%;
  font-weight: 400;
  font-style: normal;
  /*position: fixed;*/
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
  height: calc(100vh - 70px);
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
