<script setup>
import {
  Grid,
  Histogram,
  HomeFilled,
  MoreFilled,
  TrendCharts,
  TrophyBase,
  UserFilled
} from "@element-plus/icons-vue";
let name = 'USER'
</script>

<template>
  <div class="common-layou2t" style="margin-top: 20px">
    <el-container>
      <el-aside width="150px">
          <el-menu
            class="el-menu"
            mode="vertical"
            :ellipsis="false"
            router
            style="text-align: center"
          >
            <el-text style="text-align: center;">{{name}}</el-text>
            <el-menu-item class="menu-item" index="/user/message" style="justify-content: center;font-size: large">通知</el-menu-item>
            <el-menu-item class="menu-item" index="/user/profile" style="justify-content: center;font-size: large">资料</el-menu-item>
            <el-menu-item class="menu-item" index="/user/account" style="justify-content: center;font-size: large">账户</el-menu-item>
            <el-menu-item class="menu-item" index="/user/security" style="justify-content: center;font-size: large">安全</el-menu-item>
          </el-menu>
      </el-aside>
      <el-main><router-view></router-view></el-main>
    </el-container>
  </div>
</template>

<style scoped>
.flex-grow {
  flex-grow: 1;
}
.el-dropdown-link  {
  cursor: pointer;
  color: var(--el-color-primary);
  display: flex;
  align-items: center;
}
.el-dropdown-link{
  font-size: larger;
}
.el-dropdown-link:focus {
  outline: none;
}
.router-link{
  color:#606266;
}
.router-link:hover {
  background-color:unset;
}
.drop-item:focus{
  color: unset !important;
}
.menu-item:hover{
  background-color:unset !important;
}
.menu-item:focus{
  background-color:unset !important;
}
</style>