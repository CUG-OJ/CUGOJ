<template>
  <el-container>
    <el-aside style="max-width: 200px">
      <el-menu>
        <el-text style="font-size: x-large;">管理员</el-text>
        <el-menu-item index="/admin/dashboard"><el-icon><Menu /></el-icon>概览</el-menu-item>
        <el-menu-item index="/admin/setting"><el-icon><Setting /></el-icon>系统设置</el-menu-item>
        <el-menu-item index="/admin/user"><el-icon><Avatar /></el-icon>管理用户</el-menu-item>
        <el-menu-item index="/admin/problem"><el-icon><List /></el-icon>导入题目</el-menu-item>
        <el-menu-item index="/admin/createProblem"><el-icon><Plus /></el-icon>创建题目</el-menu-item>
        <el-menu-item index="/admin/createContest"><el-icon><Trophy /></el-icon>创建竞赛</el-menu-item>
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
