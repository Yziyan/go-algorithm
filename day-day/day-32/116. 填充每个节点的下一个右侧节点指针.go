// @Author: Ciusyan 1/26/24

package day_32

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	// 准备一个队列，用于层序遍历
	queue := NewMyQueue()
	queue.Add(root)

	for queue.Size() != 0 {
		// 用于串后面的节点
		var pre *Node
		// 一层一层的遍历
		size := queue.Size()
		for i := 0; i < size; i++ {
			cur := queue.Remove()
			if pre != nil {
				// 说明不是第一个打头的节点了，需要串到 pre 的后面
				pre.Next = cur
			}
			pre = cur

			// 依次添加左和右
			if cur.Left != nil {
				queue.Add(cur.Left)
			}
			if cur.Right != nil {
				queue.Add(cur.Right)
			}
		}
	}

	return root
}

// MyQueue 自己实现一个 Queue，节点就是 Node
type MyQueue struct {
	head *Node
	tail *Node
	size int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{}
}

func (q *MyQueue) Add(node *Node) {
	if q.size == 0 {
		// 说明是第一次添加节点
		q.head = node
		q.tail = node
	} else {
		// 说明之前已经有节点了，尾插到最后
		q.tail.Next = node
		q.tail = node
	}
	q.size++
}

func (q *MyQueue) Remove() *Node {
	if q.size == 0 {
		// 说明没有节点
		return nil
	}
	q.size--

	// 弹出头节点返回。
	oldHead := q.head
	q.head = q.head.Next
	// 将 oldHead 的末尾清空
	oldHead.Next = nil

	return oldHead
}

func (q *MyQueue) Size() int {
	return q.size
}
