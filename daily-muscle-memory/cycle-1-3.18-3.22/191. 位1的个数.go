// @Author: Ciusyan 3/21/24

package cycle_1_3_18_3_22

func hammingWeight(num uint32) int {
	bits := 0

	for num != 0 {
		// 能够每次消掉一个 num 的最后一个 1，能消多少次？
		bits++
		// 只留下最后一个 1
		// 0101 1010 -> lastOne = 0000 0010
		// 1010 0101 + 1 = 1010 0110
		// 0101 1010 & 1010 0110 = 0000 0010
		lastOne := num & (-num)
		// 0101 1010 ^ 0000 0010 = 0101 1000
		// 然后再对 num 进行异或运算，异或（异 1，同 0）
		num ^= lastOne
	}

	return bits
}
