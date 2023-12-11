// @Author: Ciusyan 12/11/23

package day_27

import "math"

/**
https://leetcode.cn/problems/n-queens-ii/submissions/488099115/

N 皇后问题是指在 N*N 的棋盘上要摆 N 个皇后，
要求任何两个皇后不同行、不同列， 也不在同一条斜线上
给定一个整数 n，返回 n 皇后的摆法有多少种。n=1，返回 1
n=2 或3，2 皇后和 3 皇后问题无论怎么摆都不行，返回 0
n=8，返回 92
*/

func nQueens(n int) int {
	if n < 1 {
		return 0
	}

	// 暴力枚举，递归含义：
	// 当前需要在第 row 行放皇后，record 记录者几行几列，放置了皇后。record[row] = col，
	// 返回能放置 [row ...] 成功放置的方法数
	var process func(n, row int, record []int) int
	process = func(n, row int, record []int) int {
		if row == n {
			// 说明所有行都放完了，居然没冲突，是一种放置方法
			return 1
		}

		// 总的结果数
		res := 0
		// n列 挨列尝试
		for col := 0; col < n; col++ {
			if !isValidate(record, row, col) {
				// 说明 (row, col) 位置不能放置皇后，会起冲突
				continue
			}

			// 说明能放置皇后，先记录 record
			record[row] = col
			// 再去下一行放置皇后
			res += process(n, row+1, record)
		}

		return res
	}

	record := make([]int, n)
	// 那么主函数如何调用呢？需要求解 [0 ... n-1] 行，能成功放置皇后的方法数量
	return process(n, 0, record)
}

// 验证是否能放置在 (row, col) 位置，
// record[row] = col，record[3] = 4，意味着在 3行4列 已经放置了皇后
func isValidate(record []int, row, col int) bool {

	// 需要验证 0~row-1 行，列、斜线是否会冲突
	for k := 0; k < row; k++ {
		// 需要验证 col 列是否会冲突
		if record[k] == col {
			// 说明 col 列已经有行填写过了
			return false
		}
		// 需要验证斜线是否会冲突，求斜率，是那两个点呢？
		// (row, col), (k, record[k])，斜率如果是 -1 or +1，就说明在斜线上
		// k = (row - k) / (col - record[k]) == +1 or -1，
		// => row - k == col - record[k] or row - k == -(col - record[k])
		// => row - k == col - record[k] or row - k == record[k] - col)
		// => row - k == math.Abs(col - record[k])
		if row-k == int(math.Abs(float64(col-record[k]))) {
			// 说明在斜线上，已经有地方填过了
			return false
		}
	}

	return true
}
