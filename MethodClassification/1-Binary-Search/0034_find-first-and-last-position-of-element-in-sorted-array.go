package main

import (
	"fmt"
)

// searchRange 函数旨在找出目标值在整数数组中出现的最左侧和最右侧的位置。
// 如果目标值不存在于数组中，则返回[-1, -1]。
// 该函数通过一次遍历数组来实现，遇到目标值时，记录其首次出现的位置（left）和最后一次出现的位置（right）。
func searchRange(nums []int, target int) []int {
	left, right := -1, -1
	for i, num := range nums {
		if num == target {
			if left == -1 {
				left = i
			}
			right = i
		}
	}
	return []int{left, right}
}

// searchRange1 函数同样旨在找出目标值在整数数组中出现的最左侧和最右侧的位置。
// 与searchRange不同，此函数分别调用findFirst和findLast函数来寻找目标值的起始和结束位置。
func searchRange1(nums []int, target int) []int {
	first := findFirst(nums, target)
	last := findLast(nums, target)
	return []int{first, last}
}

// findFirst 函数使用二分查找算法来寻找目标值在数组中首次出现的位置。
// 如果找到目标值，记录其位置并继续在左半部分查找，以确保找到的是目标值的最左侧位置。
func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			result = mid
			right = mid - 1 // 继续查找左半部分
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

// findLast 函数使用二分查找算法来寻找目标值在数组中最后一次出现的位置。
// 如果找到目标值，记录其位置并继续在右半部分查找，以确保找到的是目标值的最右侧位置。
func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			result = mid
			left = mid + 1 // 继续查找右半部分
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

// TestName1 测试函数用于验证searchRange函数的正确性。
// 它定义了一个已排序的数组和一个目标值，然后调用searchRange函数，并打印出目标值在数组中的索引范围。
func main() {
	nums := []int{5, 7, 7, 8, 8, 10}    // 定义一个已排序的数组
	target := 8                         // 定义目标值
	result := searchRange(nums, target) // 调用函数获取目标值的索引范围
	fmt.Println(result)                 // 输出结果
}
