package main

import "fmt"

func main() {
	// 调用 isPerfectSquare 函数并打印结果。
	result := isPerfectSquare(16)
	fmt.Println(result) // 输出结果

	// 调用 isPerfectSquare 函数并打印结果。
	result = isPerfectSquare(14)
	fmt.Println(result) // 输出结果
}

// isPerfectSquare1 判断一个整数是否为完全平方数。
func isPerfectSquare1(num int) bool {
	// 如果num小于0，则不是完全平方数，因为完全平方数非负。
	if num < 0 {
		return false
	}
	// 如果num等于1，则是完全平方数，因为1的平方根是1。
	if num == 1 {
		return true
	}
	// 从2遍历到num-1，检查是否存在一个整数的平方等于num。
	for i := 2; i < num; i++ {
		if i*i == num {
			return true
		}
	}
	// 如果没有找到任何整数的平方等于num，则num不是完全平方数。
	return false
}

// isPerfectSquare 判断一个整数是否为完全平方数。（二分查找）
func isPerfectSquare(num int) bool {
	// 如果 num 小于 0，则不是完全平方数。
	if num < 0 {
		return false
	}
	// 如果 num 等于 1，则是完全平方数。
	if num == 1 {
		return true
	}
	// 初始化左右指针，用于二分查找。
	var left, right int = 0, num
	// 当左指针小于等于右指针时，执行二分查找。
	for left <= right {
		// 计算中间值。
		mid := (left + right) / 2
		// 如果中间值的平方等于 num，则 num 是完全平方数。
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			// 如果中间值的平方小于 num，则调整左指针。
			left = mid + 1
		} else {
			// 如果中间值的平方大于 num，则调整右指针。
			right = mid - 1
		}
	}
	// 如果查找结束没有找到完全平方数，则返回 false。
	return false
}
