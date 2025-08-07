package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	nums := []int{1, 2, 2, 9}
	onceNums := onceNum(nums)

	fmt.Println(onceNums)

	// 是否是回文数
	fmt.Println(isPalindrome(2))

	// 加一
	fmt.Println(addOne_1([]int{9, 9, 9}))
	fmt.Println(addOne_2([]int{9, 9, 9}))

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

// 回文数
func isPalindrome(num int) bool {
	//处理负数 个位数并且是非0数字
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

func addOne_1(nums []int) []int {
	numsLen := len(nums)
	var num int
	for i := 0; i < numsLen; i++ {
		num += int(float64(nums[i]) * math.Pow(10, float64(numsLen-1)-float64(i)))
	}
	num += 1
	numStr := strconv.FormatInt(int64(num), 10)

	// 创建一个切片，长度与字符串长度相同
	nums1 := make([]int, 0, len(numStr))
	for _, char := range numStr {
		nums1 = append(nums1, int(char-'0'))
	}
	return nums1
}

func addOne_2(nums []int) []int {
	carry := 1 // 初始进位为 1（即要加的 1）
	for i := len(nums) - 1; i >= 0; i-- {
		sum := nums[i] + carry
		nums[i] = sum % 10 // 当前位的结果
		carry = sum / 10   // 计算新的进位
		if carry == 0 {
			break // 没有进位时，提前退出循环
		}
	}
	// 如果所有位都进位了（例如 999 + 1 = 1000）
	if carry > 0 {
		nums = append([]int{carry}, nums...)
	}
	return nums
}
