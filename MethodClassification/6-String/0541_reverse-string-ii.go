package main

import "fmt"

// 主函数
func main() {
	// 调用reverseStr函数反转字符串并打印结果
	fmt.Println(reverseStr("abcdefg", 2))
}

// reverseStr 函数用于反转字符串中的每一个至多 k 个字符的子串，并返回反转后的字符串。
// 参数 s 是待反转的字符串，k 是子串的最大长度。
func reverseStr(s string, k int) string {
	// 将字符串转换为字节切片以便进行字符操作
	arr := []byte(s)
	// 遍历字符串，每次处理2k个字符
	for i := 0; i < len(s); i = i + 2*k {
		// 确定当前需要反转的子串的结束位置
		j := min(i+k, len(s))
		// 反转当前子串
		reverse(arr[i:j])
	}
	// 将处理后的字节切片转换回字符串并返回
	return string(arr)
}

// reverse 函数用于反转一个字节切片。
// 参数 arr 是待反转的字节切片。
func reverse(arr []byte) {
	// 使用双指针技术从两端开始向中间遍历并交换元素
	i, j := 0, len(arr)-1
	for i < j {
		// 交换两端的元素
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

// min 函数用于返回两个整数中的最小值。
// 参数 a, b 是待比较的两个整数。
func min(a, b int) int {
	// 如果 a 小于 b，则返回 a，否则返回 b
	if a < b {
		return a
	}
	return b
}
