// @Author: Ciusyan 12/10/23

package day_26

/**
给定一个正数数组 arr，请把 arr 中所有的数分成两个集合
如果 arr 长度为偶数，两个集合包含数的个数要一样多
如果 arr 长度为奇数，两个集合包含数的个数必须只差一个
请尽量让两个集合的累加和接近
返回最接近的情况下，较小集合的累加和
*/

// 动态规划方法
func splitSumClosedSizeHalf(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return -1
	}

	n := len(arr)
	sum := 0
	for _, v := range arr {
		sum += v
	}
	// 除 2
	half := sum >> 1
	halfSize := n >> 1

	// 根据可变参数及其范围
	// cur ∈ [0, n]，size ∈ [0, (n+1)/2]，remain ∈ [0, sum/2]
	// 建立缓存 dp，dp[cur][size][remain] 的含义是：
	// 对 arr[cur ...] 进行分割，得到最接近 remain 的较小累加和，但是对应的集合元素必须是 size 个
	dp := make([][][]int, n+1)
	for cur := range dp {
		// 如果是奇数个，需要保证向上取整
		dp[cur] = make([][]int, (n+1)>>1+1)
		for size := range dp[cur] {
			dp[cur][size] = make([]int, half+1)
		}
	}
	// 先初始化一下，因为递归函数的无效值，使用的是 -1
	for cur := 0; cur <= n; cur++ {
		for size := 0; size <= halfSize; size++ {
			for remain := 0; remain <= half; remain++ {
				dp[cur][size][remain] = -1
			}
		}
	}

	// 根据递归基：
	for remain := 0; remain <= half; remain++ {
		// 说明没有元素可分割了，并且 size 也达标了，再分结果也只能是 0 了
		dp[n][0][remain] = 0
	}

	// 根据依赖情况，size 依赖 size-1，cur 依赖 cur+1，
	// 又因为已经填完了 o-cur-remain 平面，所以从 size 轴开始推
	for size := 1; size <= halfSize; size++ {
		for cur := n - 1; cur >= 0; cur-- {
			for remain := 0; remain <= half; remain++ {
				// 根据一般情况：有两种可能
				// 1.不选 cur
				p1 := dp[cur+1][size][remain]
				// 2.选 cur
				p2 := -1
				next := -1
				if arr[cur] <= remain {
					// 说明 remain-arr[cur] 不会小于  0
					next = dp[cur+1][size-1][remain-arr[cur]]
				}
				if next != -1 {
					// 说明得到一种有效解，可以加上当前值
					p2 = arr[cur] + next
				}

				// 那么答案就是两种情况的最大值，因为求出来的都是 size 个较小的累加和，大的就更接近 remain
				dp[cur][size][remain] = max(p1, p2)
			}
		}
	}

	// 根据递归主函数调用
	if (n & 1) == 0 {
		// 说明是偶数个
		return dp[0][halfSize][half]
	} else {
		// 说明是奇数个
		return max(dp[0][halfSize][half], dp[0][halfSize+1][half])
	}
}

// 暴力递归方法
func splitSumClosedSizeHalf1(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}

	// 先来憋一个暴力递归，含义是：
	// 对 arr[cur ...] 进行分割，返回最接近 remain 的较小累加和，但是这个集合的元素数量必须是 size 个
	var process func(arr []int, cur, size, remain int) int
	process = func(arr []int, cur, size, remain int) int {
		if cur == len(arr) {
			// 说明没有元素了
			if size == 0 {
				// 说明刚好挑了 size 个，那么最接近的累加和是 0
				return 0
			}
			// 否则说明是无效解
			return -1
		}

		// 对于一般情况，还是有两种可能：
		// 1.不选 cur
		p1 := process(arr, cur+1, size, remain)
		// 2.选 cur
		p2 := -1
		next := -1
		if arr[cur] <= remain {
			// 必须小于 remain，要不然会是负数了，
			// 选择了 arr[cur]，那么数量就少一个了，并且 [cur+1 ...] 就只需要分割出 remain-arr[cur] 了
			next = process(arr, cur+1, size-1, remain-arr[cur])
		}

		if next != -1 {
			// 因为 process 下层可能返回 -1，如果不是 -1 了，说明有解法，还需要加上当前值
			p2 = arr[cur] + next
		}

		return max(p1, p2)
	}

	// 那么对于主函数，该如何调用呢？
	// 还是先求一个累加和
	sum := 0
	for _, v := range arr {
		sum += v
	}
	half := sum >> 1 // 除 2
	n := len(arr)

	if (n & 1) == 0 {
		// 说明是偶数个，那么只需要限制累加和的集合是 n/2 个元素即可
		return process(arr, 0, n>>1, half)
	} else {
		// 说明是奇数个，那么有两种可能：1.选 n/2 个累加，2.选 n/2+1 个累加
		// 比如，n == 7，要么选 3 个累加，要么选 4 个累加，得到两个结果，挑个最大的，即是最接近 half 的
		return max(process(arr, 0, n>>1, half), process(arr, 0, (n>>1)+1, half))
	}
}
