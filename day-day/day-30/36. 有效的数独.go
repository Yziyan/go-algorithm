// @Author: Ciusyan 1/14/24

package day_30

// https://leetcode.cn/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	// 准备三个数组，分别用于记录：行、列、桶，是否出现过某个数字了
	var (
		rows    = make([][]bool, 9) // rows[5][3] = true，代表第 6 行已经有 3 这个数了
		cols    = make([][]bool, 9) // cols[2][9] = false，代表第 3 列还没有 9 这个数字
		buckets = make([][]bool, 9)
	)

	for j := 0; j < 9; j++ {
		rows[j] = make([]bool, 10)
		cols[j] = make([]bool, 10)
		buckets[j] = make([]bool, 10)
	}

	// 数独固定式 9*9 的，所以直接写了
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				// 说明当前格子没填数字，不用检验
				continue
			}

			// 如何求解 (row, col) 属于哪一个桶呢？
			bucket := 3*(row/3) + (col / 3)
			// 然后对比当前格子的数字，在相同行、列、桶中有没有出现过
			curNum := board[row][col] - '0'
			if rows[row][curNum] || cols[col][curNum] || buckets[bucket][curNum] {
				// 只要有一个存在过，就说明不是有效的数独
				return false
			}

			// 来到这里，说明之前没出现过，但是现在出现过了
			rows[row][curNum] = true
			cols[col][curNum] = true
			buckets[bucket][curNum] = true
		}
	}

	// 说明全部检验完了，有效
	return true
}
