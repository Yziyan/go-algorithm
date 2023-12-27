// @Author: Ciusyan 12/27/23

package day_29

/**
给定一个数 N，想象只由 0和1 两种字符，组成的所有长度为 N 的字符串
如果某个字符串，任何 0 字符的左边都有 1 紧挨着，认为这个字符串达标
返回有多少达标的字符串
*/

func str01Num1(n int) int {
	if n <= 2 {
		return n
	}

	var (
		first  = 1
		second = 2
	)

	for i := 3; i <= n; i++ {
		second += first
		first = second - first
	}

	return second
}

func str01Num2(n int) int {
	if n <= 2 {
		return n
	}
	// F(n) = F(n-1) + F(n-2)
	return str01Num2(n-1) + str01Num2(n-2)
}
