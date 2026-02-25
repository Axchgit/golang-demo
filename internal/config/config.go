package config

import (
	"fmt"
	"os"
)

// Config 应用配置
type Config struct {
	AppName    string
	Version    string
	DebugMode  bool
	Port       int
	GoRoutines int
}

// 默认配置
var defaultConfig = Config{
	AppName:    "golang-learn",
	Version:    "1.0.0",
	DebugMode:  true,
	Port:       8080,
	GoRoutines: 10,
}

// Load 加载配置
func Load() *Config {
	// 这里可以后续从文件/环境变量读取
	// 目前返回默认值
	cfg := defaultConfig
	
	// 可以通过环境变量覆盖
	if v := os.Getenv("APP_NAME"); v != "" {
		cfg.AppName = v
	}
	if v := os.Getenv("PORT"); v != "" {
		fmt.Sscanf(v, "%d", &cfg.Port)
	}
	
	return &cfg
}
