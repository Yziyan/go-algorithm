// @Author: Ciusyan 12/10/23

package day_26

/**
给定一个正数数组 arr，请把 arr 中所有的数分成两个集合
如果 arr 长度为偶数，两个集合包含数的个数要一样多
如果 arr 长度为奇数，两个集合包含数的个数必须只差一个
请尽量让两个集合的累加和接近
返回最接近的情况下，较小集合的累加和
*/

func splitSumClosedSizeHalf(arr []int) int {
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
