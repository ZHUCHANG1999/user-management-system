package types

// 用户注册请求
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Nickname string `json:"nickname,omitempty"`
}

// 用户注册响应
type RegisterResponse struct {
	UserId  int64  `json:"userId"`
	Message string `json:"message"`
}

// 用户登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 用户登录响应
type LoginResponse struct {
	UserId int64  `json:"userId"`
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

// 获取用户信息请求
type GetUserRequest struct {
	UserId int64 `json:"userId,optional"`
}

// 用户信息
type UserInfo struct {
	UserId    int64  `json:"userId"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Role      string `json:"role"`
	CreatedAt int64  `json:"createdAt"`
}

// 获取用户信息响应
type GetUserResponse struct {
	User UserInfo `json:"user"`
}

// 更新用户信息请求
type UpdateUserRequest struct {
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}

// 用户列表请求
type UserListRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

// 用户列表响应
type UserListResponse struct {
	Total int64      `json:"total"`
	Users []UserInfo `json:"users"`
}

// 删除用户请求
type DeleteUserRequest struct {
	UserId int64 `json:"userId"`
}

// 通用响应
type CommonResponse struct {
	Message string `json:"message"`
}
