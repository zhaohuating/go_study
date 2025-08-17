package main

import (
	"errors"
	"fmt"
	"task3/connect"
	"task3/structs"

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
func main() {
	//db := connect.GetDB()
	//db.AutoMigrate(&book{})
	db := connect.GetSqlxDB()
	task3_2_2(db)
}
