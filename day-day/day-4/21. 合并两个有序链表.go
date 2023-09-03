// @Author: Ciusyan 2023/9/3

package day_4

// https://leetcode.cn/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 准备一条新链表，将他们串起来
	dummyHead := &ListNode{}
	tail := dummyHead

	// 只要有一条链表不是 nil，就串起来
	for list1 != nil && list2 != nil {
		v := list1.Val
		if list1.Val > list2.Val {
			v = list2.Val
			list2 = list2.Next
		} else {
			list1 = list1.Next
		}

		// 新建一个节点
		node := &ListNode{Val: v}
		tail.Next = node
		tail = node
	}

	// 来到这里至少有一条链表是 nil 的，将另一条拼在 tail 后面
	if list1 == nil {
		tail.Next = list2
	} else {
		tail.Next = list1
	}

	return dummyHead.Next
}
