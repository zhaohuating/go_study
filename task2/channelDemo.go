package main

import (
	"fmt"
	"sync"
)

// 1.编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 2.实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

func sendChan(sendChan chan<- int, wg *sync.WaitGroup) {
	defer func() {
		close(sendChan)
		wg.Done()
	}()

	for i := 1; i < 10; i++ {
		fmt.Println("向通道中发送数字：", i)
		sendChan <- i
	}
}

func reciveChan(reciveChan <-chan int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	for {
		if num, ok := <-reciveChan; ok {
			fmt.Println("从通道中获取到数字：", num)
		} else {
			break
		}
	}
}

// func main() {
// 	var wg sync.WaitGroup
// 	c := make(chan int) // 没有缓冲
// 	// c := make(chan int, 10) // 有缓冲
// 	wg.Add(2)
// 	go sendChan(c, &wg)
// 	go reciveChan(c, &wg)

// 	wg.Wait()
// }
