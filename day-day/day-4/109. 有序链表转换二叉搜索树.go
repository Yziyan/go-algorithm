// @Author: Ciusyan 2023/9/3

package day_4

// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		// 说明只有一个节点
		return &TreeNode{Val: head.Val}
	}

	// 找到中间节点的前驱节点，因为之后需要将中间节点与前面的断掉后，才去构建左子树
	midPrev := midNodePrev(head)
	root := &TreeNode{Val: midPrev.Next.Val}
	// 先去构建右子树，就不用记录状态了，从中间节点的后一个开始构建
	root.Right = sortedListToBST(midPrev.Next.Next)
	// 再去构建左子树，（从头开始）但是在构建前，需要将中间节点与前面断掉
	midPrev.Next = nil
	root.Left = sortedListToBST(head)

	return root
}

// 寻找中间节点的前驱节点（快慢指针）
func midNodePrev(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var (
		slow = head
		// 如果需要找中间的前驱，那么只需要初始时，fast 多走一步即可
		fast = head.Next.Next
	)

	for fast != nil && fast.Next != nil {
		// 慢指针每次走一步，快指针每次走俩步
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 当循环退出时。慢指针就是 midPrev
	return slow
}
