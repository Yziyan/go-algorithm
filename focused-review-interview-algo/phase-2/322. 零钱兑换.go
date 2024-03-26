// @Author: Ciusyan 3/26/24

package phase_2

import "math"

func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	for remain := 1; remain <= amount; remain++ {
		dp[n][remain] = -1
	}

	for cur := n - 1; cur >= 0; cur-- {
		for remain := 0; remain <= amount; remain++ {
			// 因为每一张硬币都可以使用若干次，所以到底用几次呢？
			minRes := math.MaxInt
			for num := 0; num*coins[cur] <= remain; num++ {
				// 选 num 张，但是要确保，num 张后不超过总钱数
				next := dp[cur+1][remain-num*coins[cur]]
				if next != -1 {
					minRes = min(minRes, num+next)
				}
			}

			if minRes == math.MaxInt {
				minRes = -1
			}

			dp[cur][remain] = minRes
		}
	}

	return dp[0][amount]
}

func coinChange2(coins []int, amount int) int {
	// 利用 [cur ... ] 凑出 remain 的钱，最少的硬币数是多少
	var process func(coins []int, cur, remain int) int
	process = func(coins []int, cur, remain int) int {
		if cur == len(coins) {
			// 说明钱用完了
			if remain == 0 {
				// 说明刚好可以筹够，现在就不需要筹钱了
				return 0
			}
			return -1
		}

		// 因为每一张硬币都可以使用若干次，所以到底用几次呢？
		minRes := math.MaxInt
		for num := 0; num*coins[cur] <= remain; num++ {
			// 选 num 张，但是要确保，num 张后不超过总钱数
			next := process(coins, cur+1, remain-num*coins[cur])
			if next != -1 {
				minRes = min(minRes, num+next)
			}
		}

		if minRes == math.MaxInt {
			return -1
		}

		return minRes
	}

	return process(coins, 0, amount)
}
