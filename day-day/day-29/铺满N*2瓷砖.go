// @Author: Ciusyan 12/28/23

package day_29

/**
用 1*2 的瓷砖，把 N*2 的区域填满，返回铺瓷砖的方法数
*/

func fullBlock(n int) int {
	if n <= 2 {
		return n
	}

	// a * b
	matrixMul := func(a, b [][]int) [][]int {

		aRowL := len(a)
		bRowL := len(b)
		bColL := len(b[0])

		res := make([][]int, aRowL)
		for i := range res {
			res[i] = make([]int, bColL)
		}

		for i := 0; i < aRowL; i++ {
			for j := 0; j < bColL; j++ {
				for k := 0; k < bRowL; k++ {
					res[i][j] += a[i][k] * b[k][j]
				}
			}
		}

		return res
	}

	// matrix~n
	matrixPow := func(matrix [][]int, n int) [][]int {

		rowL := len(matrix)
		colL := len(matrix[0])
		res := make([][]int, rowL)
		for i := range res {
			res[i] = make([]int, colL)
			res[i][i] = 1
		}

		t := matrix
		for n != 0 {
			if (n & 1) != 0 {
				res = matrixMul(res, t)
			}

			t = matrixMul(t, t)
			n >>= 1
		}

		return res
	}

	// base：
	base := [][]int{
		{1, 1},
		{1, 0},
	}

	// base^(n-2)
	res := matrixPow(base, n-2)

	return res[0][0]*2 + res[0][1]
}

// 观察，或者使用递归的方式，可以知道，这个题也是：F(1) = 1 F(2) = 2 的斐波那契问题
func fullBlock1(n int) int {
	if n <= 2 {
		return n
	}

	return fullBlock1(n-1) + fullBlock1(n-2)
}
