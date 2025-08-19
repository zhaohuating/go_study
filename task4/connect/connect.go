package connect

import (
	"task4/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	MYSQL  = "mysql"
	SQLITE = "sqlite"
)

var db *gorm.DB
var err error

func init() {
	var dialector gorm.Dialector
	var driver string = config.Cfg.Database.Driver
	var dsn string = config.Cfg.Database.Dsn
	var maxOpenConns int = config.Cfg.Database.MaxOpenConns
	var maxIdleConns int = config.Cfg.Database.MaxIdleConns
	var connMaxLifetime int = config.Cfg.Database.ConnMaxLifetime
	if driver == MYSQL {
		dialector = mysql.Open(dsn)
	} else if driver == SQLITE {
		dialector = sqlite.Open(dsn)
	} else {
		dialector = mysql.Open(dsn)
	}

	levelMap := map[string]logger.LogLevel{
		"silent": logger.Silent,
		"error":  logger.Error,
		"warn":   logger.Warn,
		"info":   logger.Info,
		"debug":  logger.Info, // debug 映射到 info（输出详细日志）
	}

	level := config.Cfg.Log.Level
	logLevel, ok := levelMap[level]

	if !ok {
		logLevel = logger.Info
	}

	db, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		panic("连接数据库失败")
	}

	sqlDB, _ := db.DB()
	// 配置连接池参数
	// 1. 最大打开连接数（默认值为0，表示无限制）
	sqlDB.SetMaxOpenConns(maxOpenConns) // 根据服务器性能和数据库承载能力调整

	// 2. 最大空闲连接数（默认值为2）
	sqlDB.SetMaxIdleConns(maxIdleConns) // 建议设置为与 MaxOpenConns 接近，避免频繁创建连接

	// 3. 连接的最大存活时间（超过此时间的空闲连接会被关闭）
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second) //
}

func GetDB() *gorm.DB {
	return db
}
