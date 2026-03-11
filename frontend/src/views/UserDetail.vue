<template>
  <div class="user-detail">
    <el-card>
      <template #header>
        <span>用户详情</span>
      </template>

      <el-descriptions :column="1" border v-loading="loading">
        <el-descriptions-item label="用户 ID">{{ user.user_id }}</el-descriptions-item>
        <el-descriptions-item label="用户名">{{ user.username }}</el-descriptions-item>
        <el-descriptions-item label="邮箱">{{ user.email || '-' }}</el-descriptions-item>
        <el-descriptions-item label="昵称">{{ user.nickname || '-' }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="user.status === 1 ? 'success' : 'danger'">
            {{ user.status === 1 ? '正常' : '禁用' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ user.created_at }}</el-descriptions-item>
      </el-descriptions>

      <div style="margin-top: 20px">
        <el-button type="primary" @click="handleEdit">编辑</el-button>
        <el-button @click="handleBack">返回</el-button>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getUserDetail } from '@/api/user'
import { ElMessage } from 'element-plus'

const route = useRoute()
const router = useRouter()
const loading = ref(false)
const user = reactive({
  user_id: '',
  username: '',
  email: '',
  nickname: '',
  status: 1,
  created_at: ''
})

const loadUser = async () => {
  loading.value = true
  try {
    const res = await getUserDetail(route.params.id)
    Object.assign(user, res)
  } catch (error) {
    ElMessage.error('加载用户详情失败')
  } finally {
    loading.value = false
  }
}

const handleEdit = () => {
  router.push(`/users/${user.user_id}/edit`)
}

const handleBack = () => {
  router.back()
}

onMounted(() => {
  loadUser()
})
</script>

<style scoped>
.user-detail {
  padding: 20px;
}
</style>
