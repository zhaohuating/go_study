package middleware

import (
	"log"
	"net/http"
	"task4/errors"

	"github.com/gin-gonic/gin"
)

// ErrorHandler 全局错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 先执行后续处理函数
		c.Next()

		// 检查是否有错误发生
		if len(c.Errors) > 0 {
			// 获取最后一个错误
			err := c.Errors.Last()

			// 判断是否为自定义AppError
			if appErr, ok := err.Err.(*errors.AppError); ok {
				// 记录原始错误（用于后端排查问题）
				if appErr.Err != nil {
					log.Printf("错误: %v", appErr.Err)
				}
				// 返回标准化响应
				c.JSON(appErr.Code, gin.H{
					"code":    appErr.Code,
					"message": appErr.Message,
				})
				return
			}

			// 非自定义错误（如系统错误），统一返回500
			log.Printf("未处理的错误: %v", err.Err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "服务器内部错误",
			})
		}
	}
}
