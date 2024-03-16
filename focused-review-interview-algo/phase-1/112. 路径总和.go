// @Author: Ciusyan 3/16/24

package phase_1

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		// 没有在叶子节点返回，那么说明这里就不用看了
		return false
	}

	if root.Left == nil && root.Right == nil {
		// 说明到叶子节点了，才需要判断是否满足条件
		if root.Val == targetSum {
			// 说明刚好能凑出
			return true
		}
		return false
	}

	// 否则往左右子树去看看能否有路径
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
