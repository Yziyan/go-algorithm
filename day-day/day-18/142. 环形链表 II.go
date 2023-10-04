// @Author: Ciusyan 10/4/23

package day_18

// https://leetcode.cn/problems/linked-list-cycle-ii

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 准备快慢指针
	var (
		slow = head.Next
		fast = head.Next.Next
	)

	// 如果链表有环， fast 一定能追上 slow
	for slow != fast {
		if fast == nil || fast.Next == nil {
			// 如果 fast 到了 nil 了，说明一定没有环
			return nil
		}
		// 快指针需要追赶，需要走快一步
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 来到这里，现在的 fast 和 slow 一定相遇了，说明一定有环
	fast = head
	for slow != fast {
		// 现在就一样的走，一定会在第一个入环节点相遇
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
