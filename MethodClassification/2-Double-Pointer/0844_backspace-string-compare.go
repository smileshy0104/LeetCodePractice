package main

import (
	"fmt"
	"strings"
)

// main函数是程序的入口点。
func main() {
	// 比较两个字符串，考虑回退字符的影响
	compare := backspaceCompare("ab#c", "ad#c")
	fmt.Println(compare)

	// 比较两个不同的字符串，考虑回退字符的影响
	compare = backspaceCompare("ab##", "c#d#")
	fmt.Println(compare)

	// 比较两个截然不同的字符串，考虑回退字符的影响
	compare = backspaceCompare("a#c", "b")
	fmt.Println(compare)
}

// backspaceCompare比较两个字符串s和t，考虑回退字符的影响。
// 它使用check函数来处理每个字符串，以模拟回退字符的行为。
func backspaceCompare(s string, t string) bool {
	return check(s) == check(t)
}

// check函数处理输入字符串，模拟回退字符的行为。
// 它通过维护一个结果切片来构建最终的字符串表示形式。
func check(str string) string {
	var res []string
	for _, v := range str {
		if string(v) == "#" {
			// 如果当前字符是回退字符且结果切片不为空，则删除最后一个字符
			if len(res) > 0 {
				res = res[:len(res)-1] // 删除最后一个字符
			}
		} else {
			// 如果当前字符不是回退字符，则将其添加到结果切片中
			res = append(res, string(v))
		}
	}
	// 将结果切片连接成字符串并返回
	return strings.Join(res, "")
}
