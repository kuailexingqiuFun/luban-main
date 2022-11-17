<template>
  <div>
    <el-card>
      <div style="padding-top: 20px">
        <el-button type="primary" @click="addUserModal = true" size="small">添加用户</el-button>
      </div>
      <el-table
          :data="userData"
          style="width: 100%">
        <el-table-column
            prop="id"
            label="ID">
        </el-table-column>
        <el-table-column
            prop="username"
            label="用户名">
        </el-table-column>
        <el-table-column
            prop="email"
            label="邮箱">
        </el-table-column>
        <el-table-column
            prop="status"
            label="状态">
          <template v-slot="scope">
            <el-tag v-if="scope.row.status">激活</el-tag>
            <el-tag v-if="!scope.row.status" type="danger">禁用</el-tag>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <div>
      <el-dialog title="添加用户" :visible.sync="addUserModal" width="600px" :modal="false" :destroy-on-close="true">
        <el-form :model="addUserForm" label-width="100px" :rules="addUserRules" ref="addUserForm">
          <el-form-item label="用户名" prop="username">
            <el-input v-model="addUserForm.username" autocomplete="off" placeholder="请输入用户名"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input placeholder="请输入密码" v-model="addUserForm.password" type="password"></el-input>
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="addUserForm.email" autocomplete="off" placeholder="请输入邮箱"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="addUserModal = false" size="small">取 消</el-button>
          <el-button type="primary" @click="AddUser('addUserForm')" size="small">确 定</el-button>
        </div>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import {UserManage, UserRegister} from "@/api/user";
import {Message} from "element-ui";


export default {
  name: "UserManage",
  inject: ['reload'],
  data() {
    return {
      userData: [],
      addUserModal: false,
      addUserForm: {
        username: '',
        password: '',
        email: '',
      },
      addUserRules: {
        username: [
          {required: true, message: '请输入用户名', trigger: 'blur'}
        ],
        password: [
          {required: true, message: '请输入密码', trigger: 'blur'}
        ]
      }
    }
  },
  created() {
    this.GetUser()
  },
  methods: {
    GetUser() {
      UserManage().then((res) => {
        if (res.code ===0){
          this.userData = res.data
        } else {
          Message.error(res.msg)
        }
      })
    },
    // 添加用户
    AddUser(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          console.log(valid)
          UserRegister(this.addUserForm).then((res) => {
            if (res.code ===0) {
              Message.success("添加成功")
              this.addUserModal = false
              this.reload()
            } else {
              Message.error(res.msg)
            }
          })
        } else {
          return false
        }
      })
    },
  }
}
</script>

<style scoped>

</style>