// @Author: Ciusyan 2023/9/2

package day_3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/balanced-binary-tree/

func isBalanced(root *TreeNode) bool {
	return balance(root) != -1
}

// 返回值是 root 的高度，但若是 -1，则代表不平衡
func balance(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lH := balance(root.Left)
	rH := balance(root.Right)

	// 说明左右子树都不平衡了，那么整体肯定不平衡
	if lH == -1 || rH == -1 {
		return -1
	}

	// 说明左右子树的高度差超过 1 了，不平衡
	if lH-rH < -1 || lH-rH > 1 {
		return -1
	}

	if lH > rH {
		return lH + 1
	}

	return rH + 1
}
