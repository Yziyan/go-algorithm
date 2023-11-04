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
		{"Start from Node 1", 1, "1 2 3 5 6 7 4 8 "},
		{"Start from Node 2", 2, "2 3 6 7 4 "},
		{"Start from Node 3", 3, "3 4 7 "},
		{"Start from Non-Existent Node", 4, "4 "},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			graph.DFS(tc.startVal)
		})
	}

}
