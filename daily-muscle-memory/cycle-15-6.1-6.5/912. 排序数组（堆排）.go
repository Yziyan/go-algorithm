// @Author: Ciusyan 5/27/24

package cycle_14_5_27_5_31

// https://leetcode.cn/problems/sort-an-array/

func sortArray(nums []int) []int {

	// 建大根堆
	siftDown := func(nums []int, idx, size int) {
		downEle := nums[idx]
		leafIdx := size >> 1

		// 没有子节点后，就不用下滤了
		for idx < leafIdx {
			// 取出较大的子节点，默认左儿子最大
			childIdx := idx<<1 + 1
			child := nums[childIdx]
			rightIdx := childIdx + 1
			if rightIdx < size && nums[rightIdx] > child {
				// 说明有右孩子，并且右孩子还比左孩子大
				child = nums[rightIdx]
				childIdx = rightIdx
			}

			// 然后看看儿子是否和下滤元素的大小关系
			if downEle >= child {
				// 说明下滤可以终止了
				break
			}

			// 来到这里，说明这个儿子需要上去
			nums[idx] = child
			// 然后继续去子节点进行下滤
			idx = childIdx
		}

		// 可以将下滤元素放置在合理位置了
		nums[idx] = downEle
	}

	heapSize := len(nums)
	// 原地建堆，采用自下而上的下滤，从最后一个非叶子节点开始
	for i := heapSize - 1; i >= 0; i-- {
		siftDown(nums, i, heapSize)
	}

	// 现在 nums 就是一个大根堆了。

	for heapSize > 0 {
		heapSize--
		// 将堆顶的最大值放置最后，最后一个位置就确认了
		nums[heapSize], nums[0] = nums[0], nums[heapSize]
		// 将堆顶进行下滤，但是只下滤到 heapSize
		siftDown(nums, 0, heapSize)
	}

	return nums
}
