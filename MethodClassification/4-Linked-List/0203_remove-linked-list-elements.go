package main

import "fmt"

// 主函数
func main() {
	// 创建测试链表: 1 -> 2 -> 6 -> 3 -> 4 -> 5 -> 6
	values := []int{1, 2, 6, 3, 4, 5, 6}
	head := CreateList(values)

	fmt.Print("原链表: ")
	Print(head) // 输出: 1 2 6 3 4 5 6

	// 移除值为 6 的节点
	head = removeElements(head, 6)

	fmt.Print("移除值为 6 后的链表: ")
	Print(head) // 输出: 1 2 3 4 5

}

// ListNode 定义链表节点
type ListNode struct {
	Val  int       // 值
	Next *ListNode // 下一个节点
}

// removeElements 移除链表中所有值为 val 的节点
// 参数:
//
//	head: 链表头节点
//	val: 需要移除的节点值
//
// 返回值:
//
//	移除指定值后的新链表头节点
func removeElements(head *ListNode, val int) *ListNode {
	// 如果链表为空，则直接返回
	if head == nil {
		return nil
	}
	// 处理头节点
	if head.Val == val {
		return removeElements(head.Next, val)
	}
	// 遍历链表，找到指定值的节点
	current := head
	for current.Next != nil {
		if current.Next.Val == val {
			// 移除节点
			current.Next = current.Next.Next
		} else {
			// 继续遍历下一个节点
			current = current.Next
		}
	}
	return head
}

// CreateList 创建链表
// 参数:
//
//	values: 用于创建链表的值数组
//
// 返回值:
//
//	创建的链表头节点
func CreateList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, value := range values[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return head
}

// Print 方法打印链表的所有元素
// 参数:
//
//	head: 链表头节点
func Print(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}
