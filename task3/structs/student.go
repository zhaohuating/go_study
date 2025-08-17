package structs

import (
	"gorm.io/gorm"
)

// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
type Student struct {
	gorm.Model
	Name  string
	Age   int
	Grade string
}

func (s *Student) InserOne(db *gorm.DB) {
	db.AutoMigrate(&Student{})
	db.Create(s)
}

func (s *Student) SelectByGtAge(db *gorm.DB) []Student {
	var stuSli []Student
	db.Where("age > ?", s.Age).Find(&stuSli)

	return stuSli
}

func (s *Student) UpdateGradeByName(db *gorm.DB) {
	db.Model(s).Where("name = ?", s.Name).Update("grade", s.Grade)
}

func (s *Student) DeleteByLtAge(db *gorm.DB) {
	db.Where("age < ?", s.Age).Delete(s)
}
