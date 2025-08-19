package main

import (
	"task4/config"
	"task4/router"
	"task4/structs"
)

type User structs.User

func main() {

	//if err := config.Init("./config/config.yaml"); err != nil {
	//	log.Fatalf("配置初始化失败: %v", err)
	//}
	//connect.Init(config.Cfg.Database)
	engine := router.SetupRouter()
	engine.Run(config.Cfg.Server.Port)
	//
	//db := connect.GetDB()
	//
	//db.AutoMigrate(&structs.Comment{}, &structs.User{}, &structs.Post{})
}
