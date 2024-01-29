// @Author: Ciusyan 1/29/24

package day_32

// https://leetcode.cn/problems/pascals-triangle-ii/

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}

	// 因为索引是从 0 开始的，我们将其转换成：1 就是第 1 层，3 就是第 3 层
	rowIndex++
	// 长度就是层数。
	res := make([]int, rowIndex)
	// 从第一层，依次滚动的求下去
	for row := 1; row <= rowIndex; row++ {
		// 因为用一个数组要滚动求解，所以得从后往前求（依赖：左上+上）
		// 该层的最后一个位置，肯定是 1
		res[row-1] = 1
		for j := row - 2; j >= 1; j-- {
			res[j] = res[j] + res[j-1] // 上 + 左上
		}
		// 该层的第一个位置，肯定是 1
		res[0] = 1
	}

	return res
}
