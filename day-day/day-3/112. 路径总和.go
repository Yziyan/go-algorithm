// @Author: Ciusyan 2023/9/2

package day_3

// https://leetcode.cn/problems/path-sum/

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// 递归基：到达了叶子节点
	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			// 说明找到了一个
			return true
		}
	}

	// 将目标值减去当前节点的值，再往左右子树搜索
	nS := targetSum - root.Val

	// 往左右子树搜索
	return hasPathSum(root.Left, nS) || hasPathSum(root.Right, nS)
}
