// @Author: Ciusyan 12/7/23

package day_26

/**
给定一个正数 n，求 n 的裂开方法数，
规定：后面的数不能比前面的数小
比如 4 的裂开方法有：
1+1+1+1、1+1+2、1+3、2+2、4
5 种，所以返回 5
*/

func splitNumber(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// 根据递归参数，有两个可变参数：pre ∈ [1, n]，remain ∈ [1, n]
	// 所以准备缓存，dp[pre][remain] 代表：
	// 前一个裂开的数是 pre，裂开 remain 有多少种方法数
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 根据递归基：
	// 当 pre > remain 时都是 0，不用初始化
	for pre := 1; pre <= n; pre++ {
		// remain == 0 时，说明没有要裂开的数了，前面的选择是一种方法数
		dp[pre][0] = 1
	}

	// 根据依赖情况：可以发现两个参数都在变。first 往下走，remain 往左走，
	// 所以，dp[pre][remain] 依赖左下方某些位置的值，那么我们应该从下往上求解
	for pre := n; pre >= 1; pre-- {
		// remain 至少是从 pre 起
		for remain := pre; remain <= n; remain++ {
			ways := 0
			// 根据一般情况，有多少种方法呢？挨个尝试
			// first 从 pre 开始，但是不能比 remain 还大
			for first := pre; first <= remain; first++ {
				// 代表着，当前选择了 first，那么对于下一次尝试，first 就是前一个数，去裂开 remain-first 有的方法数
				ways += dp[first][remain-first]
			}

			dp[pre][remain] = ways
		}
	}

	// 根据递归主函数调用，
	// 前一个裂开的数是 1，裂开 n 有多少种方法数
	return dp[1][n]
}

// 暴力递归方法
func splitNumber1(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// 憋一个暴力递归，含义是：
	// 裂开的前一个数是 pre，还剩下 remain 要裂开，能得到多少种裂开的方法数
	var process func(pre int, remain int) int
	process = func(pre int, remain int) int {
		if remain == 0 {
			// 说明没有要裂开的数了，也就是说明，之前裂开的选择，就是一种方法数
			return 1
		}

		if pre > remain {
			// 前一个数不能比后面的数大，所以是一种无效解
			return 0
		}

		// 对于一般情况：返回的方法数应该如何呢？
		ways := 0
		// 当然是挨个试了，first 从 pre 开始试，但是再怎么试，first 一定不能比 remain 还大
		for first := pre; first <= remain; first++ {
			// 所有尝试得到的方法数相加
			// 代表当前选择 first 去裂开，所以下一次去分裂时，first 就作为裂开的前一个数，余下 remain-first 要裂
			ways += process(first, remain-first)
		}

		return ways
	}

	// 那么返回值怎么定呢？我们想要裂开 n，remain 好定，那 pre 呢？当然是能用于裂开的最小的数呐。
	// 裂开的前一个数是 1，还剩下 n 要裂开，能得到多少种裂开的方法数
	return process(1, n)
}
