// @Author: Ciusyan 9/23/23

package day_15

// 请对 nums 排序，nums 基本有序，若要将 nums 完全有序，每个元素至多移动 k 个距离
func sortedArrLengthK(nums []int, k int) {
	if nums == nil {
		return
	}

	// 需要一个小跟堆
	heap := NewHeap()
	idx := 0
	if k > len(nums) {
		k = len(nums)
	}
	for idx < k {
		// 先加入 K 个元素进入堆里面
		heap.Add(nums[idx])
		idx++
	}

	// 用来覆盖数组的索引
	i := 0
	for cur := idx; cur < len(nums); cur++ {
		// 先将当前元素入堆
		heap.Add(nums[cur])
		// 然后将堆顶元素覆盖到 i 位置
		nums[i] = heap.Remove()
		i++
	}

	// 最终退出了，将队里面剩余的元素全部弹出来
	for !heap.IsEmpty() {
		nums[i] = heap.Remove()
		i++
	}
}
