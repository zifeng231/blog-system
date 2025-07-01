package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	// 创建数据库连接
	// 创建自定义日志器
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到控制台
		logger.Config{
			SlowThreshold:             time.Second, // 慢查询阈值（超过 1 秒）
			LogLevel:                  logger.Info, // 日志级别：Silent, Error, Warn, Info
			IgnoreRecordNotFoundError: true,        // 忽略记录不存在的错误
			ParameterizedQueries:      false,       // 打印完整 SQL（而非占位符）
			Colorful:                  true,        // 启用彩色日志
		},
	)
	database, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect database")
	}
	DB = database

	return database
}
