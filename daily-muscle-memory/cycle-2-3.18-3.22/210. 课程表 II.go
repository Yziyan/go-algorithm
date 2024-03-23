// @Author: Ciusyan 3/20/24

package cycle_2_3_18_3_22

// https://leetcode.cn/problems/course-schedule-ii/

/*
*
思路重复
对于这个题，核心点其实就是，找到它的原型：拓扑排序。
但是为什么能找到呢？一个是经验吧，另一个点就是：找到关键词：依赖关系。
题目说了，选某个课之前，必须将某们课程已经全部学完了，那么就代表选的那门课程依赖于一门课程。
所以，找到了，关键原型，就是如何来实现这个拓扑排序了。
那么思路就是：
利用已有的信息，简单的将这幅图给建出来。
然后进行拓扑排序，核心思路是：卡恩算法，即从入度为零的顶点开始遍历，遍历完成就将其顶点从图中删除。
那么我们核心就是需要记录每个顶点的入度，然后从入度为零的顶点开始，依次往下找。
每遍历完一个节点，就将其终点的入度减一，如果到零后，就将其也加入队列中等待遍历。
*/
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
