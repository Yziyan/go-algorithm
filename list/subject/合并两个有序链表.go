// @Author: Ciusyan 2023/7/30

package subject

// https://leetcode.cn/problems/merge-two-sorted-lists/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil || list2 == nil {
		if list1 == nil {
			return list2
		}
		return list1
	}

	// 合并两条链表，先新建一条链表
	dummyHead := &ListNode{}
	tail := dummyHead

	// 两个都不为 nil，才合并，要不然肯定是已经合并完了的
	for list1 != nil && list2 != nil {
		// 先默认
		val := 0
		if list1.Val > list2.Val {
			val = list2.Val
			list2 = list2.Next
		} else {
			// 说明 list1 要小
			val = list1.Val
			list1 = list1.Next
		}

		// 构建新链表
		node := &ListNode{Val: val}
		tail.Next = node
		tail = node
	}

	// 来到这里，要么是 list1 为 nil，要么是 list2 为 nil
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummyHead.Next
}
