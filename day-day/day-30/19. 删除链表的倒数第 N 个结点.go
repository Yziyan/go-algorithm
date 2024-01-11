// @Author: Ciusyan 1/11/24

package day_30

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		return nil
	}

	var (
		// 准备快慢两个指针
		slow = head
		fast = head
	)

	// 让他们的初始距离相差 n + 1，因为要找到待删除节点的前一个节点
	for i := 0; i < n+1; i++ {
		if fast == nil && i == n {
			// 说明第一个节点就是倒数第 n 个节点
			return head.Next
		}
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 现在 slow 就是倒数第 n+1 个节点
	slow.Next = slow.Next.Next

	return head
}
