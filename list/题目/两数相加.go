// @Author: Ciusyan 2023/7/30

package 题目

// https://leetcode.cn/problems/add-two-numbers/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// 进位
	carry := 0
	dummyHead := &ListNode{}
	tail := dummyHead

	// 只要有一个不为 nil，都可以继续相加
	for l1 != nil || l2 != nil {
		// 算出 V1
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		// 算出 V2
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		if sum >= 10 {
			carry = 1
			// 取出个位数
			sum = sum % 10
		} else {
			// 说明需要设置为 0
			carry = 0
		}

		// 构建新的链表
		node := &ListNode{Val: sum}
		tail.Next = node
		tail = node
	}
	// 如果进位是 1，那么还需要一个新的节点
	if carry == 1 {
		tail.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next
}
