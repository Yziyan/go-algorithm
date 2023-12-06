// @Author: Ciusyan 12/6/23

package day_26

import "math"

/**
coins 是硬币面值数组，其中的值都是正数且没有重复。再给定一个正数 aim。
每个值都认为是一种面值，且认为张数是无限的。
返回组成 aim 的最少硬币数
*/

func minCoinsNoLimit(coins []int, aim int) int {
	if coins == nil || len(coins) == 0 || aim <= 0 {
		return math.MaxInt
	}
	n := len(coins)
	// 可以发现，就两个可变参数，cur 和 remain，他们的范围分别属于：
	// cur ∈ [0, n]，remain ∈ [0, aim]，所以准备缓存 dp，dp[cur][remain] 的含义是：
	// 当前处于 coins[cur] 这枚硬币，使用 [cur ...] 这么多枚硬币，找零 remain 的钱所需的最少张数。
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}

	// 根据递归基可知，当没有硬币的时候，只有不需要找零才要 0 枚硬币，否则都是无限枚
	dp[n][0] = 0
	for remain := 1; remain <= aim; remain++ {
		dp[n][remain] = math.MaxInt
	}

	// 根据依赖关系，cur 依赖 cur+1，所以需要从下往上求
	for cur := n - 1; cur >= 0; cur-- {
		for remain := 0; remain <= aim; remain++ {
			// 一般情况，看看使用 coins[cur] 多少张，但不管怎样，肯定不能超过 remain 的面值
			res := math.MaxInt
			for num := 0; num*coins[cur] <= remain; num++ {
				// 代表当前凑了 num*coins[cur] 的钱，剩下的 [cur+1 ...] 就只需要凑余下的钱
				next := dp[cur+1][remain-num*coins[cur]]
				if next != math.MaxInt {
					// 说明当前选择能凑出 remain，得到的张数 next，还需要 + num 枚 coins[cur]
					res = min(res, next+num)
				}
			}

			dp[cur][remain] = res
		}
	}

	// 当前是 coins[0] 这枚硬币，需要使用 [0 ...] 这些硬币，凑 aim 的钱，得到的最少硬币数量
	return dp[0][aim]
}

// 暴力递归方法，从左往右的尝试模型
func minCoinsNoLimit1(coins []int, aim int) int {
	if coins == nil || len(coins) == 0 || aim <= 0 {
		return math.MaxInt
	}

	// 先憋一个递归，从左往右的尝试模型，递归含义：
	// 当前 coins[cur] 这枚硬币，使用 [cur ...] 枚硬币，凑出正好 remain 的钱，所需要的最少硬币数量
	var process func(coins []int, cur, remain int) int
	process = func(coins []int, cur, remain int) int {
		if cur == len(coins) {
			// 说明没有钱了，
			if remain == 0 {
				// 说明没有钱需要找，那就是 0 枚硬币就可以咯
				return 0
			}
			return math.MaxInt
		}

		res := math.MaxInt
		// 那么对于一般情况，每一枚硬币，看看选几枚去尝试，但是不管选多少枚，一定不能超过 remain 的钱，要不然就多找了
		for num := 0; num*coins[cur] <= remain; num++ {
			// 代表 coins[cur] 这枚硬币选了 num 枚，剩下的 [cur+1 ...] 枚硬币，就需要凑出 remain - 当前的钱
			next := process(coins, cur+1, remain-num*coins[cur])
			if next != math.MaxInt {
				// 说明这样的选择，能够成功找零
				// 那么就用以前的结果 res，与此次选择得到的答案，取一个最小值，因为要求最小的数量嘛
				res = min(res, next+num)
			}
		}

		return res
	}

	// 那么主函数就应该：
	// 当前处于 coins[0] 这枚硬币，使用 [0 ...] 枚硬币，凑出正好 aim 的钱，所需要的最少硬币数量
	return process(coins, 0, aim)
}
