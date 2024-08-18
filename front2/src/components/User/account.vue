<script setup>
let tmp={
    form:{
        name:''
    }
}
</script>

<template>
  <el-container>
    <el-main>
      <el-text style="font-size: x-large">个人信息</el-text>  <br/><br/>
      <el-form :model="tmp" label-width="auto" style="max-width: 300px">
        <el-form-item label="学校">
          <el-input v-model="tmp.name" />
        </el-form-item>
        <el-form-item label="专业">
          <el-input v-model="tmp.name" />
        </el-form-item>
        <el-button type="primary">修改</el-button>
      </el-form>
    </el-main>
    <el-main>
      <el-text style="font-size: x-large">远程判题</el-text>  <br/><br/>
      <el-form :model="tmp" label-width="auto" style="max-width: 300px">
        <el-form-item label="Codeforces">
          <el-input v-model="tmp.name" />
        </el-form-item>
        <el-form-item label="洛谷">
          <el-input v-model="tmp.name" />
        </el-form-item>
        <el-form-item label="AtCoder">
          <el-input v-model="tmp.name" />
        </el-form-item>
        <el-form-item label="HDU">
          <el-input v-model="tmp.name" />
        </el-form-item>
        <el-form-item label="POJ">
          <el-input v-model="tmp.name" />
        </el-form-item>
        <el-form-item label="SPOJ">
          <el-input v-model="tmp.name" />
        </el-form-item>
      </el-form>
    </el-main>
  </el-container>
</template>

<style scoped>

</style>