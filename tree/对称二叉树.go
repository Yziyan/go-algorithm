// @Author: Ciusyan 2023/7/31

package tree

// https://leetcode.cn/problems/symmetric-tree/

func isSymmetric(root *TreeNode) bool {
	// 第一次把对称轴上的两个节点比对
	return isSymmetric1(root, root)
}

// 判断 left 和 right 是否相等，也就是镜像是否相等
func isSymmetric1(left, right *TreeNode) bool {
	if left == nil && right == nil {
		// 左右都为 nil，说明肯定对称了
		return true
	}

	if left == nil || right == nil {
		// 有一个不为 nil，那么说明肯定不对称
		return false
	}

	// 当前左和右的值要相等 && 左边的左边 <-> 右边的右边 && 左边的右边 <-> 右边的左边
	return left.Val == right.Val && isSymmetric1(left.Left, right.Right) && isSymmetric1(left.Right, right.Left)
}
