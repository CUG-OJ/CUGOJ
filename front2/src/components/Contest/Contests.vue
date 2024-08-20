<script setup>
import {Calendar, Clock, Search} from "@element-plus/icons-vue";
const options = [
  { value:'all', label:'全部' },
  { value:'icpc', label:'ICPC' },
  { value:'ioi', label:'IOI' },
  { value:'oi', label:'OI' },
  {value: 'other',label:'其他'}
]
let test = ''
const testData=[
  {startTime:'2021-10-01',name:'第一届篮球杯程序设计竞赛-中国地质大学（左岭）',type:'ICPC',endTime:"2021-10-02",status:'未开始'},
  {startTime:'2021-10-01',name:'第二届',type:'ICPC',endTime:"2021-10-02",status:'未开始'},
  {startTime:'2021-10-01',name:'第三届',type:'OI',endTime:"2021-10-02",status:'未开始'},
  {startTime:'2021-10-01',name:'第四届',type:'IOI',endTime:"2021-10-02",status:'未开始'},
  {startTime:'2021-10-01',name:'第五届',type:'IOI',endTime:"2021-10-02",status:'未开始'},
  {startTime:'2021-10-01',name:'第六届',type:'OI',endTime:"2021-10-02",status:'未开始'},
  {startTime:'2021-10-01',name:'第七届',type:'ICPC',endTime:"2021-10-02",status:'未开始'},
  {startTime:'2021-10-01',name:'第八届',type:'ICPC',endTime:"2021-10-02",status:'未开始'},
  {startTime:'2021-10-01',name:'第九届',type:'ICPC',endTime:"2021-10-02",status:'未开始'},
]
</script>

<template>
 <div style="margin-bottom: 20px">
   <el-cascader
       placeholder="赛制"
       :options="options"
        v-model="value"
       hoverThreshold="500"
       @change="handleChange"
       style="margin-right: 20px"
   />
    <el-input
      v-model="test1"
      style="max-width: 200px"
      placeholder="输入关键词"
    >
      <template #append>
        <el-button :icon="Search" />
      </template>
    </el-input>
  </div>
  <li v-for="item in testData" :key="item.name" class="list-item" style="list-style: none;">
    <el-card shadow="never" style="margin-top: 5px">
      <div style="float: left">
        <el-avatar :size="80" :src="circleUrl" style="float: left;margin-right: 30px" />
      </div>
    <div slot="header" style="font-size: 22px;margin-bottom: 20px">
      <el-button link style="font-size: 20px">{{item.name}}</el-button>
      <el-tag style="float: right">{{item.status}}</el-tag>
    </div>
    <div style="font-size: 18px">
      <el-icon><Calendar /></el-icon>{{item.endTime}}
      <span style="margin: 0 10px"/>
      <el-icon><Clock /></el-icon>{{item.endTime}}
      <span style="margin: 0 10px"/>
      <el-button size="small" round>{{item.type}}</el-button>
    </div>
    </el-card>
  </li>

 <el-row style="margin-top: 30px">
    <el-col :span="12"></el-col>
    <el-col :span="12">
      <el-pagination
        :page-size="20"
        :pager-count="8"
        background
        layout="prev, pager, next"
        :total="1000"
      /></el-col>
  </el-row>

<el-backtop :right="150" :bottom="150"style="font-size: 200px"/>
</template>

<style scoped>

</style>