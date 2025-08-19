package router

import (
	"task4/controller"
	"task4/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Logger())
	router.Use(gin.Recovery())
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
		private.GET("/post/:id", controller.GetPostContent)
		private.PUT("/post/:id", controller.UpdatePost)
		private.DELETE("/post/:id", controller.DeletePost)
		private.POST("/comment", controller.AddComment)
		private.GET("/comment/:postID", controller.GetPostComments)
	}

	return router
}
