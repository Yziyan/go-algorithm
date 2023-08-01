// @Author: Ciusyan 2023/7/31

package tree

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil || len(preorder) != len(inorder) {
		return nil
	}

	// 把中序遍历的所有结果，都放在 map 中
	m := make(map[int]int, len(inorder))
	for i, v := range inorder {
		m[v] = i
	}

	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1, m)
}

// 利用 [begin, end) 构建出整棵树
func build(preorder, inorder []int, l1, r1, l2, r2 int, m map[int]int) *TreeNode {
	if l1 > r1 {
		// 说明遇到了 空树
		return nil
	}

	// 构建节点
	root := &TreeNode{Val: preorder[l1]}
	if l1 == r1 {
		// 说明只有一个节点了，返回 root
		return root
	}

	// 先找到根节点在中序遍历中的位置
	idx := m[preorder[l1]]

	// 递归的构建左右子树
	root.Left = build(preorder, inorder, l1+1, l1+idx-l2, l2, idx-1, m)
	root.Right = build(preorder, inorder, l1+idx-l2+1, r1, idx+1, r2, m)

	return root
}
