// @Author: Ciusyan 11/1/23

package day_22

// https://leetcode.cn/problems/number-of-islands/description/

func numIslands(grid [][]byte) int {

	rowLen := len(grid)
	colLen := len(grid[0])
	// 定义一个染色方法，从 grid[row][col] 出发，将所有相邻的 1 都染成其他颜色
	var infact func(row, col int)
	infact = func(row, col int) {
		if row >= rowLen || row < 0 || col >= colLen || col < 0 || grid[row][col] != '1' {
			// 1.越界的情况 2.[row][col] 不是岛屿，统统不染色
			return
		}

		// 来到这里，将 [row][col] 的上下左右全染色了，但是需要将当前字符染色，要不然退不出循环
		grid[row][col] = '2'

		infact(row-1, col) // 上
		infact(row+1, col) // 下
		infact(row, col-1) // 左
		infact(row, col+1) // 右
	}

	res := 0
	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if grid[row][col] == '1' {
				// 只要触发一次染色，那么就会
				res++
				infact(row, col)
			}
		}
	}

	return res
}
