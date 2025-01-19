<script setup>
import {Search} from "@element-plus/icons-vue";
import {onMounted, reactive, ref} from 'vue'
import axios from "axios";
//url
const url='http://127.0.0.1:5174/api/problem'
//组件的隐藏显示
const addVis = ref(false)
const updateVis=ref(false)
//返回值，查找操作
const amount = defineModel('total',{ default: 0 })
const arr = defineModel('arr',{default:[]})
//表单，增加操作、更新操作
const form = reactive({
  problemID: '', //题目编号
  problemTitle: '',//题目标题
  timeLimit: '',//时间限制
  memoryLimit: '',//内存限制
  content: '',//题目内容
  input: '',//输入
  output: '',//输出
  sampleInput: '',//样例输入
  sampleOutput: '',//样例输出
  hint: '',//提示
  source: '',//来源
  testInput: '',//测试输入
  testOutput: '',//测试输出
  solution: '',//解答
  updateTime: '',//更新时间
  operateType: '',//操作类型
})
//表单，查找操作
const readForm = reactive({
  value: '',
  type: '',
  page: 1,
  size: 12,
})
//表单，修改操作-读取
function updateForm(row){
  updateVis.value=true
  form.problemID=row.problemID
  form.problemTitle=row.problemTitle
  form.timeLimit=row.timeLimit
  form.memoryLimit=row.memoryLimit
  form.content=row.content
  form.input=row.input
  form.output=row.output
  form.sampleInput=row.sampleInput
  form.sampleOutput=row.sampleOutput
  form.hint=row.hint
  form.source=row.source
  form.testInput=row.testInput
  form.testOutput=row.testOutput
  form.solution=row.solution
  form.updateTime=row.updateTime
}
//页面加载时获取设备信息
onMounted(()=>{
  read()
})
//增加操作HTTP
function create(){
  axios.post(url,form,{
    headers:{
      'Authorization':sessionStorage.getItem('token')
    }
  }).then(res=>{
    read()
    addVis.value=false
    ElNotification({
      title: '添加成功',
      message: '添加成功',
      type: 'success',
      duration: 2500
    })
  }).catch(err=>{
    ElNotification({
      title: '添加失败',
      message: err.response.data.message,
      type: 'error',
      duration: 2500
    })
  })
}
//删除操作HTTP
function remove(removeId){
  axios.post(url,{
    params:{
      newId:removeId
    },
    headers:{
      'Authorization':sessionStorage.getItem('token')
    }
  }).then(res=>{  read()
  }).catch(err=>{
    ElNotification({
      title: '删除失败',
      message: err.response.data.message,
      type: 'error',
      duration: 2500
    })
  })
}
//修改操作HTTP
function update() {
  axios.post(url, form, {
    headers: {
      'Authorization': sessionStorage.getItem('token')
    }
  }).then(res => {
    read()
    updateVis.value = false
  }).catch(err => {
    ElNotification({
      title: '添加失败',
      message: err.response.data.message,
      type: 'error',
      duration: 2500
    })
  })
}
//查找操作HTTP
function read(){
  axios.post(url,readForm,{
    headers:{
      'Authorization':sessionStorage.getItem('token')
    }
  }).then(res=>{
    amount.value=res.data.total
    arr.value=res.data.user
  }).catch(err=>{
    ElNotification({
      title: '获取用户信息失败',
      message: err.response.data.message,
      type: 'error',
      duration: 2500
    })
  })
}
</script>

<template>
  <!--  搜索框-->
  <el-row>
    <span style="flex-grow: 1"></span>
    <el-button @click="">FPS导入</el-button>
    <el-button @click="addVis=true">添加</el-button>
    <el-input @input="read" v-model="readForm.value" placeholder="输入数据搜索"
              style="width: 450px;height: auto;margin:auto 50px auto 50px">
      <template #prepend>
        <el-select @change="read()" v-model="readForm.type" placeholder="搜索类型" style="width: 100px">
          <el-option label="题目编号" value="problemID" />
          <el-option label="题目标题" value="problemTitle" />
          <el-option label="时间限制" value="timeLimit" />
          <el-option label="内存限制" value="memoryLimit" />
          <el-option label="题目内容" value="content" />
          <el-option label="输入" value="input" />
          <el-option label="输出" value="output" />
          <el-option label="样例输入" value="sampleInput" />
          <el-option label="样例输出" value="sampleOutput" />
          <el-option label="提示" value="hint" />
          <el-option label="来源" value="source" />
          <el-option label="测试输入" value="testInput" />
          <el-option label="测试输出" value="testOutput" />
          <el-option label="解答" value="solution" />
          <el-option label="更新时间" value="updateTime" />
        </el-select>
      </template>
      <template #append>
        <el-button @click="read()" :icon="Search"></el-button>
      </template>
    </el-input>
  </el-row>
  <!--  显示的表格-->
  <el-table max-height="540" v-model:data="arr" border stripe style="width: 100%;margin-top:20px">
    <el-table-column prop="problemID" label="题目编号"  />
    <el-table-column prop="problemTitle" label="题目标题"  />
    <el-table-column prop="timeLimit" label="时间限制"  />
    <el-table-column prop="memoryLimit" label="内存限制"  />
    <el-table-column prop="updateTime" label="更新时间"  />
    <el-table-column fixed="right" label="操作">
    <template #default="item">
      <el-button @click="remove(item.row.userId)" link type="primary">删除</el-button>
      <el-button @click="updateForm(item.row)" link type="primary">编辑</el-button>
      <el-button @click="" link type="primary">详情</el-button>
    </template>
  </el-table-column>
  </el-table>
  <!--  分页组件-->
  <div style="position: absolute;bottom: 15px ">
    <el-pagination
        background
        @current-change="read()"
        @size-change="read()"
        :page-sizes="[8, 12, 50, 100]"
        v-model:current-page="readForm.page"
        v-model:page-size="readForm.size"
        layout="sizes,prev, pager, next,jumper"
        :total="amount"/>
  </div>
  <!--  窗口，添加操作-->
  <el-drawer v-model="addVis" :center="true" title= "添加题目" size="50%" :with-header="true">
    <el-form :model="form" label-width="auto">
      <el-form-item label="题目标题">
        <el-input v-model="form.problemTitle" placeholder="请输入题目标题"></el-input>
      </el-form-item>
      <el-row>
        <el-form-item label="时间限制">
          <el-input v-model="form.timeLimit" placeholder="请输入时间限制"></el-input>
        </el-form-item>
        <el-form-item label="内存限制" style="margin-left: 40px">
          <el-input v-model="form.memoryLimit" placeholder="请输入内存限制"></el-input>
        </el-form-item>
      </el-row>
      <el-form-item label="题目内容">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入内容"
            v-model="form.content">
        </el-input>
      </el-form-item>
      <el-form-item label="输入">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入输入"
            v-model="form.input">
        </el-input>
      </el-form-item>
      <el-form-item label="输出">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入输出"
            v-model="form.output">
        </el-input>
      </el-form-item>
      <el-form-item label="样例输入">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入样例输入"
            v-model="form.sampleInput">
        </el-input>
      </el-form-item>
      <el-form-item label="样例输出">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入样例输出"
            v-model="form.sampleOutput">
        </el-input>
      </el-form-item>
      <el-form-item label="提示">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入提示"
            v-model="form.hint">
        </el-input>
      </el-form-item>
      <el-form-item label="来源">
        <el-input v-model="form.source" placeholder="请输入来源"></el-input>
      </el-form-item>
      <el-form-item label="测试输入">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入测试输入"
            v-model="form.testInput">
        </el-input>
      </el-form-item>
      <el-form-item label="测试输出">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入测试输出"
            v-model="form.testOutput">
        </el-input>
      </el-form-item>
      <el-form-item label="解答">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入解答"
            v-model="form.solution">
        </el-input>
      </el-form-item>
    </el-form>
    <span style="float: right">
      <el-button @click="addVis=false">取消</el-button>
      <el-button @click="create()" type="primary">添加</el-button>
    </span>
  </el-drawer>
  <!--  窗口，更新操作-->
  <el-drawer v-model="updateVis" :center="true" title= "更新题目" size="50%" :with-header="true">
    <el-form :model="form" label-width="auto">
      <el-form-item label="题目标题">
        <el-input v-model="form.problemTitle" placeholder="请输入题目标题"></el-input>
      </el-form-item>
      <el-row>
        <el-form-item label="时间限制">
          <el-input v-model="form.timeLimit" placeholder="请输入时间限制"></el-input>
        </el-form-item>
        <el-form-item label="内存限制" style="margin-left: 40px">
          <el-input v-model="form.memoryLimit" placeholder="请输入内存限制"></el-input>
        </el-form-item>
      </el-row>
      <el-form-item label="题目内容">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入内容"
            v-model="form.content">
        </el-input>
      </el-form-item>
      <el-form-item label="输入">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入输入"
            v-model="form.input">
        </el-input>
      </el-form-item>
      <el-form-item label="输出">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入输出"
            v-model="form.output">
        </el-input>
      </el-form-item>
      <el-form-item label="样例输入">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入样例输入"
            v-model="form.sampleInput">
        </el-input>
      </el-form-item>
      <el-form-item label="样例输出">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入样例输出"
            v-model="form.sampleOutput">
        </el-input>
      </el-form-item>
      <el-form-item label="提示">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入提示"
            v-model="form.hint">
        </el-input>
      </el-form-item>
      <el-form-item label="来源">
        <el-input v-model="form.source" placeholder="请输入来源"></el-input>
      </el-form-item>
      <el-form-item label="测试输入">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入测试输入"
            v-model="form.testInput">
        </el-input>
      </el-form-item>
      <el-form-item label="测试输出">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入测试输出"
            v-model="form.testOutput">
        </el-input>
      </el-form-item>
      <el-form-item label="解答">
        <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            placeholder="请输入解答"
            v-model="form.solution">
        </el-input>
      </el-form-item>
    </el-form>
    <span style="float: right">
      <el-button @click="addVis=false">取消</el-button>
      <el-button @click="update()" type="primary">添加</el-button>
    </span>
  </el-drawer>
</template>

<style scoped>

</style>
