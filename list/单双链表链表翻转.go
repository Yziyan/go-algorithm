// @Author: Ciusyan 2023/7/22

package list

// ListNode 单向链表节点
type ListNode struct {
	value int64
	next  *ListNode
}

func NewListNode(value int64) *ListNode {
	return &ListNode{value: value}
}

// ReverseList 单向链表节点
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return nil
	}

	var (
		prev *ListNode
		next *ListNode
	)

	for head != nil {
		next = head.next
		head.next = prev
		prev = head
		head = next
	}

	return prev
}

func ReverseDoubleList(head *DoubleListNode) *DoubleListNode {
	if head == nil {
		return nil
	}

	var (
		prev *DoubleListNode
		next *DoubleListNode
	)

	for head != nil {
		next = head.next
		head.next = prev
		head.prev = next
		prev = head
		head = next
	}

	return prev
}
