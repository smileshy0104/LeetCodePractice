package main

import "fmt"

// 主函数，用于测试 intersection 函数
// 通过定义两个数组 nums1 和 nums2，并调用 intersection 函数计算它们的交集
// 最后将结果打印到控制台
func main() {
	nums1 := []int{4, 9, 5}       // 定义第一个数组
	nums2 := []int{9, 4, 9, 8, 4} // 定义第二个数组
	// 调用 intersection 函数并打印结果
	fmt.Println(intersection(nums1, nums2))
}

// intersection 计算两个整数数组的交集
// 参数：
//
//	nums1 - 第一个整数数组
//	nums2 - 第二个整数数组
//
// 返回值：
//
//	包含两个数组交集元素的数组（每个元素唯一）
func intersection(nums1 []int, nums2 []int) []int {
	nums := make(map[int]bool) // 创建一个 map 用于存储 nums1 的元素
	// 遍历 nums1，将所有元素存入 map 中，值设置为 true
	for _, num := range nums1 {
		nums[num] = true
	}
	result := make([]int, 0) // 创建一个切片用于存储交集结果
	// 遍历 nums2，检查元素是否存在于 map 中
	// 如果存在，则将其添加到结果切片中，并将 map 中对应的值设置为 false，避免重复添加
	for _, num := range nums2 {
		if nums[num] {
			result = append(result, num)
			nums[num] = false
		}
	}
	return result // 返回交集结果
}
