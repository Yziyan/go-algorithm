// @Author: Ciusyan 2023/9/1

package day_2

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/merge-k-sorted-lists/

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	// 递归基
	if len(lists) == 1 {
		return lists[0]
	}

	// 采用分治的思想去求解
	mid := len(lists) >> 1
	leftList := mergeKLists(lists[:mid])
	rightList := mergeKLists(lists[mid:])

	// 最后合并两条链表
	return merge2Lists(leftList, rightList)
}

func merge2Lists(list1, list2 *ListNode) *ListNode {

	// 采用虚拟头结点的方式，将其串起来
	dummyHead := &ListNode{}
	tail := dummyHead

	for list1 != nil && list2 != nil {
		val := 0
		if list1.Val > list2.Val {
			val = list2.Val
			list2 = list2.Next
		} else {
			val = list1.Val
			list1 = list1.Next
		}

		// 新建节点
		node := &ListNode{Val: val}
		tail.Next = node
		tail = node
	}

	// 来到这里至少有一条链表走到了最后，将另外一条串起来
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummyHead.Next
}
