package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	// 边界条件检查：如果数组长度小于3，直接返回空结果
	if len(nums) < 3 {
		return [][]int{}
	}

	// 对数组进行排序，便于后续去重和双指针操作
	sort.Ints(nums)
	ret := [][]int{}

	// 遍历数组，固定一个数，使用双指针查找另外两个数
	for i := 0; i < len(nums)-2; i++ {
		// 去重：跳过相同的第一个数
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 定义左右指针
		left, right := i+1, len(nums)-1

		// 双指针查找剩余两个数
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				// 找到符合条件的三元组
				ret = append(ret, []int{nums[i], nums[left], nums[right]})

				// 去重：跳过相同的第二个数
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				// 去重：跳过相同的第三个数
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// 移动指针继续查找
				left++
				right--
			} else if sum < 0 {
				// 和小于0，移动左指针增大数值
				left++
			} else {
				// 和大于0，移动右指针减小数值
				right--
			}
		}
	}

	return ret
}
