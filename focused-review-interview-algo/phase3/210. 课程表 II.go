// @Author: Ciusyan 6/19/24

package phase3

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	ins := make([]int, numCourses)

	for _, course := range prerequisites {
		from := course[1]
		to := course[0]

		graph[from] = append(graph[from], to)
		ins[to]++
	}

	var (
		queue = make([]int, numCourses)
		tail  = 0
		head  = 0
	)

	for vertex, in := range ins {
		if in == 0 {
			queue[tail] = vertex
			tail++
		}
	}

	res := make([]int, 0, numCourses)

	for head != tail {
		from := queue[head]
		head++
		res = append(res, from)

		for _, to := range graph[from] {
			ins[to]--
			if ins[to] == 0 {
				queue[tail] = to
				tail++
			}
		}
	}

	if len(res) != numCourses {
		return nil
	}
	return res
}
