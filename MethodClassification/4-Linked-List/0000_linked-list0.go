package main

import "fmt"

// ListNode 定义链表节点
type ListNode struct {
	Val  int       // 值
	Next *ListNode // 下一个节点
}

// LinkedList 定义链表结构
type LinkedList struct {
	Head *ListNode // 链表头
}

// Add 方法在链表末尾添加元素
func (l *LinkedList) Add(value int) {
	// 创建新节点
	newNode := &ListNode{Val: value}
	// 如果链表为空，则将新节点作为头节点
	if l.Head == nil {
		l.Head = newNode
		return
	}
	// 遍历链表，找到最后一个节点
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	// 将新节点链接到链表末尾
	current.Next = newNode
}

// Remove 方法移除链表中的指定值
func (l *LinkedList) Remove(value int) {
	// 如果链表为空，则直接返回
	if l.Head == nil {
		return
	}
	// 处理头节点
	if l.Head.Val == value {
		// 头节点移除
		l.Head = l.Head.Next
		return
	}
	// 遍历链表，找到指定值的节点
	current := l.Head
	for current.Next != nil && current.Next.Val != value {
		current = current.Next
	}
	// 找到并移除节点
	if current.Next != nil {
		// 移除节点
		current.Next = current.Next.Next
	}
}

// Print 方法打印链表的所有元素
func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	list := &LinkedList{}

	// 添加元素
	list.Add(1)
	list.Add(2)
	list.Add(3)
	fmt.Print("链表元素: ")
	list.Print() // 输出: 1 2 3

	// 移除元素
	list.Remove(2)
	fmt.Print("移除元素 2 后的链表: ")
	list.Print() // 输出: 1 3

	// 移除头节点
	list.Remove(1)
	fmt.Print("移除元素 1 后的链表: ")
	list.Print() // 输出: 3
}
