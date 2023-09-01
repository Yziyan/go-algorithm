// @Author: Ciusyan 2023/9/1

package day_2

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) != len(inorder) {
		return nil
	}

	l := len(inorder)
	// 先将中序遍历打一个表，方便快速确定范围
	m := make(map[int]int, l)
	for i, v := range inorder {
		m[v] = i
	}

	// 根据范围去构建二叉树
	return build(preorder, inorder, 0, l-1, 0, l-1, m)
}

// 利用前序遍历得出根节点，再利用根节点去中序遍历中确定左右子树的范围
// 返回的 TreeNode：
//
//	在先序数组的范围是：preorder[pb ... pe]
//	在中序数组的范围是：inorder[ib ... ie]
func build(preorder, inorder []int, pb, pe int, ib, ie int, m map[int]int) *TreeNode {
	// pb 代表根节点所在位置，
	if pb > pe {
		// 说明遇到了 nil 节点，左子树肯定在右子树的前边
		return nil
	}

	rootVal := preorder[pb]
	root := &TreeNode{Val: rootVal}
	if pb == pe {
		// 说明只有一个根节点了
		return root
	}

	// 查询出根节点在中序遍历中的位置，
	//	mid 的左边就是左子树
	//	mid 的右边就是右子树
	mid := m[rootVal]

	// 左子树的范围在先序遍历中是 [pb+1, mid-ib+pb]，中序遍历中是 [ib, mid-1]
	root.Left = build(preorder, inorder, pb+1, mid-ib+pb, ib, mid-1, m)
	// 右子树的范围在先序遍历中是 [mid-ib+pb+1, pe]，中序遍历中是 [mid+1, ie]
	root.Right = build(preorder, inorder, mid-ib+pb+1, pe, mid+1, ie, m)

	return root
}
