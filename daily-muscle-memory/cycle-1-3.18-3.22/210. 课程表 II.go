// @Author: Ciusyan 3/20/24

package cycle_1_3_18_3_22

// https://leetcode.cn/problems/course-schedule-ii/

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 构建这幅图，edges[0] = {1, 2}，代表 0 这个顶点的出边是 1 和 2
	graph := make([][]int, numCourses)
	// 并统计入度
	ins := make([]int, numCourses)
	for _, course := range prerequisites {
		// 先取出这条边的两个顶点
		from := course[1]
		to := course[0]
		// 然后将这条边构建好
		graph[from] = append(graph[from], to)

		// 然后统计 to 的入度
		ins[to]++
	}

	var (
		// 准备一个队列来依次执行入度为 0 的顶点
		queue = make([]int, numCourses)
		head  = 0
		tail  = 0

		// 准备结果
		res = make([]int, 0, numCourses)
	)

	// 将入度为 0 的顶点先加入队列
	for vertex, in := range ins {
		if in == 0 {
			queue[tail] = vertex
			tail++
		}
	}

	// 当队列还有元素
	for head < tail {
		// 弹出队头元素
		from := queue[head]
		head++
		// 然后收集起来
		res = append(res, from)
		outEdges := graph[from]
		for _, to := range outEdges {
			// 将出点的入度减一
			ins[to]--
			if ins[to] == 0 {
				// 说明入度为 0 了，将其加入队列中等待调度
				queue[tail] = to
				tail++
			}
		}
	}

	if len(res) != numCourses {
		// 如果没有收集到全部的顶点，说明有环，不能进行 Top 排序
		return nil
	}

	return res
}
