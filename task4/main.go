package main

import (
	"task4/config"
	"task4/router"
	"task4/structs"
)

type User structs.User

func main() {
	// 初始化数据库
	//db := connect.GetDB()
	//db.AutoMigrate(&structs.Comment{}, &structs.User{}, &structs.Post{})
	engine := router.SetupRouter()
	engine.Run(config.Cfg.Server.Port)

}
