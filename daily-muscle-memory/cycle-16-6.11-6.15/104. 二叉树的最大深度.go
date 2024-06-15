// @Author: Ciusyan 6/14/24

package cycle_16_6_11_6_15

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/

/**
思路重复：
使用递归的方式求解即可。
root 的最大深度，等于 Left 和 Right 的最大深度 +1。
即可完成该题目，如果是 nil 树，最大深度肯定就是 0
*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
