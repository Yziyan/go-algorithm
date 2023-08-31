// @Author: Ciusyan 2023/8/31

package day_1

type ListNode struct {
	Val  int
	Next *ListNode
}

type DoubleListNode struct {
	Val  int
	Next *DoubleListNode
	Prev *DoubleListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		prev *ListNode
		next *ListNode
	)

	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func ReverseDoubleList(head *DoubleListNode) *DoubleListNode {
	if head == nil {
		return head
	}

	var (
		prev *DoubleListNode
		next *DoubleListNode
	)

	for head != nil {
		next = head.Next
		head.Next = prev
		head.Prev = next
		prev = head
		head = next
	}

	return nil
}

func RemoveListValue(head *ListNode, num int) *ListNode {
	if head == nil {
		return head
	}

	remove := func(head *ListNode, num int) *ListNode {
		// 跳过一开始重复的值
		for head != nil && head.Val == num {
			head = head.Next
		}

		cur := head
		prev := head

		for cur != nil {
			if cur.Val == num {
				prev.Next = cur.Next
			} else {
				prev = cur
			}

			cur = cur.Next
		}

		return head
	}

	return remove(head, num)
}
