// @Author: Ciusyan 12/5/23

package day_26

import "math"

/**
给定 3 个参数，N，M，K
怪兽有 N 滴血，等着英雄来砍自己
英雄每一次打击，都会让怪兽流失 [0~M] 的血量
到底流失多少？每一次在 [0~M] 上等概率的获得一个值
求 K 次打击之后，英雄把怪兽砍死的概率
*/

// KillMonster 动态规划方法，优化版本
func KillMonster(n, m, k int) float64 {
	if n < 1 || m < 1 || k < 1 {
		// 说明：怪兽本身就是死的 or 永远砍不出伤害 or 要砍的刀数为 0
		return 0
	}

	// 其他的还是一样的含义
	dp := make([][]int64, k+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}

	for remain := 0; remain <= k; remain++ {
		// 不管有没有刀数，怪兽肯定已经死掉了，那么再砍 remain 刀，都是在鞭尸，所有的砍法都是答案
		dp[remain][0] = int64(math.Pow(float64(m+1), float64(remain)))
	}

	for remain := 1; remain <= k; remain++ {
		for hp := 1; hp <= n; hp++ {
			// 根据优化推导的转移方程
			ways := dp[remain-1][hp] + dp[remain][hp-1]
			if hp-(m+1) <= 0 {
				// 说明怪兽肯定死了，剩余的 remain 都是在鞭尸，也说明索引会越界
				ways -= int64(math.Pow(float64(m+1), float64(remain-1)))
			} else {
				// 说明不越界
				ways -= dp[remain-1][hp-(m+1)]
			}
			dp[remain][hp] = ways
		}
	}

	kill := dp[k][n]
	all := math.Pow(float64(m+1), float64(k))
	return float64(kill) / all
}

// KillMonster2 动态规划方法，普通版本
func KillMonster2(n, m, k int) float64 {
	if n < 1 || m < 1 || k < 1 {
		// 说明：怪兽本身就是死的 or 永远砍不出伤害 or 要砍的刀数为 0
		return 0
	}

	// 有两个可变参数：remain 和 hp，它们的范围分别是：
	// remain ∈ [0, k]，hp ∈ [负数, n]
	// 但是对于 hp 为负数的情况，我们可以通过代码处理，所以 hp ∈ [0, n]
	// 准备缓存 dp，dp[remain][hp] 的含义是：
	// 还剩余 remain 刀要砍掉怪兽 hp 滴血，每刀能掉 0~m 滴血，砍完 remain 刀怪兽死掉的数量
	dp := make([][]int64, k+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}

	// 根据递归基：当刀数为零的时候，只有怪兽的血 <= 0 的时候才算死掉了
	dp[0][0] = 1 // dp[0][...] = 0
	for remain := 1; remain <= k; remain++ {
		// 当 hp = 0 时，怪兽已经死掉了，那么余下的 remain 刀都是在鞭尸，每一种伤害，都是结果
		dp[remain][0] = int64(math.Pow(float64(m+1), float64(remain)))
	}

	// 根据依赖关系，remain 依赖 remain-1，所以需要从上往下求
	for remain := 1; remain <= k; remain++ {
		for hp := 1; hp <= n; hp++ {
			// 对于一般情况，需要看 remain 这一刀，砍掉的血量
			ways := int64(0)
			for hurt := 0; hurt <= m; hurt++ {
				if hp-hurt <= 0 {
					// 说明砍这一刀，怪兽绝对死掉了，剩余 remain-1 刀都是在鞭尸
					ways += int64(math.Pow(float64(m+1), float64(remain-1)))
				} else {
					// 否则依赖 remain-1
					ways += dp[remain-1][hp-hurt]
				}
			}

			dp[remain][hp] = ways
		}
	}

	// 剩余 k 刀要砍掉怪兽 n 滴血，每刀能掉 0~m 滴血，砍完 k 刀怪兽死掉的数量
	kill := dp[k][n]
	all := math.Pow(float64(m+1), float64(k))
	return float64(kill) / all
}

// KillMonster1 暴力递归方法
func KillMonster1(n, m, k int) float64 {
	if n < 1 || m < 1 || k < 1 {
		// 说明：怪兽本身就是死的 or 永远砍不出伤害 or 要砍的刀数为 0
		return 0
	}

	// 先憋一个暴力递归，递归含义是：
	// 剩余 remain 刀，每一刀能砍掉 0~m 滴血，怪兽现在还剩 hp 滴血，返回砍完 remain 刀，能把怪兽砍死的总数
	var process func(remain, m, hp int) int64
	process = func(remain, m, hp int) int64 {
		if remain == 0 {
			// 说明没有刀数了
			if hp > 0 {
				// 说明怪兽没死
				return 0
			}
			return 1
		}

		if hp <= 0 {
			// 说明怪兽已经死了，现在余下的 remain 刀都在鞭尸
			return int64(math.Pow(float64(m+1), float64(remain)))
		}

		// 对于一般情况，有多少种方式呢？
		ways := int64(0)
		// 对于砍第 remain 这一刀，能够掉 hurt 点血
		// 那么砍死怪兽的总数，也就是将所有情况累加起来
		for hurt := 0; hurt <= m; hurt++ {
			ways += process(remain-1, m, hp-hurt)
		}

		return ways
	}

	// 如何求解概率呢？(砍 k 刀把怪兽砍死的数量) / 砍出的所有伤害组合的总数
	// 所以如何求解分子呢？代表：
	// 还剩余 k 刀，每刀能砍掉 0~m 的血，怪兽现在还剩 n 滴血。返回砍完 k 刀，能够把怪兽砍死的总数
	kill := process(k, m, n)
	all := math.Pow(float64(m+1), float64(k))

	return float64(kill) / all
}
