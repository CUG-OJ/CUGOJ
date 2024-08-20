<script setup>
let DeviceInfo= {
  ip: '127.0.0.1',
  os: 'Windows 10',
  browser: 'Chrome 91.0.4472.124',
  time: '2021-07-15 12:00:00',
}
let form = {
  name: 'UserName',
  password: '',
}
let form2 = {
  name: ''
}
let num=1;
</script>

<template>
  <el-row>
    <el-col :span="8">
    <el-text style="font-size: x-large">安全设置</el-text>
    <el-form :model="form" :inline="true" label-width="auto" style="max-width: 400px" >
      <el-text>30天内只可修改一次用户名</el-text>
      <el-form-item label="用户名" label-width="55">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item><el-button type="primary">修改</el-button></el-form-item>
    </el-form>
    <el-form :model="form2" :inline="true" label-width="auto" style="max-width: 400px" >
      <el-form-item label="邮件" label-width="55">
        <el-input v-model="form2.name" />
      </el-form-item>
      <el-form-item><el-button type="primary">修改</el-button></el-form-item>
    </el-form>
      <el-form :model="form" label-width="auto" style="max-width: 300px" >
      <el-form-item label="原密码" label-width="55">
        <el-input v-model="form.password" show-password />
      </el-form-item>
      <el-form-item label="新密码" label-width="55">
        <el-input v-model="form.password" show-password />
      </el-form-item>
      <el-form-item><el-button type="primary">修改</el-button></el-form-item>
    </el-form>
    </el-col>
    <el-col :span="16">
    <el-text style="font-size: x-large">登录设备&nbsp;{{num}}/3</el-text>
      <el-row>
      <el-card v-for="_ in 3" style="margin-right: 20px;max-width: 250px">
        <el-row type="flex" justify="space-between">
          <el-text class="el-text-item">{{DeviceInfo.ip}}</el-text>
          <el-tag type="success" style="font-size: medium">本机</el-tag>
        </el-row>
        <el-row><el-text class="el-text-item">{{DeviceInfo.os}}</el-text></el-row>
        <el-row><el-text class="el-text-item">{{DeviceInfo.browser}}</el-text></el-row>
        <el-row><el-text class="el-text-item">{{DeviceInfo.time}}</el-text></el-row>
      </el-card>
      </el-row>
    </el-col>
  </el-row>
  <el-text style="font-size: x-large">双因素验证</el-text>
</template>

<style scoped>
.el-text-item{
  font-size: larger;
}
</style>