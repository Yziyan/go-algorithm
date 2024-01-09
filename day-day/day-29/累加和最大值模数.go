// @Author: Ciusyan 1/2/24

package day_29

/**
给定一个非负数组 arr，和一个正数 m，返回 arr 的所有子序列中累加和 %m 之后的最大值
*/

func subArrModM(arr []int, m int) int {

	n := len(arr)
	// 1.先求出累加和
	sumArr := 0
	for _, v := range arr {
		sumArr += v
	}

	// 2.定义缓存 dp，dp[cur][sum] 的含义是：
	// 在 [0 ... cur] 这个序列中，能否挑出累加和为 sum 的子序列
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, sumArr+1)
	}

	// 初始化
	for cur := 0; cur < n; cur++ {
		// 不论序列是什么，都能凑出 0 这个累加和
		dp[cur][0] = true
	}
	// 在 [0 ... 0] 范围，第一个累加和，肯定能求出来
	dp[0][arr[0]] = true

	// 对于其余的每一个位置
	for cur := 1; cur < n; cur++ {
		for sum := 1; sum <= sumArr; sum++ {
			// 1.不选择当前 arr[cur] 的值累加
			// 那就完全由 [0 ... cur-1] 个数累加出来
			dp[cur][sum] = dp[cur-1][sum]
			// 2.选择当前 arr[cur] 的值累加
			if sum-arr[cur] >= 0 {
				// 说明前面的 cur-1 个数，只需要凑 sum - arr[cur] 的值
				dp[cur][sum] = dp[cur][sum] || dp[cur-1][sum-arr[cur]]
			}
		}
	}

	res := 0
	for sum := 0; sum <= sumArr; sum++ {
		if dp[n-1][sum] {
			// 说明使用 [0 ... n-1] 可以累加出 sum，用于求解一个最大值
			res = max(res, sum%m)
		}
	}

	return res
}
