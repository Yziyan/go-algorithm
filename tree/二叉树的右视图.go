// @Author: Ciusyan 2023/7/27

package tree

// https://leetcode.cn/problems/binary-tree-right-side-view/

func RightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	queue := &Queue[*TreeNode]{}
	queue.Push(root)

	for !queue.IsEmpty() {
		// 需要一层一层的判断，
		size := queue.Size
		for i := 0; i < size; i++ {
			node := queue.Poll()
			// 能看到的肯定是这一层中最后一个元素
			if i == size-1 {
				res = append(res, node.Val)
			}

			if node.Left != nil {
				queue.Push(node.Left)
			}

			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
	}

	return res
}

type listNode[T any] struct {
	Val  T
	Next *listNode[T]
}

type Queue[T any] struct {
	Size int
	Head *listNode[T]
	Tail *listNode[T]
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size == 0
}

func (q *Queue[T]) Push(val T) {
	node := &listNode[T]{Val: val}
	if q.Tail == nil {
		q.Head = node
	} else {
		// 尾插法
		q.Tail.Next = node
	}

	q.Tail = node
	q.Size++
}

func (q *Queue[T]) Poll() T {
	var res T
	if q.Head != nil {
		res = q.Head.Val
		q.Head = q.Head.Next

		q.Size--
	}

	if q.Head == nil {
		q.Tail = nil
	}

	return res
}
