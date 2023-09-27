// @Author: Ciusyan 9/27/23

package day_18

type ListNode struct {
	Val  int
	Next *ListNode
}

// MidOrUpMidNode 1、输入链表头节点，奇数长度返回中点，偶数长度返回上中点
func MidOrUpMidNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		// 慢指针指向头，快指针指向头结点的下一个节点
		slow = head
		fast = head.Next
	)

	for fast != nil && fast.Next != nil {
		// 慢指针每次走一步，快指针每次走两步
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

// MidOrDownMidNode 2、输入链表头节点，奇数长度返回中点，偶数长度返回下中点
func MidOrDownMidNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		// 快慢指针一开始都指向头结点
		slow = head
		fast = head
	)

	for fast != nil && fast.Next != nil {
		// 慢指针每次走一步，快指针每次走两步
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// MidOrUpMidPreNode 3、输入链表头节点，奇数长度返回中点前一个，偶数长度返回上中点前一个
func MidOrUpMidPreNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		// 慢指针指向 nil，快指针领先两个节点
		slow *ListNode
		fast = head.Next
	)

	for fast != nil && fast.Next != nil {
		// 慢指针每次走一步，快指针每次走两步
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// MidOrDownMidPreNode	4、输入链表头节点，奇数长度返回中点前一个，偶数长度返回下中点前一个
func MidOrDownMidPreNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		// 慢指针指向头结点，快指针领先一个节点
		slow *ListNode
		fast = head
	)

	for fast != nil && fast.Next != nil {
		// 慢指针每次走一步，快指针每次走两步
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
