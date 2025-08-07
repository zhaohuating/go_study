package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 4, 5, 7, 98, 5, 98}
	onceNums := onceNum(nums)

	fmt.Println(onceNums)

	// 是否是回文数
	fmt.Println(isPalindrome(2))
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

func isPalindrome(num int) bool {
	//处理负数 个位数是0的非0数字
	if num < 0 || (num%10 == 0 && num != 0) {
		return false
	}

	halfReversed := 0

	for halfReversed < num {
		lastNum := num % 10
		halfReversed = halfReversed*10 + lastNum
		num /= 10
	}

	return num == halfReversed || num == halfReversed/10
}
