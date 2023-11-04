// @Author: Ciusyan 11/3/23

package day_23

import "fmt"

// Queue模拟一个队列
type Queue []*vertex

func newQueue() Queue {
	return make([]*vertex, 0)
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) Push(vt *vertex) {
	*q = append(*q, vt)
}

func (q *Queue) Pop() *vertex {
	vt := (*q)[0]
	// 模拟队头弹出
	*q = (*q)[1:]
	return vt
}

// BFS 广度优先遍历，startVal 起点的值
func (g *Graph) BFS(startVal int) {
	start, ok := g.vertexes[startVal]
	if !ok {
		return
	}

	// 准备一个队列，并将队头入队
	queue := newQueue()
	queue.Push(start)
	// 准备一个 SET 用于记录已经访问过的节点
	visited := make(map[int]struct{}, len(g.vertexes))

	for queue.Size() != 0 {
		// 说明队列不为空，弹出队头，访问
		vt := queue.Pop()
		fmt.Print(vt.val, " ")

		// 遍历所有邻居
		for _, next := range vt.nexts {
			if _, ok = visited[next.val]; ok {
				// 说明已经访问过了
				continue
			}

			// 说明还未访问过，加到队列中去排队
			queue.Push(next)
			visited[next.val] = struct{}{}
		}
	}
}

// DFS 深度优先遍历，startVal 起点的值
func (g *Graph) DFS(startVal int) {
	start, ok := g.vertexes[startVal]
	if !ok {
		return
	}

	// 准备一一个栈，用于深度优先遍历
	stack := newStack()
	// 将根起点入栈
	stack.Push(start)

	// 准备一个 Set，用于记录已经访问过的节点
	visited := make(map[int]struct{}, len(g.vertexes))

	// 栈不为空，就继续遍历
	for stack.Size() != 0 {
		vt := stack.Pop()
		if _, ok = visited[vt.val]; ok {
			// 说明访问过了
			continue
		}

		// 访问，并标记为已访问
		fmt.Print(vt.val, " ")
		visited[vt.val] = struct{}{}

		// 找一个能往下钻的所有节点
		for _, next := range vt.nexts {
			// 将其加入 Stack
			stack.Push(next)
		}
	}
}

type Stack []*vertex

func newStack() Stack {
	return make([]*vertex, 0)
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Push(vt *vertex) {
	*s = append(*s, vt)
}

func (s *Stack) Pop() *vertex {
	last := len(*s) - 1
	res := (*s)[last]
	*s = (*s)[:last]
	return res
}
