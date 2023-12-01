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

// CoinsWayNoLimit 暴力尝试方法(从左往右的尝试方法)
func CoinsWayNoLimit(coins []int, aim int) int {
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
