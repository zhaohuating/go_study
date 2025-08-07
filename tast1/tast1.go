package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 4, 5, 7, 98, 5, 98}
	onceNums := onceNum(nums)

	fmt.Println(onceNums)
}

// 只出现一次的数字
func onceNum(nums []int) []int {
	numCount := make(map[int]int)
	onceNums := make([]int, 0)

	for _, num := range nums {
		numCount[num]++
	}

	for num, count := range numCount {
		if count == 1 {
			onceNums = append(onceNums, num)
		}
	}

	return onceNums
}
