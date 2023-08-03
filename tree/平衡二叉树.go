// @Author: Ciusyan 2023/8/3

package tree

// https://leetcode.cn/problems/balanced-binary-tree/

func isBalanced(root *TreeNode) bool {
	return getInfo(root) != -1
}

// 返回 root 的高度，并且如果高度是 -1.那么代表这棵树不平衡
func getInfo(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 先收集左右子树的信息。
	leftInfo := getInfo(root.Left)
	rightInfo := getInfo(root.Right)

	if leftInfo == -1 || rightInfo == -1 {
		// 说明左右子树肯定有一颗不平衡
		return -1
	}

	// 还需要查看，左右高度差的绝对值是否大于 1
	abs := leftInfo - rightInfo
	if abs > 1 || abs < -1 {
		return -1
	}

	// 否则就返回正确的高度
	return 1 + max(leftInfo, rightInfo)
}
