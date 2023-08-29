package Study

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
func RemoveElement(nums []int, val int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		for left <= right && nums[left] != val {
			left++
		}
		for left <= right && nums[right] == val {
			right--
		}
		for left < right {
			nums[left] = nums[right]
		}
	}
	return left
}
