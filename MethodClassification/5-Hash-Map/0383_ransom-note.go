package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aa", "aab"))
	fmt.Println(canConstruct("aaa", "aab"))
}

// canConstruct 检查是否可以通过拆解magazine中的字符来构造ransomNote字符串。
// 这个函数解决的问题是：给定两个字符串ransomNote和magazine，判断能否在不重复使用magazine中的字符的情况下，构造出ransomNote。
// 参数:
//
//	ransomNote - 需要构造的目标字符串。
//	magazine - 提供字符的源字符串。
//
// 返回值:
//
//	bool - 如果能够构造出目标字符串，则返回true；否则返回false。
func canConstruct(ransomNote string, magazine string) bool {
	// 使用map来统计magazine中每个字符的出现次数。
	value := make(map[rune]int)

	// 遍历magazine，统计每个字符的出现次数。
	for _, v := range magazine {
		value[v]++
	}

	// 遍历ransomNote，检查每个字符是否都在magazine中出现过，并且出现的次数不少于在ransomNote中的次数。
	for _, v := range ransomNote {
		// 如果ransomNote中的字符在magazine中没有出现过，则无法构造。
		if _, ok := value[v]; !ok {
			return false
		}
		// 减少该字符在magazine中的出现次数。
		value[v]--
		// 如果magazine中该字符的出现次数少于ransomNote中的次数，则无法构造。
		if value[v] < 0 {
			return false
		}
	}

	// 如果所有字符都可以在magazine中找到，并且出现次数符合要求，则返回true。
	return true
}
