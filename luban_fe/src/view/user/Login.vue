<template>
  <div class="login">
      <h3 style="z-index:1;text-align:center;"><img src="http://demo.luban.dnsjia.com/img/luban.ba1cc9b6.svg" width="120" height="80"></h3>
      <el-tabs>
        <el-tab-pane label="用户登陆">
          <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" size="small">
            <el-form-item prop="username">
              <el-input type="text" v-model="ruleForm.username" placeholder="帐号" autocomplete="off"></el-input>
            </el-form-item>

            <el-form-item prop="password">
              <el-input type="password" v-model="ruleForm.password" placeholder="密码" autocomplete="off"></el-input>
            </el-form-item>

            <el-form-item prop="remember">
              <el-checkbox v-model="ruleForm.remember" label="记住密码" name="type">
              </el-checkbox>
            </el-form-item>

            <el-form-item>
              <el-button type="primary" style="min-width: 370px" @click="submitForm('ruleForm')">登陆</el-button>
            </el-form-item>

          </el-form>
        </el-tab-pane>
      </el-tabs>
    </div>
</template>
<script>
import {UserLogin} from "@/api/user";
import {Message} from "element-ui";

export default {
  name: 'Login',
  data() {
    return {
      ruleForm: {
        username: sessionStorage.getItem("username"),
        password: sessionStorage.getItem("password"),
        remember: false
      },
      rules: {
        username: [
          {  required: true, message: '请输入用户',  trigger: 'change', }
        ],
        password: [
          {  required: true, message: '请输入密码',  trigger: 'change', }
        ],
      }
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (this.ruleForm.remember) {
            sessionStorage.setItem("username", this.ruleForm.username)
            sessionStorage.setItem("password", this.ruleForm.password)
          } else {
            sessionStorage.removeItem("username")
            sessionStorage.removeItem("password")
          }
          UserLogin(this.ruleForm).then((res) => {
            if (res.code === 0) {
              Message.success("登陆成功")
              localStorage.setItem("onLine", true)
             this.$router.push("/")
            } else {
              Message.error(res.msg)
            }

          })
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
  }
}
</script>

<style scoped>
.login {
  justify-content: center;
  width: 370px;
  margin: 180px auto;
  padding: 10px 20px 10px 20px;
  border-radius: 20px;
  box-shadow: 0 0 20px #DCDFE6;
}
</style>