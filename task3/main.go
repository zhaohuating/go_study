package main

import (
	"errors"
	"fmt"
	"math/rand"
	"task3/connect"
	"task3/structs"
	"time"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
}

func (p *Person) TableName() string {
	return "person"
}

func InsetStu(db *gorm.DB) {
	stu := structs.Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}

	stu.InserOne(db)

}

func SelectByGtAge(db *gorm.DB) {
	stu := structs.Student{
		Name:  "张三",
		Age:   12,
		Grade: "三年级",
	}

	stuSli := stu.SelectByGtAge(db)

	for _, stu := range stuSli {
		fmt.Println(stu)
	}
}

func UpdateGradeByName(db *gorm.DB) {
	stu := structs.Student{
		Name:  "张三",
		Grade: "四年级",
	}

	stu.UpdateGradeByName(db)
}

func DeleteByLtAge(db *gorm.DB) {
	stu := structs.Student{
		Age: 30,
	}

	stu.DeleteByLtAge(db)
}

type Accounts struct {
	gorm.Model
	Name    string
	Balance float32
}
type Transactions struct {
	gorm.Model
	FromAccountID uint
	ToAccountID   uint
	Amount        float32
}

type employee struct {
	gorm.Model
	Name       string
	Department string
	Salary     float32
}
type book struct {
	ID     uint
	Title  string
	Author string
	Price  float32
}

func Task3_1_2(db *gorm.DB) {
	db.AutoMigrate(&Accounts{}, &Transactions{})
	from := "B"
	to := "A"
	num := 100

	//db.Create(&Accounts{
	//	Name:    from,
	//	Balance: 200,
	//})
	//db.Create(&Accounts{
	//	Name:    to,
	//	Balance: 200,
	//})

	err := db.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Select("balance").Where("balance >= ? and name = ?", num, from).Take(&Accounts{}).Count(&count)
		if count == 0 {
			return errors.New("转账余额不足！！！")
		}

		tx.Model(&Accounts{}).Where("name = ?", from).Update("balance", gorm.Expr("balance - ?", num))
		tx.Model(&Accounts{}).Where("name = ?", to).Update("balance", gorm.Expr("balance + ?", num))

		var AID uint
		var BID uint

		tx.Model(Accounts{}).Select("id").Where(&Accounts{Name: from}).Scan(&AID)
		tx.Model(Accounts{}).Select("id").Where(&Accounts{Name: to}).Scan(&BID)
		tx.Create(&Transactions{
			FromAccountID: AID,
			ToAccountID:   BID,
			Amount:        float32(num),
		})

		return nil
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("转账完成")
	}
}

func task3_1_1(db *gorm.DB) {
	// InsetStu(db)
	// SelectByGtAge(db)
	// UpdateGradeByName(db)
	// DeleteByLtAge(db)
}

func task3_2_1(db *sqlx.DB) {
	//insertDate(db)
	//selectDate1(db)
	selectDate2(db)
}

func task3_2_2(db *sqlx.DB) {
	//insertBookDate(db)
	selectBookDate(db)
}

func selectBookDate(db *sqlx.DB) {
	sql := "select * from books where price > ?"
	var books []book
	err := db.Select(&books, sql, 50)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(books)
}

func insertBookDate(db *sqlx.DB) {
	sql := "insert into books (id, title, author,price) values (?,?,?,?)"

	_, err := db.Exec(sql, 1, "书名1", "作者1", 59.9)
	if err != nil {
		fmt.Println(err)
	}
	db.Exec(sql, 2, "书名2", "作者2", 29.9)
	db.Exec(sql, 3, "书名3", "作者3", 49.9)
	db.Exec(sql, 4, "书名4", "作者4", 79.9)
	db.Exec(sql, 5, "书名5", "作者5", 19.9)
	db.Exec(sql, 6, "书名6", "作者6", 89.9)
}

func selectDate1(db *sqlx.DB) {
	fmt.Println(db.DriverName())
	sql := "select id, name,department,salary from employees where department = ? order by salary desc limit 1"
	var emp employee
	err := db.Get(&emp, sql, "技术部")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(emp)
}

func selectDate2(db *sqlx.DB) {
	fmt.Println(db.DriverName())
	sql := "select id, name,department,salary from employees order by salary desc limit 1"
	var emp employee
	err := db.Get(&emp, sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(emp)
}

func insertDate(db *sqlx.DB) {
	sqlInsert := "insert into employees (name, department,salary) values (?, ?, ?)"
	db.Exec(sqlInsert, "李四", "技术部", 13200)
	db.Exec(sqlInsert, "王武", "运营部", 8700)
	db.Exec(sqlInsert, "刘欢", "运营部", 8500)
	db.Exec(sqlInsert, "张梅", "人事部", 5300)
}

func insertUserAndPost(db *gorm.DB) {
	db.AutoMigrate(&structs.User{}, &structs.Comment{}, &structs.Post{})

	posts := []structs.Post{
		structs.Post{
			Title:   "文章的标题1",
			Content: "内容一",
			Summary: "摘要一",
		},
		structs.Post{
			Title:   "文章的标题2",
			Content: "内容二",
			Summary: "摘要二",
		},
		structs.Post{
			Title:   "文章的标题3",
			Content: "内容三",
			Summary: "摘要三",
		},
		structs.Post{
			Title:   "文章的标题3",
			Content: "内容三",
			Summary: "摘要三",
		},
		structs.Post{
			Title:   "文章的标题4",
			Content: "内容4",
			Summary: "摘要4",
		},
	}

	//comments := []structs.Comment{
	//	structs.Comment{
	//		Content: "评论1",
	//	},
	//	structs.Comment{
	//		Content: "评论2",
	//	},
	//	structs.Comment{
	//		Content: "评论3",
	//	},
	//	structs.Comment{
	//		Content: "评论4",
	//	},
	//}
	user := structs.User{
		Username: "李华",
		Email:    "lihua@qq.com",
		Password: "23344",
	}

	db.Create(&user)
	err := db.Model(&user).Association("Posts").Append(&posts)
	if err != nil {
		fmt.Println("++++++++++++", err)
	}
}
func insertComment(db *gorm.DB) {
	rand.Seed(time.Now().UnixNano())
	users := []structs.User{}
	posts := []structs.Post{}

	db.Find(&users)
	db.Find(&posts)

	for _, user := range users {
		for _, post := range posts {
			num := rand.Intn(10)
			commentSli := []structs.Comment{}
			for i := 0; i < num; i++ {
				comment := structs.Comment{
					UserID:  user.ID,
					PostID:  post.ID,
					Content: "这是一条评论",
				}
				commentSli = append(commentSli, comment)
			}

			db.CreateInBatches(&commentSli, 10)
		}
	}
}

// 用户发布的所有文章及其对应的评论信息。
func selectPostAndComments(db *gorm.DB) {
	userID := 1
	user := structs.User{}
	posts := []structs.Post{}

	// 用户所有的文章
	db.Where("id = ?", userID).Preload("Posts").Take(&user)
	// 对应文章的评论
	db.Where("user_id = ?", userID).Preload("Comments").Find(&posts)
	for _, post := range posts {
		fmt.Println("+++++", post)
	}
	fmt.Println(user)
}

// 查询评论数量最多的文章信息
func GetPostWithMostComments(db *gorm.DB) {
	var post structs.Post
	subQuery := db.Model(&structs.Comment{}).Select("count(*) as count,post_id").Group("post_id").Order("count desc").Limit(1)
	err := db.Joins("join (?) as commentNum on commentNum.post_id = posts.id", subQuery).Preload("User").First(&post).Error

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(post)
}

func insertPost(db *gorm.DB) {
	db.AutoMigrate(&structs.Post{})
	post := structs.Post{
		Title:   "文章的标题1",
		Content: "内容一",
		Summary: "摘要一",
		UserID:  2,
	}

	db.Create(&post)
}
func main() {
	db := connect.GetDB()
	//db := connect.GetSqlxDB()
	comment := structs.Comment{
		UserID: 1,
		PostID: 2,
	}

	err := db.Where("user_id = ? and post_id = ?", comment.UserID, comment.PostID).Delete(&comment).Error
	if err != nil {
		fmt.Println("&&&&&&&", err)
	}

}
