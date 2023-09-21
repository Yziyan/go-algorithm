// @Author: Ciusyan 9/21/23

package day_14

// https://leetcode.cn/problems/sort-an-array/description/

// 使用堆排序
func sortArray(nums []int) []int {
	if nums == nil {
		return nil
	}

	heapSize := len(nums)
	if heapSize < 2 {
		return nums
	}
	// 对 nums 进行快速排序
	// 	原地建堆，这里建大根堆：采用自下而上的下滤，那么从最后一个非叶子节点开始其实就可以了
	for i := (heapSize >> 1) - 1; i >= 0; i-- {
		siftDown(nums, i, heapSize)
	}

	// 建堆之后，开始排序
	// 堆里只剩一个元素时，不用排了；
	for heapSize > 0 {
		// 将堆顶的最大值放置到堆尾, 然后缩小堆的范围，代表最后一个位置已经固定是最大值了
		heapSize--
		nums[0], nums[heapSize] = nums[heapSize], nums[0]
		// 但是此时还需要维护堆的性质
		siftDown(nums, 0, heapSize)
	}

	return nums
}

// 下滤操作，这里需要将大的放置在 parentIdx
// @ parentIdx：从什么位置开始下滤，size：堆现在的大小。
func siftDown(nums []int, parentIdx, size int) {
	// 先计算出叶子节点的数量
	leafSize := size >> 1
	// 先将父节点的值保存，找到合适的位置后，再交换即可
	parent := nums[parentIdx]
	// 当 parentIdx 大于叶子节点的数量，说明没有叶子节点，也就不用比较了
	for parentIdx < leafSize {
		// 找到左右孩子最大值对应的索引，现在假装最大值是左孩子
		childIdx := (parentIdx << 1) + 1
		child := nums[childIdx]

		// 算出右孩子的索引
		rightIdx := childIdx + 1
		if rightIdx < size && nums[rightIdx] > child {
			// 能来到这里，说明有右孩子，并且右孩子比左孩子大
			childIdx = rightIdx
			child = nums[rightIdx]
		}
		// 来到这里，child 肯定是最大的孩子，将其与传入的父节点比较
		if child <= parent {
			// 说明子节点小，不能往下走了，可以终止了
			break
		}
		// 否则将子孩子设置到现在的父节点的位置，然后现在的子节点变成父节点
		nums[parentIdx] = child
		parentIdx = childIdx
	}
	// 最后，还需要将父节点最终停留的位置，设置上正确的值
	nums[parentIdx] = parent
}
