// @Author: Ciusyan 3/21/24

package cycle_1_3_18_3_22

import "testing"

func TestHammingWeight(t *testing.T) {

	weight := hammingWeight(0b00000000000000010000000001001011)
	weight2 := hammingWeight(0b00000000000000000000000010000000)
	weight3 := hammingWeight(0b11111111111111111111111111111101)

	t.Log(weight)
	t.Log(weight2)
	t.Log(weight3)
}
