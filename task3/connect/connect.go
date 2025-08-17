package connect

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var connectType string = "sqlx"
var db *gorm.DB
var sqlxDB *sqlx.DB
var err error

func init() {
	if connectType == "gorm" {
		db, err = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else if connectType == "sqlx" {
		sqlxDB, err = sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	}

	if err != nil {
		panic("连接数据库失败")
	}
}

func GetSqlxDB() *sqlx.DB {
	return sqlxDB
}

func GetDB() *gorm.DB {
	return db
}
