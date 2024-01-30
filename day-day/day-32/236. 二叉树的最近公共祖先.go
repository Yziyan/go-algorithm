// @Author: Ciusyan 1/30/24

package day_32

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		// 如果没有根节点，公共祖先也没有
		// 否则只要 p or q 有一个是根节点，那么最低公共祖先必然是 root
		return root
	}
	// 看看最近公共最先是否存在于左右子树
	existLeft := lowestCommonAncestor(root.Left, p, q)
	existRight := lowestCommonAncestor(root.Right, p, q)

	if existLeft != nil && existRight != nil {
		// 说明既存在于左子树，又存在于右子树，说明只能是 根节点。
		return root
	}

	// 来到这里，有三种情况：1.在左 2.在右 3.都不存在
	if existLeft != nil {
		// 说明在左边
		return existLeft
	}

	// 说明在右 or 不存在
	return existRight
}
