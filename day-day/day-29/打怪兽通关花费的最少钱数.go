// @Author: Ciusyan 12/30/23

package day_29

import "math"

/**
int[] d，d[i]：i号怪兽的能力
int[] p，p[i]：i号怪兽要求的钱
开始时你的能力是0，你的目标是从0号怪兽开始，通过所有的怪兽。
如果你当前的能力，小于i号怪兽的能力，你必须付出p[i]的钱，贿赂这个怪兽，然后怪兽就会加入你
他的能力直接累加到你的能力上；如果你当前的能力，大于等于i号怪兽的能力
你可以选择直接通过，你的能力并不会下降，你也可以选择贿赂这个怪兽，然后怪兽就会加入你
他的能力直接累加到你的能力上
返回通过所有的怪兽，需要花的最小钱数
*/

// abilities[i] 代表 i 号怪兽拥有的能力，coins[i] 代表贿赂 i 号怪兽需要花的钱
// 方法1：使用能力作为突破口
func minMoney(abilities, coins []int) int {

	// 憋一个暴力递归，含义是：
	// 目前拥有 ability 的能力，打 [cur ... n-1] 的怪兽，通关所需要花费的最少钱数
	var process func(abilities, coins []int, cur int, ability int) int
	process = func(abilities, coins []int, cur int, ability int) int {
		if cur == len(abilities) {
			// 代表已经打通过了，不需要花钱了
			return 0
		}

		// 当前 cur 号怪兽，有两种情况：
		// 1. 贿赂当前怪兽 2. 不贿赂当前怪兽
		if ability < abilities[cur] {
			// 代表当前能力不够，只能贿赂
			// 贿赂花费 coins[cur]，并且贿赂完后，用增长后的能力，去打剩下的 [cur+1 ... n-1] 个怪兽
			return coins[cur] + process(abilities, coins, cur+1, ability+abilities[cur])
		}

		// 来到这里，说明可以贿赂当前怪兽，也可以不贿赂，两种方式选择花费最少的一种
		// 2.1 直接通过，无需贿赂
		p1 := process(abilities, coins, cur+1, ability)
		// 2.2 能通过也贿赂
		p2 := coins[cur] + process(abilities, coins, cur+1, ability+abilities[cur])

		return min(p1, p2)
	}

	// 代表：目前拥有 0 的能力，打 [0 ... n-1] 的怪兽，通过所需要花费的最少钱数
	return process(abilities, coins, 0, 0)
}

func minMoneyDp(abilities, coins []int) int {
	// 有两个可变参数，他们的范围分别是：
	// cur: [0, n]
	// ability：[0, sum(abilities)]
	n := len(abilities)
	sumAbility := 0
	for _, v := range abilities {
		sumAbility += v
	}

	// 建立 dp 缓存，dp[cur][ability] 含义是：
	// 当前拥有 ability 能力，打 [cur, n-1] 的怪兽所需的最小花费
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, sumAbility+1)
	}

	// 根据依赖关系可知，cur 依赖 cur+1，所以从后往前求
	for cur := n - 1; cur >= 0; cur-- {
		for ability := 0; ability <= sumAbility; ability++ {

			// 两种选择：
			// 1.贿赂当前怪兽
			if ability < abilities[cur] {
				// 说明必须贿赂，因为打不过
				dp[cur][ability] = coins[cur] + dp[cur+1][ability+abilities[cur]]
			} else {
				// 说明可以贿赂，也可以不贿赂
				// 不贿赂，直接通过
				p1 := dp[cur+1][ability]
				// 贿赂，增长能力，但是得控制能力不能操过最大能力，要不然干嘛买它
				p2 := math.MaxInt
				if ability+abilities[cur] <= sumAbility {
					p2 = coins[cur] + dp[cur+1][ability+abilities[cur]]
				}

				// 两种方式种，花费最少的一种方式
				dp[cur][ability] = min(p1, p2)
			}
		}
	}

	// 根据递归的主函数调用可得
	return dp[0][0]
}
