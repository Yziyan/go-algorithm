// @Author: Ciusyan 4/12/24

package cycle_5_4_12_4_16

// https://leetcode.cn/problems/powx-n/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n == -1 {
		return 1 / x
	}

	pow := n
	if pow < 0 {
		pow = -pow
	}

	res := 1.0

	for pow != 0 {
		if (pow & 1) == 1 {
			// 说明当前末尾二进制为 1，对结果有收益
			res *= x
		}

		x *= x
		pow >>= 1
	}

	if n < 0 {
		return 1 / res
	}
	return res
}
