// @Author: Ciusyan 12/1/23

package day_25

/**
arr 是面值数组，其中的值都是正数且没有重复。再给定一个正数 aim。
每个值都认为是一种面值，且认为张数是无限的。
返回组成 aim 的方法数
例如：arr = {1,2}，aim = 4
方法如下：1+1+1+1、1+1+2、2+2
一共就 3 种方法，所以返回 3
*/

// CoinsWayNoLimit 动态规划方法
func CoinsWayNoLimit(coins []int, aim int) int {
	if coins == nil || len(coins) == 0 || aim <= 0 {
		return 0
	}

	n := len(coins)
	// 可变参数是 cur 和 remain，建立缓存
	// dp[cur][remain] 的含义是：
	// 使用 coins[cur ...] 枚硬币，想要筹够 remain 的钱，有多少种方法
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}

	// 根据递归基可知，当没钱可用时，钱也筹够才有一种方法数
	dp[n][0] = 1

	// 根据依赖关系，cur 依赖 cur+1，所以需要从下往上求
	for cur := n - 1; cur >= 0; cur-- {
		for remain := 0; remain <= aim; remain++ {

			ways := 0
			// 要使用 cur 这枚硬币多少次，但不管使用多少枚，都不能比余额还大
			for num := 0; num*coins[cur] <= remain; num++ {
				ways += dp[cur+1][remain-num*coins[cur]]
			}

			dp[cur][remain] = ways
		}
	}

	// 返回 dp[0][aim]，代表使用 coins[0 ...] 枚硬币，筹够 aim 钱，有多少种方法
	return dp[0][aim]
}

// CoinsWayNoLimit1 暴力尝试方法(从左往右的尝试方法)
func CoinsWayNoLimit1(coins []int, aim int) int {
	if coins == nil || len(coins) == 0 || aim <= 0 {
		return 0
	}

	// 先憋一个暴力递归方法，递归含义是：
	// 使用 coins[cur ...] 枚硬币，需要筹够 remain，返回有多少种筹钱方法
	var process func(coins []int, cur, remain int) int
	process = func(coins []int, cur, remain int) int {
		if cur == len(coins) {
			// 说明最后一张钱都用完了
			if remain != 0 {
				// 说明还没有筹够，那也没方案了
				return 0
			}

			return 1
		}

		// 对于一般情况，每一枚硬币都可以使用无限次，那么使用多少次呢？
		ways := 0
		// 要用多少张，需要一张一张的试，但是呢？试也不能超过 remain
		for num := 0; num*coins[cur] <= remain; num++ {
			// 那么接下来，就是使用：coins[cur+1 ...] 枚硬币，筹出 remain-num*coins[cur]
			ways += process(coins, cur+1, remain-num*coins[cur])
		}

		return ways
	}

	// 那么主函数应该如何调用呢？
	// 使用 coins[0 ...] 枚硬币，需要筹够 aim，返回多少种方法
	return process(coins, 0, aim)
}
