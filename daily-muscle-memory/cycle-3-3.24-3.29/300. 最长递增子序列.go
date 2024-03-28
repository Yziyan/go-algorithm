// @Author: Ciusyan 3/28/24

package cycle_3_3_24_3_29

// https://leetcode.cn/problems/longest-increasing-subsequence/description/

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)

	// 准备缓存，dp[cur] 代表，从 [0 ... cur) 的 LIS 长度
	dp := make([]int, n)
	dp[0] = 1
	maxRes := 1

	for cur := 1; cur < n; cur++ {
		// 每一个位置，默认肯定有自己的长度
		dp[cur] = 1
		// 但是还需要看，到 cur 前面，有多少满足上升子序列的
		for pre := 0; pre < cur; pre++ {
			if nums[pre] >= nums[cur] {
				// 说明不满足上升子序列了
				continue
			}
			// 否则满足，尝试更新最大值
			dp[cur] = max(dp[cur], dp[pre]+1)
		}
		// 更新最大值
		maxRes = max(maxRes, dp[cur])
	}

	return maxRes
}
