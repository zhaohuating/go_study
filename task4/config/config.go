package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// 全局配置变量
var Cfg *Config

// 配置结构体，与yaml文件对应
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	JWT      JWTConfig      `yaml:"jwt"`
	Log      LogConfig      `yaml:"log"`
}

// 服务器配置
type ServerConfig struct {
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

// 数据库配置
type DatabaseConfig struct {
	Driver          string `yaml:"driver"`
	Dsn             string `yaml:"dsn"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

// JWT配置
type JWTConfig struct {
	Secret      string `yaml:"secret"`
	ExpireHours int    `yaml:"expire_hours"`
}

// 日志配置
type LogConfig struct {
	Level      string `yaml:"level"`
	Encoding   string `yaml:"encoding"`
	Mode       string `yaml:"mode"`
	OutputPath string `yaml:"output_path"` // 日志路径
}

// 初始化配置
func init() {
	// 读取配置文件
	data, _ := os.ReadFile("./config/config.yaml")

	// 解析yaml到结构体
	Cfg = &Config{}
	yaml.Unmarshal(data, Cfg)
}
