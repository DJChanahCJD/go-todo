package main

import (
	"fmt"
	"log"
	"os"

	"github.com/djchanahcjd/go-todo/config"
	"github.com/djchanahcjd/go-todo/database"
	"github.com/djchanahcjd/go-todo/routes"
	"gopkg.in/yaml.v3"
)

func main() {
	// 读取配置文件
	configFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 解析配置文件
	var cfg config.Config
	if err := yaml.Unmarshal(configFile, &cfg); err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	// 初始化数据库连接
	if err := database.InitDB(&cfg); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// 设置路由
	router := routes.SetupRouter()

	// 启动服务器
	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Server starting on%s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
