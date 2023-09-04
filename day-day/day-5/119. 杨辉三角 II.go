// @Author: Ciusyan 2023/9/4

package day_5

// https://leetcode.cn/problems/pascals-triangle-ii/

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}

	res := make([]int, rowIndex+1)
	// 先将第一行添加了
	res[0] = 1

	// 从上往下计算
	for i := 1; i <= rowIndex; i++ {

		// res 就是上一层的值，我们想要复用内存的话，
		// 需要从后往前计算，否则会覆盖掉原先的值，那么现将最后一个位置设置为 1
		res[i] = 1
		for j := i - 1; j > 0; j-- {
			res[j] = res[j] + res[j-1]
		}

	}

	return res
}
