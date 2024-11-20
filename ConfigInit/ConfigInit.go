package ConfigInit

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config 结构体用于解析 YAML 配置文件
type Config struct {
	TokenKey      string `yaml:"token-key"`
	DriverName    string `yaml:"sql-driver-name"`
	DriverCommand string `yaml:"sql-driver-command"`
}

// ReadConfigFile 读取根目录下的 config.yaml 文件并解析配置
func ReadConfigFile() (*Config, error) {
	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// 构建配置文件路径
	configFilePath := currentDir + "/config.yaml"

	// 读取配置文件内容
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	// 解析 YAML 配置文件
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	// 返回配置
	return &config, nil
}
