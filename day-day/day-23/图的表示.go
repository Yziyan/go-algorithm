// @Author: Ciusyan 11/3/23

package day_23

import "fmt"

// 这里推荐一种比较好懂的图的表示方法
//	核心思路是：
//	1.顶点需要什么？2.边需要什么？3.图需要什么？

// vertex 图的顶点
type vertex struct {
	// 顶点的值
	val int
	// 入度、出度
	in, out int
	// 直接能够到达的顶点
	nexts []*vertex
	// 从 this 出发的边
	edges []*edge
}

func newVertex(val int) *vertex {
	return &vertex{
		val: val,
	}
}

// edge 图的边
type edge struct {
	// 权重
	weight int
	// 起点
	from *vertex
	// 终点
	to *vertex
}

func (e *edge) key() string {
	// 唯一标识一个 Key
	return fmt.Sprintf("%d_%d", e.from.val, e.to.val)
}

func newEdge(weight int, from, to *vertex) *edge {
	return &edge{
		weight: weight,
		from:   from,
		to:     to,
	}
}

// Graph 图
type Graph struct {
	// 所有的顶点
	vertexes map[int]*vertex
	// 所有的边，key = "fromVal_toVal"，因为 Go 不支持重写 hash + equal
	edges map[string]*edge
}

func NewGraph() *Graph {
	return &Graph{
		vertexes: make(map[int]*vertex),
		edges:    make(map[string]*edge),
	}
}
