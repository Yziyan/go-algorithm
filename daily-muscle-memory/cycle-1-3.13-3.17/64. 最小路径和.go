// @Author: Ciusyan 3/17/24

package cycle_1_3_13_3_17

import "math"

// 动态规划，一维数组
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var (
		rowL = len(grid)
		colL = len(grid[0])
	)

	// dp[row][col] 代表从 (row, col) 出发，到达右下角时，能够得到的最小路径和
	dp := make([]int, colL+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}

	// 因为 next 依赖 row+1 和 col+1，所以都是从后往前求解
	for row := rowL - 1; row >= 0; row-- {
		for col := colL - 1; col >= 0; col-- {
			next := min(dp[col], dp[col+1])
			if next == math.MaxInt {
				next = 0
			}
			dp[col] = grid[row][col] + next
		}
	}

	// 从 左上角 -> 右下角，得到的最小路径和
	return dp[0]
}

// 动态规划
func minPathSum2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var (
		rowL = len(grid)
		colL = len(grid[0])
	)

	// dp[row][col] 代表从 (row, col) 出发，到达右下角时，能够得到的最小路径和
	dp := make([][]int, rowL+1)
	for i := range dp {
		dp[i] = make([]int, colL+1)
	}

	// 先设置最后一列
	for row := 0; row < rowL; row++ {
		dp[row][colL] = math.MaxInt
	}
	// 最后一行
	for col := 0; col < colL; col++ {
		dp[rowL][col] = math.MaxInt
	}

	// 因为 next 依赖 row+1 和 col+1，所以都是从后往前求解
	for row := rowL - 1; row >= 0; row-- {
		for col := colL - 1; col >= 0; col-- {
			next := min(dp[row+1][col], dp[row][col+1])
			if next == math.MaxInt {
				next = 0
			}
			dp[row][col] = grid[row][col] + next
		}
	}

	// 从 左上角 -> 右下角，得到的最小路径和
	return dp[0][0]
}

// 暴力递归
func minPathSum1(grid [][]int) int {

	// 从 (row, col) 出发，到达右下角时，能够得到的最小路径和
	var process func(grid [][]int, row, col int) int
	process = func(grid [][]int, row, col int) int {
		if row == len(grid) || col == len(grid[0]) {
			// 越界了，得到的路径是 0
			return math.MaxInt
		}

		if row == len(grid)-1 && col == len(grid[0]) {
			// 说明刚好到达右下角
			return grid[row][col]
		}

		next := min(process(grid, row+1, col), process(grid, row, col+1))
		if next == math.MaxInt {
			next = 0
		}

		// 否则往左 or 往下，选一个最小值，加上当前的值，即可
		return grid[row][col] + next
	}

	// 从 左上角 -> 右下角，得到的最小路径和
	return process(grid, 0, 0)
}
