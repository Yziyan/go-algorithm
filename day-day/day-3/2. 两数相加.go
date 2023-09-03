// @Author: Ciusyan 2023/9/2

package day_3

// https://leetcode.cn/problems/add-two-numbers/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		if l1 == nil {
			return l2
		}
		return l1
	}

	// 准备一条新的链表，将他们串起来
	dummyHead := &ListNode{}
	tail := dummyHead
	carry := 0

	for l1 != nil || l2 != nil {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		v := v1 + v2 + carry
		if v >= 10 {
			v %= 10
			carry = 1
		} else {
			carry = 0
		}

		// 新建一个节点
		node := &ListNode{Val: v}
		tail.Next = node
		tail = node
	}

	// 如果进位还有数字，那么说明还需要加一个节点到末尾
	if carry != 0 {
		tail.Next = &ListNode{Val: 1}
	}

	return dummyHead.Next
}
