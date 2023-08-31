// @Author: Ciusyan 2023/9/1

package day_2

// https://leetcode.cn/problems/same-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q == nil) {
		// 如果只有一棵树的根节点为 nil，那么肯定不相等
		return false
	}

	// 来到这里要么都为空，要么都不为空
	if p == nil && q == nil {
		return true
	}

	// 来到这里，看看根节点的值是否相等，再看看左子树和右子树是否相等。
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
