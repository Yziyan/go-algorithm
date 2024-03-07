// @Author: Ciusyan 2/1/24

package day_33

// https://leetcode.cn/problems/surrounded-regions/

const (
	SRC = 'O'
	DIC = 'X'

	TMP = 'T'
)

func solve(board [][]byte) {
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return
	}

	var (
		rowL = len(board)
		colL = len(board[0])
	)

	// 先尝试将边界的 SRC 感染成 TMP 字符
	for col := 0; col < colL; col++ {
		if board[0][col] == SRC {
			// 说明是第一行边界的 SRC，将其能联通的 SRC 都感染成 TMP
			infact(board, 0, col, SRC, TMP)
		}

		if board[rowL-1][col] == SRC {
			// 说明是最后一行边界的 SRC，将其能联通的 SRC 都感染成 TMP
			infact(board, rowL-1, col, SRC, TMP)
		}
	}

	for row := 0; row < rowL; row++ {
		if board[row][0] == SRC {
			// 说明是第一列边界的 SRC，将其能联通的 SRC 都感染成 TMP
			infact(board, row, 0, SRC, TMP)
		}

		if board[row][colL-1] == SRC {
			// 说明是最后一列边界的 SRC，将其能联通的 SRC 都感染成 TMP
			infact(board, row, colL-1, SRC, TMP)
		}
	}

	// 最后遍历 board，将被修改成 TMP 的换成 SRC，将还是 SRC 的改成 DIC
	for row := 0; row < rowL; row++ {
		for col := 0; col < colL; col++ {
			if board[row][col] == SRC {
				// 说明是被包围的 SRC，需要改成 DIC
				board[row][col] = DIC
			} else if board[row][col] == TMP {
				// 说明原先是边界的 SRC，没有闭口，不能被改成 DIC，需要还原
				board[row][col] = SRC
			}
		}
	}
}

// 从 board[row][col] 位置开始，进行感染。将能联通的 src 全部变成 dic
func infact(board [][]byte, row, col int, src, dic byte) {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		// 说明越界了
		return
	}

	if board[row][col] != src {
		// 说明都不是 src 字符，不需要感染
		return
	}

	// 来到这里，说明需要感染，先将当前位置感染了
	board[row][col] = dic
	// 再去上下左右四个方向感染
	infact(board, row-1, col, src, dic) // 上
	infact(board, row+1, col, src, dic) // 下
	infact(board, row, col-1, src, dic) // 左
	infact(board, row, col+1, src, dic) // 右
}
