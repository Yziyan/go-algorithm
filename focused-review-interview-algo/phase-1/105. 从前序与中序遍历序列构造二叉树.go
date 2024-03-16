// @Author: Ciusyan 3/16/24

package phase_1

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}
	l := len(preorder)
	cache := make(map[int]int, l)
	for i, v := range inorder {
		cache[v] = i
	}

	return build(preorder, inorder, 0, l-1, 0, l-1, cache)
}

func build(preorder, inorder []int, l1, r1 int, l2, r2 int, cache map[int]int) *TreeNode {
	if l1 > r1 {
		return nil
	}

	root := &TreeNode{Val: preorder[l1]}
	if l1 == r1 {
		return root
	}

	idx := cache[preorder[l1]]

	// 构建左右子树
	root.Left = build(preorder, inorder, l1+1, idx-l2+l1, l2, idx-1, cache)
	root.Right = build(preorder, inorder, idx-l2+l1+1, r1, idx+1, r2, cache)

	return root
}
