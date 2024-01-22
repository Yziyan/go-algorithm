// @Author: Ciusyan 1/22/24

package day_32

// https://leetcode.cn/problems/word-search/

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 {
		return false
	}

	// 定义一个深度优先遍历的递归函数
	// 当前处于 board 的 (row, col) 位置，想要能否搜索出 word[cur ...] 的单词？
	var process func(board [][]byte, word []byte, row, col, cur int) bool
	process = func(board [][]byte, word []byte, row, col, cur int) bool {
		if cur == len(word) {
			// 说明搜索到达末尾了
			return true
		}

		if (row < 0 || row == len(board)) || (col < 0 || col == len(board[0])) {
			// 说明越界了
			return false
		}

		// 否则准备往上下左右四个方向去搜索 word[cur+1 ...]
		temp := board[row][col]
		if temp != word[cur] {
			// 如果开始搜索的字符，和单词的当前字符都不相等，没必要搜素了
			return false
		}

		// 不能走回头路，先将其设置为无效字符
		board[row][col] = 0

		if process(board, word, row-1, col, cur+1) || // 向上搜索
			process(board, word, row+1, col, cur+1) || // 向下搜索
			process(board, word, row, col-1, cur+1) || // 向左搜索
			process(board, word, row, col+1, cur+1) { // 向右搜索
			// 只要有一个方向能搜索出来，word[cur+1 ...] 就说明能找到 word
			return true
		}

		// 来到这里说明四个方向都没有搜索到，退出时还原现场
		board[row][col] = temp

		return false
	}

	var (
		rowL  = len(board)
		colL  = len(board[0])
		chars = []byte(word)
	)

	for row := 0; row < rowL; row++ {
		for col := 0; col < colL; col++ {
			// 尝试从每一个位置开始搜索 word[0 ...] 的单词
			if process(board, chars, row, col, 0) {
				// 只要有一条路径能搜索出来，就说明存在 word
				return true
			}
		}
	}

	// 所有位置都不能搜索出 word
	return false
}
