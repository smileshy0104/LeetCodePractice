package main

import "fmt"

// ListNode 表示链表中的一个节点，包含一个整数值 Val 和指向下一个节点的指针 Next。
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle 检查链表中是否存在环。
// 使用哈希表 valueMap 来存储已经访问过的节点，如果遇到已经存在于 valueMap 中的节点，则说明存在环。
// 参数 head 是链表的头节点指针。
// 返回值为布尔类型，如果链表中存在环则返回 true，否则返回 false。
func hasCycle(head *ListNode) bool {
	// 如果链表为空或者只有一个节点，则不存在环，直接返回 false。
	if head == nil || head.Next == nil {
		return false
	}
	// 初始化一个 map 来存储访问过的节点。
	valueMap := make(map[*ListNode]bool)
	// 遍历链表。
	for head != nil {
		// 如果当前节点已经存在于 map 中，说明遇到了环，返回 true。
		if valueMap[head] {
			return true
		}
		// 将当前节点添加到 map 中，表示已经访问过。
		valueMap[head] = true
		// 移动到下一个节点。
		head = head.Next
	}
	// 如果遍历结束都没有遇到环，返回 false。
	return false
}

func main() {
	// 构造一个带有环的链表：3 -> 2 -> 0 -> -4 -> 2 ...
	first := ListNode{Val: 3}
	second := ListNode{Val: 2}
	third := ListNode{Val: 0}
	fourth := ListNode{Val: -4}
	first.Next = &second
	second.Next = &third
	third.Next = &fourth
	fourth.Next = &second

	// 输出环的起始节点
	fmt.Println(hasCycle(&first))
}
