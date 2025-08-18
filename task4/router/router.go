package router

import (
	"task4/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	//公开路由 不需要验证用户的
	public := router.Group("/api")
	{
		public.POST("/register", controller.Register)
		public.POST("/login", controller.Login)
	}

	//需要验证用户的
	private := router.Group("/api/user")
	{
		private.POST("/register", controller.Register)
	}

	return router
}
