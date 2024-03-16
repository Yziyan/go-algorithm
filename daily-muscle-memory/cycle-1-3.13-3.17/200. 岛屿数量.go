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

/*
*
思路重复，伪代码
这里使用感染的方式来做，
我们去检查相邻岛屿的时候，我们可以在检查的时候，就将相邻的岛屿直接感染成其他（一会别干扰已经连成岛屿的位置）
比如我们从 (0, 0) 出发，将是岛屿的位置，进行感染，全部变成 岛屿`，
那么和 (0, 0) 相邻的岛屿，全被变成 岛屿` 了，之后就不会被识别为岛屿了，就不会被感染了。
那么我们进行感染的次数，就是相邻岛屿的数量。

那么感染如何实现呢？
使用递归即可，从 (row, col) 出发，只要保证 row 和 col 不越位，并且这个位置是岛屿，那就先将其变成 岛屿`，
然后去 (row, col) 的上下左右都进行感染。直至感染结束。即可完成
*/
func numIslands2(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var infect2 func(grid [][]byte, row, col int)
	infect2 = func(grid [][]byte, row, col int) {
		if row == -1 || row == len(grid) || col == -1 || col == len(grid[0]) {
			// 保证不能越位
			return
		}

		if grid[row][col] != '1' {
			// 说明不是岛屿
			return
		}
		// 改变自己为岛屿`
		grid[row][col] = '2'
		// 去感染上下左右
		infect2(grid, row-1, col)
		infect2(grid, row+1, col)
		infect2(grid, row, col-1)
		infect2(grid, row, col+1)
	}

	// 挨个位置进行检测，是否需要感染
	res := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == '1' {
				// 说明需要感染
				res++
				infect(grid, row, col)
			}
		}
	}

	return res
}
