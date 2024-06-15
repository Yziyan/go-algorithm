// @Author: Ciusyan 6/15/24

package cycle_16_6_11_6_15

// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/description/

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
