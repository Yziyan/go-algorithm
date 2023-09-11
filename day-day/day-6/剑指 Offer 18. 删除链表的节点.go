// @Author: Ciusyan 2023/9/11

package day_6

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	// 头节点的值是相等的
	if head.Val == val {
		return head.Next
	}

	// 来到这里，头节点肯定不是 val 了，遍历链表

	cur := head.Next
	prev := head

	for cur != nil {
		if cur.Val == val {
			// 改变指向
			prev.Next = cur.Next
			// 因为只有一个节点，指向后即可返回
			break
		}

		prev = cur
		cur = cur.Next
	}

	return head
}
