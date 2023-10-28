// @Author: Ciusyan 10/27/23

package day_21

// 小跟堆
type minHeap struct {
	elements []int
	size     int
}

const defaultCapacity = 10

func NewHeap() *minHeap {
	return &minHeap{
		elements: make([]int, defaultCapacity),
		size:     0,
	}
}

// Add 元素入堆
func (h *minHeap) Add(ele int) {
	// 检查容量是否足够
	h.ensureCapacity(h.size + 1)

	// 然后在末尾添加一个新的元素
	h.elements[h.size] = ele
	// 然后进行上滤操作
	h.siftUp(h.size)
	h.size++
}

// Remove 删除堆顶元素
func (h *minHeap) Remove() int {
	if h.size == 0 {
		panic("堆是空的")
	}
	h.size--
	ele := h.elements[0]
	// 将堆尾放在堆顶
	h.elements[0] = h.elements[h.size]
	h.elements[h.size] = 0

	// 然后对堆顶进行下滤操作
	h.siftDown(0)

	return ele
}

func (h *minHeap) Get() int {
	if h.size == 0 {
		panic("堆为空")
	}

	return h.elements[0]
}

func (h *minHeap) Size() int {
	return h.size
}

// 确保容量足够
func (h *minHeap) ensureCapacity(minCapacity int) {
	oldCapacity := len(h.elements)
	if oldCapacity >= minCapacity {
		return
	}

	// 说明容量不够了
	newCapacity := oldCapacity + (oldCapacity >> 1)
	newElements := make([]int, newCapacity)

	// 挨个拷贝旧数组的内容
	for i := 0; i < h.size; i++ {
		newElements[i] = h.elements[i]
	}

	// 改变引用
	h.elements = newElements
}

func (h *minHeap) siftUp(idx int) {

	// 上滤就是不断用上滤元素不断与父节点比较，放置在合适的位置
	upEle := h.elements[idx]

	// 只要还有父节点，就进行操作
	for idx > 0 {

		// 取出父节点的 idx
		parentIdx := (idx - 1) >> 1
		parentEle := h.elements[parentIdx]

		// 小的要放上面
		if upEle >= parentEle {
			break
		}

		// 说明这里需要将 parent 移到下面来
		h.elements[idx] = parentEle
		// 并且继续上滤
		idx = parentIdx
	}

	// 现在的 idx 就是要放置的位置
	h.elements[idx] = upEle
}

func (h *minHeap) siftDown(idx int) {
	downEle := h.elements[idx]
	// 从有孩子的地方开始遍历
	half := h.size >> 1 // 叶子节点的数量，那么 idx < half，就一定有孩子
	for idx < half {
		// 取较小的子节点与自己比较，默认是左孩子
		childIdx := (idx << 1) + 1
		child := h.elements[childIdx]

		// 看看右边是否比左边还小
		rightIdx := childIdx + 1

		// 但是得看还有没有右子树
		if rightIdx < h.size && h.elements[rightIdx] < child {
			// 说明右边小
			childIdx = rightIdx
			child = h.elements[rightIdx]
		}

		// 看看下滤元素是否能往下走
		if child >= downEle {
			// 说明最小的子节点都比下滤元素大了，就没必要往下走了
			break
		}

		// 说明当前位置应该要换成子节点
		h.elements[idx] = child
		// 记得更好索引
		idx = childIdx
	}

	// 说明下滤节点找到合适的位置了
	h.elements[idx] = downEle
}
