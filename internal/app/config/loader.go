// Package config 用于解析配置文件config.yaml
package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3" // 第三方YAML解析库，用于将YAML数据解析为Go结构体
)

// LoadConfig 加载并解析配置文件
func LoadConfig(configPath string) (*Config, error) {

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 替换环境变量
	data = replaceEnvVariables(data)

	// 解析 YAML 数据到Config结构体
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	// 验证配置
	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("配置验证失败: %w", err)
	}

	return &config, nil
}

// replaceEnvVariables 替换配置中的环境变量（格式: ${VAR_NAME} 或 ${VAR_NAME:-default}）
func replaceEnvVariables(data []byte) []byte {
	// os.Expand：该函数接收一个字符串和一个替换函数，会自动识别字符串中${key}格式的占位符，将key传入替换函数并替换为返回值
	return []byte(os.Expand(string(data), func(key string) string {
		// 处理带默认值的情况: ${VAR:-default}
		if strings.Contains(key, ":-") {
			parts := strings.SplitN(key, ":-", 2)      // 按" :-"分割为[变量名, 默认值]
			if val, ok := os.LookupEnv(parts[0]); ok { // 检查环境变量是否存在
				return val // 存在则返回环境变量值
			}
			return parts[1] // 不存在则返回默认值
		}

		// 普通环境变量: ${VAR}
		return os.Getenv(key) // 直接返回环境变量值（不存在则返回空字符串）
	}))
}

// validateConfig 验证配置的有效性
func validateConfig(config *Config) error {
	// 验证应用配置
	if err := validateAppConfig(&config.App); err != nil {
		return fmt.Errorf("应用配置验证失败: %w", err)
	}

	// 验证数据库配置
	if err := validateDatabaseConfig(&config.Database, config.App.Env); err != nil {
		return fmt.Errorf("数据库配置验证失败: %w", err)
	}

	return nil
}

// validateAppConfig 验证应用配置
func validateAppConfig(appConfig *AppConfig) error {
	// 检查环境配置
	validEnvs := map[string]bool{"development": true, "testing": true, "production": true}
	if !validEnvs[appConfig.Env] {
		return fmt.Errorf("无效的应用环境: '%s'，有效值为 'development', 'testing', 'production'", appConfig.Env)
	}

	// 检查端口配置
	if appConfig.Port <= 0 || appConfig.Port > 65535 {
		return fmt.Errorf("无效的应用端口: %d，端口范围应为 1-65535", appConfig.Port)
	}

	return nil
}

// validateDatabaseConfig 验证数据库配置
func validateDatabaseConfig(dbConfig *DatabaseConfig, appEnv string) error {
	// 检查必要的数据库配置
	if dbConfig.Driver == "" {
		return fmt.Errorf("数据库驱动(driver)不能为空")
	}
	if dbConfig.Host == "" {
		return fmt.Errorf("数据库主机(host)不能为空")
	}
	if dbConfig.Port <= 0 || dbConfig.Port > 65535 {
		return fmt.Errorf("无效的数据库端口: %d，端口范围应为 1-65535", dbConfig.Port)
	}
	if dbConfig.Username == "" {
		return fmt.Errorf("数据库用户名(username)不能为空")
	}
	if dbConfig.DBName == "" {
		return fmt.Errorf("数据库名(dbname)不能为空")
	}

	// 在生产环境中，密码不能为空
	if appEnv == "production" && dbConfig.Password == "" {
		return fmt.Errorf("在生产环境中，数据库密码(password)不能为空")
	}

	// 检查连接池配置
	if dbConfig.MaxOpenConnections > 0 && dbConfig.MaxIdleConnections > dbConfig.MaxOpenConnections {
		return fmt.Errorf("最大空闲连接数(max_idle_connections)不能大于最大连接数(max_open_connections)")
	}

	return nil
}
