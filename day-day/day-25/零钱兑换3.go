// @Author: Ciusyan 12/3/23

package day_25

/**
arr 是货币数组，其中的值都是正数。再给定一个正数 aim。
每个值都认为是一张货币，
认为值相同的货币没有任何不同，
返回组成 aim 的方法数
例如：arr = {1,2,1,1,2,1,2}，aim = 4
方法：1+1+1+1、1+1+2、2+2
一共就 3 种方法，所以返回 3
*/

func CoinsWaySameValueSamePaper(coins []int, aim int) int {
	if coins == nil || len(coins) == 0 || aim <= 0 {
		return 0
	}

	// 先憋一个暴力递归，含义是：
	// coins[cur] 枚硬币有 nums[cur] 个，使用 coins[cur ...] 凑出 remain 面值，拥有的方法数
	var process func(coins, nums []int, cur, remain int) int
	process = func(coins, nums []int, cur, remain int) int {
		if cur == len(coins) {
			// 说明没有硬币可以用了
			if remain != 0 {
				// 代表面值凑完了，都还没凑出 remain，那也没办法了
				return 0
			}
			return 1
		}

		// 对于一般情况，对于每一枚硬币，看看能使用多少枚来凑，挨个尝试
		ways := 0 // 总的方法数就是都尝试相加
		// 但是呢？尝试也不能多找零，还有 coins[cur] 只有 nums[cur] 枚，肯定找完就不能使用了
		for num := 0; num*coins[cur] <= remain && num <= nums[cur]; num++ {
			// 下一种尝试方案就是：使用 coins[cur+1 ...]，去凑够使用 cur 这枚硬币找零的钱
			ways += process(coins, nums, cur+1, remain-num*coins[cur])
		}

		return ways
	}

	// 先构建出硬币的信息
	info := newCoinsInfo(coins)
	// 那么主函数就应该这样调用，代表：使用 coins[0 ...] 凑出 aim 面值，拥有的方法数
	return process(info.coins, info.nums, 0, aim)
}

type coinsInfo struct {
	coins []int // 硬币数组，已去重
	nums  []int // 每个硬币的个数
	// 与 coins[i] 一一对应，比如 coins[2] = 3，nums[2] = 5，代表 3 这枚硬币有 5 枚
}

// 构建出 CoinsInfo
func newCoinsInfo(oldCoins []int) *coinsInfo {
	maxL := len(oldCoins)
	// <coin, num>，统计每一枚硬币的数量
	coinsMap := make(map[int]int, maxL)
	for _, coin := range oldCoins {
		coinsMap[coin]++
	}

	newCoins := make([]int, 0, maxL)
	nums := make([]int, 0, maxL)

	for coin, num := range coinsMap {
		newCoins = append(newCoins, coin)
		nums = append(nums, num)
	}

	return &coinsInfo{
		coins: newCoins,
		nums:  nums,
	}
}
