import api from './request'

// 获取用户列表
export function getUserList(params) {
  return api.get('/users', { params })
}

// 获取用户详情
export function getUserDetail(userId) {
  return api.get(`/users/${userId}`)
}

// 创建用户
export function createUser(data) {
  return api.post('/users', data)
}

// 更新用户
export function updateUser(userId, data) {
  return api.put(`/users/${userId}`, {
    email: data.email,
    nickname: data.nickname,
    status: data.status
  })
}

// 删除用户
export function deleteUser(userId) {
  return api.delete(`/users/${userId}`)
}
