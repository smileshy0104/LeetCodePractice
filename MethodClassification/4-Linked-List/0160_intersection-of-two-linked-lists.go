package main

import "fmt"

func main() {
	a := ListNode{Val: 10}
	b := ListNode{Val: 99}
	c := ListNode{Val: 100}
	d := ListNode{Val: 80}

	a.Next = &b
	d.Next = &b
	b.Next = &c
	fmt.Println(getIntersectionNode0(&a, &d))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// getIntersectionNode 用于寻找两个链表的交点节点。
// 参数:
//
//	headA - 第一个链表的头节点 (*ListNode)
//	headB - 第二个链表的头节点 (*ListNode)
//
// 返回值:
//
//	如果两个链表有交点，返回交点的第一个公共节点；如果没有交点，返回 nil。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil { // 如果任一链表为空，则直接返回 nil
		return nil
	}

	// 计算链表 headA 的长度
	a := headA
	b := headB
	aLength := 0
	bLength := 0
	for a.Next != nil {
		aLength++
		a = a.Next
	}

	// 计算链表 headB 的长度
	for b.Next != nil {
		bLength++
		b = b.Next
	}

	// 确保两个链表从相同的起始位置开始遍历（消除长度差）
	step := 0
	a0 := headA
	b0 := headB
	if aLength > bLength {
		step = aLength - bLength
		for step > 0 {
			a0 = a0.Next
			step--
		}
	} else {
		step = bLength - aLength
		for step > 0 {
			b0 = b0.Next
			step--
		}
	}

	// 同时遍历两个链表，直到找到相同的节点或到达链表末尾
	for a0 != b0 {
		a0 = a0.Next
		b0 = b0.Next
	}

	// 返回交点节点（如果不存在交点，则返回 nil）
	return a0
}

// getIntersectionNode0 是一个函数，用于找到两个链表的交叉点。
// 它接受两个参数 headA 和 headB，分别代表链表 A 和链表 B 的头节点。
// 函数返回的是两个链表交叉的那个节点，如果没有交叉，则返回 nil。
func getIntersectionNode0(headA, headB *ListNode) *ListNode {
	// 使用一个哈希表 values 来记录链表 A 中节点的值。
	values := map[int]bool{}
	// 遍历链表 A，将每个节点的值存入哈希表中。
	for headA != nil {
		values[headA.Val] = true
		headA = headA.Next
	}
	// 遍历链表 B，检查每个节点的值是否在哈希表中。
	for headB != nil {
		// 如果在哈希表中找到了相同的值，说明这个节点是两个链表的交叉点。
		if values[headB.Val] {
			return headB
		}
		headB = headB.Next
	}
	// 如果没有找到交叉点，返回 nil。
	return nil
}
