package main

import (
	"fmt"
)

// main函数是程序的入口点。
func main() {
	// 测试用例1
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	// 调用minSubArrayLen函数来找出和至少为target的最小子数组长度
	length := minSubArrayLen(target, nums)
	fmt.Println(length)

	// 测试用例2
	nums = []int{1, 4, 4}
	target = 4
	// 调用minSubArrayLen函数来找出和至少为target的最小子数组长度
	length = minSubArrayLen(target, nums)
	fmt.Println(length)

	// 测试用例3
	nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	target = 11
	// 调用minSubArrayLen函数来找出和至少为target的最小子数组长度
	length = minSubArrayLen(target, nums)
	fmt.Println(length)
}

// 方法一：滑动窗口（双指针）
// minSubArrayLen函数用于寻找和至少为target的最小子数组长度。
// 参数target是目标和，nums是输入的整数数组。
// 返回值是满足条件的最小子数组长度，如果不存在则返回0。
func minSubArrayLen(target int, nums []int) int {
	// 输入校验
	if len(nums) == 0 {
		return 0
	}

	// 初始化变量
	minLength := len(nums) + 1 // 设置一个不可能达到的初始值
	left, sum := 0, 0

	// 滑动窗口遍历数组
	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		// 当窗口内的和大于等于 target 时，尝试缩小窗口
		for sum >= target {
			currentLength := right - left + 1
			if currentLength < minLength {
				minLength = currentLength
			}
			// 收缩窗口
			sum -= nums[left]
			left++
		}
	}

	// 如果 minLength 未被更新，说明没有找到符合条件的子数组
	if minLength == len(nums)+1 {
		return 0
	}
	return minLength
}
