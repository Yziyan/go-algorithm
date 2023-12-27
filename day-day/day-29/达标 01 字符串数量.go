// @Author: Ciusyan 12/27/23

package day_29

/**
给定一个数 N，想象只由 0和1 两种字符，组成的所有长度为 N 的字符串
如果某个字符串，任何 0 字符的左边都有 1 紧挨着，认为这个字符串达标
返回有多少达标的字符串
*/

func str01Num(n int) int {
	if n <= 2 {
		return n
	}

	// F(1) = 1 F(2) = 2 F(n) = F(n-1) + F(n-2)
	// 可得出
	base := [][]int{
		{1, 1},
		{1, 0},
	}

	// |Fn Fn-1| = base^(n-2) * |F2|
	//							|F1|
	res := matrixPow_Sn(base, n-2)

	// Fn = 2*a + 1*b
	return 2*res[0][0] + res[0][1]
}

// _Sn 标识 str01Num matrix^n
func matrixPow_Sn(matrix [][]int, n int) [][]int {

	rowL := len(matrix)
	colL := len(matrix[0])

	// 初始化一个单位矩阵
	unitRes := make([][]int, rowL)
	for i := range unitRes {
		unitRes[i] = make([]int, colL)
		unitRes[i][i] = 1
	}

	// 使用快速幂的求解方法

	// 默认是 t^1
	t := matrix
	for n != 0 {
		if (n & 1) != 0 {
			// 说明要作用于结果上
			unitRes = matrixMul_Sn(unitRes, t)
		}

		// 不管怎样，都要自己和自己乘一下
		t = matrixMul_Sn(t, t)

		n >>= 1
	}

	return unitRes
}

// a * b
func matrixMul_Sn(a, b [][]int) [][]int {

	aRowL := len(a)
	bRowL := len(b)
	bColL := len(b[0])

	// 结果一定是一个 aRowL * bColL 的矩阵
	res := make([][]int, aRowL)
	for i := range res {
		res[i] = make([]int, bColL)
	}

	// 矩阵乘法
	for i := 0; i < aRowL; i++ {
		for j := 0; j < bColL; j++ {
			for k := 0; k < bRowL; k++ {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return res
}

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
