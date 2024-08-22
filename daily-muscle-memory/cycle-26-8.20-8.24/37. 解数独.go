// @Author: Ciusyan 2024/8/22

package cycle_26_8_20_8_24

// https://leetcode.cn/problems/sudoku-solver/

/**
思路重复：
我们要解数独，就得先知道数独现在的情况是怎样的。
所以我们可以便利每一个格子，将其状态保存至 rows、cols、buckets 中。
然后使用 dfs 的方式，从 (row, col) 位置开始，利用 rows、cols、buckets 来尝试填写当前的格子。
如果当前格子已经填写过了，直接去到下一个格子。否则从 1～9 开始尝试，挨个尝试填写，但是填写的时候，
必须要保重三个数组中没有当前数字。
填写完后，去下一个格子开始搜索，如果能填完，那么直接返回即可。否则说明之前某个选择不合适。
还原现场后，回到当时的格子，去到下一个数字进行尝试。直至成功。除非全部便利完都没有成功，那么说明无解，直接返回即可
*/

func solveSudoku(board [][]byte) {
	rows, cols, buckets := initBoard(board)
	dfs(board, rows, cols, buckets, 0, 0)
}

// 记录棋盘现有的情况
func initBoard(board [][]byte) (rows, cols, buckets [9][10]bool) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}
			bucket := 3*(row/3) + (col / 3)
			num := board[row][col] - '0'

			rows[row][num] = true
			cols[col][num] = true
			buckets[bucket][num] = true
		}
	}

	return
}

// 基于 rows, cols, buckets，从 (row, col) 位置出发，去解数独，能解出返回 true
func dfs(board [][]byte, rows, cols, buckets [9][10]bool, row, col int) bool {
	if row == 9 {
		// 说明全部填完了
		return true
	}
	// 先搞清楚下一个位置去哪里，从上往下，从左往右走
	nextRow, nextCol := row, col+1
	if col == 8 {
		// 说明要换行了
		nextRow = row + 1
		nextCol = 0
	}

	if board[row][col] != '.' {
		// 说明这个位置已经填过了，直接去下一个位置解
		return dfs(board, rows, cols, buckets, nextRow, nextCol)
	}

	bucket := 3*(row/3) + (col / 3)

	// 说明这里需要填写数字，从 1～9 挨个尝试
	for num := 1; num <= 9; num++ {
		if rows[row][num] || cols[col][num] || buckets[bucket][num] {
			// 说明这个数字用过了，直接用下一个数字
			continue
		}

		// 说明可以使用当前数字，先用为敬
		board[row][col] = byte(num) + '0'
		rows[row][num] = true
		cols[col][num] = true
		buckets[bucket][num] = true

		// 用完直接去下一个位置填写
		if dfs(board, rows, cols, buckets, nextRow, nextCol) {
			// 如果能解出来，说明当初的选择是正确的，直接返回了
			return true
		}

		// 否则需要还原现场
		board[row][col] = '.'
		rows[row][num] = false
		cols[col][num] = false
		buckets[bucket][num] = false
	}

	// 说明解不出来
	return false
}
