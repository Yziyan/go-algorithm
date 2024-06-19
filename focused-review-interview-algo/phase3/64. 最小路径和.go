// @Author: Ciusyan 6/19/24

package phase3

import "math"

// https://leetcode.cn/problems/minimum-path-sum/
func minPathSum(grid [][]int) int {

	rowL := len(grid)
	colL := len(grid[0])

	dp := make([]int, colL+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}

	for row := rowL - 1; row >= 0; row-- {
		for col := colL - 1; col >= 0; col-- {
			next := min(dp[col], dp[col+1])
			if next == math.MaxInt {
				next = 0
			}
			dp[col] = next + grid[row][col]
		}
	}

	return dp[0]
}
