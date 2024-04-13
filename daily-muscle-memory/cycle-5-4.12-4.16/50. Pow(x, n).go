// @Author: Ciusyan 4/12/24

package cycle_5_4_12_4_16

// https://leetcode.cn/problems/powx-n/

/**
思路重复：
对于快速幂的求解，其实核心就是采用分治的思想：
将 pow 次幂，分为二进制的幂次方相乘：
myPow(3, 5)
3 的二进制：0101
3^5 = 3^1 * 3^4
3^(2^0) * 3^(2^2)
所以对于二进制而言：0101，如果对应位置是 1，那么就对结果有帮助，是零对结果就没有帮助。
也就是说：
    0       1       0       1
   3^8     3^4     3^2     3^1
 3^(2^3) 3^(2^2) 3^(2^1) 3^(2^0)
那么结果就是：
	1 *	  3^4  *   1  *  3^(2^0)

*/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	pow := n
	if pow < 0 {
		pow = -pow
	}

	res := 1.0
	for pow != 0 {
		if (pow & 1) == 1 {
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
