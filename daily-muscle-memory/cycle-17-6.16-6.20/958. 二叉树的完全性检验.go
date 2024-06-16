// @Author: Ciusyan 6/15/24

package cycle_16_6_11_6_15

// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
思路重复：
可以通过层序遍历的方式，挨层每个节点进行判断。
因为我们完全二叉树的定义是：只有最后一层和倒数第二层可以不是满的节点，并且不满的节点只能左对齐。

所以当我们挨个节点进行层序遍历的时候，就可以看这个节点，是不是左对齐的。
对于完全二叉树，当遇到第一个叶子节点，或者遇到第一个左对齐的节点时，之后再遇到的每一个节点，都必须要是叶子节点了，否者说明不满足性质。

所以我们可以准备一个开关，这个开关用于标识，什么时候开启，叶子节点模式，这样，当所有节点都遍历完成时，就说明是完全二叉树的了
*/

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var (
		queue = make([]*TreeNode, 0, 20)
		leaf  = false // 用于标记，是否开启了叶子节点
	)

	queue = append(queue, root)

	for len(queue) > 0 {
		nd := queue[0]
		queue = queue[1:]

		if leaf && (nd.Left != nil || nd.Right != nil) {
			// 说明开启了叶子节点模式，但是还有儿子，肯定不完全
			return false
		}

		if nd.Left == nil {
			if nd.Right != nil {
				// 说明有右，无左，不完全
				return false
			}
			// 说明左右都没有，也需要开启叶子节点模式了
			leaf = true
		} else if nd.Right == nil {
			// 说明左有，如果右没有，就要开启叶子节点模式了
			leaf = true
		}

		if nd.Left != nil {
			queue = append(queue, nd.Left)
		}

		if nd.Right != nil {
			queue = append(queue, nd.Right)
		}
	}

	return true
}
