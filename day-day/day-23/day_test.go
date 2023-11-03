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
