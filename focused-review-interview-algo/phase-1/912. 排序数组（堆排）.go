// @Author: Ciusyan 3/21/24

package phase_1

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	heapSize := len(nums)

	// 先原地建堆
	// 采用自下而上的下滤，从第一个非叶子节点开始
	for i := heapSize>>1 - 1; i >= 0; i-- {
		// 下滤操作
		siftDown(nums, i, heapSize)
	}

	// 然后挨个取出最大的堆顶，到末尾交换，然后进行下滤操作
	for heapSize > 0 {
		heapSize--
		// 交换
		nums[0], nums[heapSize] = nums[heapSize], nums[0]
		// 让后不算被排好序的元素，直接去进行下滤操作
		siftDown(nums, 0, heapSize)
	}

	return nums
}

func siftDown(nums []int, idx int, size int) {
	// 下滤到第一个叶子节点，就别下滤了
	downEle := nums[idx]
	leafSize := size >> 1

	for idx < leafSize {
		// 取出最大的孩子，跟下滤节点比较
		childIdx := idx<<1 + 1 // 先默认左子节点最大
		child := nums[childIdx]

		rightIdx := childIdx + 1
		if rightIdx < size && child < nums[rightIdx] {
			// 说明有右子节点，并且比左子节点还大
			child = nums[rightIdx]
			childIdx = rightIdx
		}

		// 来到这里，看看需不需进行下滤
		if child <= downEle {
			// 说明已经不需要下滤了
			break
		}

		// 到这里说明需要下滤，将子节点放上来，让自己接着下去遍历
		nums[idx] = child
		idx = childIdx
	}
	// 将下滤元素放在合适位置
	nums[idx] = downEle
}
