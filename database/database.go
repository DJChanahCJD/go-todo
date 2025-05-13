package database

import (
	"fmt"
	"log"

	"github.com/djchanahcjd/go-todo/config"
	"github.com/djchanahcjd/go-todo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB(config *config.Config) error {
	// 构建MySQL DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySQL.Username,
		config.MySQL.Password,
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.Database,
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(100)

	// 自动迁移数据库表
	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Printf("Failed to auto migrate database: %v", err)
		return err
	}

	DB = db
	return nil
}
