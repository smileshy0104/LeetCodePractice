package main_test

import (
	"fmt"
	"testing"
)

// mySqrt 函数计算并返回一个整数的平方根。
// 参数 x: 需要计算平方根的整数。
// 返回值: x 的平方根，也是一个整数。
func mySqrt(x int) int {
	// 当输入为0时，直接返回0。
	if x == 0 {
		return 0
	}
	// 当输入为1时，直接返回1。
	if x == 1 {
		return 1
	}
	// 初始化左右指针，用于二分查找。
	var left, right int = 0, x
	// 当左指针小于等于右指针时，执行二分查找。
	for left <= right {
		// 计算中间值，避免整数溢出。
		mid := left + (right-left)/2
		// 如果中间值的平方等于x，找到平方根，返回中间值。
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			// 如果中间值的平方小于x，调整左指针。
			left = mid + 1
		} else {
			// 如果中间值的平方大于x，调整右指针。
			right = mid - 1
		}
	}
	// 返回最接近平方根的值。
	return right
}

// TestName2 测试 mySqrt 函数。
func TestName2(t *testing.T) {
	// 调用 mySqrt 函数并打印结果。
	result := mySqrt(8)
	fmt.Println(result) // 输出结果
}
