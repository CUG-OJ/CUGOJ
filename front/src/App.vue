<script setup>
import Header from "@/components/header.vue";
</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-header style="user-select: none"><Header/></el-header>
      <el-main>
          <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<style scoped>
.common-layout {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
</style>
