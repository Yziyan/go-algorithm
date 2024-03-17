// @Author: Ciusyan 3/16/24

package phase_1

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	l := len(lists)

	if l == 1 {
		// 说明只有一个链表
		return lists[0]
	}

	mid := l >> 1

	// 分治思路
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])

	return mergeList(left, right)
}

// 合并左右两条升序链表
func mergeList(left, right *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tail := dummyHead

	for left != nil && right != nil {
		curVal := 0
		if left.Val > right.Val {
			curVal = right.Val
			right = right.Next
		} else {
			curVal = left.Val
			left = left.Next
		}

		// 新建节点接入
		newNode := &ListNode{Val: curVal}
		tail.Next = newNode
		tail = newNode
	}

	if left != nil {
		tail.Next = left
	} else {
		tail.Next = right
	}

	return dummyHead.Next
}
