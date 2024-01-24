// @Author: Ciusyan 1/24/24

package day_32

// https://leetcode.cn/problems/validate-binary-search-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var (
		preNode *TreeNode // 遍历的前一个节点
	)

	var inorder func(root *TreeNode) bool
	inorder = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		// 先中遍历左子树
		isBst := inorder(root.Left)
		if !isBst {
			// 说明左子树不是 BST
			return false
		}

		// 再处理中序逻辑
		if preNode != nil {
			// 遍历的不是第一个节点才进行判断
			if preNode.Val >= root.Val {
				// 说明违反 BST 的特点了
				return false
			}
		}
		// 走的时候别忘记将当前节点给 preNode
		preNode = root

		// 最后中遍历右子树
		return inorder(root.Right)
	}

	// 进行中序遍历，遍历结束后，就知道了
	return inorder(root)
}
