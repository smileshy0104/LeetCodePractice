package main

import "fmt"

// ListNode 定义链表节点
type ListNode struct {
	Val  int       // 值
	Next *ListNode // 下一个节点
}

// removeNthFromEnd 从链表中删除倒数第n个节点
// head *ListNode: 链表头节点
// n int: 倒数第n个节点
// 返回 *ListNode: 删除节点后的链表头节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 创建一个临时节点，方便处理头节点删除的情况
	temp := &ListNode{Next: head}
	// 初始化快慢指针
	fast, slow := temp, temp
	// 快指针先移动n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 快慢指针同时移动，直到快指针到达链表末尾
	// 这样慢指针就指向了倒数第n个节点的前一个节点
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 删除倒数第n个节点
	slow.Next = slow.Next.Next
	// 返回处理后的链表头节点
	return temp.Next
}

// removeNthFromEnd0 从单链表中删除倒数第n个节点。
// 它首先使用两个遍历来确定链表的总长度，然后找到需要删除的节点的前一个节点，并进行删除操作。
// 参数:
//
//	head - 单链表的头节点指针。
//	n - 倒数第n个节点，即要删除的节点的位置。
//
// 返回值:
//
//	*ListNode - 删除节点后的链表头节点指针。
func removeNthFromEnd0(head *ListNode, n int) *ListNode {
	// 创建一个临时节点，指向头节点，简化边界情况处理。
	temp := &ListNode{Next: head}
	// 初始化当前节点指针为临时节点，用于遍历链表。
	cur := temp
	// 初始化链表总长度计数器。
	total := 0
	// 第一次遍历链表，计算链表总长度。
	for cur.Next != nil {
		cur = cur.Next
		total++
	}
	// 重置当前节点指针为临时节点，准备第二次遍历。
	cur = temp
	// 初始化已遍历节点计数器。
	count := 0
	// 第二次遍历链表，寻找要删除的节点的前一个节点。
	for cur.Next != nil {
		// 当找到要删除的节点的前一个节点时，进行删除操作。
		if total-n == count {
			cur.Next = cur.Next.Next
			break
		}
		// 更新已遍历节点计数器。
		cur = cur.Next
		count++
	}
	// 返回删除节点后的链表头节点指针。
	return temp.Next
}

func main() {
	// 初始化链表
	first := &ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	// 删除倒数第3个节点
	first = removeNthFromEnd0(first, 3)
	// 遍历并打印链表
	for first != nil {
		fmt.Println(first.Val)
		first = first.Next
	}
}
