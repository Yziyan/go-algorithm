// @Author: Ciusyan 10/5/23

package day_19

// https://leetcode.cn/problems/binary-tree-postorder-traversal/description/

// 非递归实现版本 1：利用前序遍历的方式
func postorderTraversal(root *TreeNode) []int {
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
