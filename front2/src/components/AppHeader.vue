<template>
  <el-menu
    class="el-menu"
    mode="horizontal"
    :ellipsis="false"
    @select="handleSelect"
    router
  >
    <el-text style="font-size: xx-large;margin-right: 30px">CUGOJ</el-text>
    <el-menu-item class="menu-item" index="/" style="font-size: large"><el-icon><HomeFilled /></el-icon>主页</el-menu-item>
    <el-menu-item class="menu-item" index="/problem" style="font-size: large"><el-icon><Grid /></el-icon>问题</el-menu-item>
    <el-menu-item class="menu-item" index="/contests" style="font-size: large"><el-icon><TrophyBase /></el-icon>竞赛</el-menu-item>
    <el-menu-item class="menu-item" index="/status" style="font-size: large"><el-icon><TrendCharts /></el-icon>状态</el-menu-item>
    <el-menu-item class="menu-item" index="/rank" style="font-size: large"><el-icon><Histogram /></el-icon>排名</el-menu-item>
    <el-menu-item class="menu-item" index="/group" style="font-size: large"><el-icon><UserFilled /></el-icon>班级</el-menu-item>
    <el-menu-item class="menu-item" index="/help" style="font-size: large"><el-icon><MoreFilled /></el-icon>帮助</el-menu-item>
    <el-menu-item class="menu-item" index="/test" style="font-size: large">题目</el-menu-item>
    <el-menu-item class="menu-item" index="/admin" style="font-size: large">管理员</el-menu-item>
    <div class="flex-grow" />
<!--    如果本地存在token就不显示登录按钮-->
    <el-menu-item class="menu-item" v-show="isLogin()">
    <el-dropdown>
      <span class="el-dropdown-link">{{usrname}}<el-icon><arrow-down /></el-icon></span>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item><router-link to="/user/profile" class="router-link">个人中心</router-link></el-dropdown-item>
          <el-dropdown-item class="drop-item" @click="logout()">&nbsp;退出登录</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
    </el-menu-item>
    <el-button v-show="!isLogin()" @click="loginVis=true" round style="font-size: large;margin-top: auto;margin-bottom: auto">登录</el-button>
    <el-button v-show="!isLogin()" @click="regisVis=true" round style="font-size: large;margin-top: auto;margin-bottom: auto">注册</el-button>
    <el-dialog v-model="loginVis" :center="true" title="登录" width="300">
       <el-form :model="form" label-width="auto" style="max-width: 600px">
          <el-form-item label="用户名">
            <el-input v-model="form.username" placeholder="请输入用户名"></el-input>
          </el-form-item>
          <el-form-item label="密码">
            <el-input v-model="form.password" placeholder="请输入密码" show-password></el-input>
          </el-form-item>
          <el-form-item style="text-align: center;  justify-content: center;">
            <el-button @click="login" style="width: 300px" type="primary">登录</el-button>
          </el-form-item>
       </el-form>
    </el-dialog>
    <el-dialog v-model="regisVis" :center="true" title="注册" width="300">
       <el-form :model="form" label-width="auto" style="max-width: 600px">
          <el-form-item label="用户名">
            <el-input v-model="form.username" placeholder="请输入用户名"></el-input>
          </el-form-item>
          <el-form-item label="密码" style="text-align: center">
            <el-input v-model="form.password" placeholder="请输入密码" show-password></el-input>
          </el-form-item>
          <el-form-item label="邮箱">
            <el-input v-model="form.mail" placeholder="请输入邮箱" show-password></el-input>
          </el-form-item>
          <el-form-item label="确认密码">
            <el-input v-model="form.password" placeholder="请输入确认密码" show-password></el-input>
          </el-form-item>
         <el-form-item label="人机验证">
            <el-input v-model="form.captcha" placeholder="请输入验证码" show-password></el-input>
          </el-form-item>
          <el-form-item style="text-align: center;  justify-content: center;">
            <el-button style="width: 300px" type="primary">注册</el-button>
          </el-form-item>
       </el-form>
    </el-dialog>
  </el-menu>
</template>

<script lang="ts" setup>
import {ref} from 'vue'
import {
  ArrowDown,
  Grid,
  Histogram,
  HomeFilled,
  MoreFilled,
  TrendCharts,
  TrophyBase, UserFilled,
} from "@element-plus/icons-vue";
import axios from "axios";
let usrname=ref('USER')
let loginVis = ref(false)
let regisVis = ref(false)
let form = ref({
  username: '',
  password: '',
  mail: '',
  captcha: ''
})
function logout() {
  localStorage.removeItem('token')
  location.reload()
}
function isLogin() {
  // 判断token有没有过期，过期就删除
  return localStorage.getItem('token')
}
// axios登录函数
function login() {
  axios.post('http://127.0.0.1:5174/api/login', {
    username: form["username"],
    password: form["password"]
  }).then(res => {
    localStorage.setItem('token', res.data.token)
    location.reload()
  }).catch(err => {
    console.log(err)
  })
//   刷新页面

}
// axios注册函数

const activeIndex = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
}
</script>

<style>
.flex-grow {
  flex-grow: 1;
}
.el-dropdown-link  {
  cursor: pointer;
  color: var(--el-color-primary);
  display: flex;
  align-items: center;
}
.el-dropdown-link{
  font-size: larger;
}
.el-dropdown-link:focus {
  outline: none;
}
.router-link{
  color:#606266;
}
.router-link:hover {
  background-color:unset;
}
.drop-item:focus{
  color: unset !important;
}
.menu-item:hover{
  background-color:unset !important;
}
.menu-item:focus{
  background-color:unset !important;
}
</style>