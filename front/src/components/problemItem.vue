<script lang="ts" setup>
let code = ref('')
import  { ref } from 'vue'
import {Codemirror } from "vue-codemirror"
import { cpp } from '@codemirror/lang-cpp'
import { oneDark } from '@codemirror/theme-one-dark'
</script>
<template style="overflow-y: hidden">
<div class="left" style=" float: left;height: 500px;width: 49.5%;overflow-y: auto;">
  A+B problem
  <br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br>
</div>
<div class="right" style=" float: right;height: 500px;width: 49.5%; overflow-y: auto;">
  <codemirror v-model="code"
            class="codemirror"
            placeholder="Code here..."
            :theme="oneDark"
            :matchBrackets="true"
            :extensions="[cpp()]"
            :tab-size="4" />
</div>

</template>
<style>

</style>

