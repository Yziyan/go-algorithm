// @Author: Ciusyan 5/9/24

package cycle_10_5_7_5_11

// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	// 找到中点前一个节点
	midPrev := &ListNode{Next: head}
	fast := head
	for fast != nil && fast.Next != nil {
		midPrev = midPrev.Next
		fast = fast.Next.Next
	}

	// 中点就是根节点
	root := &TreeNode{Val: midPrev.Next.Val}

	// 中点后方，就是右子树，先构建右子树就不需要记录节点了
	root.Right = sortedListToBST(midPrev.Next.Next)

	// 中点前方，就是左子树
	// 将中点前方断掉，就代表左子树的范围只到 mid
	midPrev.Next = nil
	root.Left = sortedListToBST(head)

	return root
}
