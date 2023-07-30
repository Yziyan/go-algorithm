// @Author: Ciusyan 2023/7/30

package 题目

// https://leetcode.cn/problems/lMSNwu/

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 先翻转两条链表
	l1 = reverseList(l1)
	l2 = reverseList(l2)

	// 准备一个虚拟头结点
	dummyHead := &ListNode{}
	tail := dummyHead

	// 代表进位
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

		sum := v1 + v2 + carry
		if sum >= 10 {
			carry = 1
			// 只取个位数
			sum %= 10
		} else {
			// 说明需要将进位置为 0
			carry = 0
		}

		// 新建节点
		node := &ListNode{Val: sum}
		tail.Next = node
		tail = node
	}
	// 再判断最后一次相加，进位是否为1
	if carry == 1 {
		tail.Next = &ListNode{Val: carry}
	}

	// 将结果也需要逆序
	return reverseList(dummyHead.Next)
}

// 翻转链表
func reverseList(head *ListNode) *ListNode {

	var (
		prev *ListNode
		next *ListNode
	)

	// 利用头插法，进行逆序
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
