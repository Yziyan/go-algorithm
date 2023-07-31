// @Author: Ciusyan 2023/7/31

package 题目

// https://leetcode.cn/problems/merge-k-sorted-lists/

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	// 递归基
	if len(lists) == 1 {
		return lists[0]
	}

	// 采用二分，将链表一分为二，分别去合并两组链表
	mid := len(lists) >> 1
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])

	// 最后再合并两条链表
	return mergeList(left, right)
}

// 合并 left 和 right 两条链表
func mergeList(left, right *ListNode) *ListNode {

	// 采用虚拟头几点的方式，去拼接新的链表
	dummyHead := &ListNode{}
	tail := dummyHead

	// 两条都不为 nil，才值得往下拼接
	for left != nil && right != nil {
		val := 0
		if left.Val > right.Val {
			val = right.Val
			right = right.Next
		} else {
			val = left.Val
			left = left.Next
		}

		// 新建一条链表
		node := &ListNode{Val: val}
		tail.Next = node
		tail = node
	}

	// 来到这里，肯定有一条拼接完了，把未拼接的那一条拼接到末尾
	if left != nil {
		tail.Next = left
	} else {
		tail.Next = right
	}

	return dummyHead.Next
}
