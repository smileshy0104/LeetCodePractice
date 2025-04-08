package main

import (
	"fmt"
)

// searchInsert 通过线性扫描查找目标值在排序数组中的插入位置。
// 如果目标值存在于数组中，则返回其索引；如果不存在，则返回目标值应插入的位置索引。
// 参数:
//
//	nums: 排序后的整数数组。
//	target: 需要查找的目标值。
//
// 返回值:
//
//	目标值所在或应插入的索引位置。
func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}

// searchInsert2 使用二分查找法确定目标值在排序数组中的插入位置。
// 如果目标值存在于数组中，则返回其索引；如果不存在，则返回目标值应插入的位置索引。
// 参数:
//
//	nums: 排序后的整数数组。
//	target: 需要查找的目标值。
//
// 返回值:
//
//	目标值所在或应插入的索引位置。
func searchInsert2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	// 使用二分查找法缩小搜索范围
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 当未找到目标值时，返回左指针作为插入位置
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}             // 定义一个已排序的数组
	target := 5                           // 定义目标值
	result := searchInsert2(nums, target) // 调用二分查找函数获取目标值的索引或插入位置
	fmt.Println(result)                   // 输出结果
}
