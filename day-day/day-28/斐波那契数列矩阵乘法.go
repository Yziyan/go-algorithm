// @Author: Ciusyan 12/23/23

package day_28

// 斐波那契数列矩阵乘法方式的实现

// 非递归实现，从前往后递推
func fibonacciProblem2(n int) int {
	if n <= 0 {
		return 0
	}

	var (
		// 准备前一个数和后一个数
		first  = 1
		second = 1
	)

	for i := 3; i <= n; i++ {
		// 先计算出后一项
		second += first
		// 再减去加上的 first 的值，就是原始的 second
		first = second - first
	}

	return second
}

// 最简单的实现
func fibonacciProblem1(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		// f(1) = f(2) = 1
		return 1
	}

	// f(n) = f(n-1) + f(n-2)
	return fibonacciProblem1(n-1) + fibonacciProblem1(n-2)
}
