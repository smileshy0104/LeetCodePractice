package main

import (
	"fmt"
	"sort"
)

// main函数是程序的入口点。
func main() {
	// 测试用例1
	nums := []int{-4, -1, 0, 3, 10}
	// 调用sortedSquares1函数来计算数组中每个元素的平方，并保持数组有序
	squares := sortedSquares(nums)
	fmt.Println(squares)

	// 测试用例2
	nums = []int{-7, -3, 2, 3, 11}
	// 调用sortedSquares1函数来计算数组中每个元素的平方，并保持数组有序
	squares = sortedSquares1(nums)
	fmt.Println(squares)
}

// sortedSquares函数接收一个整数数组nums，计算每个元素的平方，并保持数组有序。
// 它首先对原数组元素进行平方计算，然后对结果进行排序。
// 参数: nums []int - 输入的整数数组
// 返回值: []int - 计算平方后并排序的数组
func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

// sortedSquares1函数接收一个整数数组nums，计算每个元素的平方，并保持数组有序。
// 与sortedSquares不同，该函数使用双指针法来避免排序，从而提高效率。
// 参数: nums []int - 输入的整数数组
// 返回值: []int - 计算平方后并排序的数组
func sortedSquares1(nums []int) []int {
	// 初始化一个新的数组newNums来存储结果
	newNums := make([]int, len(nums))
	// indexF和indexS分别是新数组的前向和后向索引
	indexF := 0
	indexS := len(nums) - 1
	// index用于跟踪新数组中当前要填充的位置
	index := len(nums) - 1
	// 使用双指针法填充新数组
	for indexF <= indexS {
		// 比较两端元素的平方值，将较大的平方值放入新数组的当前位置
		if nums[indexF]*nums[indexF] < nums[indexS]*nums[indexS] {
			newNums[index] = nums[indexS] * nums[indexS]
			indexS--
		} else {
			newNums[index] = nums[indexF] * nums[indexF]
			indexF++
		}
		index--
	}
	return newNums
}
