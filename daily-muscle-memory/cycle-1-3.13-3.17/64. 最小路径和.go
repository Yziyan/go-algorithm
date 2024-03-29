// @Author: Ciusyan 3/17/24

package cycle_1_3_13_3_17

import "math"

/*
*
思路重复：
要求解从 (0, 0) -> (rowL-1, colL-1) 的最小路径和，
很容易联想到范围尝试模型，即来到 (row, col) 时，可以选择往 (row+1, col)，也可以选择往 (row, col+1)
哪边最小，就往那边走，然后路径就是 grid[row][col] + min(往下, 往右) 的值。
所以我们可以定义一个递归函数就是：process(row, col int) int，代表从 (row, col) -> (rowL-1, colL-1) 的最小路径和
当走越界时，我们可以设置一个无效值。
然后去看看 process(row+1, col) 和 process(row, col+1) 的最小值。如果还是无效值，那么说明后面的路径不可用，使用 grid[row][col] 即可
否则就是 grid[row][col] + min(往下, 往右)
也就可以写成动态规划了
*/
func minPathSum22(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var (
		rowL = len(grid)
		colL = len(grid[0])
	)

	// 准备 dp，使用滚动数组，因为只依赖，右和下的值
	dp := make([]int, colL+1)
	for i := range dp {
		// 先都标记为无效值
		dp[i] = math.MaxInt
	}

	// 然后从下往上，从右往左求
	for row := rowL - 1; row >= 0; row-- {
		for col := colL - 1; col >= 0; col-- {
			// 此时 dp[col] 就代表下，dp[col+1] 就代表右
			next := min(dp[col], dp[col+1])
			if next == math.MaxInt {
				// 说明后面的路径是无效值
				next = 0
			}
			dp[col] = grid[row][col] + next
		}
	}

	// 求解到左上角了
	return dp[0]
}

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
