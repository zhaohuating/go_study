package structs

import (
	"fmt"

	"gorm.io/gorm"
)

// User 博客系统的用户实体
type User struct {
	gorm.Model
	Username  string `gorm:"size:50;uniqueIndex;not null" json:"username"` // 用户名（唯一）
	Email     string `gorm:"size:100;uniqueIndex;not null" json:"email"`   // 邮箱（唯一）
	Password  string `gorm:"size:100;not null" json:"-"`                   // 密码（JSON序列化时忽略）
	Postcount uint   `gorm:"not null;default:0" json:"postcount"`
	// 关联：一个用户可以有多个文章
	Posts []Post `gorm:"foreignKey:UserID" json:"posts,omitempty"`
	//
	//// 关联：一个用户可以有多个评论
	//Comments []Comment `gorm:"foreignKey:UserID" json:"comments,omitempty"`
}

// Post 博客系统的文章实体
type Post struct {
	gorm.Model
	Title        string `gorm:"size:200;not null" json:"title"`    // 文章标题
	Content      string `gorm:"type:text;not null" json:"content"` // 文章内容
	Summary      string `gorm:"size:500" json:"summary"`           // 文章摘要
	UserID       uint   `gorm:"not null;index" json:"userId"`      // 作者ID（外键）
	CommentState string `gorm:"size:50" json:"commentState"`       // 评论状态

	// 关联：文章属于一个用户
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`

	// 关联：一篇文章可以有多个评论
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
}

// Comment 博客系统的评论实体
type Comment struct {
	gorm.Model
	Content string `gorm:"size:500;not null" json:"content"` // 评论内容
	UserID  uint   `gorm:"not null;index" json:"userId"`     // 评论者ID（外键）
	PostID  uint   `gorm:"not null;index" json:"postId"`     // 文章ID（外键）

	// 关联：评论属于一个用户
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`

	// 关联：评论属于一篇文章
	Post Post `gorm:"foreignKey:PostID" json:"post,omitempty"`
}

// 创建时自动更新用户的文章数量统计字段
func (p *Post) AfterCreate(db *gorm.DB) error {
	userid := p.UserID
	var num int64
	db.Model(p).Where("user_id = ?", userid).Count(&num)
	db.Model(User{}).Where("id = ?", userid).Update("postcount", num)
	fmt.Println("post AfterCreate runnint :::", num)
	return nil
}

// 评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterDelete(db *gorm.DB) error {
	userid := c.UserID
	postID := c.PostID
	var num int64
	db.Debug().Model(c).Where("user_id = ? and post_id = ?", userid, postID).Count(&num)
	if num == 0 {
		db.Model(Post{}).Where("id = ?", postID).Update("comment_state", "无评论")
	}
	fmt.Println("comment AfterDelete runnint :::", num)
	return nil
}
