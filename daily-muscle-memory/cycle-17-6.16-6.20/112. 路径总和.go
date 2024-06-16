// @Author: Ciusyan 6/17/24

package cycle_16_6_11_6_15

// https://leetcode.cn/problems/path-sum/

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	remain := targetSum - root.Val
	if root.Left == nil && root.Right == nil {
		// 说明是叶子节点，是叶子节点才判断，是否满足 target 了
		if remain == 0 {
			return true
		}
		// 说明还有 remain
		return false
	}

	// 去左右搜索查看 remain
	return hasPathSum(root.Left, remain) || hasPathSum(root.Right, remain)
}
