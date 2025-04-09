package main

import "fmt"

// main函数是程序的入口点。
func main() {
	// 测试用例1
	nums := []int{0, 1, 0, 3, 12}
	// 调用moveZeroes函数将nums中的所有0移动到数组的末尾，并打印处理后的nums
	moveZeroes(nums)
	fmt.Println(nums)

	// 测试用例2
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	// 调用moveZeroes函数将nums中的所有0移动到数组的末尾，并打印处理后的nums
	moveZeroes(nums)
	fmt.Println(nums)

	// 测试用例3
	nums = []int{0}
	// 调用moveZeroes函数将nums中的所有0移动到数组的末尾，并打印处理后的nums
	moveZeroes(nums)
	fmt.Println(nums)
}

// moveZeroes函数将数组中的所有0移动到数组的末尾。（双指针+补0）
// 它通过两个步骤实现：
// 1. 遍历数组，将所有非0元素按顺序移动到数组的前部。
// 2. 将数组剩余的位置填充为0。
func moveZeroes(nums []int) {
	i := 0
	for _, num := range nums {
		if num != 0 {
			nums[i] = num
			i++
		}
	}
	for i < len(nums) {
		nums[i] = 0
		i++
	}
}
