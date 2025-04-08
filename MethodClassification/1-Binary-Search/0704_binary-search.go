package main

// search 使用二分查找算法在排序数组中寻找目标值。闭区间 【left, right】
// 它接受一个排序后的整数数组 nums 和一个目标值 target 作为参数。
// 如果目标值存在于数组中，返回其索引；如果不存在，返回 -1。
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// search2 使用二分查找算法在排序数组中寻找目标值。开区间 【left, right）
func search2(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid // 因为其中是开区间，所以是 mid
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	//target := 2
	target := 9
	result := search(nums, target)
	println(result)
}
