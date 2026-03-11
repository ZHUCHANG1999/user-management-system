import api from './request'

// 登录
export function login(data) {
  return api.post('/auth/login', data)
}

// 注册
export function register(data) {
  return api.post('/auth/register', data)
}

// 登出
export function logout() {
  return api.post('/auth/logout', {
    token: localStorage.getItem('token')
  })
}

// 刷新 Token
export function refreshToken(token) {
  return api.post('/auth/refresh', { token })
}
