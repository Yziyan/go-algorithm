// @Author: Ciusyan 9/22/23

package day_15

// 手写堆结构

type IHeap interface {
	Size() int
	IsEmpty() bool
	Clear()
	Add(ele int)
	Get() int
	Remove() int
	Replace(ele int) int
}

const (
	DEFAULT_SIZE = 10
)

type Heap struct {
	// 这里直接当做数组来使用。就不使用动态数组来用了
	elements []int
	size     int
}

func NewHeap() *Heap {
	return &Heap{
		elements: make([]int, DEFAULT_SIZE),
	}
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

func (h *Heap) Clear() {
	// 可以不清空内存，可以复用
	for i := 0; i < h.size; i++ {
		h.elements[i] = 0
	}
	h.size = 0
}

func (h *Heap) Add(ele int) {
	// 要添加，最少要能保证的容量是当前容量能装入新加入的一个元素
	h.ensureCapacity(h.size + 1)
	// 加入元素到末尾
	h.elements[h.size] = ele
	// 然后对数组的末尾元素进行上滤操作
	h.siftUp(h.size)
	// 添加了一个元素，别忘记维护 Size
	h.size++
}

func (h *Heap) Get() int {
	h.emptyCheck()
	// 返回堆顶元素即可
	return h.elements[0]
}

// Remove 删除堆顶元素
func (h *Heap) Remove() int {
	h.emptyCheck()
	ele := h.elements[0]
	h.size--
	// 将末尾元素与放在堆顶
	h.elements[0] = h.elements[h.size]
	h.elements[h.size] = 0
	// 然后将堆顶元素进行下滤操作
	h.siftDown(0)

	// 返回被删除的堆顶
	return ele
}

func (h *Heap) Replace(ele int) int {
	oldEle := h.elements[0]

	// 先覆盖堆顶
	h.elements[0] = ele
	if h.size == 0 {
		h.size++
	} else {
		// 说明以前有元素，对堆顶进行下滤操作即可，因为新换了一个数进来
		h.siftDown(0)
	}

	return oldEle
}

func (h *Heap) emptyCheck() {
	if h.size == 0 {
		panic("Heap is empty!!!")
	}
}

// 确保容量足够，不够时扩容
func (h *Heap) ensureCapacity(capacity int) {
	oldCapacity := len(h.elements)
	if oldCapacity >= capacity {
		// 说明不需要扩容
		return
	}

	// 需要扩容，新容量为原来的 1.5 倍
	newCapacity := oldCapacity + (oldCapacity >> 1)
	newElements := make([]int, newCapacity)
	// 挨个拷贝元素
	for i := 0; i < h.size; i++ {
		newElements[i] = h.elements[i]
	}

	// 然后将指向改变
	h.elements = newElements
}

// 进行上滤操作，
// @ index：需要上滤的节点
func (h *Heap) siftUp(index int) {
	// 上滤操作就是将当前节点不断往上与父节点比较
	child := h.elements[index]

	// 有父节点才上滤
	for index > 0 {
		// 先计算父节点的索引
		parentIdx := (index - 1) >> 1
		parent := h.elements[parentIdx]

		if parent <= child {
			// 如果子节点不比父节点小，就直接退出了
			break
		}

		// 来到这里，至少需要先将当前节点的值换成父节点
		h.elements[index] = parent
		// 然后将父节点变成上滤节点
		index = parentIdx
	}
	// 来到这里，再将上滤到的终点赋值
	h.elements[index] = child
}

// 进行下滤操作，
// @ index：需要下滤的节点
func (h *Heap) siftDown(index int) {
	// 先计算出叶子节点的数量
	leafSize := h.size >> 1

	parent := h.elements[index]

	// 如果下滤的节点本身就是叶子节点了，就没必要下滤了
	for index < leafSize {
		// 先假设左孩子是最小值
		childIdx := (index << 1) + 1
		child := h.elements[childIdx]

		rightIdx := childIdx + 1
		if rightIdx < h.size && h.elements[rightIdx] < child {
			// 说明有右孩子，并且右边大一些
			child = h.elements[rightIdx]
			childIdx = rightIdx
		}

		// 来到这里，最小的子节点肯定找出来了，需要看看是否能下滤
		if parent <= child {
			// 如果父节点不比子节点还小，就没必要下滤了
			break
		}

		// 看来需要换位置了
		h.elements[index] = child
		// 换完过后，咱们需要
		index = childIdx
	}
	// 最后不能下滤了
	h.elements[index] = parent
}
