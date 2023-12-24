// @Author: Ciusyan 12/24/23

package day_28

/**
奶牛生小牛问题
第一年农场有 1 只成熟的母牛 A，往后的每年：
1）每一只成熟的母牛都会生一只母牛
2）每一只新出生的母牛都在出生的第三年成熟
3）每一只母牛永远不会死
返回 N 年后牛的数量
*/

func cowProblem(n int) int {
	if n <= 3 {
		// 1 2 3 4 6 9 ...
		return n
	}

	// 递推式：F(n) = F(n-1) + F(n-3)
	return cowProblem(n-1) + cowProblem(n-3)
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
