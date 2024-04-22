// @Author: Ciusyan 4/18/24

package cycle_6_4_17_4_21

import "testing"

func TestReverse(t *testing.T) {

	area := maxArea([]int{1, 1})

	t.Log(area)
}

func TestLongestCommonPrefix(t *testing.T) {

	prefix := longestCommonPrefix([]string{"cir", "car"})

	t.Log(prefix)
}

func TestPermute(t *testing.T) {

	res := permute([]int{1, 2, 3})
	t.Log(res)
}
