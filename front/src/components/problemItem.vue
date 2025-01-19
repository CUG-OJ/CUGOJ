<script lang="ts" setup>
import {reactive, ref} from 'vue'
import {Codemirror } from "vue-codemirror"
import { cpp } from '@codemirror/lang-cpp'
import { oneDark } from '@codemirror/theme-one-dark'
import axios from "axios";
import {ElNotification} from "element-plus";
//变量
const submitForm = {
  code: '',
}
let result = ref('')
//自测运行
function submit(){
  axios.post('http://120.79.32.206:8080/api/submit',submitForm,{
    headers:{
      'Content-Type':"application/x-www-form-urlencoded"}
  }).then(res => {
    ElNotification({
      title: '完成提交',
      message: res.data.message,
      type: 'info',
      duration: 2500
    })
    result.value = res.data.message
  }).catch(err => {
    ElNotification({
      title: '运行失败',
      message: err.message,
      type: 'error',
      duration: 2500
    })
  })
}
function testSubmit(){
  axios.post('http://120.79.32.206:8080/api/test',submitForm,{
    headers:{
      'Content-Type':"application/x-www-form-urlencoded"}
  }).then(res => {
    ElNotification({
      title: '完成自测',
      message: "测试不计入分数",
      type: 'info',
      duration: 2500
    })
    result.value = res.data.message
  }).catch(err => {
    ElNotification({
      title: '运行失败',
      message: err.message,
      type: 'error',
      duration: 2500
    })
  })
}
</script>
<template style="overflow-y: hidden;">
<div style=" float: left;height: 100%;width: 50%;overflow-y:scroll">
  A+B problem
  <el-divider></el-divider>
  <p>描述</p>
  输入两个整数a和b，输出它们的和。
  <p>输入</p>
  n (1 <= n <= 1000)
  接下来n行，每行包含两个整数a和b，表示a+b的值。
  <p>输出</p>
  一个整数，即a+b的值。
  <p>样例输入</p>
  4 <br/>
  1 2<br/>
  3 4<br/>
  5 6<br/>
  7 8<br/>
  <p>样例输出</p>
  3<br/>
  7<br/>
  11<br/>
  15<br/>
</div>
<div style="height: 100%">
  <div style=" float: right;height: 70%;width: 50%; ">
      <codemirror v-model="submitForm.code"
                style="height: 100%;width: 100%"
                placeholder="Code here..."
                :autofocus="true"
                :matchBrackets="true"
                :extensions="[cpp(), oneDark]"
                :tab-size="4" />
  </div>
  <div style="float: right;width: 49.9%;height: 30%;padding: 1%;border-top: #cdd0d6 1px solid">
<!--    <el-button>运行结果</el-button>-->
    <el-button disabled>自测输入</el-button>
    <el-button @click="testSubmit()" >自测运行</el-button>
    <el-button style="float: right;" @click="submit()">提交</el-button>
    <div style="white-space: pre-line;">{{result}}</div>
  </div>
</div>
</template>
