package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func lockDemo1() {
	var mt sync.Mutex
	var wg sync.WaitGroup
	num := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mt.Lock()
			for i := 0; i < 1000; i++ {
				num += 1
			}
			mt.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(num)
}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func lockDemo2() {
	var num int32 = 0
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		index := i
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				fmt.Printf("第 %d 号协程正在操作,当前数据：%d\n", index+1, atomic.LoadInt32(&num))
				atomic.AddInt32(&num, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Print(num)
}

// func main() {
// 	lockDemo1()
// 	lockDemo2()
// }
