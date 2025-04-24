package main

import (
	"fmt"
	"sort"
)

// 主函数，用于测试fourSum函数
func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}

// fourSum函数寻找数组中所有不重复的四个元素组合，使其总和等于给定的目标值
// nums: 输入的整数数组
// target: 目标总和值
// 返回值: 一个二维整数数组，包含所有满足条件的四个元素组合
func fourSum(nums []int, target int) [][]int {
	// 如果数组长度小于4，不可能找到满足条件的组合，直接返回空结果
	if len(nums) < 4 {
		return [][]int{}
	}
	// 对数组进行排序，以便后续使用双指针技术
	sort.Ints(nums)
	// 初始化结果集
	res := make([][]int, 0)
	// 遍历数组，固定两个元素，使用双指针技术寻找其他两个元素
	for i := 0; i < len(nums)-3; i++ {
		// 跳过重复元素，避免重复组合
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			// 跳过重复元素，避免重复组合
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 初始化双指针
			left, right := j+1, len(nums)-1
			// 寻找其他两个元素的组合
			for left < right {
				// 计算当前组合的总和
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				// 根据总和与目标值的比较，调整指针位置或记录当前组合
				if sum == target {
					// 记录当前组合
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					// 跳过重复元素
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					// 调整指针位置
					left++
					right--
				} else if sum < target {
					// 总和小于目标值，增加左指针
					left++
				} else {
					// 总和大于目标值，减少右指针
					right--
				}
			}
		}
	}
	// 返回结果集
	return res
}
