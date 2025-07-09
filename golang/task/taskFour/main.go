package main

import (
	"github.com/spf13/viper"
	"golang-homework/task/taskFour/app"
	"golang-homework/task/taskFour/db"
	"golang-homework/task/taskFour/router"
	"log"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

func main() {
	// 解析配置文件
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	// 初始化数据库
	database, err := db.InitDB(config.Database.DSN)
	if err != nil {
		log.Fatalf("Unable to connect to database, %v", config.Database.DSN)
	}
	// 初始化应用
	application := app.Initialize(database)
	// 设置路由
	r := router.SetupRouter(application)
	// 启动服务器
	if err := r.Run("0.0.0.0:" + config.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
