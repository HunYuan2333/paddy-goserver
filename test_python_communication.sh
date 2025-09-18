#!/bin/bash

echo "🚀 开始测试Go服务与Python服务的通信..."
echo ""

# 进入Go项目目录
cd /Users/2333hunyuan/Documents/paddy-goserver

# 确保依赖已安装
echo "📦 检查Go模块依赖..."
go mod tidy

# 运行测试
echo ""
echo "🧪 运行Python服务通信测试..."
echo ""

# 运行单独的测试工具
go run test/python_service_test.go

echo ""
echo "✅ 测试完成!"
echo ""
echo "💡 使用说明:"
echo "   1. 确保Python服务正在运行: cd /Users/2333hunyuan/PycharmProjects/paddy-sever && python main_new.py"
echo "   2. 如果测试失败，检查Python服务是否在 http://127.0.0.1:5050 运行"
echo "   3. 原始Go代码可以直接使用，无需修改"
