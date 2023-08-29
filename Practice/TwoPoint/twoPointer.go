package TwoPoint

import "fmt"

/*
27. 移除元素
提示
简单
1.9K
相关企业
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/

// 使用双指针（前后指针）
func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1

	for left < right {
		// 从左边找到第一个等于 val 的元素（找到第一个等于的时候，跳出）
		for left < right && nums[left] != val {
			left++
		}
		// 从右边找到第一个不等于 val 的元素（找到第一个不等于的时候，跳出）
		for left < right && nums[right] == val {
			right--
		}
		// 交换左右指针所指向的元素
		nums[left], nums[right] = nums[right], nums[left]
	}

	// 判断最后一个元素是否为 val
	if nums[left] == val {
		return left
	}
	// 返回新的数组长度
	return left + 1
}

// 快慢指针法
func RemoveElement1(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 0
	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] != val {
			nums[result] = nums[i]
			result++
		}
	}
	nums = nums[:result]
	fmt.Println("the nums is:", nums)
	return result
}
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}
