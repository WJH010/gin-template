// Config 定义配置结构
package config

import (
	"time"
)

// Config 主配置结构
type Config struct {
	App      AppConfig      `yaml:"app"` // yaml 标签:用于告诉解析器在解析 YAML 文件时，如何将 YAML 文件中的键名映射到 Go 结构体的字段名。
	Database DatabaseConfig `yaml:"database"`
}

// AppConfig 应用配置
type AppConfig struct {
	Name           string   `yaml:"name"`
	Env            string   `yaml:"env"`
	Port           int      `yaml:"port"`
	Debug          bool     `yaml:"debug"`
	AllowedOrigins []string `yaml:"cors.allowed_origins"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver                string        `yaml:"driver"`
	Host                  string        `yaml:"host"`
	Port                  int           `yaml:"port"`
	Username              string        `yaml:"username"`
	Password              string        `yaml:"password"`
	DBName                string        `yaml:"dbname"`
	MaxOpenConnections    int           `yaml:"max_open_connections"`
	MaxIdleConnections    int           `yaml:"max_idle_connections"`
	ConnectionMaxLifetime time.Duration `yaml:"connection_max_lifetime"`
}
