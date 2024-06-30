// @Author: Ciusyan 2024/7/1

package cycle_18_7_1_7_5

// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	res := make([][]int, 0, 10)
	flag := true // true 代表正向

	// 层序遍历，一层层的遍历即可。
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
		// 一层一层的收集，一次性收集一层
		if flag {
			// 正向，那就从小收集到末尾
			for i := 0; i < size; i++ {
				nd := queue[0]
				queue = queue[1:]
				level[i] = nd.Val

				if nd.Left != nil {
					queue = append(queue, nd.Left)
				}

				if nd.Right != nil {
					queue = append(queue, nd.Right)
				}
			}
		} else {
			// 反向，从大到校收集
			for i := size - 1; i >= 0; i-- {
				nd := queue[0]
				queue = queue[1:]
				level[i] = nd.Val

				if nd.Left != nil {
					queue = append(queue, nd.Left)
				}

				if nd.Right != nil {
					queue = append(queue, nd.Right)
				}
			}

		}

		res = append(res, level)
		flag = !flag // 转变方向
	}

	return res
}
