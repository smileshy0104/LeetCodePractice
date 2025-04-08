package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}             // 定义一个已排序的数组
	target := 5                           // 定义目标值
	result := searchInsert2(nums, target) // 调用二分查找函数获取目标值的索引或插入位置
	fmt.Println(result)                   // 输出结果

	nums = []int{1, 3, 5, 6}             // 定义一个已排序的数组
	target = 2                           // 定义目标值
	result = searchInsert2(nums, target) // 调用二分查找函数获取目标值的索引或插入位置
	fmt.Println(result)                  // 输出结果

	nums = []int{1, 3, 5, 6}             // 定义一个已排序的数组
	target = 7                           // 定义目标值
	result = searchInsert2(nums, target) // 调用二分查找函数获取目标值的索引或插入位置
	fmt.Println(result)                  // 输出结果
}

// searchInsert 函数在有序数组中查找目标值的插入位置。
// 如果目标值在数组中，返回其索引；如果目标值不在数组中，返回它应该插入的位置以保持数组有序。
// 参数 nums 是一个按非递减顺序排列的整数数组。
// 参数 target 是需要查找插入位置的目标值。
// 返回值是目标值在数组中应该插入的索引位置。
func searchInsert(nums []int, target int) int {
	// 遍历数组 nums，比较每个元素 num 与目标值 target。
	for i, num := range nums {
		// 如果当前元素 num 大于等于目标值 target，说明找到了插入位置。
		if num >= target {
			// 返回当前索引 i 作为插入位置。
			return i
		}
	}
	// 如果数组中所有元素都小于目标值 target，目标值应插入数组末尾。
	// 返回数组长度作为插入位置。
	return len(nums)
}

// searchInsert 函数在有序数组中查找目标值的插入位置。（通过二分查找）
// 如果目标值在数组中，返回其索引；如果目标值不在数组中，返回它应该插入的位置以保持数组有序。
// 参数 nums 是一个按非递减顺序排列的整数数组。
// 参数 target 是需要查找插入位置的目标值。
// 返回值是目标值在数组中应该插入的索引位置。
func searchInsert2(nums []int, target int) int {
	// 初始化左指针和右指针
	left := 0
	right := len(nums) - 1

	// 使用二分查找法缩小搜索范围
	for left <= right {
		// 计算中间位置
		mid := (left + right) / 2
		// 比较中间值与目标值
		if nums[mid] == target {
			// 如果找到目标值，返回其索引
			return mid
		} else if nums[mid] < target {
			// 如果中间值小于目标值，调整左指针
			left = mid + 1
		} else {
			// 如果中间值大于目标值，调整右指针
			right = mid - 1
		}
	}
	// 当未找到目标值时，返回左指针作为插入位置
	return left
}
