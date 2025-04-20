package main

import "fmt"

// 主函数
func main() {
	// 测试isAnagram函数
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}

// isAnagram检查两个字符串是否是字母异位词
// s和t是需要比较的两个字符串
// 返回值表示两个字符串是否是字母异位词
func isAnagram(s string, t string) bool {
	// 如果两个字符串长度不同，则它们不可能是字母异位词
	if len(s) != len(t) {
		return false
	}

	// 使用sMap记录字符串s中每个字符的出现次数
	sMap := make(map[rune]int)
	for _, v := range s {
		sMap[v]++
	}

	// 遍历字符串t，减少sMap中相应字符的计数
	// 如果sMap中某字符的计数为0，说明t中有字符不在s中，返回false
	for _, v := range t {
		if sMap[v] == 0 {
			return false
		}
		sMap[v]--
	}

	// 如果所有字符计数都减为0，说明s和t是字母异位词
	return true
}
