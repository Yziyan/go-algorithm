// @Author: Ciusyan 11/6/23

package day_23

// TopologicalSort 拓扑排序，适用于有向无环图，可以利用卡恩算法，
func (g *Graph) TopologicalSort() []int {
	l := len(g.vertexes)
	if l == 0 {
		return nil
	}

	// 统计入度
	ins := make(map[int]int, l)

	queue := newQueue()

	for key, val := range g.vertexes {
		if val.in != 0 {
			// 说明统计入度
			ins[key] = val.in
			continue
		}

		// 否则加入队列等待遍历
		queue.Push(val)
	}

	res := make([]int, 0, l)

	// 只要队列不为空，就说明还没遍历完
	for queue.Size() != 0 {
		// 弹出便利
		vt := queue.Pop()
		res = append(res, vt.val)

		// 将所有邻居的入度减一
		for _, next := range vt.nexts {
			// 入度都要减
			ins[next.val]--
			if ins[next.val] == 0 {
				// 说明下一次遍历它
				queue.Push(next)
			}
		}
	}

	if len(res) != l {
		panic("该图存在环，无拓扑序")
	}

	return res
}
