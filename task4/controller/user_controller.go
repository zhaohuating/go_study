package controller

import (
	"net/http"
	"strconv"
	"task4/connect"
	"task4/errors"
	"task4/pagination"
	"task4/structs"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User structs.User
type Post structs.Post

var db = connect.GetDB()

func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser User
	if err := db.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// 剩下的逻辑...
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
		"user": gin.H{
			"id":       storedUser.ID,
			"username": storedUser.Username,
			// 可以添加其他非敏感用户信息
		},
	})
}

// 新建文章
func AddPost(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UserID, _ := c.Get("userID")
	post.UserID = UserID.(uint)

	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
	}
}

func GetPostList(c *gin.Context) {
	var post []Post
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	var pageParam = pagination.Param{
		Page:     page,
		PageSize: pageSize,
	}
	UserID, _ := c.Get("userID")
	tx := db.Model(post).Where("user_id = ?", UserID)
	paginate, err := pagination.Paginate(tx, pageParam, &post)

	if err != nil {
		serverErrorCode := http.StatusInternalServerError
		c.AbortWithError(serverErrorCode, errors.NewError(serverErrorCode, "查询失败", err))
		return
	}

	c.JSON(http.StatusOK, paginate)
}
