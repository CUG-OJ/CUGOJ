<script lang="ts" setup>
let code = ref('')
import  { ref } from 'vue'
import {Codemirror } from "vue-codemirror"
import { cpp } from '@codemirror/lang-cpp'
import { oneDark } from '@codemirror/theme-one-dark'
</script>
<template style="overflow-y: hidden;">
<div style=" float: left;height: 100%;width: 50%;overflow-y:scroll">
  A+B problem
  <br><br><br><br><br><br><br><br><br><br><br><br><br><br>aaa<br><br><br><br><br>aaa
  <br><br><br><br><br><br><br><br><br><br>aaa<br><br><br><br><br>aaa
</div>
<div style="height: 100%">
  <div style=" float: right;height: 70%;width: 50%; ">
      <codemirror v-model="code"
                style="height: 100%;width: 100%"
                placeholder="Code here..."
                :autofocus="true"
                :matchBrackets="true"
                :extensions="[cpp(), oneDark]"
                :tab-size="4" />
  </div>
  <div style="float: right;width: 49.9%;height: 30%;padding: 1%;border-top: #cdd0d6 1px solid">
    <el-button>运行结果</el-button>
    <el-button>自测输入</el-button>
    <el-button>自测运行</el-button>
      <el-button style="float: right;">提交</el-button>
  </div>
</div>
</template>

