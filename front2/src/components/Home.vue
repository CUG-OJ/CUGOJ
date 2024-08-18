<script setup>
import { marked } from 'marked'
let notes=[
  {
    title: '广告位招租',
    content: 'welcome to pay for the advertisement'
  },
  {
    title: '个人资料',
    content: '姓名：张三\n年龄：18\n性别：男'
  }
]
const tableData = [
  {
    //三三三三三三三三三三三三三三
    number: 1,
    name: '张四',
    address: '18'
  },
  {
    number: 2,
    name: '李四',
    address: '20'
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
    number: 15,
    name: '宋十七',
    address: '46'
  }
]

let markdownText = `
# 个人资料
## 个人信息
- 姓名：张三
- 年龄：18
- 性别：男
`
</script>

<template>
  <div class="common-layout2">
    <el-container>
      <el-main>
          <el-card style="max-width: 900px;background-color:#f2f2f2">
            <template #header style="align-content: center;line-height: 1px">广告位招租</template>
            welcome to pay for the advertisement
            <div v-html="marked(markdownText)"></div>
          </el-card>
          <el-card v-for="item in notes" style="max-width: 900px;margin-top:10px">
            <template #header style="align-content: center;line-height: 1px">{{item.title}}</template>
            {{item.content}}
          </el-card>
      </el-main>
      <el-aside width="400px" style="margin-right:50px">
<!--      解析markdown-->
        <div>
            <el-table :data="tableData" stripe border style="width: 100%"
                      :cell-style="{'text-align':'center'}"
                      :header-cell-style="{'color':'black','text-align':'center'}">
              <el-table-column width="100px" prop="number" label="#"/>
              <el-table-column width="200px" prop="name" label="用户"/>
              <el-table-column width="100px" prop="address" label="Rating" />
            </el-table>
        </div>
      </el-aside>
    </el-container>
  </div>
</template>

<style scoped>

</style>