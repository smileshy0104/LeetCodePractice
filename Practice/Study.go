package Study

/*
35. 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，
返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。



示例 1:

输入: nums = [1,3,5,6], target = 5
输出: 2
示例 2:

输入: nums = [1,3,5,6], target = 2
输出: 1
示例 3:

输入: nums = [1,3,5,6], target = 7
输出: 4
*/
// 35. 搜索插入位置
func SearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

/*
34. 在排序数组中查找元素的第一个和最后一个位置
中等
2.4K
相关企业
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]
*/
func SearchRange0(nums []int, target int) []int {
	var index0, index1 = -1, -1
	for left := 0; left < len(nums)-1; left++ {
		if index0 == -1 {
			if nums[left] == target {
				index0 = left
			}
		}
		if index1 == -1 {
			if nums[len(nums)-1-left] == target {
				index1 = len(nums) - 1 - left
			}
		}
	}
	index := []int{index0, index1}
	return index
}

func SearchRange(nums []int, target int) []int {
	//找第一个满足条件的位置
	start := findFirstPosition(nums, target)
	if start == -1 {
		return []int{-1, -1}
	}
	//找最后一个满足条件的位置
	end := findLastPosition(nums, target)
	return []int{start, end}
}

// 找第一个满足条件的位置
func findFirstPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
			if nums[mid] == target {
				result = mid
			}
		} else {
			left = mid + 1
		}
	}
	return result
}

// 找最后一个满足条件的位置
func findLastPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
			if nums[mid] == target {
				result = mid
			}
		} else {
			right = mid - 1
		}
	}
	return result
}

/*
69. x 的平方根
提示
简单
1.4K
相关企业
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。



示例 1：

输入：x = 4
输出：2
示例 2：

输入：x = 8
输出：2
解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。

*/

// 69. x 的平方根
func MySqrt(x int) int {
	left := 1
	right := x
	index := 0
	for left <= right {
		middle := left + (right-left)/2
		if middle*middle == x {
			return middle
		} else if middle*middle > x {
			right = middle - 1
		} else {
			left = middle + 1
			index = middle
		}
	}
	return index
}

/*
367. 有效的完全平方数
简单
526
相关企业
给你一个正整数 num 。如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

完全平方数 是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。

不能使用任何内置的库函数，如  sqrt 。

示例 1：

输入：num = 16
输出：true
解释：返回 true ，因为 4 * 4 = 16 且 4 是一个整数。
示例 2：

输入：num = 14
输出：false
解释：返回 false ，因为 3.742 * 3.742 = 14 但 3.742 不是一个整数。
*/

// 367. 有效的完全平方数
func IsPerfectSquare(num int) bool {
	left := 0
	right := num
	for left <= right {
		middle := left + (right-left)/2
		if middle*middle == num {
			return true
		} else if middle*middle > num {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return false
}

/**
 * 范围查询规律
 * 初始化:
 *   int left = 0;
 *   int right = nums.length - 1;
 * 循环条件
 *   left <= right
 * 右边取值
 *   right = mid - 1
 * 左边取值
 *   left = mid + 1
 * 查询条件
 *   >= target值, 则 nums[mid] >= target时, 都减right = mid - 1
 *   >  target值, 则 nums[mid] >  target时, 都减right = mid - 1
 *   <= target值, 则 nums[mid] <= target时, 都加left = mid + 1
 *   <  target值, 则 nums[mid] <  target时, 都加left = mid + 1
 * 结果
 *   求大于(含等于), 返回left
 *   求小于(含等于), 返回right
 * 核心思想: 要找某个值, 则查找时遇到该值时, 当前指针(例如right指针)要错过它, 让另外一个指针(left指针)跨过他(体现在left <= right中的=号), 则找到了
 */

// （版本一）左闭右闭区间
func BinarySearch1(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

// （版本二）左闭右开区间
func BinarySearch2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right { //此处不同
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return -1
}

// 使用递归进行二分查找
func BinarySearch3(nums []int, left, right, target int) int {
	if left > right {
		return left
	}
	middle := left + (left-right)/2
	if nums[middle] >= target {
		//保证找到最左边
		return BinarySearch3(nums, left, middle-1, target)
	} else {
		return BinarySearch3(nums, middle+1, right, target)
	}
}
