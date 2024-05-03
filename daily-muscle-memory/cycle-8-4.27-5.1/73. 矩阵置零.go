// @Author: Ciusyan 4/29/24

package cycle_8_4_27_5_1

// https://leetcode.cn/problems/set-matrix-zeroes/

/*
*
思路重复：
若 (row, col) 为 0，我们想将 row 行和 col 列都置为 0。
一种思路是我们可以使用一个辅助数组，记录哪一行、哪一列都需要置零。
但是我们还可以复用以前的第一行和第一列，用这两个位置来记录要不要被覆盖为零。
1.先记录第一行和第一列是否需要被置为 0
2.然后遍历每一个位置，若为零，将 (row, 0) 和 (0, col) 置为 0，
3.将需要置零的行和列全部置为 0（被标记了的）
*/

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	rowL := len(matrix)
	colL := len(matrix[0])

	firstRow0 := false
	firstCol0 := false

	// 找到第一行是否为 0
	for col := 0; col < colL; col++ {
		if matrix[0][col] == 0 {
			firstRow0 = true
		}
	}
	// 找到第一列是否为 0
	for row := 0; row < rowL; row++ {
		if matrix[row][0] == 0 {
			firstCol0 = true
		}
	}

	for row := 1; row < rowL; row++ {
		for col := 1; col < colL; col++ {
			if matrix[row][col] == 0 {
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}

	// 挨行检查，需要置零的行
	for row := 1; row < rowL; row++ {
		if matrix[row][0] == 0 {
			for col := 1; col < colL; col++ {
				matrix[row][col] = 0
			}
		}
	}

	// 挨列检查，需要置零的列
	for col := 1; col < colL; col++ {
		if matrix[0][col] == 0 {
			for row := 1; row < rowL; row++ {
				matrix[row][col] = 0
			}
		}
	}

	// 第一行是否要置零
	if firstRow0 {
		for col := 0; col < colL; col++ {
			matrix[0][col] = 0
		}
	}
	// 第一列是否要置零
	if firstCol0 {
		for row := 0; row < rowL; row++ {
			matrix[row][0] = 0
		}
	}
}
