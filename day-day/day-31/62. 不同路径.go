// @Author: Ciusyan 1/19/24

package day_31

// https://leetcode.cn/problems/unique-paths/

// 动态规划方法
func uniquePaths(m int, n int) int {

	// 准备一个 dp 缓存，dp[row][col] 代表：
	// 当前处于 (row, col) 位置，走到 (m-1, n-1) 位置，拥有的方法数
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 当位于最后一行或者最后一列时，只能有一种方式行走
	for col := 0; col < n; col++ {
		// 说明位于最后一行
		dp[m-1][col] = 1
	}
	for row := 0; row < m; row++ {
		// 说明位于最后一列
		dp[row][n-1] = 1
	}

	// 因为 (row, col) 依赖下和右，所以得：从下往上、从右往左求解
	for row := m - 2; row >= 0; row-- {
		for col := n - 2; col >= 0; col-- {
			// (row, col) 是右边和下边格子的值相加
			dp[row][col] = dp[row+1][col] + dp[row][col+1]
		}
	}

	// 代表从 (0, 0) 位置开始，走到 (m-1, n-1) 得到的方法数
	return dp[0][0]
}

// 暴力递归的方法
func uniquePaths2(m int, n int) int {

	// 递归含义是：当前处于 (row, col) 位置，走到 (m-1, n-1) 位置，拥有的方法数
	var process func(row, col int, m, n int) int
	process = func(row, col int, m, n int) int {
		if row == m-1 || col == n-1 {
			// 当位于最后一行或者最后一列时，智能有一种方式行走
			return 1
		}

		// 来到这里，要么往下走、要么往左走，得到的结果相加返回即可
		return process(row+1, col, m, n) + process(row, col+1, m, n)
	}

	return process(0, 0, m, n)
}

// 数学排列组合的方法
func uniquePaths1(m int, n int) int {
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
