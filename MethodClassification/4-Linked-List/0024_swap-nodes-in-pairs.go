package main

import "fmt"

func main() {
	first := &ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	first = swapPairs(first)
	for first != nil {
		fmt.Println(first.Val)
		first = first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// swapPairs 交换链表中的相邻节点
// 此函数接收一个链表的头节点指针作为输入，并返回交换相邻节点后的链表的头节点指针
// 该算法通过一系列指针操作来交换节点，无需额外数据结构，节省空间
func swapPairs(head *ListNode) *ListNode {
	// 创建一个临时节点，指向链表头部，作为遍历的起点和处理边界情况的便利
	temp := &ListNode{Next: head}
	// prev 用于记录当前处理的节点的前一个节点，初始为临时节点
	prev := temp
	// 遍历链表，每次处理一对相邻节点
	for head != nil && head.Next != nil {
		// 将一对相邻节点分别命名为 first 和 second，便于后续操作
		first, second := head, head.Next
		// 将前一个节点的 Next 指向 second，建立新的链接关系
		prev.Next = second
		// 交换 first 和 second 的位置，通过改变它们的 Next 指针
		first.Next, second.Next = second.Next, first
		// 更新 prev 和 head 指针，为下一次迭代做准备
		prev, head = first, first.Next
	}
	// 返回新的头节点，即临时节点的 Next 指向的节点
	return temp.Next
}
