// @Author: Ciusyan 5/9/24

package cycle_10_5_7_5_11

// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
思路重复：
要对这个有序的列表排序成 BST，思路还是一样的。
找到中点，即代表根节点
中点往前，代表根节点的左子树
中点往后，代表根节点的右子树

因为这是链表，所以我们可以通过快慢指针寻找到中点。
中点左侧构建左子树，右侧构建右子树。
即：head -> ... -> midPrev -> mid -> midNext -> ...
root: mid.Val
root.Right = build(midNext, tail)

root.Left = build(head, midPrev)
*/

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
