// @Author: Ciusyan 2023/7/27

package tree

import "math"

// https://leetcode.cn/problems/minimum-distance-between-bst-nodes/description/

func MinDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}

	inorder(root)

	if res == math.MaxInt {
		return 0
	}

	return res
}

var (
	prev *TreeNode
	res  = math.MaxInt
)

func inorder(root *TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left)

	if prev != nil {
		res = min(res, root.Val-prev.Val)
	}

	prev = root

	inorder(root.Right)
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
