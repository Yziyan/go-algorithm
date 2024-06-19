// @Author: Ciusyan 6/19/24

package phase3

// https://leetcode.cn/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	var defact func(grid [][]byte, row, col int)
	defact = func(grid [][]byte, row, col int) {
		if row == -1 || row == len(grid) || col == -1 || col == len(grid[0]) {
			return
		}

		cur := grid[row][col]
		if cur != '1' {
			return
		}

		grid[row][col] = '2'

		defact(grid, row-1, col)
		defact(grid, row+1, col)
		defact(grid, row, col-1)
		defact(grid, row, col+1)
	}

	res := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == '1' {
				res++
				defact(grid, row, col)
			}
		}
	}

	return res
}
