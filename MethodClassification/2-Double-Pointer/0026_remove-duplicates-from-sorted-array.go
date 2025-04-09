package main

import "fmt"

// main函数是程序的入口点。
func main() {
	// 测试用例1
	nums := []int{1, 1, 2}
	result := removeDuplicates(nums)
	fmt.Println(nums)
	fmt.Println(result)

	// 测试用例2
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result = removeDuplicates(nums)
	fmt.Println(nums)
	fmt.Println(result)
}

// removeDuplicates函数用于移除切片中的重复元素。（HashMap）
// 它接受一个整数切片作为参数，并返回移除重复元素后的切片长度。
// 该函数使用了一个map来记录出现过的元素，从而保证每个元素只出现一次。
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numsMap := make(map[int]bool)
	i := 0
	for _, num := range nums {
		if !numsMap[num] {
			numsMap[num] = true
			nums[i] = num
			i++
		}
	}
	return i
}

// removeDuplicates2函数是removeDuplicates函数的改进版。（利用双指针）
// 它同样用于移除切片中的重复元素，但不使用额外的map，而是通过双指针技术来实现。
// 该函数接受一个整数切片作为参数，并返回移除重复元素后的切片长度。
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
