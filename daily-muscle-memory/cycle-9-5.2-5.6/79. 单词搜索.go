// @Author: Ciusyan 5/3/24

package cycle_9_5_2_5_6

// https://leetcode.cn/problems/word-search/

func exist(board [][]byte, word string) bool {

	// 从 (row, col) 开始搜，若搜索到 word[cur ...] 的字符，则返回 true
	var dfs func(board [][]byte, word string, row, col, cur int) bool
	dfs = func(board [][]byte, word string, row, col, cur int) bool {
		if cur == len(word) {
			// 说明之前搜索到了
			return true
		}

		if row == -1 || row == len(board) || col == -1 || col == len(board[0]) {
			// 说明越界了
			return false
		}

		temp := board[row][col]
		if temp != word[cur] {
			// 说明字符不相等
			return false
		}

		// 来到这里说明相等，先把当前位置暂时设置为无效字符。
		board[row][col] = 0
		// 然后去 dfs，
		next := dfs(board, word, row-1, col, cur+1) || dfs(board, word, row+1, col, cur+1) ||
			dfs(board, word, row, col-1, cur+1) || dfs(board, word, row, col+1, cur+1)

		// 然后还原现场
		board[row][col] = temp

		return next
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			// 每一个位置都尝试一次，直到搜到为止
			if dfs(board, word, row, col, 0) {
				return true
			}
		}
	}

	return false
}
