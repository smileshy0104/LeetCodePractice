package main

import "fmt"

// 主函数
func main() {
	// 初始化一个整数数组
	numbers := []int{2, 7, 11, 15}
	// 目标值
	target := 9
	// 调用twoSum函数并打印结果
	fmt.Println(twoSum(numbers, target))
}

// twoSum函数用于寻找数组中两个数之和等于目标值的索引
// 参数numbers为一个已排序的整数数组，target为目标值
// 返回值为一个包含这两个数的索引的切片，索引从1开始计数
// 如果没有找到这样的两个数，则返回[-1, -1]
func twoSum(numbers []int, target int) []int {
	// 初始化左指针
	left := 0
	// 初始化右指针
	right := len(numbers) - 1
	// 当左指针小于右指针时循环
	for left < right {
		// 计算当前左右指针所指元素之和
		sum := numbers[left] + numbers[right]
		// 如果和等于目标值，返回当前指针的索引（从1开始）
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			// 如果和小于目标值，左指针右移
			left++
		} else {
			// 如果和大于目标值，右指针左移
			right--
		}
	}
	// 如果没有找到满足条件的两个数，返回[-1, -1]
	return []int{-1, -1}
}
