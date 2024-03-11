// @Author: Ciusyan 3/11/24

package day_34

func trailingZeroes(n int) int {
	if n <= 0 {
		return 0
	}

	res := 0
	// res = n/5 + n/5^2 + n/5^3 ...
	for n != 0 {
		n /= 5
		res += n
	}

	return res
}
