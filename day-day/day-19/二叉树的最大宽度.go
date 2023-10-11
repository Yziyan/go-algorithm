// @Author: Ciusyan 10/11/23

package day_19

// 这个题要求求出二叉树的最大宽度，（哪一层节点最多，那一层就最宽）

func MaxWidth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var (
		maxWight = 1
	)

	// 准备一个队列，用于层序遍历，并将根节点入队
	queue := NewQueue()
	queue.Offer(root)

	for queue.Size() != 0 {

		// 先获取当前层有多少节点（宽度）
		curWight := queue.Size()
		if curWight > maxWight {
			maxWight = curWight
		}
		
		for i := 0; i < curWight; i++ {
			// 一层一层的处理
			// 弹出当前节点
			cur := queue.Poll()
			// 处理左边
			if cur.Left != nil {
				queue.Offer(cur.Left)
			}
			// 处理右边
			if cur.Right != nil {
				queue.Offer(cur.Right)
			}
		}
	}

	return maxWight
}

func MaxWidth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var (
		maxWight = 1
		curWight = 0

		curEnd  = root
		nextEnd *TreeNode
	)

	// 准备一个队列，用于层序遍历，并将根节点入队
	queue := NewQueue()
	queue.Offer(root)

	for queue.Size() != 0 {
		// 弹出当前节点
		cur := queue.Poll()

		// 处理左边
		if cur.Left != nil {
			queue.Offer(cur.Left)
			// 先假设是最后
			nextEnd = cur.Left
		}

		// 处理右边
		if cur.Right != nil {
			queue.Offer(cur.Right)
			// 先假设是最后
			nextEnd = cur.Right
		}

		// 当前层的宽度 +1
		curWight++
		if cur == curEnd {
			if curWight > maxWight {
				maxWight = curWight
			}
			// 将 curWight 置零，计算下一层的宽度
			curWight = 0
			// 说明即将要去下一层了，把推到最远的节点给 curEnd
			curEnd = nextEnd
		}
	}

	return maxWight
}
