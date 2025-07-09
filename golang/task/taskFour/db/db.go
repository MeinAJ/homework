package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info), // 生产环境可改为 logger.Error
		PrepareStmt: true,                                // 开启预编译语句缓存
	})

	if err != nil {
		return nil, err
	}

	// 获取通用数据库对象 sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 连接池配置
	sqlDB.SetMaxIdleConns(10)           // 空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          // 打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可复用的最大时间

	return db, nil
}
