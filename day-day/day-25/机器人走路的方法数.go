// @Author: Ciusyan 11/12/23

package day_25

/*
假设有排成一行的 N 个位置记为 1~N，N 一定大于或等于2
开始时机器人在其中的 M 位置上( M 一定是 1~N 中的一个)
如果机器人来到 1 位置，那么下一步只能往右来到 2 位置；
如果机器人来到 N 位置，那么下一步只能往左来到 N-1 位置；
如果机器人来到中间位置，那么下一步可以往左走或者往右走；
规定机器人必须走 K 步，最终能来到P位置( P 也是 1~N 中的一个) 的方法有多少种
给定四个参数 N、M、K、P，返回方法数
*/

// 暴力递归方法
func ways1(n, start, aim, k int) int {

	// 总共有 n 个格子，要走到 aim 处，当前在 cur，还剩余 remain 步要走
	// 返回能走的方法数量
	var process func(n, aim, cur, remain int) int

	process = func(n, aim, cur, remain int) int {
		if remain == 0 {
			// 说明没有步数要走了
			if cur == aim {
				// 说明有一种方法，已经走到了 aim 位置
				return 1
			}
			// 说明没步数了，都还没有走到 aim 位置，ways 就是 0
			return 0
		}

		if cur == 1 {
			// 这只能往右走
			return process(n, aim, cur+1, remain-1)
		}

		if cur == n {
			// 这只能往前走
			return process(n, aim, cur-1, remain-1)
		}

		// 来到这里，既可以选择往左，也可以选择往右，如何才能看到求出最多的方法数呢？
		// 那肯定是相加咯
		return process(n, aim, cur+1, remain-1) + process(n, aim, cur-1, remain-1)
	}

	// 最开始是从 start 开始，还剩余 k 步
	return process(n, aim, start, k)
}

// 傻缓存法
func ways2(n, start, aim, k int) int {

	// 其余含义不变，dp 就是一个缓存，dp[i][j] 代表 process(n, aim, i, j) 的答案
	var process func(n, aim, cur, remain int, dp [][]int) int
	process = func(n, aim, cur, remain int, dp [][]int) int {
		if dp[cur][remain] != -1 {
			// 说明已经计算过了，直接返回缓存的结果
			return dp[cur][remain]
		}
		// 否则说明以前没计算过

		res := 0
		if remain == 0 {
			if cur == aim {
				res = 1
			}
		} else if cur == 1 {
			res = process(n, aim, cur+1, remain-1, dp)
		} else if cur == n {
			res = process(n, aim, cur-1, remain-1, dp)
		} else {
			// 说明左右都可走
			res = process(n, aim, cur+1, remain-1, dp) + process(n, aim, cur-1, remain-1, dp)
		}

		// 返回前需要设置缓存
		dp[cur][remain] = res

		return res
	}

	// 主方法如何调用呢？
	// 准备缓存
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			// 初始化为 -1
			dp[i][j] = -1
		}
	}

	return process(n, aim, start, k, dp)
}
