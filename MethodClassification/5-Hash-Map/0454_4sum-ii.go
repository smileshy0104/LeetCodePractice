package main

import "fmt"

func main() {
	fmt.Println(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}

// fourSumCount 计算四个数组中，选取的任意四个数相加等于0的组合数量。
// 通过预先计算nums1和nums2中任意两数之和的出现次数，存储在valueMap中，
// 然后遍历nums3和nums4，寻找与之前计算的两数之和相反数的出现次数，
// 从而统计所有满足条件的组合数量。
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 初始化一个map来存储nums1和nums2中任意两数之和的出现次数
	valueMap := make(map[int]int)
	// 初始化计数器，用于统计满足条件的组合数量
	count := 0

	// 遍历nums1和nums2，计算所有可能的两数之和，并记录每个和值的出现次数
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			valueMap[v1+v2]++
		}
	}

	// 遍历nums3和nums4，寻找与之前计算的两数之和相反数的出现次数
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			// 如果找到了相反数，说明存在一个组合使得四个数之和为0
			if value, ok := valueMap[0-v3-v4]; ok {
				// 将这个相反数的出现次数加到计数器上，代表这些组合都是有效的
				count += value
			}
		}
	}

	// 返回满足条件的组合数量
	return count
}
