// @Author: Ciusyan 9/28/23

package day_18

// https://leetcode.cn/problems/palindrome-linked-list/description/

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 先找到链表中点的前一个节点
	midPrev := getMiddlePrev(head)
	// 对中间节点进行翻转
	rightHead := reverse(midPrev.Next)

	var (
		// 左边的头和右边的头
		leftH  = head
		rightH = rightHead

		res = true
	)

	// 挨个比较
	for leftH != nil && rightH != nil {
		if leftH.Val != rightH.Val {
			res = false
			// 先别着急返回，还得去还原链表
			break
		}
		leftH = leftH.Next
		rightH = rightH.Next
	}

	// 还原链表
	midPrev.Next = reverse(rightHead)

	return res
}

func getMiddlePrev(head *ListNode) *ListNode {
	var (
		slow = head
		fast = head.Next.Next
	)

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// 翻转链表
func reverse(head *ListNode) *ListNode {

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
