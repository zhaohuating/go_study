package controller

import (
	errors2 "errors"
	"net/http"
	"strconv"
	"task4/errors"
	"task4/pagination"
	"task4/structs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User structs.User
type Post structs.Post

// 新建文章
func AddPost(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	UserID, _ := c.Get("userID")
	post.UserID = UserID.(uint)

	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create post"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})

}

// 获取博客列表(不包含内容)
func GetPostList(c *gin.Context) {
	var post []Post
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	var pageParam = pagination.Param{
		Page:     page,
		PageSize: pageSize,
	}
	UserID, _ := c.Get("userID")
	tx := db.Model(post).Select("id, created_at, updated_at, title").Where("user_id = ?", UserID)
	paginate, err := pagination.Paginate(tx, pageParam, &post)

	if err != nil {
		serverErrorCode := http.StatusInternalServerError
		c.AbortWithError(serverErrorCode, errors.NewError(serverErrorCode, "查询失败", err))
		return
	}

	c.JSON(http.StatusOK, paginate)
}

// 单个文章的详细信息。
func GetPostContent(c *gin.Context) {
	post := Post{}
	id := c.Param("id")
	UserID, _ := c.Get("userID")
	db.Where("user_id = ? and id = ?", UserID, id).Take(&post)

	c.JSON(http.StatusOK, post)
}

// 实现文章的更新功能，只有文章的作者才能更新自己的文章。
func UpdatePost(c *gin.Context) {
	post := Post{}
	if err := c.ShouldBindBodyWithJSON(&post); err != nil {
		serverErrorCode := http.StatusInternalServerError
		c.AbortWithError(serverErrorCode, errors.NewError(serverErrorCode, "数据转换失败", err))
		return
	}

	val, exit := c.Get("userID")
	if !exit {
		serverErrorCode := http.StatusForbidden
		c.AbortWithError(serverErrorCode, errors.NewError(serverErrorCode, "token过期，请重新登录", errors2.New("token过期，请重新登录")))
		return
	}

	db.Where("id = ? and user_id=? ", c.Param("id"), val.(uint)).Updates(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

// 实现文章的删除功能，只有文章的作者才能删除自己的文章。
func DeletePost(c *gin.Context) {
	post := Post{}
	val, exit := c.Get("userID")
	if !exit {
		serverErrorCode := http.StatusForbidden
		c.AbortWithError(serverErrorCode, errors.NewError(serverErrorCode, "token过期，请重新登录", errors2.New("token过期，请重新登录")))
		return
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		// 删除博客
		db.Where("id = ? and user_id=?", c.Param("id"), val.(uint)).Delete(&post)
		// 删除评论
		db.Where("user_id = ? and post_id", val.(uint), c.Param("id")).Delete(&structs.Comment{})
		c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
		return nil
	})

	if err != nil {
		serverErrorCode := http.StatusInternalServerError
		c.AbortWithError(serverErrorCode, errors.NewError(serverErrorCode, "服务器内部错误", err))
	}
}
