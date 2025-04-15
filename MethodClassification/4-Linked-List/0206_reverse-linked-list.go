package main

import "fmt"

func main() {
	first := new(ListNode)
	first.Val = 1
	second := new(ListNode)
	second.Val = 2
	third := new(ListNode)
	third.Val = 6
	fourth := new(ListNode)
	fourth.Val = 3
	fifth := new(ListNode)
	fifth.Val = 4
	sixth := new(ListNode)
	sixth.Val = 5
	seventh := new(ListNode)
	seventh.Val = 6
	first.Next = second
	second.Next = third
	third.Next = fourth
	fourth.Next = fifth
	fifth.Next = sixth
	sixth.Next = seventh
	first = reverseList(first)
	for {
		fmt.Println(first.Val)
		if first.Next == nil {
			break
		}
		first = first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList 反转一个单链表。
// 该函数采用迭代的方式，通过改变链表中每个节点的Next指针来实现链表的反转。
// 参数:
//
//	head *ListNode: 链表的头节点。
//
// 返回值:
//
//	*ListNode: 反转后的链表的头节点。
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// leetcode206_反转链表
// reverseList0 反转一个单链表。
// 该函数使用迭代的方式，通过创建新节点并将其指向已有节点的方式来实现链表的反转。
// 参数:
//
//	head *ListNode: 链表的头节点。
//
// 返回值:
//
//	*ListNode: 反转后的链表的头节点。
func reverseList0(head *ListNode) *ListNode {
	var res *ListNode
	for {
		if head == nil {
			break
		}
		res = &ListNode{head.Val, res}
		head = head.Next
	}
	return res
}
