// @Author: Ciusyan 6/19/24

package phase3

func hammingWeight(n int) int {
	bits := 0

	for n != 0 {
		bits++
		lastOne := (-n) & n
		n ^= lastOne
	}

	return bits
}
