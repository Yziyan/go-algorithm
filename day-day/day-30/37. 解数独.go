// @Author: Ciusyan 1/15/24

package day_30

// https://leetcode.cn/problems/sudoku-solver/

func solveSudoku(board [][]byte) {

	var (
		// 先准备三个数组，去收集每一行、每一列、每一个桶，已经填写的信息
		rows    = make([][]bool, 9) // rows[3][6] = true 代表：第四行已经填写过 6 了
		cols    = make([][]bool, 9)
		buckets = make([][]bool, 9)
	)
	for i := 0; i < 9; i++ {
		rows[i] = make([]bool, 10)
		cols[i] = make([]bool, 10)
		buckets[i] = make([]bool, 10)
	}

	// 初始化以前填写数字的信息
	initMap(board, rows, cols, buckets)

	// 告诉 rows, cols, buckets 这些信息，从 (row, col) 位置开始，去填写 board
	var dfs func(row, col int, board [][]byte, rows, cols, buckets [][]bool) bool
	dfs = func(row, col int, board [][]byte, rows, cols, buckets [][]bool) bool {
		if row == 9 {
			// 说明九行都填写完成了
			return true
		}

		var (
			// 下一次要去的格子是 (nextRow, nextCol)，从左往右去填写
			nextRow = row     // 默认还是在这一行
			nextCol = col + 1 // 默认往后走一个格子
		)

		if col == 8 {
			// 如果到达一行的末尾了，要跳转到下一行了
			nextRow = row + 1
			nextCol = 0
		}

		if board[row][col] != '.' {
			// 说明已经填写过了，直接去下一个位置
			return dfs(nextRow, nextCol, board, rows, cols, buckets)
		}

		bucket := 3*(row/3) + (col / 3)
		// 来到这里，说明还没有填写，需要填写了再去填下一个格子，尝试 1~9 的所有可能
		for num := 1; num <= 9; num++ {
			// 但是在用之前，需要保证和 rows、cols、buckets 没有冲突
			if rows[row][num] || cols[col][num] || buckets[bucket][num] {
				// 只要有一个冲突，都不能选了
				continue
			}

			board[row][col] = byte(num) + '0'
			// 但是要标记这一个数字用过了
			rows[row][num] = true
			cols[col][num] = true
			buckets[bucket][num] = true
			// 然后去下一个格子填写
			if dfs(nextRow, nextCol, board, rows, cols, buckets) {
				// 说明这次尝试 num，最终让所有格子都填写完毕了，直接返回
				return true
			}

			// 否则就还原下现场，换下一个数字尝试
			board[row][col] = '.' // 这个可以不还原，反正下一个数字会覆盖掉它
			rows[row][num] = false
			cols[col][num] = false
			buckets[bucket][num] = false
		}
		// 说明不能完成这个数独，没有答案
		return false
	}

	// 主函数从 (row, col) 位置开始，去深度优先遍历
	dfs(0, 0, board, rows, cols, buckets)
}

// 初始化填写 rows、cols、buckets
func initMap(board [][]byte, rows, cols, buckets [][]bool) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				// 说明是待填写的格子
				continue
			}

			// 先计算出桶号
			bucket := 3*(row/3) + (col / 3)
			// 当前格子的数字
			num := board[row][col] - '0'

			// 在三个数组都要记录，num 这个数字出现过了。
			rows[row][num] = true
			cols[col][num] = true
			buckets[bucket][num] = true
		}
	}

}
