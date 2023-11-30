// @Author: Ciusyan 11/30/23

package day_25

/**
arr 是货币数组，其中的值都是正数。再给定一个正数 aim。
每个值都认为是一张货币，
即便是值相同的货币也认为每一张都是不同的，
返回组成 aim 的方法数
例如：arr = {1,1,1}，aim = 2
第 0个 和第 1个 能组成 2，第 1个 和第 2个 能组成 2，第 0个 和第 2个 能组成 2
一共就 3 种方法，所以返回 3
*/

// CoinsWayEveryPaperDifferent 动态规划方法（一维数组）
func CoinsWayEveryPaperDifferent(coins []int, aim int) int {
	if coins == nil || len(coins) == 0 || aim <= 0 {
		return 0
	}

	n := len(coins)

	// 因为 cur 只依赖 cur+1，并且后面只依赖前面的某个值，所以可以简化二维数组
	dp := make([]int, aim+1)
	// 代表 cur == n 时，remain == 0，有一种方法数
	dp[0] = 1

	// 下面依赖上面，从下往上求解
	for cur := n - 1; cur >= 0; cur-- {
		// 因为后面依赖前面，所以从后往前求解
		for remain := aim; remain >= 0; remain-- {
			if remain-coins[cur] < 0 {
				// 说明使用当前零钱会多找，没必要。
				continue
			}
			// 不选当前零钱，那么方法数就是以前的 dp[remain]
			// 选当前零钱，并且选了不会多找，那么方法数就加上 dp[remain-coins[cur]]
			dp[remain] += dp[remain-coins[cur]]
		}
	}

	return dp[aim]
}

// CoinsWayEveryPaperDifferent2 动态规划方法（二维数组）
func CoinsWayEveryPaperDifferent2(coins []int, aim int) int {
	if coins == nil || len(coins) == 0 || aim <= 0 {
		return 0
	}

	n := len(coins)
	// 发现可变参数就俩：cur ∈ [0, n] remain ∈ [0, aim]
	// 准备缓存 dp，dp[cur][remain] 的含义是：使用 coins[cur ...] 找零 remain，拥有多少种方案
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	// 根据递归基可知，当 cur == n 时，只有 remain == 0，才有一种方案，其余的全是 0
	dp[n][0] = 1
	// 根据依赖关系，cur 依赖 cur+1，所以应该从下往上求解
	for cur := n - 1; cur >= 0; cur-- {
		for remain := 0; remain <= aim; remain++ {
			// 有两种方案，要么选 cur，要么不选 cur
			p1 := 0
			if remain-coins[cur] >= 0 {
				// 要保证选择当前零钱后，不能多找钱才算
				p1 = dp[cur+1][remain-coins[cur]]
			}
			// 不选
			p2 := dp[cur+1][remain]

			// 两种方案累加起来
			dp[cur][remain] = p1 + p2
		}
	}

	// 那么主函数应该代表：使用 coins[0 ...] 找零 aim，拥有多少种方案
	return dp[0][aim]
}

// CoinsWayEveryPaperDifferent1 每一张钱都不同，只有 len(coins) 张钱，需要找零 aim，有多少种找零方案
// 暴力递归方法（从左往右的尝试模型）
func CoinsWayEveryPaperDifferent1(coins []int, aim int) int {
	if coins == nil || len(coins) == 0 || aim <= 0 {
		return 0
	}

	// 憋一个递归，递归含义是：当前拿到的零钱是第 cur 张，需要使用 coins[cur ...] 找零 remain 的钱，有多少种找零方案。
	var process func(coins []int, cur int, remain int) int
	process = func(coins []int, cur int, remain int) int {
		if remain < 0 {
			// 说明这种方案会多找，不可行
			return 0
		}
		if cur == len(coins) {
			// 如果用完了最后一张钱，需要看看还有 remain 吗
			if remain == 0 {
				// 说明没有剩余，之前的找零方案可行
				return 1
			}
			return 0
		}

		// 那么一般情况呢？就是要么选当前这张钱来找，要么不选，总的找零方案就是两者相加
		p1 := process(coins, cur+1, remain-coins[cur]) // 选
		p2 := process(coins, cur+1, remain)            // 不选
		return p1 + p2
	}

	// 那么主函数就应该这样调用：使用 coins[0 ...] 找零 aim 的钱，有多少种找零方案
	return process(coins, 0, aim)
}
