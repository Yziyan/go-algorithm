// @Author: Ciusyan 1/17/24

package day_31

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{1, 0, 1, 1, 1}

	b := search(nums, 0)
	t.Log(b)
}

func TestUniquePaths(t *testing.T) {

	paths := uniquePaths2(3, 7)
	paths2 := uniquePaths(3, 7)
	t.Log(paths)
	t.Log(paths2)
}
