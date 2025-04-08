package main_test

import (
	"fmt"
	"testing"
)

// isPerfectSquare 检查一个整数是否是完全平方数。
// 参数:
//
//	num - 需要检查的整数。
//
// 返回值:
//
//	如果 num 是完全平方数，则返回 true；否则返回 false。
func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}

	if num == 1 {
		return true
	}

	// 使用开方公式计算平方根
	sqrt := int(float64(num) / 2.0)

	// 使用二分查找法查找平方根
	for sqrt*sqrt > num {
		sqrt = (sqrt + num/sqrt) / 2
	}

	// 检查平方根是否为整数
	return sqrt*sqrt == num
}

// isPerfectSquare1 检查一个整数是否是完全平方数，使用二分查找法。
// 参数:
//
//	num - 需要检查的整数。
//
// 返回值:
//
//	如果 num 是完全平方数，则返回 true；否则返回 false。
func isPerfectSquare1(num int) bool {
	if num < 0 {
		return false
	}
	// 初始化左右指针，用于二分查找。
	var left, right int = 0, num
	// 当左指针小于等于右指针时，执行二分查找。
	for left <= right {
		// 计算中间值，避免整数溢出。
		mid := left + (right-left)/2
		// 如果中间值的平方等于x，找到平方根，返回中间值。
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			// 如果中间值的平方小于x，调整左指针。
			left = mid + 1
		} else {
			// 如果中间值的平方大于x，调整右指针。
			right = mid - 1
		}
	}
	// 返回最接近平方根的值。
	return false

}

// TestName3 测试 isPerfectSquare1 函数。
func TestName3(t *testing.T) {
	result := isPerfectSquare1(14)
	fmt.Println(result) // 输出结果
}
