// @Author: Ciusyan 11/29/23

package day_25

/**
给定一个二维数组 matrix，一个人必须从左上角出发，最后到达右下角
沿途只可以向下或者向右走，沿途的数字都累加就是距离累加和
返回最小距离累加和
*/

func minPathSum(matrix [][]int) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	// 准备缓存，可变参数是 row, col
	m, n := len(matrix), len(matrix[0])
	// dp[row][col] 代表，从 (0, 0) -> (row, col) 最小的累加和
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 根据递归基
	dp[0][0] = matrix[0][0]

	// 第一行
	for col := 1; col < n; col++ {
		// 只依赖左边
		dp[0][col] = dp[0][col-1] + matrix[0][col]
	}
	// 第一列
	for row := 1; row < m; row++ {
		// 只依赖上面
		dp[row][0] = dp[row-1][0] + matrix[row][0]
	}

	// 其余位置
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			// 一般位置，就依赖左边和右边的最小值
			dp[row][col] = min(dp[row][col-1], dp[row-1][col]) + matrix[row][col]
		}
	}

	// 从 (0, 0) -> (m, n) 的最小累加和
	return dp[m-1][n-1]
}

// 暴力递归方法
func minPathSum1(matrix [][]int) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	// 从 (0, 0) 位置 -> (row, col) 位置，得到的最小累加和
	var process func(matrix [][]int, row, col int) int
	process = func(matrix [][]int, row, col int) int {
		if row == 0 && col == 0 {
			return matrix[0][0]
		}
		if row == 0 {
			// 说明只能往左到 (0, col)，那么应该是 (0,0) -> (0, col-1) 的最小累加和 + (0, col) 的值
			return process(matrix, 0, col-1) + matrix[0][col]
		}

		if col == 0 {
			// 说明只能往上面到 (row, 0)，那么应该是 (0, 0) -> (row-1, 0) 的最小累加和 + (row, 0) 的值
			return process(matrix, row-1, 0) + matrix[row][0]
		}

		// 否则一般情况是：上边和左边的最小累加和的最小值 + (row, col) 的值
		return min(process(matrix, row-1, col), process(matrix, row, col-1)) + matrix[row][col]
	}

	m, n := len(matrix), len(matrix[0])
	// 从 (0, 0) 位置 -> (m, n) 位置，得到的最小累加和
	return process(matrix, m-1, n-1)
}
