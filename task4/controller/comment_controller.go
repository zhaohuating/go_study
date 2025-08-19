package controller

import (
	errors2 "errors"
	"net/http"
	"strconv"
	"task4/errors"
	"task4/pagination"
	"task4/structs"

	"github.com/gin-gonic/gin"
)

type Comment structs.Comment

// 实现评论的创建功能，已认证的用户可以对文章发表评论。
func AddComment(c *gin.Context) {
	var comment Comment

	if err := c.ShouldBindBodyWithJSON(&comment); err != nil {
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
	comment.UserID = val.(uint)

	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create comment"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment created successfully"})
}

// 实现评论的读取功能，支持获取某篇文章的所有评论列表 分页查询
func GetPostComments(c *gin.Context) {
	postIDStr := c.Param("postID")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	var pageParam = pagination.Param{
		Page:     page,
		PageSize: pageSize,
	}
	postID, _ := strconv.Atoi(postIDStr)
	tx := db.Model(Comment{}).Where("post_id = ?", postID)
	paginate, err := pagination.Paginate(tx, pageParam, &[]Comment{})

	if err != nil {
		serverErrorCode := http.StatusInternalServerError
		c.AbortWithError(serverErrorCode, errors.NewError(serverErrorCode, "查询失败", err))
		return
	}

	c.JSON(http.StatusOK, paginate)
}
