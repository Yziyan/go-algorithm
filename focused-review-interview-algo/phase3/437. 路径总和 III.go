// @Author: Ciusyan 6/19/24

package phase3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/path-sum-iii/
func pathSum(root *TreeNode, targetSum int) int {
	var prs func(root *TreeNode, remain int) int
	prs = func(root *TreeNode, remain int) int {
		if root == nil {
			return 0
		}

		res := 0
		remain -= root.Val
		if remain == 0 {
			res++
		}

		res += prs(root.Left, remain)
		res += prs(root.Right, remain)

		return res
	}

	if root == nil {
		return 0
	}

	rootVal := prs(root, targetSum)
	rootVal += pathSum(root.Left, targetSum)
	rootVal += pathSum(root.Right, targetSum)

	return rootVal
}
