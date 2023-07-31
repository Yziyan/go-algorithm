// @Author: Ciusyan 2023/7/31

package tree

// https://leetcode.cn/problems/same-tree/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		// 说明两个都是空树，肯定相等
		return true
	}

	if p == nil || q == nil {
		// 说明有一个是空树，肯定不相等
		return false
	}

	// 判断当前节点是否相等 && 左子树是否相等 && 右子树是否相等
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
