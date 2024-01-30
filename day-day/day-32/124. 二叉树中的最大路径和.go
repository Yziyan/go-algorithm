// @Author: Ciusyan 1/30/24

package day_32

import "math"

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/

func maxPathSum(root *TreeNode) int {
	// 使用二叉树的递归套路，需要收集的信息
	type info struct {
		maxPathSum         int // 整棵树的最大路径和
		maxPathSumFromRoot int // 从根节点出发的最大路径和
	}

	var process func(root *TreeNode) *info
	process = func(root *TreeNode) *info {
		if root == nil {
			return &info{
				maxPathSum: math.MinInt,
			}
		}

		// 先跟左右子树要 Info 信息
		leftInfo := process(root.Left)
		rightInfo := process(root.Right)

		var (
			// 需要构造当前这颗数的 Info 信息，默认为根节点的值
			maxPathSum         = root.Val
			maxPathSumFromRoot = root.Val
		)
		// 从头出发
		maxPathSumFromRoot = max(
			root.Val,                              // 只包含根节点
			root.Val+leftInfo.maxPathSumFromRoot,  // 包含根节点和左子树
			root.Val+rightInfo.maxPathSumFromRoot, // 包含根节点和右子树
		)

		// 求解出最大路径和
		maxPathSum = max(root.Val, maxPathSumFromRoot, leftInfo.maxPathSum, rightInfo.maxPathSum,
			root.Val+leftInfo.maxPathSumFromRoot+rightInfo.maxPathSumFromRoot, // 包含左+根+右
		)

		return &info{
			maxPathSum:         maxPathSum,
			maxPathSumFromRoot: maxPathSumFromRoot,
		}
	}

	// 求解出 root 的最大路径和
	return process(root).maxPathSum
}
