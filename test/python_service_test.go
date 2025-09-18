package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// PredictRequest 预测请求结构体
type PredictRequest struct {
	ImageId string `json:"imageid"`
	ModelId string `json:"modelid"`
}

// PredictResponse 预测响应结构体
type PredictResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// TestPythonService 测试Python预测服务
func TestPythonService() {
	fmt.Println("=== Python预测服务通信测试 ===")

	// 基础URL
	baseURL := "http://127.0.0.1:5050"

	// 测试健康检查
	fmt.Println("\n1. 测试健康检查...")
	testHealthCheck(baseURL)

	// 测试模型信息
	fmt.Println("\n2. 测试模型信息...")
	testModelInfo(baseURL)

	// 测试生长期预测
	fmt.Println("\n3. 测试生长期预测...")
	testGrowPrediction(baseURL)

	// 测试病害预测
	fmt.Println("\n4. 测试病害预测...")
	testDiseasePrediction(baseURL)

	// 测试上传端点
	fmt.Println("\n5. 测试上传端点...")
	testUploadEndpoints(baseURL)
}

// testHealthCheck 测试健康检查
func testHealthCheck(baseURL string) {
	url := baseURL + "/health"
	resp, err := httpGet(url)
	if err != nil {
		fmt.Printf("❌ 健康检查失败: %v\n", err)
		return
	}

	fmt.Printf("✅ 健康检查成功: %s\n", resp)
}

// testModelInfo 测试模型信息
func testModelInfo(baseURL string) {
	url := baseURL + "/model_info"
	resp, err := httpGet(url)
	if err != nil {
		fmt.Printf("❌ 模型信息获取失败: %v\n", err)
		return
	}

	fmt.Printf("✅ 模型信息获取成功: %s\n", resp)
}

// testGrowPrediction 测试生长期预测
func testGrowPrediction(baseURL string) {
	url := baseURL + "/predict_image"

	// 创建测试请求
	request := PredictRequest{
		ImageId: "test_grow.jpg",
		ModelId: "1",
	}

	resp, err := httpPostJSON(url, request)
	if err != nil {
		fmt.Printf("❌ 生长期预测测试失败: %v\n", err)
		return
	}

	// 解析响应
	var predictResp PredictResponse
	if err := json.Unmarshal([]byte(resp), &predictResp); err != nil {
		fmt.Printf("❌ 生长期预测响应解析失败: %v\n", err)
		fmt.Printf("原始响应: %s\n", resp)
		return
	}

	fmt.Printf("✅ 生长期预测测试完成\n")
	fmt.Printf("   响应码: %s\n", predictResp.Code)
	fmt.Printf("   消息: %s\n", predictResp.Message)
	if predictResp.Code == "200" {
		fmt.Printf("   🎉 预测成功!\n")
	} else {
		fmt.Printf("   ⚠️  预测返回错误码: %s\n", predictResp.Code)
	}
}

// testDiseasePrediction 测试病害预测
func testDiseasePrediction(baseURL string) {
	url := baseURL + "/predict_image"

	// 创建测试请求
	request := PredictRequest{
		ImageId: "test_disease.jpg",
		ModelId: "2",
	}

	resp, err := httpPostJSON(url, request)
	if err != nil {
		fmt.Printf("❌ 病害预测测试失败: %v\n", err)
		return
	}

	// 解析响应
	var predictResp PredictResponse
	if err := json.Unmarshal([]byte(resp), &predictResp); err != nil {
		fmt.Printf("❌ 病害预测响应解析失败: %v\n", err)
		fmt.Printf("原始响应: %s\n", resp)
		return
	}

	fmt.Printf("✅ 病害预测测试完成\n")
	fmt.Printf("   响应码: %s\n", predictResp.Code)
	fmt.Printf("   消息: %s\n", predictResp.Message)
	if predictResp.Code == "200" {
		fmt.Printf("   🎉 预测成功!\n")
	} else {
		fmt.Printf("   ⚠️  预测返回错误码: %s\n", predictResp.Code)
	}
}

// testUploadEndpoints 测试上传端点
func testUploadEndpoints(baseURL string) {
	endpoints := []string{
		"/upload_grow_image",
		"/upload_disease_image",
	}

	for _, endpoint := range endpoints {
		url := baseURL + endpoint
		resp, err := httpGet(url)
		if err != nil {
			fmt.Printf("❌ %s 端点测试失败: %v\n", endpoint, err)
		} else {
			fmt.Printf("✅ %s 端点可达\n", endpoint)
			fmt.Printf("   响应: %s\n", resp)
		}
	}
}

// httpGet 发送GET请求
func httpGet(url string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	return string(body), nil
}

// httpPostJSON 发送JSON POST请求
func httpPostJSON(url string, data interface{}) (string, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("JSON序列化失败: %v", err)
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	return string(body), nil
}

// TestOriginalGoFunction 测试原始Go函数的完整流程
func TestOriginalGoFunction() {
	fmt.Println("\n=== 原始Go函数完整流程测试 ===")

	// 模拟原始Go服务中PredictImage函数的调用
	testData := PredictRequest{
		ImageId: "test_image.jpg",
		ModelId: "1",
	}

	jsonData, err := json.Marshal(testData)
	if err != nil {
		fmt.Printf("❌ JSON序列化失败: %v\n", err)
		return
	}

	fmt.Printf("📤 发送请求数据: %s\n", string(jsonData))

	// 直接调用Python API（模拟原始Go代码的行为）
	pythonURL := "http://127.0.0.1:5050/predict_image"

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Post(pythonURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ 与Python API通信失败: %v\n", err)
		fmt.Printf("   请确保Python服务正在运行在 http://127.0.0.1:5050\n")
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ 读取Python API响应失败: %v\n", err)
		return
	}

	fmt.Printf("📥 Python API响应: %s\n", string(responseBody))
	fmt.Printf("   状态码: %d\n", resp.StatusCode)

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("✅ 原始Go函数流程测试成功!\n")
		fmt.Printf("   🎉 Go服务可以正常与Python服务通信\n")
	} else {
		fmt.Printf("⚠️  响应状态码不是200: %d\n", resp.StatusCode)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("🚀 开始测试Go与Python服务通信...")

	// 基础服务测试
	TestPythonService()

	// 原始函数流程测试
	TestOriginalGoFunction()

	fmt.Println("\n=== 测试完成 ===")
	fmt.Println("💡 提示:")
	fmt.Println("   - 如果看到错误，请确保Python服务正在运行")
	fmt.Println("   - 启动Python服务: cd /Users/2333hunyuan/PycharmProjects/paddy-sever && python main_new.py")
	fmt.Println("   - Python服务地址: http://127.0.0.1:5050")
}
