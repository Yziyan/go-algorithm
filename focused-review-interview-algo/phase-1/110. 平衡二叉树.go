// @Author: Ciusyan 3/16/24

package phase_1

func isBalanced(root *TreeNode) bool {

	// 返回 root 的高度，如果返回 -1，代表不平衡
	var process func(root *TreeNode) int
	process = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		// 先取出左右子树的高
		leftH := process(root.Left)
		rightH := process(root.Right)
		if leftH == -1 || rightH == -1 {
			// 只要左右有一边是不平衡的，整体就是不平衡的
			return -1
		}

		div := leftH - rightH
		if div < -1 || div > 1 {
			// 说明高度差大于 1 了，不平衡了
			return -1
		}

		// 否则返回 root 的高
		return 1 + max(leftH, rightH)
	}

	// 如果返回了正常的高度，说明这棵树平衡
	return process(root) != -1
}
