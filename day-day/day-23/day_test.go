// @Author: Ciusyan 11/3/23

package day_23

import "testing"

func TestGenerateGraph(t *testing.T) {

	args := [][]int{
		{4, 3, 9},
		{4, 3, 9},
		{4, 3, 9},
		{5, 3, 9},
	}

	graph := GenerateGraph(args)

	t.Log(graph)
}

func TestGraph_BFS(t *testing.T) {
	// 创建一个图
	args := [][]int{
		{4, 1, 2},
		{4, 1, 3},
		{4, 2, 3},
		{4, 3, 4},
		{4, 1, 5},
		{4, 2, 6},
		{4, 2, 7},
		{4, 3, 7},
		{4, 5, 8},
	}

	graph := GenerateGraph(args)

	// 添加节点和边

	// 创建测试用例表
	testCases := []struct {
		name     string
		startVal int
		expected string
	}{
		{"Start from Node 1", 1, "1 2 3 5 6 7 4 8 "},
		{"Start from Node 2", 2, "2 3 6 7 4 "},
		{"Start from Node 3", 3, "3 4 7 "},
		{"Start from Non-Existent Node", 4, "4 "},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			graph.BFS(tc.startVal)
		})
	}
}

func TestGraph_DFS(t *testing.T) {
	// 创建一个图
	args := [][]int{
		{4, 1, 2},
		{4, 1, 3},
		{4, 2, 3},
		{4, 3, 4},
		{4, 1, 5},
		{4, 2, 6},
		{4, 2, 7},
		{4, 3, 7},
		{4, 5, 8},
	}

	graph := GenerateGraph(args)

	// 添加节点和边

	// 创建测试用例表
	testCases := []struct {
		name     string
		startVal int
		expected string
	}{
		{"Start from Node 1", 1, "1 5 8 3 7 4 2 6 "},
		{"Start from Node 2", 2, "2 7 6 3 4 "},
		{"Start from Node 3", 3, "3 7 4 "},
		{"Start from Non-Existent Node", 4, "4 "},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			graph.DFS(tc.startVal)
			graph.DFS1(tc.startVal)
		})
	}

}

func TestGraph_TopologicalSort(t *testing.T) {
	// 创建一个图
	args := [][]int{
		{4, 1, 2},
		{4, 1, 3},
		{4, 3, 4},
		{4, 3, 9},
		{4, 1, 5},
		{4, 2, 6},
		{4, 2, 7},
		{4, 2, 10},
		{4, 5, 8},
		{4, 6, 12},
		{4, 7, 12},
		{4, 8, 12},
	}

	graph := GenerateGraph(args)
	got := graph.TopologicalSort()
	t.Log(got)
}

func TestGraph_KruskalBST(t *testing.T) {
	// 创建一个无向图
	args := [][]int{
		{4, 1, 2},
		{4, 2, 1},
		{6, 1, 3},
		{6, 3, 1},
		{9, 1, 4},
		{9, 4, 1},
		{9, 2, 3},
		{9, 3, 2},
		{10, 2, 5},
		{10, 5, 2},
		{4, 2, 4},
		{4, 4, 2},
		{1, 3, 4},
		{1, 4, 3},
		{8, 3, 6},
		{8, 6, 3},
		{12, 3, 5},
		{12, 5, 3},
		{15, 4, 6},
		{15, 6, 4},
		{12, 5, 6},
		{12, 6, 5},
	}

	graph := GenerateGraph(args)
	got := graph.KruskalBST()
	got2 := graph.PrimBst()
	t.Log(got)
	t.Log(got2)
}

func TestGraph_DijkstraShortPath(t *testing.T) {
	// 创建一个有向，无负权环的图
	args := [][]int{
		{1, 1, 2},
		{6, 1, 3},
		{7, 1, 4},
		{2, 2, 3},
		{10, 2, 5},
		{3, 3, 4},
		{6, 3, 5},
		{2, 4, 5},
	}

	graph := GenerateGraph(args)

	shortPath := graph.DijkstraShortPath(1)

	// map[1:0 2:1 3:3 4:6 5:8]
	t.Log(shortPath)
}
