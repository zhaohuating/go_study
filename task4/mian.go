package main

import (
	"task4/router"
	"task4/structs"
)

type User structs.User

func main() {

	engine := router.SetupRouter()
	engine.Run(":8082")
	//
	//db := connect.GetDB()
	//
	//db.AutoMigrate(&structs.Comment{}, &structs.User{}, &structs.Post{})
}
