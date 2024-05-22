// @Author: Ciusyan 5/20/24

package cycle_12_5_17_5_21

// https://leetcode.cn/problems/unique-paths/description/

/*
*
思路重复：
这种格子求 xx 最值的问题，很容易想到使用动态规划来实现，
但是呢？我们这里这个题，要求解的是从左上角到右下角位置，有多少种走法。
因为每一步都可以选择往下或者往右，并且无论怎么走，要往下，往右走多少步是一定的。
那么就可以使用排列组合来求解，比如：往右走 n-1、往下走 m-1
总共走 m+n-2，那么最终结果就是：C all 右 || C all 下
并且由于这个 C 的公式，需要算 all! / 右! * (all-右)!
防止溢出，我们可以从小往大算，并且每计算一次，就可以进行约一次分。约掉分子分母的最大公约数。
那么如何求解分子分母的最大公约数呢？可以使用辗转相除法求解。
*/

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
