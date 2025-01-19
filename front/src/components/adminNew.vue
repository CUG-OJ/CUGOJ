<script setup>
import {Search} from "@element-plus/icons-vue";
import {onMounted, reactive, ref} from 'vue'
import axios from "axios";
import {ElNotification} from "element-plus";
//url
const url='http://127.0.0.1:5174/api/new'
//组件的隐藏显示
const addVis = ref(false)
const updateVis = ref(false)
//返回值，查找操作
const total = defineModel('total', {default: 0})
const arr = defineModel('arr', {default: []})
//表单，增加操作、更新操作
const form = reactive({
  ID: '', //新闻编号
  importance: "false",//是否置顶
  title: '',//标题
  author: '',//作者
  content: '',//内容
  updateTime: '',//更新时间
})
//表单，查找操作
const readForm = reactive({
  value: '',
  type: '',
  page: 1,
  size: 12,
})

//表单，修改操作
function updateForm(row) {
  updateVis.value = true
  form.ID = row.ID
  form.title = row.title
  form.author = row.author
  form.content = row.content
  form.updateTime = row.updateTime
}

//页面加载时获取信息
onMounted(() => {
  read()
})

//增加操作HTTP
function create() {
  axios.post(url, form, {
    headers: {
      'Authorization': sessionStorage.getItem('token')
    }
  }).then(res => {
    read()
    addVis.value = false
    ElNotification({
      title: '添加成功',
      message: '添加成功',
      type: 'success',
      duration: 2500
    })
  }).catch(err => {
    ElNotification({
      title: '添加失败',
      message: err.response.data.message,
      type: 'error',
      duration: 2500
    })
  })
}

//删除操作HTTP
function remove(removeId) {
  axios.post(url, {
    params: {
      ID: removeId
    },
    headers: {
      'Authorization': sessionStorage.getItem('token')
    }
  }).then(res => {
    read()
    ElNotification({
      title: '删除成功',
      message: '删除成功',
      type: 'success',
      duration: 2500
    })
  }).catch(err => {
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
      title: '修改失败',
      message: err.response.data.message,
      type: 'error',
      duration: 2500
    })
  })
}

//查找操作HTTP
function read() {
  axios.post(url, readForm, {
    headers: {
      'Authorization': sessionStorage.getItem('token')
    }
  }).then(res => {
    total.value = res.data.total
    arr.value = res.data.arr
  }).catch(err => {
    ElNotification({
      title: '获取失败',
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
    <el-button @click="addVis=true">添加</el-button>
    <el-input @input="read" v-model="readForm.value" placeholder="输入数据搜索"
              style="width: 450px;height: auto;margin:auto 50px auto 50px">
      <template #prepend>
        <el-select @change="read()" v-model="readForm.type" placeholder="搜索类型" style="width: 100px">
          <el-option label="新闻编号" value="ID"/>
          <el-option label="标题" value="title"/>
          <el-option label="作者" value="author"/>
          <el-option label="内容" value="content"/>
          <el-option label="更新时间" value="updateTime"/>
        </el-select>
      </template>
      <template #append>
        <el-button @click="read()" :icon="Search"></el-button>
      </template>
    </el-input>
  </el-row>
  <!--  显示的表格-->
  <el-table max-height="540" v-model:data="arr" border stripe style="width: 100%;margin-top:20px">
    <el-table-column prop="newId" label="新闻编号"/>
    <el-table-column prop="title" label="标题"/>
    <el-table-column prop="author" label="作者"/>
    <el-table-column prop="content" label="内容"/>
    <el-table-column prop="updateTime" label="更新时间"/>
    <el-table-column label="操作">
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
        :total="total"/>
  </div>
<!--  窗口，更新操作-->
  <el-drawer v-model="addVis" :center="true" title= "添加新闻" size="50%" :with-header="true">
    <el-form :model="form" label-width="auto">
      <el-form-item label="新闻标题">
        <el-input v-model="form.title" placeholder="请输入标题"></el-input>
      </el-form-item>
      <el-form-item label="是否置顶">
        <el-radio-group v-model="form.importance">
          <el-radio size="large" value="true">是</el-radio>
          <el-radio size="large" value="false">否</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="新闻内容">
        <el-input
            type="textarea"
            :autosize="{ minRows: 10}"
            placeholder="请输入内容"
            v-model="form.content">
        </el-input>
      </el-form-item>
    </el-form>
    <span style="float: right">
      <el-button @click="addVis=false">取消</el-button>
      <el-button @click="create()" type="primary">添加</el-button>
    </span>
  </el-drawer>
  <!--  窗口，更新操作-->
  <el-drawer v-model="updateVis" :center="true" title= "更新新闻" size="50%" :with-header="true">
    <el-form :model="form" label-width="auto">
      <el-form-item label="新闻标题">
        <el-input v-model="form.title" placeholder="请输入标题"></el-input>
      </el-form-item>
      <el-form-item label="是否置顶">
        <el-radio-group v-model="form.importance">
          <el-radio size="large" value="true">是</el-radio>
          <el-radio size="large" value="false">否</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="新闻内容">
        <el-input
            type="textarea"
            :autosize="{ minRows: 10}"
            placeholder="请输入内容"
            v-model="form.content">
        </el-input>
      </el-form-item>
    </el-form>
    <span style="float: right">
      <el-button @click="updateVis=false">取消</el-button>
      <el-button @click="update()" type="primary">添加</el-button>
    </span>
  </el-drawer>
</template>

<style scoped>

</style>
