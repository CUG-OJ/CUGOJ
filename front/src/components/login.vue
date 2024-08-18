<script setup>
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import axios from "axios";
import {ElNotification} from "element-plus";
const router = useRouter()
const loginForm = reactive({
  username: '',
  password: '',
})
function toRegister(){
  router.push({path:'/register'})
}
function login(){
  axios.post('http://120.79.32.206:8080/api/login',loginForm,{
    headers:{
      'Content-Type':"application/x-www-form-urlencoded"
    }}).then(res => {
      localStorage.setItem('username',loginForm.username)
      localStorage.setItem('token', res.data.token)
      location.reload()
    }).catch(err => {
      ElNotification({
              title: '登录失败',
              message: err.message,
              type: 'error',
              duration: 2500
      })
    })
}
</script>

<template>
  <el-card shadow="hover" style="width: 300px;height: 300px;float: right;margin: 150px 150px auto auto">
    <el-text style="font-size: x-large;margin-left: 100px">登录</el-text>
    <el-input class="item" v-model="loginForm.username" placeholder="请输入账号" style="width: 250px;"></el-input>
    <el-input class="item" v-model="loginForm.password" placeholder="请输入密码" style="width: 250px;"></el-input>
    <el-button class="item" link @click="">忘记密码?</el-button>
    <el-button class="item" link @click="toRegister" style="float: right">注册账号</el-button>
    <br/>
    <el-button @click="login" style="width: 250px">登录</el-button>
  </el-card>
</template>

<style scoped>
.item{
  margin: 10px 0;
  font-size: 15px;
}
</style>
