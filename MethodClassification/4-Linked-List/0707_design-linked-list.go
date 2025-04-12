package main

import "fmt"

// ListNode 定义链表节点
type ListNode struct {
	Val  int       // 节点存储的值
	Next *ListNode // 指向下一个节点的指针
}

// MyLinkedList 实现了一个单链表的数据结构
type MyLinkedList struct {
	head *ListNode // 头节点指针
}

// Constructor 创建并返回一个新的链表实例
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// Get 获取链表中第 index 个节点的值。如果索引无效，则返回 -1
func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}
	current := this.head // 使用一个新的指针来遍历链表
	for i := 0; i < index; i++ {
		if current == nil { // 检查当前节点是否为空
			return -1
		}
		current = current.Next
	}
	if current == nil { // 再次检查当前节点是否为空
		return -1
	}
	return current.Val // 返回当前节点的值
}

// AddAtHead 在链表的第一个元素之前添加一个值为 val 的新节点
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{Val: val, Next: this.head}
	this.head = newNode
}

// AddAtTail 在链表的最后一个元素之后添加一个值为 val 的新节点
func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{Val: val}
	if this.head == nil {
		this.head = newNode
		return
	}
	current := this.head
	// 遍历链表，找到最后一个节点
	for current.Next != nil {
		current = current.Next
	}
	// 将新节点链接到链表末尾
	current.Next = newNode
}

// AddAtIndex 在链表中的第 index 个节点之前添加一个值为 val 的新节点
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}
	newNode := &ListNode{Val: val}
	// 如果索引为0，则将新节点添加到链表头部
	if index == 0 {
		newNode.Next = this.head
		this.head = newNode
		return
	}
	current := this.head
	// 遍历链表，找到第 index-1 个节点
	for i := 0; i < index-1; i++ {
		if current == nil {
			return
		}
		current = current.Next
	}
	// 如果当前节点为空，则返回
	if current == nil {
		return
	}
	// 将新节点链接到当前节点之后
	newNode.Next = current.Next
	current.Next = newNode
}

// DeleteAtIndex 删除链表中的第 index 个节点
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}
	// 如果索引为0，则将头节点指向下一个节点
	if index == 0 {
		if this.head != nil {
			this.head = this.head.Next
		}
		return
	}
	current := this.head
	// 遍历链表，找到第 index-1 个节点
	for i := 0; i < index-1; i++ {
		if current == nil {
			return
		}
		current = current.Next
	}
	// 如果当前节点为空或当前节点的下一个节点为空，则返回
	if current == nil || current.Next == nil {
		return
	}
	// 将当前节点的下一个节点链接到当前节点的下下个节点
	current.Next = current.Next.Next
}

// Print 打印链表
func (this *MyLinkedList) Print() {
	current := this.head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

// 测试函数
func main() {
	// 创建链表
	linkedList := Constructor()

	// 添加节点
	linkedList.AddAtHead(1)     // 链表: 1
	linkedList.AddAtTail(3)     // 链表: 1 -> 3
	linkedList.AddAtIndex(1, 2) // 链表: 1 -> 2 -> 3
	linkedList.Print()          // 输出: 1 2 3

	// 获取节点值
	fmt.Println(linkedList.Get(1)) // 输出: 2
	fmt.Println(linkedList.Get(0)) // 输出: 1
	fmt.Println(linkedList.Get(2)) // 输出: 3
	fmt.Println(linkedList.Get(3)) // 输出: -1 (不存在)

	// 删除节点
	linkedList.DeleteAtIndex(1) // 链表: 1 -> 3
	linkedList.Print()          // 输出: 1 3

	linkedList.DeleteAtIndex(0) // 链表: 3
	linkedList.Print()          // 输出: 3

	linkedList.DeleteAtIndex(0) // 链表: 空
	linkedList.Print()          // 输出: (空)
}
