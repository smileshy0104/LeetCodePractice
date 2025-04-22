package main

import "fmt"

// 主函数
func main() {
	// 调用twoSum函数并打印结果
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

// twoSum函数旨在找出数组中两个数之和等于特定目标值的那两个数的索引
// 它接受一个整数数组nums和一个目标值target作为参数
// 返回一个包含这两个数的索引的整数数组，如果找不到这样的两个数，则返回一个空数组
func twoSum(nums []int, target int) []int {
	// 创建一个map用于存储数组中的值和对应的索引
	hash := make(map[int]int)
	// 遍历数组nums
	for i, v := range nums {
		// 检查目标值与当前值的差是否已存在于hash中
		if index, ok := hash[target-v]; ok {
			// 如果存在，说明找到了两个数，它们的和等于目标值
			// 返回这两个数的索引
			return []int{index, i}
		}
		// 将当前值和它的索引存入hash中，以便后续查找
		hash[v] = i
	}
	// 如果遍历结束后仍未找到符合条件的两个数，返回一个空数组
	return []int{}
}

// twoSum0 寻找两个数，使它们的和等于目标值。
// 参数 nums 是一个整数数组。
// 参数 target 是需要两数之和达到的目标值。
// 返回值是包含这两个数的索引的切片，如果不存在这样的两个数，则返回空切片。
func twoSum0(nums []int, target int) []int {
	// 遍历数组中的每个元素。
	for i := 0; i < len(nums); i++ {
		// 对当前元素之后的每个元素进行遍历。
		for j := i + 1; j < len(nums); j++ {
			// 检查当前的两个元素之和是否等于目标值。
			if nums[i]+nums[j] == target {
				// 如果找到符合条件的两个数，返回它们的索引。
				return []int{i, j}
			}
		}
	}
	// 如果没有找到符合条件的两个数，返回空切片。
	return []int{}
}
