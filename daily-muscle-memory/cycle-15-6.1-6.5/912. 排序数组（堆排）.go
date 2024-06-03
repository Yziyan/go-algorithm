// @Author: Ciusyan 5/27/24

package cycle_14_5_27_5_31

// https://leetcode.cn/problems/sort-an-array/

/**
思路重复：
堆排序，核心步骤是：
1.先原地建堆
2.然后依次将堆顶元素放置到合适位置。
如果是大根堆，就将堆顶元素依次和末尾元素交换，然后对堆顶进行下滤操作，以修复堆的性质。
如果是小跟堆，操作相反，但是操作数组前面位置不太合适，况且我们需要原地建堆，所以采用上面的方式。

那么如何原地建堆呢？一般有两种方式：
1.自下而上的下滤
2.自上而下的上滤

第一种比第二种要优，所以我们选择第一种方式，况且也需要使用下滤操作，来修复堆顶的性质。
下滤是什么操作呢？
其实就是从将下滤节点，最大的儿子来和自己比较，如果比自己还大，那么和儿子交换位置，然后继续从儿子的位置进行下滤操作。
直至没有儿子比自己还大，就停止下滤操作。
*/

func heapSortArray(nums []int) []int {

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
