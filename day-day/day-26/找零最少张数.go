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
