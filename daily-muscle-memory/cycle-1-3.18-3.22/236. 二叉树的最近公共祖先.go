// @Author: Ciusyan 3/18/24

package cycle_1_3_18_3_22

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 从根节点出发，查找 p和q 的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}

	// 从 Left 出发，查找 p和q 的最低公共祖先
	left := lowestCommonAncestor(root.Left, p, q)
	// 从 Right 出发，查找 p和q 的最低公共祖先
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		// 说明既存在左树，又存在右树，只能说明在 root
		return root
	}

	// 来到这里，至少有一边为 nil 了，返回对方，或者返回 nil
	if left == nil {
		return right
	}
	return left
}
