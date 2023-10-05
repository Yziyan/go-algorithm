// @Author: Ciusyan 10/5/23

package day_19

// https://leetcode.cn/problems/binary-tree-postorder-traversal/description/

// 非递归实现版本 2：硬写一个
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	// 准备一个栈
	stack := NewStack()
	// 准备一个 cur，用于遍历树
	cur := root
	// 准备一个 prev，用于标记前一个真正被访问的节点
	var prev *TreeNode

	// 只要有一个不为 nil，就说明没遍历完
	for cur != nil || stack.Size() != 0 {
		if cur != nil {
			// 说明还可以往左走。但在走之前，需要将当前节点和右子树入栈
			// 但是必须先入当前的节点
			stack.Push(cur)
			if cur.Right != nil {
				// 有右入右，再入右
				stack.Push(cur.Right)
			}
			// 最后往左走
			cur = cur.Left
		} else {
			// 说明栈不为空，弹出栈顶元素，但是不能直接访问
			node := stack.Pop()

			// 为什么不能直接访问呢？因为还得看看 node 有没有后顾之忧
			// node 如果是叶子节点，说明没有后顾之忧
			// 如果 node 的子节点都已经访问过来了，那么也没有后顾之忧了
			if isLeaf(node) || isChild(node, prev) {
				// 来到这里，说明没有后顾之忧了，可以正真访问了
				res = append(res, node.Val)
				prev = node
			} else {
				// 来到这里，说明还有子节点没有被访问，那么交给 cur，去遍历按次序加入栈
				cur = node
			}
		}
	}

	return res
}

// 查看 x 是否是叶子节点
func isLeaf(x *TreeNode) bool {
	return x.Left == nil && x.Right == nil
}

// 查看 y 是否是 x 的子节点
func isChild(x *TreeNode, y *TreeNode) bool {
	return y != nil && (x.Left == y || x.Right == y)
}

// 非递归实现版本 1：利用前序遍历的方式
func postorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	// 准备一个栈和一个 cur，先使用前序遍历收集结果
	stack := NewStack()
	cur := root

	// 只要有一个不为空，就别停
	for cur != nil || stack.Size() != 0 {
		if cur != nil {
			res = append(res, cur.Val)
			// 注意我们这里需要变成：头 右 左 的方式
			if cur.Left != nil {
				// 如果左子树不为 nil，加入栈
				stack.Push(cur.Left)
			}
			// 往右走
			cur = cur.Right
		} else {
			// 说明栈不为空，弹出一个继续给 cur 去遍历
			cur = stack.Pop()
		}
	}

	// 来到这里，收集的结果全部是按照：头 右 左 收集的，逆序交换即可变成：左 右 头
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}

	return res
}

// 递归实现版本
func postorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	postorder(root, &res)

	return res
}

// 后序遍历
func postorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	// 先对左右子树进行后序遍历
	postorder(root.Left, res)
	postorder(root.Right, res)
	// 搞事情
	*res = append(*res, root.Val)
}
