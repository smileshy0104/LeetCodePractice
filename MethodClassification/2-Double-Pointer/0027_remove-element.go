package main

import "fmt"

// main函数是程序的入口点。
func main() {
	// 测试用例1
	nums := []int{3, 2, 2, 3}
	target := 3
	// 调用removeElement函数移除nums中的target值，并打印处理后的nums和返回的结果
	result := removeElement(nums, target)
	fmt.Println(nums)
	fmt.Println(result)

	// 测试用例2
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	target = 2
	// 调用removeElement函数移除nums中的target值，并打印处理后的nums和返回的结果
	result = removeElement(nums, target)
	fmt.Println(nums)
	fmt.Println(result)
}

// removeElement函数用于移除切片nums中所有值为val的元素(利用双指针前移)
// 它通过迭代nums中的每个元素，如果元素值不为val，则将其移动到切片的前部。
func removeElement(nums []int, val int) int {
	i := 0
	for _, num := range nums {
		if num != val {
			nums[i] = num
			i++
		}
	}
	return i
}

// removeElement2函数是removeElement的另一种实现，使用两个指针进行操作。(利用双指针前移)
// 它通过迭代nums中的每个元素，如果元素值不为val，则将其移动到切片的前部。
func removeElement2(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
