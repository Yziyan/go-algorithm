// @Author: Ciusyan 5/20/24

package cycle_12_5_18_5_22

// https://leetcode.cn/problems/unique-paths/description/

func uniquePaths(m int, n int) int {
	// 辗转相除法求最大公约数
	var gcd func(m, n int) int
	gcd = func(m, n int) int {
		if n == 0 {
			return m
		}
		return gcd(n, m%n)
	}

	// 用排列组合来求解的话，那么总共需要走：m+n-2 步，其中有 m-1 步是往下走，or n-1 步是往右走
	// 那么答案就是：C all (m-1) or C all (n-1)
	all := m + n - 2
	right := n - 1 // 往右走的步数
	if right < m-1 {
		right = m - 1 // 找出大的可以算小的
	}

	top := 1
	bottom := 1

	// 比如求：C 8 2，其实就是求 8*7 / 2*1
	perTop := right + 1 // 分子开始乘的数，就是 right+1
	perBottom := 1      // 分母从 1 开始求

	// 要乘到 all，才能求解完成
	for perTop <= all {
		top *= perTop
		bottom *= perBottom

		// 可以求一个最大公约数，
		g := gcd(top, bottom)
		top /= g
		bottom /= g

		perTop++
		perBottom++
	}
	return top / bottom
}
