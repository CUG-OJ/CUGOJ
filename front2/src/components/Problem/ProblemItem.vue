<script setup>
import { ref, onMounted } from 'vue';
import { json } from '@codemirror/lang-json';
import {python } from '@codemirror/lang-python';
import { oneDark } from '@codemirror/theme-one-dark'
import {marked} from "marked";
const initJson = {
  name: `maybaby`,
  year: 25,
  weight: 45,
  height: 165
}
// 初始化
let codeVal = ref('');
// 转成json字符串并格式化
codeVal.value = JSON.stringify(initJson, null, '\t')
// json
const lang = json();

onMounted(() => {
})
//拓展
let extensions=[oneDark]
let markdownText = `
# 题目描述
aaab \n
aac
`

</script>

<template>
  <el-container>
    <el-main>
      <div v-html="marked(markdownText)"></div>
    </el-main>
    <el-main style="min-height: 600px">
      <div class="main" style="min-height: 600px">
        <code-mirror
          v-model="codeVal"
          :extensions="extensions"
          :lang="lang"
          :theme="oneDark"
          :options="{lineNumbers: true}"
          style="height: 100%; width: 100%"
        />
    </div>
    </el-main>
  </el-container>
</template>

<style scoped lang="css">
</style>