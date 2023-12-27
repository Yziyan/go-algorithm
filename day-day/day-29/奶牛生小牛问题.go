// @Author: Ciusyan 12/24/23

package day_29

/**
奶牛生小牛问题
第一年农场有 1 只成熟的母牛 A，往后的每年：
1）每一只成熟的母牛都会生一只母牛
2）每一只新出生的母牛都在出生的第三年成熟
3）每一只母牛永远不会死
返回 N 年后牛的数量
*/

// 因为递推式是：F(n) = F(n-1) + F(n-3)
// 所以，可以通过三阶递推，使用 O(logN) 的方法求出答案
// 即：|Fn, Fn-1, Fn-2| = base^(n-3) * |F3|
//	                                  |F2|
//	                                  |F1|
//
// 其中 base 为三阶矩阵，可以通过前面几项求出来
// 那么如果求解出 base^(n-3)：|a, b, c|
//		                   |d, e, f|
//		                   |g, h, i|
//
//
// 那么 Fn = a*F3 + b*F2 + c*F1

func cowProblem(n int) int {
	if n <= 3 {
		// 1 2 3 4 6 9 ...
		return n
	}

	base := [][]int{
		{1, 0, 1},
		{1, 0, 0},
		{0, 1, 0},
	}

	res := matrixPow(base, n-3)
	// 如果求解出res base^(n-3)：|a, b, c|
	//		                   |d, e, f|
	//		                   |g, h, i|
	//
	//
	// 那么 Fn = a*F3 + b*F2 + c*F1
	return res[0][0]*3 + res[0][1]*2 + res[0][2]
}

func cowProblem1(n int) int {
	if n <= 3 {
		return n
	}

	var (
		first  = 1
		second = 2
		third  = 3
	)

	for i := 4; i <= n; i++ {
		third += first
		oldFirst := first
		first = second
		second = third - oldFirst
	}

	return third
}

func cowProblem2(n int) int {
	if n <= 3 {
		// 1 2 3 4 6 9 ...
		return n
	}

	// 递推式：F(n) = F(n-1) + F(n-3)
	return cowProblem(n-1) + cowProblem(n-3)
}
