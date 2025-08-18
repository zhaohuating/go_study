package router

import (
	"task4/controller"
	"task4/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 错误处理
	router.Use(middleware.ErrorHandler())
	//公开路由 不需要验证用户的
	public := router.Group("/api")
	{
		public.POST("/register", controller.Register)
		public.POST("/login", controller.Login)
	}

	//需要验证用户的
	private := router.Group("/api")
	private.Use(middleware.JWTAuthMiddleware())
	{
		private.POST("/post", controller.AddPost)
		private.GET("/post", controller.GetPostList)
	}

	return router
}
