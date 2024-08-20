<script setup>

</script>

<template>
<el-container>
  <el-aside>
    <el-card>
      <el-text>note</el-text>
      <el-divider/>
      <el-text>vis:</el-text>
      <el-divider/>
      <el-text>加群方式:</el-text>
      <el-divider/>
      <el-text>群主:</el-text>
      <el-divider/>
      <el-text>管理员</el-text>
      <el-divider/>
      <el-text>比赛数量</el-text>
      <el-divider/>
      <el-text>member number:</el-text>
      <el-divider/>
      <el-text>create time</el-text>
    </el-card>
  </el-aside>
  <el-main>
    <el-row>
      <el-text>name</el-text>
      <el-button type="danger">leave</el-button>
    </el-row>
    <el-card>Board</el-card>
    <el-card>contests</el-card>
    <el-card>
      <el-text style="font-size: x-large">Members</el-text>
      <el-table stripe border style="width: 100%">
        <el-table-column prop="date" label="Date" width="180" />
        <el-table-column prop="name" label="Name" width="180" />
        <el-table-column prop="address" label="Address" />
      </el-table>
    </el-card>
  </el-main>
</el-container>
</template>

<style scoped>

</style>