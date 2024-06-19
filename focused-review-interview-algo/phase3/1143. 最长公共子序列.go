// @Author: Ciusyan 6/19/24

package phase3

func longestCommonSubsequence(text1 string, text2 string) int {
	rowL := len(text1)
	colL := len(text2)

	dp := make([][]int, rowL+1)
	for i := range dp {
		dp[i] = make([]int, colL+1)
	}

	for row := 1; row <= rowL; row++ {
		for col := 1; col <= colL; col++ {
			p1 := dp[row-1][col]
			p2 := dp[row][col-1]
			p3 := 0
			if text1[row-1] == text2[col-1] {
				p3 = dp[row-1][col-1] + 1
			}

			dp[row][col] = max(p1, p2, p3)
		}
	}

	return dp[rowL][colL]
}
