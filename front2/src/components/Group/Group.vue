<script setup>

</script>

<template>
  <span style="display: flex; justify-content: flex-end;">
    <el-text style="font-size: xx-large">我的班级</el-text>
    <span style="flex-grow: 1"></span>
    <el-input style="max-width: 200px;max-height: 32px" placeholder="班级名称或群号">aa</el-input>
    <el-button>搜索</el-button>
    <el-button>创建班级 </el-button>
  </span>
  <el-row>
      <el-card v-for="i in 20" style="width: 350px;height: 200px;margin:30px 60px">
          <template #header>
            <div class="card-header">
              <span>班级名称{{i}}</span>
            </div>
          </template>
        <p>成员数量</p>
        <p>群主</p>
        <p>群号</p>
      </el-card>
  </el-row>
</template>

<style scoped>

</style>