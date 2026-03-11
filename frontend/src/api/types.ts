export interface RegisterRequest {
  username: string
  password: string
  email: string
  nickname?: string
}

export interface LoginRequest {
  username: string
  password: string
}

export interface LoginResponse {
  userId: number
  token: string
  expire: number
}

export interface UserInfo {
  userId: number
  username: string
  email: string
  nickname: string
  avatar?: string
  role: string
  createdAt: number
}

export interface GetUserResponse {
  user: UserInfo
}

export interface UpdateUserRequest {
  nickname?: string
  email?: string
  avatar?: string
}

export interface UserListRequest {
  page: number
  pageSize: number
}

export interface UserListResponse {
  total: number
  users: UserInfo[]
}

export interface CommonResponse {
  message: string
}
