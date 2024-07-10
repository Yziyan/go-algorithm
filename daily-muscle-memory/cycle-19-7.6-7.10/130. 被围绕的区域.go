// @Author: Ciusyan 2024/7/10

package cycle_19_7_6_7_10

// https://leetcode.cn/problems/surrounded-regions/description/

func solve(board [][]byte) {
	const (
		DIC = 'X'
		SRC = 'O'
		TMP = 'T'
	)

	var infect func(board [][]byte, row, col int, src, dic byte)
	infect = func(board [][]byte, row, col int, src, dic byte) {
		if row == -1 || row == len(board) || col == -1 || col == len(board[0]) {
			return
		}

		if board[row][col] != src {
			return
		}

		board[row][col] = dic
		// 修改完当前位置后，再对当前位置的上下左右进行感染
		infect(board, row-1, col, src, dic)
		infect(board, row+1, col, src, dic)
		infect(board, row, col-1, src, dic)
		infect(board, row, col+1, src, dic)
	}

	rowL := len(board)
	colL := len(board[0])

	// 1.先对边界进行感染，将边界能到达的 O，感染成 T
	for row := 0; row < rowL; row++ {
		// 第一列
		if board[row][0] == SRC {
			infect(board, row, 0, SRC, TMP)
		}

		// 最后一列
		if board[row][colL-1] == SRC {
			infect(board, row, colL-1, SRC, TMP)
		}
	}

	for col := 0; col < colL; col++ {
		// 第一行
		if board[0][col] == SRC {
			infect(board, 0, col, SRC, TMP)
		}
		// 最后一行
		if board[rowL-1][col] == SRC {
			infect(board, rowL-1, col, SRC, TMP)
		}
	}

	// 2.感染完成后，将 TMP -> SRC | SRC -> DIC 即可
	for row := 0; row < rowL; row++ {
		for col := 0; col < colL; col++ {
			cur := board[row][col]
			if cur == SRC {
				cur = DIC
			} else if cur == TMP {
				cur = SRC
			}
			board[row][col] = cur
		}
	}
}
