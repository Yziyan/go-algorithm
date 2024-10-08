// @Author: Ciusyan 2024/8/12

package cycle_25_8_10_8_14

// https://leetcode.cn/problems/valid-sudoku/description/

/**
思路重复：
有效的数独就是：行、列、桶 没有相同的数字。

那么我们准备三个东西来标记这个三个方向中都有什么数字。
比如 rows[2][5] 说明，第三行中已经有 5 这个数字了。
那么我们在便利数独的时候，如果是数字，就先检查标记中有没有违规的，如果没有就标记后进入下一个字符的检验。
*/

func isValidSudoku(board [][]byte) bool {
	// 准备三个数组，分别用来标识：行、列、桶 中已有的元素
	var (
		rows    = [9][10]bool{} // rows[3][1] 代表第四行有 1 这个数字了
		cols    = [9][10]bool{}
		buckets = [9][10]bool{} // buckets[4][3] 代表第五个桶，有 3 这个数字了
	)

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] == '.' {
				// 说明还没填数字，直接跳过
				continue
			}

			// 桶的计算公式
			bucket := 3*(row/3) + (col / 3)
			num := board[row][col] - '0'
			if rows[row][num] || cols[col][num] || buckets[bucket][num] {
				// 只要 行、列、桶 三个地方有一个地方有这个数字了，都说明是无效的数独
				return false
			}
			// 来到这里，将其行列桶标记当前数字
			rows[row][num] = true
			cols[col][num] = true
			buckets[bucket][num] = true
		}
	}

	return true
}
