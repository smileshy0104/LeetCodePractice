package main

import "fmt"

// 主函数入口
func main() {
	// 调用isHappy函数判断数字2是否为快乐数
	fmt.Println(isHappy(2))
}

// isHappy判断一个整数是否为快乐数
// 快乐数定义：当一个数字的每位平方和最终能够等于1时，该数字为快乐数
// 参数n: 待判断的整数
// 返回值: 如果n是快乐数返回true，否则返回false
func isHappy(n int) bool {
	// 循环直到n为1或4，其中1表示快乐数，4表示非快乐数
	for n != 1 && n != 4 {
		// 计算n的每位数字平方和
		n = getSum(n)
	}
	// 如果n最终等于1，则为快乐数
	if n == 1 {
		return true
	}
	// 如果n最终等于4，则非快乐数
	return false
}

// getSum计算一个整数的每位数字平方和
// 参数n: 需要计算的整数
// 返回值: n的每位数字平方和
func getSum(n int) int {
	sum := 0
	// 遍历n的每一位数字
	for n > 0 {
		// 将每位数字的平方加到sum上
		sum += (n % 10) * (n % 10)
		// 移除已经处理过的最后一位数字
		n = n / 10
	}
	// 返回所有位数字平方和
	return sum
}
