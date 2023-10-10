// @Author: Ciusyan 10/10/23

package day_19

// Encode N叉树 -> 二叉树
func Encode(nRoot *NTreeNode) *TreeNode {
	if nRoot == nil {
		return nil
	}

	// 新建二叉树的根节点
	root := &TreeNode{Val: nRoot.Val}
	// 将所有 children 挂载到左子树的右边
	root.Left = en(nRoot.children)
	return root
}

func en(children []*NTreeNode) *TreeNode {

	var (
		root *TreeNode
		cur  *TreeNode
	)

	for _, child := range children {
		// 先将当前 child 转换成二叉树的节点
		node := &TreeNode{Val: child.Val}
		if root == nil {
			// 先建立第一个左子树的节点
			root = node
		} else {
			// 说明现在全部是挂在左子树的右子树上
			cur.Right = node
		}
		// 先深度优先把 cur 的子节点全部挂载到 cur.left
		cur = node
		cur.Left = en(child.children)
	}

	return root
}

// Decode 二叉树 -> N叉树
func Decode(root *TreeNode) *NTreeNode {
	if root == nil {
		return nil
	}

	return &NTreeNode{
		Val: root.Val,
		// root 的子节点全部在左子树的右边
		children: de(root.Left),
	}
}

// 返回 root 的子节点
func de(root *TreeNode) []*NTreeNode {

	children := make([]*NTreeNode, 0, 1)
	for root != nil {
		// 先把打头的子节点建好（也要先携带上它的所有子节点）
		child := &NTreeNode{
			Val:      root.Val,
			children: de(root.Left),
		}
		// 加入 children 中
		children = append(children, child)
		// 然后往右边去加入下一个 child
		root = root.Right
	}

	return children
}
