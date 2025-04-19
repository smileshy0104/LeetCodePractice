package main

import "fmt"

// ListNode 表示链表中的一个节点，包含一个整数值 Val 和指向下一个节点的指针 Next。
type ListNode struct {
	Val  int
	Next *ListNode
}

// detectCycle 检测链表中是否存在环，并返回环的起始节点。
// 参数:
//   - head: 链表的头节点 (*ListNode)
//
// 返回值:
//   - 如果存在环，则返回环的起始节点；如果不存在环，则返回 nil。
//
// 该方法使用哈希表存储访问过的节点，时间复杂度为 O(n)，空间复杂度为 O(n)。
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	valueMap := make(map[*ListNode]bool)
	for head != nil {
		// 如果当前节点已经存在于哈希表中，则说明存在环，且当前节点为环的起始节点。
		if valueMap[head] {
			return head
		}
		valueMap[head] = true
		head = head.Next
	}
	return nil
}

// detectCycle0 检测链表中是否存在环，并返回环的起始节点。
// 参数:
//   - head: 链表的头节点 (*ListNode)
//
// 返回值:
//   - 如果存在环，则返回环的起始节点；如果不存在环，则返回 nil。
//
// 该方法使用 Floyd 判圈算法（快慢指针法），时间复杂度为 O(n)，空间复杂度为 O(1)。
func detectCycle0(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head

	// 使用快慢指针遍历链表，如果存在环，则快慢指针最终会相遇。
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	// 如果快指针到达链表末尾或其下一个节点为空，则说明不存在环。
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 将其中一个指针重新指向头节点，然后两个指针以相同速度前进，再次相遇时即为环的起始节点。
	for slow != head {
		slow = slow.Next
		head = head.Next
	}
	return slow
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
	fmt.Println(detectCycle0(&first))
}
