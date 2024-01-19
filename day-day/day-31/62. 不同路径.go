// @Author: Ciusyan 1/19/24

package day_31

// https://leetcode.cn/problems/unique-paths/

func uniquePaths(m int, n int) int {
	// 机器人总共要往右走 n-1 步，往下走 m-1 步。总共要走 m+n-2 步
	// 对于任意一步，都可以选择往左 or 往右，所以可以使用排列组合做
	// 即  C all right or C all button，其中 all 是下标，right、button 是上标
	var (
		all   = m + n - 2 // 走的总步数
		right = n - 1     // 往右走的步数
	)
	if m-1 > right {
		// 为了尽量少计算，因为两个组合数是相等的即：C 10 6 == C 10 4
		right = m - 1
	}

	var (
		// 防止溢出，我们每乘一个数，就先除以他们的最大公约数
		member      = 1 // 分子
		denominator = 1 // 分母
		// 比如算 C 10 6
		// 分子本来是  6! * 7 * 8 * 9 * 10
		// 分母本来是  6! * 1 * 2 * 3 * 4
		// 所以将分子和分母先最大可能约掉最多的 6!，那么就只需要算后面的数了

		// 分子和分母每次需要乘的数
		o1 = right + 1 // 分子从 right+1 开始乘
		o2 = 1         // 分母从 1 开始乘
	)

	for o1 <= all {
		// 分子和分母每次都依次乘一个数
		member *= o1
		denominator *= o2

		// 计算分子和分母的最大公约数
		g := gcd(member, denominator)

		member /= g
		denominator /= g

		// 下一次需要乘的数
		o1++
		o2++
	}

	// 最终分母肯定被约完了
	return member
}

// 求 m 和 n 的最大公约数
func gcd(m, n int) int {
	if n == 0 {
		return m
	}
	// 利用辗转相除法
	return gcd(n, m%n)
}
