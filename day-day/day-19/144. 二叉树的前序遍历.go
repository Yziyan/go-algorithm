// @Author: Ciusyan 10/5/23

package day_19

// https://leetcode.cn/problems/binary-tree-preorder-traversal/

// 非递归实现
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	// 准备一个遍历用的节点和一个栈
	cur := root
	stack := NewStack()

	// 只要 cur 或者 栈 不为空，就需要做操作
	for cur != nil || stack.Size() != 0 {
		if cur != nil {
			// 只要 cur 不为 nil，上来就操作 cur
			res = append(res, cur.Val)
			// 要是有右节点，将右节点入栈
			if cur.Right != nil {
				stack.Push(cur.Right)
			}

			// 然后一路向左
			cur = cur.Left
		} else {
			// 这里上来说明 cur 为 nil了，但是栈不为 nil，将栈里面的元素给 cur 继续遍历下去
			cur = stack.Pop()
		}
	}

	return res
}

// 非递归实现：类似层序遍历
func preorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	// 准备一个栈，并现将根结点入栈
	stack := NewStack()
	stack.Push(root)

	// 当栈不为空时，就去操作
	for stack.Size() != 0 {
		node := stack.Pop()
		// 收集结果
		res = append(res, node.Val)

		// 如果左右边不为空，将其加入栈中，如果要 头 左 右，就得先处理右边
		if node.Right != nil {
			stack.Push(node.Right)
		}

		if node.Left != nil {
			stack.Push(node.Left)
		}
	}

	return res
}

// 递归实现
func preorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	preorder(root, &res)

	return res
}

// 递归法实现的前序遍历
func preorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	// 先做事情
	*res = append(*res, root.Val)
	// 递归的前序遍历左右子树
	preorder(root.Left, res)
	preorder(root.Right, res)
}
