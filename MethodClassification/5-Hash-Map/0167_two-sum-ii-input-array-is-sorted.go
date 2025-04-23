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
	valueMap := make(map[int]int)
	for i, num := range numbers {
		if index, ok := valueMap[target-num]; ok {
			return []int{index + 1, i + 1}
		}
		valueMap[num] = i
	}
	return []int{-1, -1}
}
