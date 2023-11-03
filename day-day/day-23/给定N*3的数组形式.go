// @Author: Ciusyan 11/3/23

package day_23

// GenerateGraph 根据 N*3 的图，表示成我们熟悉的图
// matrix 所有的边
// N*3 的矩阵
// [weight, from节点上面的值，to节点上面的值]
//
// [ 5 , 0 , 7]
// [ 3 , 0,  1]
func GenerateGraph(matrix [][]int) *Graph {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) != 3 {
		return nil
	}

	// 建立一个图
	graph := NewGraph()
	// 遍历每一条边
	for i := range matrix {
		weight := matrix[i][0]
		fromVal := matrix[i][1]
		toVal := matrix[i][2]
		// 构建起点、终点、边

		// 起点
		from, ok := graph.vertexes[fromVal]
		if !ok {
			from = newVertex(fromVal)
			graph.vertexes[fromVal] = from
		}

		// 终点
		to, ok := graph.vertexes[toVal]
		if !ok {
			to = newVertex(toVal)
			graph.vertexes[toVal] = to
		}
		// 边
		eg := newEdge(weight, from, to)
		eKey := eg.key()
		if _, ok = graph.edges[eKey]; ok {
			// 说明这条边存在过了，直接重新赋值，跳过就好了，因为 weight 可能改变
			graph.edges[eKey] = eg
			continue
		}

		// 否则说明不存在边
		graph.edges[eg.key()] = eg

		// 入度+1
		to.in++
		// 出度+1
		from.out++
		// 出边
		from.nexts = append(from.nexts, to)
		// 出点
		from.edges = append(from.edges, eg)

	}

	return graph
}
