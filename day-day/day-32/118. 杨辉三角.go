// @Author: Ciusyan 1/29/24

package day_32

// https://leetcode.cn/problems/pascals-triangle/

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}

	// 某层的长度，就是层数
	res := make([][]int, numRows)

	// 从第 1 层开始填写
	for row := 1; row <= numRows; row++ {
		curLevel := make([]int, row)
		// 第一个位置，肯定是 1
		curLevel[0] = 1
		// 填中间的位置
		for j := 1; j < row-1; j++ {
			// 是上一层的：上+左上
			curLevel[j] = res[row-2][j] + res[row-2][j-1]
		}
		// 最后一个位置，肯定是 1
		curLevel[row-1] = 1

		res[row-1] = curLevel
	}

	return res
}
