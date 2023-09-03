// @Author: Ciusyan 2023/9/3

package day_4

// https://leetcode.cn/problems/word-search/

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	// 遍历所有字符
	chars := []byte(word)
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			// 从第一个字符开始去 board 中搜索，并且从 (row, col) 开始
			if dfs(0, board, chars, row, col) {
				return true
			}
		}
	}

	return false
}

// level 代表第几个字符
func dfs(level int, board [][]byte, word []byte, row, col int) bool {
	if level == len(word) {
		// 说明单词的所有字符都搜完了，找到了一个结果
		return true
	}

	// 如果越界了还没搜到，那就说明搜不到了
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return false
	}

	// 如果 (row, col) 这个位置的字符与单词的第 level 个字符都匹配不上，也没必要搜了
	if word[level] != board[row][col] {
		return false
	}

	// 到这里尝试从上下左右去搜索，
	//  但是在搜索前，需要清除 (row, col) 位置的字符，因为不能往回搜索
	temp := board[row][col]
	board[row][col] = 0

	leftExist := dfs(level+1, board, word, row, col-1)
	rightExist := dfs(level+1, board, word, row, col+1)
	topExist := dfs(level+1, board, word, row-1, col)
	buttomExist := dfs(level+1, board, word, row+1, col)

	// 如果上下左右四个方向有一个方向搜到了，那么就说明搜完了
	if leftExist || rightExist || topExist || buttomExist {
		return true
	}

	// 但是在搜完后，需要还原现场
	board[row][col] = temp

	return false
}
