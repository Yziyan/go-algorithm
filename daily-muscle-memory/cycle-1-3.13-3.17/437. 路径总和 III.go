// @Author: Ciusyan 3/15/24

package cycle_1_3_13_3_17

// https://leetcode.cn/problems/path-sum-iii/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// 先计算从根节点出发，能有的结果数
	res := process(root, targetSum)
	// 但是还需要看看，左子树、右子树中有多少满足条件的解法
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)

	return res
}

// 从 root 出发，能有多少种解法
func process(root *TreeNode, targetSum int) int {
	if root == nil {
		// 说明没有节点
		return 0
	}
	res := 0
	val := root.Val
	if val == targetSum {
		// 说明到目前为止，已经有一种方案可以找出路径和了
		res++
		// 但是还得继续往下寻找，因为节点的值可能有负数
	}
	// 去左右子树进行搜索
	res += process(root.Left, targetSum-val)  // 往左树找 targetSum-val
	res += process(root.Right, targetSum-val) // 往右树找 targetSum-val

	return res
}
