// @Author: Ciusyan 2023/8/4

package tree

// https://leetcode.cn/problems/path-sum/

var isSum bool

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// 防止全局变量干扰，
	isSum = false
	pathDfs(root, 0, targetSum)

	return isSum
}

func pathDfs(root *TreeNode, preSum, targetSum int) {
	if root.Left == nil && root.Right == nil {
		// 代表是叶子节点，需要退出递归了，但是需要看是否满足条件
		if preSum+root.Val == targetSum {
			// 只要进入一次，那么就说明有这样的和
			isSum = true
		}
		return
	}

	// 来到这里，需要将路径和添加上
	preSum += root.Val
	if root.Left != nil {
		pathDfs(root.Left, preSum, targetSum)
	}

	if root.Right != nil {
		pathDfs(root.Right, preSum, targetSum)
	}
}
