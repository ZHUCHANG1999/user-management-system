<template>
  <div class="layout">
    <el-header class="header">
      <div class="header-content">
        <h1 class="logo">用户管理系统</h1>
        <div class="user-info">
          <el-dropdown @command="handleCommand">
            <span class="user-name">
              <el-avatar :size="32" :icon="User" />
              {{ user?.nickname || user?.username }}
              <el-icon class="el-icon--right"><arrow-down /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="profile">个人信息</el-dropdown-item>
                <el-dropdown-item command="logout" divided>退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </div>
    </el-header>
    <el-main class="main">
      <router-view />
    </el-main>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { User, ArrowDown } from '@element-plus/icons-vue'
import { logout } from '@/api/auth'
import { ElMessage, ElMessageBox } from 'element-plus'

const router = useRouter()
const user = ref(null)

onMounted(() => {
  const userData = localStorage.getItem('user')
  if (userData) {
    user.value = JSON.parse(userData)
  }
})

const handleCommand = async (command) => {
  if (command === 'logout') {
    await handleLogout()
  } else if (command === 'profile') {
    router.push(`/users/${user.value.user_id}`)
  }
}

const handleLogout = async () => {
  try {
    await ElMessageBox.confirm('确定要退出登录吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    await logout()
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    ElMessage.success('已退出登录')
    router.push('/login')
  } catch (error) {
    if (error !== 'cancel') {
      console.error('Logout error:', error)
    }
  }
}
</script>

<style scoped>
.layout {
  min-height: 100vh;
  background: #f5f7fa;
}

.header {
  background: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  padding: 0 20px;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100%;
  max-width: 1400px;
  margin: 0 auto;
}

.logo {
  font-size: 20px;
  font-weight: 600;
  color: #409eff;
  margin: 0;
}

.user-info {
  display: flex;
  align-items: center;
}

.user-name {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  color: #606266;
}

.user-name:hover {
  color: #409eff;
}

.main {
  padding: 20px;
  max-width: 1400px;
  margin: 0 auto;
}
</style>
