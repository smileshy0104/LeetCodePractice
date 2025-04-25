package main

import "fmt"

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
}

func reverseStr(s string, k int) string {
	var res []byte
	for i := 0; i < len(s); i += 2 * k {

	}
	return string(res)
}

func reverse(arr []byte) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
