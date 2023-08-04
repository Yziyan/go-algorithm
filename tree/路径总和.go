// @Author: Ciusyan 2023/8/4

package tree

// https://leetcode.cn/problems/path-sum/

// 递归解法
func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// 为叶子节点的时候，
	if root.Left == nil && root.Right == nil {
		// 如果最终有路径的值被减到了 root.Val，那么就不用遍历了，直接返回即可
		return targetSum == root.Val
	}

	// 去调用左右子树前，将路径和减掉当前的值
	return hasPathSum1(root.Left, targetSum-root.Val) || hasPathSum1(root.Right, targetSum-root.Val)
}

var isSum bool

// DFS 做法
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
