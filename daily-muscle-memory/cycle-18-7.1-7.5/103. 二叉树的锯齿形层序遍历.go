// @Author: Ciusyan 2024/7/1

package cycle_18_7_1_7_5

// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
层序遍历，抓住两个核心点：
1.一层一层遍历
2.一层层遍历，每一个节点都要被遍历到。

一层批次遍历还好说，可是如何控制方向呢？
用一个标记位去记录方向，如果正向遍历，那就从小到大填充结果，如果是需要逆序遍历，就从后往前填充结果即可
*/

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
