package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT密钥（实际项目中建议从环境变量或配置文件读取）
var jwtSecret = []byte("your_secret_key")

// 自定义JWT声明结构体
type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// JWTAuthMiddleware 是一个Gin中间件，用于验证JWT令牌
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取Authorization字段
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供Authorization头"})
			c.Abort() // 终止请求处理链
			return
		}

		// 检查Authorization格式是否为"Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization格式错误"})
			c.Abort()
			return
		}

		// 解析JWT令牌
		tokenString := parts[1]
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		// 验证令牌有效性
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的令牌"})
			c.Abort()
			return
		}

		// 将用户信息存入上下文，供后续处理函数使用
		c.Set("userID", claims.ID)
		c.Set("username", claims.Username)

		// 继续处理请求
		c.Next()
	}
}
