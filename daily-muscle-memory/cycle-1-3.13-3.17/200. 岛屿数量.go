// @Author: Ciusyan 3/16/24

package cycle_1_3_13_3_17

// https://leetcode.cn/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var (
		rowLen = len(grid)
		colLen = len(grid[0])
	)

	// 记录触发染色的次数
	res := 0
	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if grid[row][col] == SRC {
				// 说明要触发一次染色
				res++
				infect(grid, row, col)
			}
		}
	}

	return res
}

const (
	SRC = '1'
	DIC = '2'
)

// 从 (row, col) 位置开始，去将 grid 中的 SRC -> DIC
func infect(grid [][]byte, row, col int) {
	if row < 0 || row == len(grid) || col < 0 || col == len(grid[0]) {
		// 都说明越界了
		return
	}

	if grid[row][col] != SRC {
		// 说明不是 SRC
		return
	}

	// 说明是 SRC，将其感染成 DIC，然后去感染它的上下左右
	grid[row][col] = DIC
	infect(grid, row-1, col) // 上
	infect(grid, row+1, col) // 下
	infect(grid, row, col-1) // 左
	infect(grid, row, col+1) // 右
}
