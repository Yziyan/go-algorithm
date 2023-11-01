// @Author: Ciusyan 11/1/23

package day_22

// https://leetcode.cn/problems/number-of-islands/description/

// 并查集做法
func numIslands(grid [][]byte) int {

	rowLen := len(grid)
	colLen := len(grid[0])

	uf := NewUnionFind[*byte]()
	// 将所有岛屿都加入并查集中
	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if grid[row][col] == '1' {
				// 这里使用 map，键不一样才行，所以用 grid[row][col] 的地址
				uf.MakeSets(&grid[row][col])
			}
		}
	}

	// 现在总共有那么多的岛屿，但是挨个查看，是不是相邻的，如果是相邻的，那么我们将其合并
	for col := 1; col < colLen; col++ {
		// 先合并第一行，只需要看左边即可
		if grid[0][col-1] == '1' && grid[0][col] == '1' {
			uf.Union(&grid[0][col-1], &grid[0][col])
		}
	}
	for row := 1; row < rowLen; row++ {
		// 再合并第一列，只需要看上边即可
		if grid[row-1][0] == '1' && grid[row][0] == '1' {
			uf.Union(&grid[row-1][0], &grid[row][0])
		}
	}

	for row := 1; row < rowLen; row++ {
		for col := 1; col < colLen; col++ {
			// 最后再来看其他的，上面和左边
			if grid[row][col] == '1' {
				if grid[row][col-1] == '1' {
					// 上面
					uf.Union(&grid[row][col], &grid[row][col-1])
				}

				if grid[row-1][col] == '1' {
					// 左边
					uf.Union(&grid[row][col], &grid[row-1][col])
				}
			}
		}
	}

	return uf.GetSize()
}

// 染色做法
func numIslands2(grid [][]byte) int {

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
