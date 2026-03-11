#!/bin/bash

# 用户管理系统 - 快速启动脚本

set -e

echo "🚀 用户管理系统 - 快速启动"
echo "================================"

# 检查 MySQL
echo "📊 检查 MySQL 连接..."
if ! command -v mysql &> /dev/null; then
    echo "⚠️  MySQL 客户端未安装，请手动初始化数据库"
else
    read -p "MySQL root 密码：" -s MYSQL_PASSWORD
    echo
    echo "创建数据库..."
    mysql -u root -p"$MYSQL_PASSWORD" -e "CREATE DATABASE IF NOT EXISTS user_management DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
    echo "✅ 数据库创建完成"
fi

# 启动后端
echo ""
echo "🔧 启动后端服务..."
cd backend

if [ ! -f "go.mod" ]; then
    echo "❌ 后端目录错误"
    exit 1
fi

echo "下载 Go 依赖..."
go mod tidy

echo "编译后端..."
go build -o user-api main.go

echo "启动后端服务（端口 8888）..."
./user-api -f config/user.yaml &
BACKEND_PID=$!
echo "✅ 后端服务已启动 (PID: $BACKEND_PID)"

cd ..

# 启动前端
echo ""
echo "🎨 启动前端服务..."
cd frontend

if [ ! -f "package.json" ]; then
    echo "❌ 前端目录错误"
    exit 1
fi

if [ ! -d "node_modules" ]; then
    echo "安装前端依赖..."
    npm install
fi

echo "启动前端开发服务器（端口 3000）..."
npm run dev &
FRONTEND_PID=$!
echo "✅ 前端服务已启动 (PID: $FRONTEND_PID)"

# 完成
echo ""
echo "================================"
echo "✅ 服务启动完成！"
echo ""
echo "📱 前端访问：http://localhost:3000"
echo "🔌 后端 API: http://localhost:8888"
echo ""
echo "按 Ctrl+C 停止所有服务"
echo ""

# 等待中断信号
trap "kill $BACKEND_PID $FRONTEND_PID 2>/dev/null; echo '服务已停止'; exit" INT

wait
