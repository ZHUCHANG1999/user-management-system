import axios from 'axios'
import type {
  RegisterRequest,
  RegisterResponse,
  LoginRequest,
  LoginResponse,
  GetUserResponse,
  UpdateUserRequest,
  UserListRequest,
  UserListResponse,
  CommonResponse
} from './types'

const api = axios.create({
  baseURL: '/api',
  timeout: 10000
})

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export const userApi = {
  // 用户注册
  register(data: RegisterRequest): Promise<RegisterResponse> {
    return api.post('/user/register', data)
  },

  // 用户登录
  login(data: LoginRequest): Promise<LoginResponse> {
    return api.post('/user/login', data)
  },

  // 获取用户信息
  getUserInfo(): Promise<GetUserResponse> {
    return api.get('/user/info')
  },

  // 更新用户信息
  updateUser(data: UpdateUserRequest): Promise<CommonResponse> {
    return api.post('/user/update', data)
  },

  // 获取用户列表
  getUserList(params: UserListRequest): Promise<UserListResponse> {
    return api.get('/user/list', { params })
  },

  // 删除用户
  deleteUser(userId: number): Promise<CommonResponse> {
    return api.post('/user/delete', { userId })
  }
}
