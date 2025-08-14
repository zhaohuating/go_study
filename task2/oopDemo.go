package main

import "fmt"

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

type Circle struct {
}

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (emp Employee) PrintInfo() {
	fmt.Printf("员工信息：姓名：%s，年龄：%d，部门ID：%d", emp.Name, emp.Age, emp.EmployeeID)
}

func (rectangle Rectangle) Area() {
	fmt.Println("Rectangle的Area方法执行")
}

func (rectangle Rectangle) Perimeter() {
	fmt.Println("Rectangle的Perimeter方法执行")
}

func (circle Circle) Area() {
	fmt.Println("Circle的Area方法执行")
}

func (circle Circle) Perimeter() {
	fmt.Println("Circle的Perimeter方法执行")
}

// func main() {
// 	rec := Rectangle{}
// 	cir := Circle{}

// 	rec.Area()
// 	rec.Perimeter()
// 	cir.Area()
// 	cir.Perimeter()

// 	p := Person{
// 		Name: "张三",
// 		Age:  20,
// 	}
// 	emp := Employee{
// 		Person:     p,
// 		EmployeeID: 0,
// 	}

// 	emp.PrintInfo()
// }
