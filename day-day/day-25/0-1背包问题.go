// @Author: Ciusyan 11/15/23

package day_25

/**
给定两个长度都为 N 的数组 weights和values，weights[i]和values[i] 分别代表 i 号物品的重量和价值
给定一个正数 bag，表示一个载重 bag 的袋子，装的物品不能超过这个重量
返回能装下的最大价值
*/

// 暴力递归方法：weights values 是物品的重量和价值，bag 是背包的容量，返回能够装取的最大价值
func maxValue1(weights, values []int, bag int) int {
	if weights == nil || values == nil || len(weights) != len(values) || len(weights) == 0 || bag < 0 {
		return 0
	}

	// 递归的含义：
	// 	当前是第 cur 件物品，还可以选择的物品是[cur ...]
	// 	重量是 weights[cur] 价值是 values[cur]，背包还剩余 remain 的重量
	// 	在这样的情形下，从 [cur ...] 这些可选物品中，返回能够装取的最大价值
	var process func(weights, values []int, cur, remain int) int
	process = func(weights, values []int, cur, remain int) int {
		if remain < 0 {
			// 当背包没有剩余了，就不可装了
			// 并且如果背包已经超重了，说明选的上一件有问题，说明那一件不可选，所以得返回一个无效的价值
			return -1
		}

		if cur == len(weights) {
			// 说明现在已经没有物品可以选择了
			return 0
		}

		// 那么能拥有的可能性，
		// 1.选择 cur 这件物品，最大价值 = [cur+1 ...] 挑 remain 得到的最大价值
		p1 := process(weights, values, cur+1, remain)

		// 2.不选择 cur 这件物品，最大价值 = [cur+1 ...] 挑 remain-weights[cur] 得到的最大价值 + values[cur]
		p2 := 0
		next := process(weights, values, cur+1, remain-weights[cur])
		if next != -1 {
			// 说明当前这件可以选，选择后并不超重
			p2 += values[cur]
		}

		// 当然返回的是两种可能的最大值呐
		return max(p1, p2)
	}

	// 那么主函数就应该，可选物品是 [0 ...]，剩余重量是 bag
	return process(weights, values, 0, bag)
}
