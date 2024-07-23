// @Author: Ciusyan 2024/7/23

package cycle_22_7_21_7_25

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0, 10)
	res := make([]int, 0, 10)
	// 说明还有元素没遍历完成
	for root != nil || len(stack) != 0 {
		if root != nil {
			// 直接访问根
			res = append(res, root.Val)
			if root.Right != nil {
				// 有右子树，入栈
				stack = append(stack, root.Right)
			}

			// 然后一路向左
			root = root.Left
		} else {
			// 这里说明已经到达最左端了，但是栈里还有元素需要访问
			last := len(stack) - 1
			root = stack[last]
			stack = stack[:last]
		}
	}

	return res
}
