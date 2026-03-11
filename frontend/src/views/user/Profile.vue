<template>
  <div class="profile">
    <el-card>
      <template #header>
        <span>个人设置</span>
      </template>

      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px" style="max-width: 500px">
        <el-form-item label="用户 ID">
          <el-input :model-value="userStore.userInfo?.userId" disabled />
        </el-form-item>

        <el-form-item label="用户名">
          <el-input :model-value="userStore.userInfo?.username" disabled />
        </el-form-item>

        <el-form-item label="昵称" prop="nickname">
          <el-input v-model="form.nickname" placeholder="请输入昵称" />
        </el-form-item>

        <el-form-item label="邮箱" prop="email">
          <el-input v-model="form.email" placeholder="请输入邮箱" />
        </el-form-item>

        <el-form-item label="头像 URL" prop="avatar">
          <el-input v-model="form.avatar" placeholder="请输入头像 URL（可选）" />
        </el-form-item>

        <el-form-item label="角色">
          <el-tag :type="userStore.userInfo?.role === 'admin' ? 'danger' : 'info'">
            {{ userStore.userInfo?.role === 'admin' ? '管理员' : '普通用户' }}
          </el-tag>
        </el-form-item>

        <el-form-item label="注册时间">
          <span>{{ formatDate(userStore.userInfo?.createdAt) }}</span>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" :loading="loading" @click="handleUpdate">保存修改</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import { userApi } from '@/api'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const formRef = ref<FormInstance>()
const loading = ref(false)

const form = reactive({
  nickname: '',
  email: '',
  avatar: ''
})

const rules: FormRules = {
  email: [
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ]
}

const formatDate = (timestamp?: number) => {
  if (!timestamp) return '-'
  return new Date(timestamp).toLocaleString('zh-CN')
}

const handleUpdate = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (!valid) return

    loading.value = true
    try {
      await userApi.updateUser({
        nickname: form.nickname || undefined,
        email: form.email || undefined,
        avatar: form.avatar || undefined
      })
      
      // 更新本地用户信息
      if (userStore.userInfo) {
        userStore.userInfo.nickname = form.nickname
        userStore.userInfo.email = form.email
        userStore.userInfo.avatar = form.avatar
      }
      
      ElMessage.success('保存成功')
    } catch (error: any) {
      ElMessage.error(error.response?.data?.message || '保存失败')
    } finally {
      loading.value = false
    }
  })
}

onMounted(() => {
  if (userStore.userInfo) {
    form.nickname = userStore.userInfo.nickname
    form.email = userStore.userInfo.email
    form.avatar = userStore.userInfo.avatar || ''
  }
})
</script>

<style scoped>
.profile {
  height: 100%;
}
</style>
