// @Author: Ciusyan 3/16/24

package phase_1

func isSymmetric(root *TreeNode) bool {

	var process func(left, right *TreeNode) bool
	process = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			// 两个都是 nil
			return true
		}

		if left == nil || right == nil {
			// 只有一个是 nil
			return false
		}

		return left.Val == right.Val && process(left.Left, right.Right) && process(left.Right, right.Left)
	}

	return process(root, root)
}
