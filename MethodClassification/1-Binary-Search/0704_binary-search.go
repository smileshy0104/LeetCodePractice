package main

import "fmt"

func main() {
	// 测试用例1
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	result := search(nums, target)
	fmt.Println(result)

	// 测试用例2
	nums = []int{-1, 0, 3, 5, 9, 12}
	target = 2
	result = search2(nums, target)
	fmt.Println(result)
}

// 闭区间 [ left, right ]
// search 函数在有序数组 nums 中查找目标值 target。
// 如果找到目标值，返回其在数组中的索引；如果没有找到，则返回 -1。
// 该函数使用二分查找算法，时间复杂度为 O(log n)。
func search(nums []int, target int) int {
	// 初始化左边界和右边界。
	left := 0
	right := len(nums) - 1

	// 当左边界不超过右边界时，执行二分查找。
	for left <= right {
		// 计算中间位置。
		mid := (left + right) / 2

		// 如果中间位置的值等于目标值，返回中间位置的索引。
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			// 如果中间位置的值小于目标值，调整左边界。
			left = mid + 1
		} else {
			// 如果中间位置的值大于目标值，调整右边界。
			right = mid - 1
		}
	}

	// 如果循环结束时未找到目标值，返回 -1。
	return -1
}

// 闭区间 [ left, right ) ( left, right ] ( left, right )
// search2 函数在有序数组 nums 中查找目标值 target。
// 如果找到目标值，返回其在数组中的索引；如果没有找到，则返回 -1。
// 该函数使用二分查找算法，时间复杂度为 O(log n)。
func search2(nums []int, target int) int {
	// 初始化左指针和右指针，右指针是开区间的，即不包含右指针本身。
	left := 0
	right := len(nums)

	// 当左指针小于右指针时，执行二分查找。
	for left < right {
		// 计算中间位置。
		mid := (left + right) / 2

		// 根据目标值与中间位置的值比较，调整搜索范围。
		if nums[mid] == target {
			// 如果中间位置的值等于目标值，返回中间位置的索引。
			return mid
		} else if nums[mid] < target {
			// 如果中间位置的值小于目标值，调整左指针到中间位置的右侧。
			left = mid + 1
		} else {
			// 如果中间位置的值大于目标值，调整右指针到中间位置。
			// 因为右指针是开区间的，所以这里使用 mid 而不是 mid - 1。
			right = mid
		}
	}

	// 如果循环结束还没有找到目标值，返回-1。
	return -1
}
