<template>
  <el-container>
    <el-aside style="max-width: 150px">
      <el-menu style="user-select: none" router>
        <el-text style="font-size: x-large;">助理</el-text>
        <el-menu-item index="/admin/dashboard"><el-icon><Menu /></el-icon>概览</el-menu-item>
        <el-menu-item index="/admin/new"><el-icon><List /></el-icon>发布新闻</el-menu-item>
        <el-menu-item index="/admin/user"><el-icon><Avatar /></el-icon>管理用户</el-menu-item>
        <el-menu-item index="/admin/problem"><el-icon><List /></el-icon>导入题目</el-menu-item>
        <el-menu-item index="/admin/createProblem"><el-icon><Plus /></el-icon>创建题目</el-menu-item>
      </el-menu>
    </el-aside>
    <el-main>
    <router-view></router-view>
    </el-main>
  </el-container>
</template>
<script lang="ts" setup>
import {Avatar, List, Menu, Plus, Setting, Trophy} from "@element-plus/icons-vue";
</script>
