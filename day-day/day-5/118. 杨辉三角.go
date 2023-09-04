// @Author: Ciusyan 2023/9/4

package day_5

// https://leetcode.cn/problems/pascals-triangle/

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}

	res := make([][]int, numRows)
	// 先把第一层加入
	res[0] = []int{1}

	for i := 1; i < numRows; i++ {
		// 取出上一层的值
		prevLevel := res[i-1]
		// 创建新一层的结果，并且将第一个数添加到里面
		curLevel := make([]int, i+1)
		curLevel[0] = 1

		for j := 1; j < i; j++ {
			curLevel[j] = prevLevel[j-1] + prevLevel[j]
		}
		// 将最后一个数也设置为 1
		curLevel[i] = 1

		res[i] = curLevel
	}

	return res
}
