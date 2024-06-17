// @Author: Ciusyan 6/17/24

package cycle_16_6_11_6_15

// https://leetcode.cn/problems/path-sum/

/*
*
思路重复：
从头一路搜索到叶子节点即可，没走一个节点，就将 targetSum 减去当前节点的值。
如果到达了最终的叶子节点，值减为零了，就说明有这样的路径总和。
如果搜索完整棵树，都没有一个和，那么说明没有这样的路径、
*/

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
