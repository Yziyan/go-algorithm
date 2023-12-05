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

func KillMonster(n, m, k int) float64 {
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
