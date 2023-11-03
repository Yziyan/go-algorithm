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
