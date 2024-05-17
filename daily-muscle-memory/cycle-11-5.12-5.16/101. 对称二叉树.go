// @Author: Ciusyan 5/16/24

package cycle_11_5_12_5_16

// https://leetcode.cn/problems/symmetric-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*
一棵树是否是对称二叉树，就是从根节点开始，比较左子树、和右子树是否是对称的。
并且对于每一个节点，都需要考虑左右子树是否是对称的。所以使用递归的思想即可。
传入两个节点，如果左右子树都是空的，说明是对称的。要是只有一个是 nil 的，就不是对称的。
若两个节点都不为 nil，首先比较这俩节点的值是否是相等的，相等后再看看，左左和右右 左右和右左 是否相等。
*/
func isSymmetric(root *TreeNode) bool {
	var process func(left, right *TreeNode) bool
	process = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			// 左右同时为 nil 说明对称
			return true
		}

		if left == nil || right == nil {
			// 要是只有其中一个为 nil 了。肯定不对称了
			return false
		}
		// 来到这里，说明左右都不为 nil 了，看看左右的值，并且递归的判断，「左左 和 右右」、「左右 和 右左」 是否是对称二叉树
		return left.Val == right.Val && process(left.Right, right.Left) && process(left.Left, right.Right)
	}

	return process(root, root)
}
