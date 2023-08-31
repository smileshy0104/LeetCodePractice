package main

import (
	Array "LeetCode/Practice/Array"
	TwoPointer "LeetCode/Practice/TwoPoint"
	"fmt"
)

// 数组的题目
func array() {
	nums := []int{1, 3, 5, 6}
	target := 5
	// 704. 二分查找
	index0 := Array.BinarySearch1(nums, target)
	fmt.Println("The index is:", index0)

	// 35. 搜索插入位置
	index := Array.SearchInsert(nums, target)
	fmt.Println("The index is:", index)

	// 34. 在排序数组中查找元素的第一个和最后一个位置
	nums1 := []int{5, 7, 7, 8, 8, 10}
	target1 := 8
	//target2 := 6
	searchRange := Array.SearchRange(nums1, target1)
	fmt.Println("The index is:", searchRange)

	// 69. x 的平方根
	x := 8
	sqrt := Array.MySqrt(x)
	fmt.Println("The index is:", sqrt)

	// 367. 有效的完全平方数
	Isbool := Array.IsPerfectSquare(14)
	fmt.Println("The bool is:", Isbool)
}

func twoPointer() {
	nums0 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val0 := 2
	length := TwoPointer.RemoveElement(nums0, val0)
	fmt.Println("the length is:", length)

	nums := []int{3, 2, 2, 3}
	val := 3
	length0 := TwoPointer.RemoveElement(nums, val)
	fmt.Println("the length is:", length0)

	s := "ab#c"
	t := "ad#c"
	Isbool := TwoPointer.BackspaceCompare(s, t)
	fmt.Println("The bool is:", Isbool)

	nums1 := []int{-4, -1, 0, 3, 10}
	sored := TwoPointer.SortedSquares1(nums1)
	fmt.Println("The sored is:", sored)
}
func main() {
	//关于数组的函数
	array()
	//关于双指针的函数
	twoPointer()
}
