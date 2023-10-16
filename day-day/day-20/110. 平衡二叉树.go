// @Author: Ciusyan 10/16/23

package day_20

// https://leetcode.cn/problems/balanced-binary-tree/description/

func isBalanced(root *TreeNode) bool {

	var getInfo func(root *TreeNode) int
	getInfo = func(root *TreeNode) int {
		if root == nil {
			// 空树是平衡二叉树，高度是 0
			return 0
		}

		// 搜集左右子树的信息
		leftInfo := getInfo(root.Left)
		rightInfo := getInfo(root.Right)

		// 利用左右子树的信息，构建 root 所需要的信息
		if leftInfo == -1 || rightInfo == -1 {
			// 这里代表左右子树只要其中之一不是平衡二叉树，那么 root 也不是
			return -1
		}

		// 还有左右子树的高度差不能大于 1
		if leftInfo-rightInfo < -1 || leftInfo-rightInfo > 1 {
			return -1
		}

		// 来到这里，说明可以计算高度， 先默认是左子树的高度
		info := leftInfo
		if info < rightInfo {
			info = rightInfo
		}

		// Max(左高, 右高) + 1
		return info + 1
	}

	// 那么最终如果高度不是 -1，那么就代表是平衡二叉树
	return getInfo(root) != -1
}
