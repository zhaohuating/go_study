package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func printOod(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println("正在输出奇数：", i)
		}
	}
}

func printEven(wg *sync.WaitGroup) {
	wg.Done()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("正在输出偶数：", i)
		}
	}
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go printOod(&wg)
// 	go printEven(&wg)

// 	wg.Wait()
// }
