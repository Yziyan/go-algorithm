// @Author: Ciusyan 4/13/24

package cycle_5_4_12_4_16

// https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/description/

func iceBreakingGame(num int, target int) int {
	// dp(1) = 0
	dp := 0
	for i := 2; i <= num; i++ {
		// dp(n, m) = dp(n-1, m) % n
		dp = (dp + target) % i
	}
	return dp
}
