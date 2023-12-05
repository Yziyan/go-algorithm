// @Author: Ciusyan 12/4/23

package day_26

import "math"

/**
给定 5 个参数，N，M，row，col，k
表示在 N*M 的区域上，醉汉 Bob 初始在 (row,col) 位置
Bob 一共要迈出 k 步，且每步都会等概率向上下左右四个方向走一个单位
任何时候 Bob 只要离开 N*M 的区域，就直接死亡
返回 k 步之后，Bob 还在 N*M 的区域的概率
*/

// 动态规划方法
func livePossibility(row, col int, k int, n, m int) float64 {

	// 我们发现，这个题有三个可变参数：row、col、remain，那么是一个三维的动态规划
	// dp[row][col][remain] 代表：当前处于(row, col) 位置，剩余 remain 步要走，能生存下来的总数
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}
	// 根据递归基，当 remain = 0 时，只要每越界都是生存的
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if row < n && col < m {
				dp[x][y][0] = 1
			}
		}
	}

	picker := func(dp [][][]int, x, y, z int, n, m int) int {
		if x < 0 || x >= n || y < 0 || y >= m {
			// 说明越界了
			return 0
		}

		return dp[x][y][z]
	}

	// 根据依赖情况，remain 层依赖 remain-1 层，所以需要从下往上求解
	for z := 1; z <= k; z++ {
		for x := 0; x < n; x++ {
			for y := 0; y < m; y++ {
				// 有四个方向，
				p1 := picker(dp, x-1, y, z-1, n, m) // 上
				p2 := picker(dp, x+1, y, z-1, n, m) // 下
				p3 := picker(dp, x, y-1, z-1, n, m) // 左
				p4 := picker(dp, x, y+1, z-1, n, m) // 右

				dp[x][y][z] = p1 + p2 + p3 + p4
			}
		}
	}

	// 那么 dp[row][col][k] 代表当前处于(row, col)，走完 k 步，能生存下来的总数，
	// 再除以总的走法，就是生存概率
	return float64(dp[row][col][k]) / math.Pow(4, float64(k))
}

// 暴力递归方法
func livePossibility1(row, col int, k int, n, m int) float64 {

	// 憋一个暴力递归，递归含义是：
	// 当前处于 (row, col) 位置，走完 remain 步，有多少种生存的可能，
	var process func(row, col int, remain int, n, m int) int64
	process = func(row, col int, remain int, n, m int) int64 {
		if row < 0 || row >= n || col < 0 || col >= m {
			// 说明已经越界了
			return 0
		}

		if remain == 0 {
			// 说明最后一步走完了，都还没越界
			return 1
		}

		// 否则可以有四种可能，朝上、下、左、右走
		p1 := process(row-1, col, remain-1, n, m) // 上
		p2 := process(row+1, col, remain-1, n, m) // 下
		p3 := process(row, col-1, remain-1, n, m) // 左
		p4 := process(row, col+1, remain-1, n, m) // 右

		// 最终结果就是朝四个方向走，能生存的可能性相加起来
		return p1 + p2 + p3 + p4
	}

	// 那么主函数怎么调用嗯？先求出总的生存可能
	// 当前处于 (row, col) 位置，走完 k 步，有多少种生存的可能
	// 那么还在 N*M 区域的概率就是：还在这个区域的数量 / 走的总步数
	return float64(process(row, col, k, n, m)) / math.Pow(4, float64(k))
}
