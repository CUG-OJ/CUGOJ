<script setup>
import AppHeader from "@/components/AppHeader.vue";
</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-header style="user-select: none"><AppHeader/></el-header>
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
