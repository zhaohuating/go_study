package main

// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func add(num *int) int {
	return *num + 10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func mul2(intSli []int) {
	for index, _ := range intSli {
		intSli[index] *= 2
	}
}

// func main() {
// 	num := 10
// 	intSli := []int{2, 5, 6, 32, 41, 2}
// 	fmt.Println(add(&num))
// 	mul2(intSli)
// 	for _, num := range intSli {
// 		fmt.Println(num)
// 	}
// }
