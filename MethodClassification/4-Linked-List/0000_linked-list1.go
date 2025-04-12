package main

import "fmt"

// DListNode 定义双链表节点
type DListNode struct {
	Val  int        // 值
	Prev *DListNode // 前驱节点
	Next *DListNode // 后继节点
}

// DoublyLinkedList 定义双链表结构
type DoublyLinkedList struct {
	Head *DListNode // 链表头
	Tail *DListNode // 链表尾
}

// Add 方法在链表末尾添加元素
func (dll *DoublyLinkedList) Add(value int) {
	newNode := &DListNode{Val: value}
	// 如果链表为空，则将新节点作为头节点
	if dll.Head == nil {
		// 将新节点作为头节点和尾节点
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		// 将新节点链接到链表末尾
		dll.Tail.Next = newNode
		// 更新新节点的前驱节点
		newNode.Prev = dll.Tail
		// 更新新节点为尾节点
		dll.Tail = newNode
	}
}

// Remove 方法移除链表中的指定值
func (dll *DoublyLinkedList) Remove(value int) {
	// 如果链表为空，则直接返回
	if dll.Head == nil {
		return
	}
	// 遍历链表，找到指定值的节点
	current := dll.Head
	for current != nil {
		// 找到并移除节点
		if current.Val == value {
			// 当前节点的前驱节点不为空，则更新前驱节点的后继节点
			if current.Prev != nil {
				// 更新前驱节点的后继节点
				current.Prev.Next = current.Next
			} else {
				// 移除头节点
				dll.Head = current.Next
			}
			// 当前节点的后继节点不为空，则更新后继节点的前驱节点
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				// 移除尾节点
				dll.Tail = current.Prev
			}
			return
		}
		// 继续遍历下一个节点
		current = current.Next
	}
}

// Print 方法打印链表的所有元素
func (dll *DoublyLinkedList) Print() {
	current := dll.Head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

// PrintReverse 方法反向打印链表的所有元素
func (dll *DoublyLinkedList) PrintReverse() {
	current := dll.Tail
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Prev
	}
	fmt.Println()
}

func main() {
	// 创建一个双链表
	dll := &DoublyLinkedList{}

	// 添加元素
	dll.Add(1)
	dll.Add(2)
	dll.Add(3)
	fmt.Print("双链表元素: ")
	dll.Print() // 输出: 1 2 3

	// 移除元素
	dll.Remove(2)
	fmt.Print("移除元素 2 后的双链表: ")
	dll.Print() // 输出: 1 3

	// 移除头节点
	dll.Remove(1)
	fmt.Print("移除元素 1 后的双链表: ")
	dll.Print() // 输出: 3

	// 添加更多元素
	dll.Add(4)
	dll.Add(5)
	fmt.Print("添加元素 4 和 5 后的双链表: ")
	dll.Print() // 输出: 3 4 5

	// 反向打印
	fmt.Print("反向打印双链表: ")
	dll.PrintReverse() // 输出: 5 4 3
}
