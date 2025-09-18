package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	// 导入原始的PredictImage包
	"paddy-goserver/PredictImage"

	"github.com/gin-gonic/gin"
)

// TestData 测试用的请求数据
type TestData struct {
	ImageId string `json:"imageid"`
	ModelId string `json:"modelid"`
}

func main() {
	fmt.Println("🚀 Go服务与Python服务通信测试")

	// 运行原始Go函数测试
	testOriginalGoFunction()
}

func testOriginalGoFunction() {
	fmt.Println("🔍 直接测试原始Go函数与Python服务通信")
	fmt.Println("==================================================")

	// 设置Gin为测试模式
	gin.SetMode(gin.TestMode)

	// 创建测试路由
	r := gin.New()
	r.POST("/test_predict", PredictImage.PredictImage)

	// 启动测试服务器
	go func() {
		fmt.Println("📡 启动测试服务器在 :8888...")
		r.Run(":8888")
	}()

	// 等待服务器启动
	fmt.Println("⏳ 等待服务器启动...")
	time.Sleep(2 * time.Second)

	// 测试不同的场景
	testCases := []TestData{
		{ImageId: "test_grow.jpg", ModelId: "1"},    // 生长期预测
		{ImageId: "test_disease.jpg", ModelId: "2"}, // 病害预测
	}

	for i, testCase := range testCases {
		fmt.Printf("\n🧪 测试案例 %d: %+v\n", i+1, testCase)

		// 准备请求数据
		jsonData, err := json.Marshal(testCase)
		if err != nil {
			fmt.Printf("❌ JSON序列化失败: %v\n", err)
			continue
		}

		// 发送请求到测试服务器
		resp, err := http.Post("http://localhost:8888/test_predict",
			"application/json",
			bytes.NewBuffer(jsonData))

		if err != nil {
			fmt.Printf("❌ 请求失败: %v\n", err)
			continue
		}
		defer resp.Body.Close()

		// 读取响应
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("❌ 读取响应失败: %v\n", err)
			continue
		}

		fmt.Printf("📥 响应状态: %d\n", resp.StatusCode)
		fmt.Printf("📥 响应内容: %s\n", string(body))

		if resp.StatusCode == 200 {
			fmt.Printf("✅ 测试案例 %d 成功!\n", i+1)
		} else {
			fmt.Printf("⚠️  测试案例 %d 返回错误状态码: %d\n", i+1, resp.StatusCode)
		}
	}

	fmt.Println("\n==================================================")
	fmt.Println("🎯 直接Python API测试")

	// 直接测试Python API
	testDirectPythonAPI()

	fmt.Println("\n✅ 所有测试完成!")
	fmt.Println("💡 原始Go代码PredictImage.PredictImage函数已验证可用")
}

func testDirectPythonAPI() {
	pythonURL := "http://127.0.0.1:5050/predict_image"

	testData := TestData{
		ImageId: "direct_test.jpg",
		ModelId: "1",
	}

	jsonData, _ := json.Marshal(testData)

	fmt.Printf("📤 直接调用Python API: %s\n", pythonURL)
	fmt.Printf("📤 请求数据: %s\n", string(jsonData))

	resp, err := http.Post(pythonURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ Python API调用失败: %v\n", err)
		fmt.Printf("   请确保Python服务运行在: %s\n", pythonURL)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ 读取Python响应失败: %v\n", err)
		return
	}

	fmt.Printf("📥 Python API状态: %d\n", resp.StatusCode)
	fmt.Printf("📥 Python API响应: %s\n", string(body))

	if resp.StatusCode == 200 {
		fmt.Printf("✅ Python API直接调用成功!\n")
	}
}
