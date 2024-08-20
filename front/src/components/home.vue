<script setup>
let notes=[
  {
    title: '使用方法',
    content: '正常注册使用即可:heart:'
  },
]
const tableData = [
  {
    number: 1,
    name: '张三张三张三张三张三张三张三张三张三张三张三张三张三张三张三张三张三张三张三张三',
    address: '18'
  },
  {
    number: 2,
    name: '李四',
    address: '20000000000000'
  },
  {
    number: 3,
    name: '王五',
    address: '22'
  },
  {
    number: 4,
    name: '赵六',
    address: '24'
  },
  {
    number: 5,
    name: '孙七',
    address: '26'
  },
  {
    number: 6,
    name: '周八',
    address: '28'
  },
  {
    number: 7,
    name: '吴九',
    address: '30'
  },
  {
    number: 8,
    name: '郑十',
    address: '32'
  },
  {
    number: 9,
    name: '钱十一',
    address: '34'
  },
  {
    number: 10,
    name: '孔十二',
    address: '36'
  },
  {
    number: 11,
    name: '曹十三',
    address: '38'
  },
  {
    number: 12,
    name: '秦十四',
    address: '40'
  },
  {
    number: 13,
    name: '汉十五',
    address: '42'
  },
  {
    number: 14,
    name: '唐十六',
    address: '44'
  },
  {
    number: 1555555555555,
    name: '宋十七',
    address: '46'
  }
]
import Markdown from 'vue3-markdown-it';
</script>
<template>
  <el-container style="margin: 1.2%">
    <el-main>
      <el-card shadow="always" style="width: 100%;margin-top: -1%;background-color:#f4f4f5">
        <template #header style="align-content: center;line-height: 1px">
          公告
        </template>
        <Markdown :source="'置顶公告的样式'" />
      </el-card>
      <el-card shadow="hover" v-for="item in notes" style="width: 100%;margin-top:2%">
        <template #header style="align-content: center;line-height: 1px">
          {{item.title}}
        </template>
        <Markdown :source="item.content" />
      </el-card>
    </el-main>
    <el-aside>
      <el-table :data="tableData" stripe border
                show-overflow-tooltip
                :cell-style="{'text-align':'center'}"
                :header-cell-style="{'color':'black','text-align':'center'}">
        <el-table-column prop="number" label="#"/>
        <el-table-column prop="name" label="用户"/>
        <el-table-column prop="address" label="Rating" />
      </el-table>
    </el-aside>
  </el-container>
</template>

<style scoped>

</style>
