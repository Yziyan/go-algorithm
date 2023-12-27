// @Author: Ciusyan 12/23/23

package day_29

// 斐波那契数列矩阵乘法方式的实现 O(logN)
func fibonacciProblem(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	// |F3, F2| = base * |F2| -> |2, 1| = base * |1|
	//					 |F1|	                 |1|

	// |F4, F3| = base * |F3| -> |3, 2| = base * |2|
	//					 |F2|                    |1|
	// -> base：
	base := [][]int{
		{1, 1},
		{1, 0},
	}

	// 如果继续递推下去，所有式子相乘，可得：|Fn, Fn-1| = base^(n-2) * |F2|
	// 															 |F1|
	res := matrixPow(base, n-2)
	// base^(n-2) = |x, y|
	//				|z, w|
	// Fn = x*F2 + y*F1
	return res[0][0] + res[0][1]
}

// matrix^(n)
func matrixPow(matrix [][]int, n int) [][]int {
	// 初始化结果，并将其设置为单位矩阵
	res := make([][]int, len(matrix))
	for i := range res {
		res[i] = make([]int, len(matrix[0]))
		res[i][i] = 1
	}

	// 二进制快速求解幂的思想
	t := matrix // 矩阵的 1 次方
	for n != 0 {
		if (n & 1) != 0 {
			// 说明当前 n 的二进制位为 1，计入结果
			res = matrixMul(res, t)
		}

		// 不管算不算结果，t 都要自己和自己相乘
		t = matrixMul(t, t)

		n >>= 1
	}

	return res
}

// 矩阵乘法：a * b
func matrixMul(a, b [][]int) [][]int {
	// 矩阵乘法
	aRowL := len(a)
	aColL := len(a[0])
	bColL := len(b[0])

	// 结果是 aRowL * bColL 的矩阵
	res := make([][]int, aRowL)
	for i := range res {
		res[i] = make([]int, bColL)
	}

	for i := 0; i < aRowL; i++ {
		for j := 0; j < bColL; j++ {
			// res[i][j] = a[i][...] * b[...][j]
			for k := 0; k < aColL; k++ {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return res
}

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
