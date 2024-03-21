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

/*
*
思路重复：
要求 num 的二进制有多少个 1，如果我们每次都消掉 num 二进制的最后一个 1，
那么消完所有 1，num 也就等于 0 了，即消了多少次，其实就有多少个 1
但是怎么每次都能消掉最后一个 1 呢？
我们每次先取出 num 二进制最后一个 1，lastOne
然后将 num 与 lastOne 进行异或，即可消掉最后一个 1.
那么又怎么取出 lastOne 呢？
其实就是将 num 与上 （num 取反 +1）即可
而 （num 取反 +1） 其实就是 -num
所以 lastOne = num & (-num)
*/
func hammingWeight2(num uint32) int {
	bits := 0
	for num != 0 {
		bits++
		lastOne := num & (-num)
		num ^= lastOne
	}
	return bits
}
