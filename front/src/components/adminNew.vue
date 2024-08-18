<script setup>
import {Search} from "@element-plus/icons-vue";
import {onMounted, reactive, ref} from 'vue'
import axios from "axios";
//组件的隐藏显示
const addVis = ref(false)
const editVis=ref(false)
//GET请求参数
const searchValue = defineModel('input',{ default: '' })
const searchType=defineModel('searchType',{default:'deviceID'})
const page = defineModel('page',{ default: 1 })
const pageSize = defineModel('pageSize',{ default: 12 })
//GET返回参数
const total = defineModel('total',{ default: 0 })
const devices=defineModel('devices',{default:[]})
//POST请求参数
const postDeviceForm = reactive({
  deviceID: '',
  deviceName: '',
  Manufacturer: '',
  ManufactureDate: ''
})
//PUT请求参数
const editDeviceForm = reactive({
  deviceID: '',
  newDeviceID: '',
  deviceName: '',
  Manufacturer: '',
  ManufactureDate: ''
})
//加载设备信息到编辑框
function pre_edit_device(row){
  editVis.value=true
  editDeviceForm.newDeviceID=row.deviceID
  editDeviceForm.deviceID=row.deviceID
  editDeviceForm.deviceName=row.deviceName
  editDeviceForm.Manufacturer=row.Manufacturer
  editDeviceForm.ManufactureDate=row.ManufactureDate
}
//页面加载时获取设备信息
onMounted(()=>{
  get_device()
})
//携带token进行GET请求
function get_device(){
  axios.get('http://127.0.0.1:5174/api/device',{
    headers:{
      'Authorization':sessionStorage.getItem('token')
    },
    params:{
      searchValue:searchValue.value,
      searchType:searchType.value,
      page:page.value,
      size:pageSize.value
    }
  }).then(res=>{
    total.value=res.data.total
    devices.value=res.data.devices
  }).catch(err=>{
    ElNotification({
      title: '获取设备失败',
      message: err.response.data.message,
      type: 'error',
      duration: 2500
    })
  })
}
//携带token进行POST请求
function post_device(){
  axios.post('http://127.0.0.1:5174/api/device',postDeviceForm,{
    headers:{
      'Content-Type':"application/x-www-form-urlencoded",
      'Authorization':sessionStorage.getItem('token')
    }
  }).then(res=>{
    get_device()
    addVis.value=false
    ElNotification({
      title: '添加设备成功',
      message: '添加设备成功',
      type: 'success',
      duration: 2500
    })
  }).catch(err=>{
      ElNotification({
        title: '添加设备失败',
        message: err.response.data.message,
        type: 'error',
        duration: 2500
      })
  })
}
//携带token进行删除请求，根据编号删除设备
function delete_device(number){
  axios.delete('http://127.0.0.1:5174/api/device',{
    params:{
      device_id:Number(number)
    },
    headers:{
      'Content-Type':"application/x-www-form-urlencoded",
      'Authorization':sessionStorage.getItem('token')
    }
  }).then(res=>{  get_device()
  }).catch(err=>{
    ElNotification({
      title: '删除设备失败',
      message: err.response.data.message,
      type: 'error',
      duration: 2500
    })
  })
}
//携带token进行PUT请求，根据编号修改设备信息
function edit_device() {
  axios.put('http://127.0.0.1:5174/api/device', editDeviceForm, {
    headers: {
      'Content-Type': "application/x-www-form-urlencoded",
      'Authorization': sessionStorage.getItem('token')
    }
  }).then(res => {
    get_device()
    editVis.value = false
  }).catch(err => {
    ElNotification({
      title: '添加设备失败',
      message: err.response.data.message,
      type: 'error',
      duration: 2500
    })
  })
}
//检查是否为数字
function notNumber(val){
  //如果为空则不显示提示
  if(val===''||val==null){
    return false
  }
  //如果不是数字则显示提示
  return isNaN(val);
}
</script>

<template>
  <el-row>
  <span style="flex-grow: 1"></span>
  <el-button @click="addVis=true">添加</el-button>
  <el-tooltip :visible="notNumber(searchValue)&&searchType==='deviceID'" content="请输入数字" placement="top">
  <el-input @input="get_device" v-model="searchValue" placeholder="输入数据搜索" style="width: 350px;height: auto;margin:auto 50px auto 50px">
    <template #prepend>
      <el-select @change="get_device()" v-model="searchType" placeholder="搜索类型" style="width: 100px">
        <el-option label="新闻编号" value="deviceID" />
        <el-option label="发布时间" value="post_time" />
        <el-option label="标题" value="title" />
        <el-option label="作者" value="writer_name" />
        <el-option label="优先级" value="importance" />
      </el-select>
    </template>
    <template #append>
      <el-button @click="get_device" :icon="Search"></el-button>
    </template>
  </el-input>
  </el-tooltip>
</el-row>
  <el-table max-height="540" v-model:data="devices" border stripe style="width: 100%;margin-top:20px">
    <el-table-column prop="news_id" label="新闻编号" width="100" />
    <el-table-column prop="post_time" label="发布时间" width="180"  />
    <el-table-column prop="title" label="标题" />
    <el-table-column prop="writer_name" label="作者" />
    <el-table-column prop="importance" label="优先级" />
    <el-table-column fixed="right" label="操作" width="120">
      <template #default="item">
        <el-button @click="delete_device(item.row.deviceID)" link type="primary">伪删除</el-button>
        <el-button @click="pre_edit_device(item.row)" link type="primary">编辑</el-button>
        <el-button @click="" link type="primary">查看</el-button>
      </template>
    </el-table-column>
  </el-table>
  <div style="position: absolute;bottom: 15px ">
    <el-pagination
      background
      @current-change="get_device()"
      @size-change="get_device()"
      :page-sizes="[8, 12, 50, 100]"
      v-model:current-page="page"
      v-model:page-size="pageSize"
      layout="sizes,prev, pager, next,jumper"
      :total="total"/>
    </div>
  <el-dialog v-model="addVis" :center="true" title="添加新闻" style="width: 300px;height: 320px">
     <el-form :model="postDeviceForm" label-width="auto">
        <el-form-item label="设备编码">
              <el-tooltip :visible="notNumber(postDeviceForm.deviceID)&&addVis" content="请输入数字" placement="right-start">
                <el-input v-model="postDeviceForm.deviceID" placeholder="请输入设备编码"></el-input>
              </el-tooltip>
        </el-form-item>
        <el-form-item label="设备名称">
          <el-input v-model="postDeviceForm.deviceName" placeholder="请输入设备名称"></el-input>
        </el-form-item>
       <el-form-item label="生产厂家">
          <el-input v-model="postDeviceForm.Manufacturer" placeholder="请输入生产厂家"></el-input>
       </el-form-item>
        <el-form-item label="生产日期">
              <el-date-picker
                v-model="postDeviceForm.ManufactureDate"
                type="datetime"
                value-format="YYYY-MM-DD HH:mm:ss"
                placeholder="选择日期"
                style="width: 100%"
              ></el-date-picker>
        </el-form-item>
        <el-form-item style="text-align: center;  justify-content: center;">
          <el-button @click="post_device()" style="width: 300px" type="primary">添加</el-button>
        </el-form-item>
     </el-form>
  </el-dialog>
  <el-dialog v-model="editVis" :center="true" title="编辑新闻" style="width: 300px;height: 320px">
     <el-form :model="editDeviceForm" label-width="auto">
        <el-form-item label="设备编码">
          <el-tooltip :visible="notNumber(editDeviceForm.newDeviceID)&&editVis" content="请输入数字" placement="right-start">
            <el-input v-model="editDeviceForm.newDeviceID" placeholder="请输入设备编码"></el-input>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="设备名称">
          <el-input v-model="editDeviceForm.deviceName" placeholder="请输入设备名称"></el-input>
        </el-form-item>
       <el-form-item label="生产厂家">
          <el-input v-model="editDeviceForm.Manufacturer" placeholder="请输入生产厂家"></el-input>
       </el-form-item>
        <el-form-item label="生产日期">
              <el-date-picker
                v-model="editDeviceForm.ManufactureDate"
                type="datetime"
                value-format="YYYY-MM-DD HH:mm:ss"
                placeholder="选择日期"
                style="width: 100%"
              ></el-date-picker>
        </el-form-item>
        <el-form-item style="text-align: center;  justify-content: center;">
          <el-button @click="edit_device()" style="width: 300px" type="primary">更改</el-button>
        </el-form-item>
     </el-form>
  </el-dialog>
</template>

<style scoped>

</style>
