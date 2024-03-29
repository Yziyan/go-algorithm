// @Author: Ciusyan 3/28/24

package cycle_3_3_24_3_29

// https://leetcode.cn/problems/longest-increasing-subsequence/description/

/*
*
思路重复：
对于这样一个序列 nums[0 ... cur]，如果我们能求出以 cur 结尾的 LIS
那么之后也就能利用这个信息，去求出可能更大的 LIS。
比如对于 [0 ... cur+1] 位置，
如果在 [0...cur]、[0... cur-1] ...[0 ... 0] 位置的最长上升子序列是 3、4、1
那么 [0...cur+1]，如果 cur+1 满足上升子序列的规则，那么它就可能是 3+1、4+1、1+1
所以，其实也是使用 dp 的思想，核心点就是限制最右位置。明确以它结尾，求出结果后，再去求另外的解
*/
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
