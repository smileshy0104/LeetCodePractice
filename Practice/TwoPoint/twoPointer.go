package TwoPoint

import (
	"fmt"
	"sort"
)

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

/*
844. 比较含退格的字符串
简单
662
相关企业
给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。

注意：如果对空文本输入退格字符，文本继续为空。

示例 1：

输入：s = "ab#c", t = "ad#c"
输出：true
解释：s 和 t 都会变成 "ac"。
示例 2：

输入：s = "ab##", t = "c#d#"
输出：true
解释：s 和 t 都会变成 ""。
示例 3：

输入：s = "a#c", t = "b"
输出：false
解释：s 会变成 "c"，但 t 仍然是 "b"。

提示：

1 <= s.length, t.length <= 200
s 和 t 只含有小写字母以及字符 '#'
*/

// 844. 比较含退格的字符串
func BackspaceCompare(s string, t string) bool {
	return buildModifiedString(s) == buildModifiedString(t)
}

func buildModifiedString(s string) string {
	slow := 0
	fast := 0
	str := []byte(s)
	for fast < len(str) {
		if str[fast] != '#' {
			str[slow] = str[fast]
			slow++
		} else {
			if slow > 0 {
				slow--
			}
		}
		fast++
	}
	fmt.Println("the string is:", string(str[:slow]))

	return string(str[:slow])
}

func buildModifiedString1(s string) string {
	str := []byte(s)
	slow := 0

	for fast := 0; fast < len(str); fast++ {
		if str[fast] != '#' {
			str[slow] = str[fast]
			slow++
		} else if slow > 0 {
			slow--
		}
	}

	return string(str[:slow])
}

/*
977. 有序数组的平方
简单
873
相关企业
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：

输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]
示例 2：

输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
*/

// 977. 有序数组的平方（非双指针）
func SortedSquares(nums []int) []int {
	for k, v := range nums {
		nums[k] = v * v
	}
	sort.Ints(nums)
	fmt.Println(nums)
	return nums
}

// 双指针实现
func SortedSquares1(nums []int) []int {
	left := 0
	right := len(nums) - 1
	index := len(nums) - 1
	result := make([]int, len(nums))
	for left <= right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			result[index] = nums[right] * nums[right]
			right--
		} else {
			result[index] = nums[left] * nums[left]
			left++
		}
		index--
	}
	return result
}
