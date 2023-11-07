// @Author: Ciusyan 11/3/23

package day_23

import "fmt"

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

// DFS 深度优先遍历，startVal 起点的值（方法一）
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
	fmt.Println()
}

// DFS1 深度优先遍历，startVal 起点的值（方法二）
func (g *Graph) DFS1(startVal int) {
	start, ok := g.vertexes[startVal]
	if !ok {
		return
	}

	// 准备一个 SET，用于记录已访问的节点
	visited := make(map[int]struct{}, len(g.vertexes))
	// 准备一一个栈，用于深度优先遍历
	stack := newStack()
	stack.Push(start)
	// 并且压入就访问
	fmt.Print(start.val, " ")
	visited[start.val] = struct{}{}

	// 栈不为空，就说明还没访问完成
	for stack.Size() != 0 {
		// 弹出队头
		vt := stack.Pop()

		// 看看能否访问一个邻居
		for _, next := range vt.nexts {
			if _, ok = visited[next.val]; ok {
				// 说明访问过了，跳过这个
				continue
			}

			// 说明还没访问，
			fmt.Print(next.val, " ")
			visited[next.val] = struct{}{}

			// 然后加入栈，但是加入前，还需要将当前节点加入，方便回溯
			stack.Push(vt)
			stack.Push(next)
			// 但是这一次只要一个邻居就可以了
			break
		}
	}

}
