// @Author: Ciusyan 2023/7/27

package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// 求解最小值
func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// 求解最大值
func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}
