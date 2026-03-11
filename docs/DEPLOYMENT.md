# 部署指南

## 环境准备

### 1. MySQL 数据库

```bash
# 安装 MySQL 8.0
# Ubuntu/Debian
sudo apt-get install mysql-server-8.0

# CentOS/RHEL
sudo yum install mysql-server

# 或使用 Docker
docker run -d --name mysql \
  -e MYSQL_ROOT_PASSWORD=your_password \
  -p 3306:3306 \
  mysql:8.0
```

### 2. 创建数据库

```bash
mysql -u root -p
```

```sql
CREATE DATABASE IF NOT EXISTS `user_management` 
DEFAULT CHARACTER SET utf8mb4 
COLLATE utf8mb4_unicode_ci;
```

### 3. Go 环境

```bash
# 安装 Go 1.21+
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### 4. Node.js 环境

```bash
# 使用 nvm 安装 Node.js 18+
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
nvm install 18
nvm use 18
```

---

## 后端部署

### 1. 配置数据库连接

编辑 `backend/config/user.yaml`:

```yaml
Name: user-api
Host: 0.0.0.0
Port: 8888

Database:
  DataSource: root:your_password@tcp(localhost:3306)/user_management?charset=utf8mb4&parseTime=true&loc=Local

JWT:
  Secret: your-secret-key-change-in-production
  Expire: 86400
```

### 2. 编译运行

```bash
cd backend

# 下载依赖
go mod tidy

# 编译
go build -o user-api main.go

# 运行
./user-api -f config/user.yaml
```

### 3. 验证服务

```bash
curl http://localhost:8888/api/v1/users
```

---

## 前端部署

### 1. 安装依赖

```bash
cd frontend
npm install
```

### 2. 开发环境

```bash
npm run dev
```

访问 http://localhost:3000

### 3. 生产环境构建

```bash
# 构建
npm run build

# 预览
npm run preview
```

构建产物在 `dist/` 目录，可部署到 Nginx 或其他 Web 服务器。

---

## Nginx 配置示例

```nginx
server {
    listen 80;
    server_name your-domain.com;

    # 前端静态文件
    location / {
        root /path/to/user-management-system/frontend/dist;
        try_files $uri $uri/ /index.html;
    }

    # 后端 API 代理
    location /api/ {
        proxy_pass http://localhost:8888;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

---

## Docker 部署（可选）

### 1. 创建 Dockerfile

**backend/Dockerfile:**
```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o user-api main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/user-api .
COPY --from=builder /app/config ./config
EXPOSE 8888
CMD ["./user-api", "-f", "config/user.yaml"]
```

**frontend/Dockerfile:**
```dockerfile
FROM node:18-alpine AS builder

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

### 2. Docker Compose

创建 `docker-compose.yml`:

```yaml
version: '3.8'

services:
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: your_password
      MYSQL_DATABASE: user_management
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql
      - ./backend/sql/init.sql:/docker-entrypoint-initdb.d/init.sql

  backend:
    build: ./backend
    ports:
      - "8888:8888"
    depends_on:
      - mysql
    environment:
      - DB_HOST=mysql
      - DB_PASSWORD=your_password

  frontend:
    build: ./frontend
    ports:
      - "3000:80"
    depends_on:
      - backend

volumes:
  mysql_data:
```

运行：
```bash
docker-compose up -d
```

---

## 常见问题

### 1. 数据库连接失败

检查 `user.yaml` 中的数据库配置：
- 用户名密码是否正确
- MySQL 服务是否运行
- 防火墙是否开放 3306 端口

### 2. 端口被占用

修改配置文件中的端口：
- 后端：修改 `user.yaml` 中的 `Port`
- 前端：修改 `vite.config.js` 中的 `port`

### 3. 跨域问题

开发环境已配置代理，生产环境使用 Nginx 反向代理解决。

---

## 监控与日志

### 日志查看

go-zero 默认输出到 stdout，可使用以下方式查看：

```bash
# 直接运行
./user-api -f config/user.yaml 2>&1 | tee app.log

# 使用 journalctl (systemd)
journalctl -u user-api -f
```

### 健康检查

```bash
curl http://localhost:8888/api/v1/users
```

返回 200 表示服务正常。
