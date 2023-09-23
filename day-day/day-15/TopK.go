// @Author: Ciusyan 9/23/23

package day_15

// TopK 挑选出 nums 数组中，最大的 K 个数
func TopK(nums []int, k int) []int {
	if nums == nil || len(nums) <= k {
		return nums
	}

	// 准备一个最小堆，
	heap := NewHeap()
	//	先加入 K 个数，保持堆里面有 k 个数
	idx := 0
	for ; idx < k; idx++ {
		heap.Add(nums[idx])
	}

	// 从第 K+1 个数开始，每个数都拿出来和堆顶的最小值比较一下
	for ; idx < len(nums); idx++ {
		if nums[idx] <= heap.Get() {
			continue
		}

		// 如果遇到的元素比堆顶大，就把堆顶替换了
		heap.Replace(nums[idx])
	}

	// 最终剩余在 heap 里面的 k 个数，就是最大的 k 个数
	res := make([]int, 0, k)
	for !heap.IsEmpty() {
		res = append(res, heap.Remove())
	}

	return res
}
