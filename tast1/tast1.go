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

	//有效括号

	fmt.Println("是否是有效括号：", isValid("({])"))
	//最大公共前缀
	fmt.Println("最大公共前缀", longestCommonPrefix([]string{"we5qq1eweqw1rerwerfsd", "we6qq1e", "we6qq1eweqe1rwe"}))

	//删除有序数组中的重复项
	fmt.Println(removeDuplicates([]int{1, 1, 2, 3, 3, 4, 4, 4, 5}))
	//合并区间
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {4, 10}, {15, 18}}))

	// 加一
	fmt.Println(addOne_1([]int{9, 9, 9}))
	fmt.Println(addOne_2([]int{9, 9, 9}))

	// 两数之和
	fmt.Println(twoSum([]int{1, 3, 34, 2, 6, 9}, 5))

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

// 有效括号
func isValid(s string) bool {
	if len(s)%2 != 0 || len(s) < 2 {
		return false
	}

	strMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if p, ok := strMap[s[i]]; ok {
			// 遇到右括号
			if len(stack) == 0 || p != stack[len(stack)-1] {
				return false
			}
			// 出栈
			stack = stack[:len(stack)-1]
		} else {
			// 左括号
			//入栈
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {

	if len(strs) == 1 {
		return strs[0]
	}

	if len(strs) == 0 {
		return ""
	}

	firstStr := strs[0]
	for i := 0; i < len(firstStr); i++ {
		c := firstStr[i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}

	return firstStr
}

// 删除有序数组中的重复项
func removeDuplicates(nums []int) []int {

	index := 1
	base := nums[0]
	len := len(nums)
	for i := 1; i < len; i++ {
		if base == nums[i] {
			continue
		}

		nums[index] = nums[i]
		base = nums[i]
		index++
	}

	nums = nums[0:index]

	return nums
}

// 合并区间
func merge(intervals [][]int) [][]int {
	newIntervals := [][]int{}
	newIntervals = append(newIntervals, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if newIntervals[len(newIntervals)-1][1] >= intervals[i][0] {
			newIntervals[len(newIntervals)-1][1] = intervals[i][1]
		} else {
			newIntervals = append(newIntervals, intervals[i])
		}
	}

	return newIntervals
}

// 加一
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

// 加一
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

// 两数之和
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, num := range nums {
		subNum := target - num
		if num, ok := numMap[subNum]; ok {
			return []int{index, num}
		}

		numMap[num] = index
	}

	return nil
}
