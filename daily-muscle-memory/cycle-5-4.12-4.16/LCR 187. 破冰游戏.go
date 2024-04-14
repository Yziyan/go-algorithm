// @Author: Ciusyan 4/13/24

package cycle_5_4_12_4_16

// https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/description/

/**
思路重复：
核心就是推导出这样一个公式：
f(n, m) = f(n-1, m) % n
如何推导的呢？明确两个前提：
1.一轮游戏中，不论剩几个人完，最终活下来的人都一是一个人
2.这是一个环，每次数 m 个，可能会越界，所以当越界后，又回到第一个人，也就是要对总人数取模

那么可以推导几个：
(0) 1 2 3 4 5		f(6, 3) = 0
3 4 5 (0) 1 		f(5, 3) = 3
(0) 1 3 4			f(4, 3) = 0
4 (0) 1				f(3, 3) = 1
4 (0) 				f(2, 3) = 1
0					f(1, 3) = 0
*/

func iceBreakingGame(num int, target int) int {
	// dp(1) = 0
	dp := 0
	for i := 2; i <= num; i++ {
		// dp(n, m) = dp(n-1, m) % n
		dp = (dp + target) % i
	}
	return dp
}
