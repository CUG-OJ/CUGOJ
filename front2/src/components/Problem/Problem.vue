<script setup>
import {ArrowDown, Search} from "@element-plus/icons-vue";

const tableData = [
  {
    number: '1001',
    name: 'A+B',
    difficulty: '简单',
    problem: '计算两个整数的和',
    submitNumber: '100',
    ACrate: '80%',
    tag: '数学',
  },
  {
    number: '1002',
    name: 'A-B',
    difficulty: '简单',
    problem: '计算两个整数的差',
    submitNumber: '100',
    ACrate: '80%',
    tag: '数学'
  }
]
</script>

<template>
  <span style="display: flex; justify-content: flex-end;">
  <el-text style="font-size: x-large">问题列表</el-text>
  <span style="flex-grow: 1;" />
  <el-dropdown>
    <span class="el-dropdown-link2" style="color: black;margin-top:-10px">难度<el-icon><arrow-down /></el-icon></span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item>入门</el-dropdown-item>
        <el-dropdown-item>普及</el-dropdown-item>
        <el-dropdown-item>提高</el-dropdown-item>
        <el-dropdown-item>竞赛</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
  <el-switch
    size="large"
    style="height: 32px;margin: 0 50px"
    inline-prompt
    active-text="显示标签"
    inactive-text="隐藏标签"/>
  <el-input placeholder="请输入题目名称" style="width: 200px;height: 32px;margin-right: 50px">
      <template #append>
        <el-button :icon="Search" />
      </template>
  </el-input>
  <el-button>刷新</el-button>
</span>
<el-table :data="tableData" style="width: 100%">
  <el-table-column prop="number" label="#" />
  <el-table-column prop="name" label="题目" />
  <el-table-column prop="difficulty" label="难度" />
  <el-table-column prop="problem" label="问题" />
  <el-table-column prop="submitNumber" label="提交数" />
  <el-table-column prop="ACrate" label="通过率" />
  <el-table-column prop="tag" label="标签" />
</el-table>
</template>

<style scoped>
.el-dropdown-link2  {
  cursor: pointer;
  color: var(--el-color-primary);
  display: flex;
  align-items: center;
}
.el-dropdown-link2{
  font-size: larger;
}
.el-dropdown-link2:focus {
  outline: none;
}
</style>