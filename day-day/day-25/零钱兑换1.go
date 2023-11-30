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

// CoinsWayEveryPaperDifferent 每一张钱都不同，只有 len(coins) 张钱，需要找零 aim，有多少种找零方案
func CoinsWayEveryPaperDifferent(coins []int, aim int) int {
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
