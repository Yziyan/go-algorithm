// @Author: Ciusyan 2024/7/23

package cycle_22_7_21_7_25

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*
前序遍历最简单的及时使用递归的方式去写。
但是我们可以使用 stack 来替代递归的方式。
两种方式：
1.将层序遍历使用的 queue 换成 stack，其余不变，便可以无缝衔接变成前序遍历
2.使用 stack，从根节点开始遍历，只要 root 不是 nil 或者 stack 有元素，
如果 root 有值，就直接收集结果，然后看看有没有右子树，有的话就入栈，然后往坐走
否则就从 stack 弹出元素，给 root，继续进行上述操作
*/

func preorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0, 10)
	res := make([]int, 0, 10)

	for root != nil || len(stack) != 0 {
		if root != nil {
			res = append(res, root.Val)
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			root = root.Left
		} else {
			last := len(stack) - 1
			root = stack[last]
			stack = stack[:last]
		}
	}

	return res
}
