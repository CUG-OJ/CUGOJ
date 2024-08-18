<script setup>
import {reactive, ref} from 'vue'
import { useRouter } from 'vue-router'
import axios from "axios";
import {ElNotification} from "element-plus";
const router = useRouter()
const regisForm = reactive({
  username: '',
  password: '',
  mail: '',
  code:''
})
function toLogin(){
  router.push({path:'/login'})
}
function register(){
  axios.post('http://120.79.32.206:8080/api/register',{
    username:regisForm.username,
    password:regisForm.password,
    mail:regisForm.mail,
    code:regisForm.code,
  },{
    headers:{
      'Content-Type':"application/x-www-form-urlencoded"
    }}).then(res => {
      localStorage.setItem('name',regisForm.username)
      localStorage.setItem('token', res.data.token)
      location.reload()
    }).catch(err => {
      ElNotification({
              title: '注册失败',
              message: err,
              type: 'error',
              duration: 2500
      })
    })
}
function get_code(){
  console.log(regisForm.mail)
  axios.post('http://120.79.32.206:8080/api/mail',{
    mail:regisForm.mail
  },{
    headers:{
      'Content-Type':"application/x-www-form-urlencoded"
    }
  }).then(res=>{
    ElNotification({
      title: '验证码已发送',
      message: '验证码已发送',
      type: 'success',
      duration: 2500
    })
  }).catch(err=>{
    ElNotification({
      title: '验证码发送失败',
      message: err,
      type: 'error',
      duration: 2500
    })
  })
}
</script>

<template>
  <el-card shadow="hover" style="width: 300px;height: 350px;float: right;margin: 150px 150px auto auto">
    <el-text style="font-size: x-large;margin-left: 100px">注册</el-text>
    <el-input class="item" v-model="regisForm.username" placeholder="请输入用户名" style="width: 250px;"></el-input>
    <el-input class="item" show-password v-model="regisForm.password" placeholder="请输入密码" style="width: 250px;"></el-input>
    <el-input class="item" v-model="regisForm.mail" placeholder="请输入邮箱" style="width: 250px;"></el-input>
    <el-input class="item" v-model="regisForm.code" placeholder="请输入验证码" style="width: 250px;">
      <template #append>
        <el-button @click="get_code" style="height: 32px">获取验证码</el-button>
      </template>
    </el-input>
    <el-button class="item" link @click="toLogin">已有账号？点击登录</el-button>
    <br/>
    <el-button @click="register" style="width: 250px;">注册</el-button>
  </el-card>
</template>

<style scoped>
.item{
  margin: 8px 0;
  font-size: 15px;
}
</style>
