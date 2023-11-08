// @Author: Ciusyan 11/7/23

package day_23

// Stack 模拟一个栈
type Stack []*vertex

func newStack() Stack {
	return make([]*vertex, 0)
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Push(vt *vertex) {
	*s = append(*s, vt)
}

func (s *Stack) Pop() *vertex {
	last := len(*s) - 1
	res := (*s)[last]
	*s = (*s)[:last]
	return res
}

// Queue 模拟一个队列
type Queue []*vertex

func newQueue() Queue {
	return make([]*vertex, 0)
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) Push(vt *vertex) {
	*q = append(*q, vt)
}

func (q *Queue) Pop() *vertex {
	vt := (*q)[0]
	// 模拟队头弹出
	*q = (*q)[1:]
	return vt
}

// MinHeap 最小堆，用于获取最小权边
type MinHeap struct {
	elements []*edge
	size     int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		// 默认给 10 个空间
		elements: make([]*edge, 10),
	}
}

func (h *MinHeap) Add(edge *edge) {
	// 检查容量
	h.ensureCapacity(h.size + 1)

	// 添加至最后
	h.elements[h.size] = edge
	// 但是要进行上滤操作
	h.siftUp(h.size)
	// 记得增加 size
	h.size++
}

func (h *MinHeap) Size() int {
	return h.size
}

func (h *MinHeap) Remove() *edge {
	// 优先取出堆顶
	res := h.elements[0]

	h.size--
	// 然后将堆尾换至堆顶
	h.elements[0] = h.elements[h.size]
	h.elements[h.size] = nil
	// 然后将堆顶进行下滤操作
	h.siftDown(0)

	return res
}

func (h *MinHeap) ensureCapacity(minCapacity int) {
	oldCapacity := len(h.elements)
	if oldCapacity >= minCapacity {
		// 说明容量够，不用扩容
		return
	}

	// 说明需要扩容
	newCapacity := oldCapacity + (oldCapacity >> 1)
	newElements := make([]*edge, newCapacity)

	for i := 0; i < h.size; i++ {
		newElements[i] = h.elements[i]
	}

	h.elements = newElements
}

func (h *MinHeap) siftUp(idx int) {
	// 上滤节点
	upEle := h.elements[idx]

	// 有爹才上滤
	for idx > 0 {
		// 取出爹
		parentIdx := (idx - 1) >> 1
		parent := h.elements[parentIdx]

		if parent.weight <= upEle.weight {
			// 爹本来就比自己小了，就没必要往上了
			break
		}

		// 否则需要将爹下移
		h.elements[idx] = parent
		idx = parentIdx
	}

	h.elements[idx] = upEle
}

func (h *MinHeap) siftDown(idx int) {
	// 进行下滤
	downEle := h.elements[idx]

	// 有儿子才下滤
	leafSize := h.size >> 1

	// 在叶子节点前出现，就说明有儿子
	for idx < leafSize {
		// 至少有左孩子
		childIdx := (idx << 1) + 1
		child := h.elements[childIdx]

		rightIdx := childIdx + 1
		if rightIdx < h.size && h.elements[rightIdx].weight <= child.weight {
			// 说明右边还要小
			childIdx = rightIdx
			child = h.elements[rightIdx]
		}

		// 然后和下滤节点比较
		if child.weight >= downEle.weight {
			// 说明儿子本来就比自己还大了，就没必要往下了
			break
		}

		// 否则需要往下
		h.elements[idx] = child
		idx = childIdx
	}

	h.elements[idx] = downEle
}

// UnionFind 并查集，用于判断是否会形成环
type UnionFind struct {
	sets map[int]*node
}

type node struct {
	val    int
	parent *node
	rank   int
}

func newNode(val int) *node {
	nd := &node{
		val:  val,
		rank: 1,
	}
	nd.parent = nd
	return nd
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		sets: make(map[int]*node),
	}
}

func (u *UnionFind) MakeSets(vertexes ...*vertex) {
	for _, vt := range vertexes {
		nd := newNode(vt.val)
		u.sets[vt.val] = nd
	}
}

func (u *UnionFind) IsSame(from, to *vertex) bool {

	fromRoot := u.findRoot(from.val)
	toRoot := u.findRoot(to.val)

	if fromRoot != toRoot {
		return false
	}

	return true
}

func (u *UnionFind) Union(from, to *vertex) {

	fromRoot := u.findRoot(from.val)
	toRoot := u.findRoot(to.val)

	if fromRoot == toRoot {
		// 说明本来就在一个集合
		return
	}

	// 否则需要根据 rank 合并
	if fromRoot.rank < toRoot.rank {
		// 将矮的挂高的
		fromRoot.parent = toRoot
	} else if fromRoot.rank > toRoot.rank {
		toRoot.parent = fromRoot
	} else {
		fromRoot.parent = toRoot
		toRoot.rank++
	}
}

func (u *UnionFind) findRoot(v int) *node {
	nd, ok := u.sets[v]
	if !ok {
		return nil
	}

	// 否则一种往上寻找
	for nd != nd.parent {
		// 将自己挂载到祖父身上，路径减半
		nd.parent = nd.parent.parent
		nd = nd.parent
	}

	return nd
}

// EnhanceMinHeap 加强的小根堆
type EnhanceMinHeap struct {
	// 堆里的元素
	elements []*vertex
	// 堆里元素的索引位置
	eleIdx map[*vertex]int
	// 最短路径
	shortDistance map[*vertex]int

	// 元素
	size int
}

func NewEnhanceMinHeap(capacity int) *EnhanceMinHeap {
	return &EnhanceMinHeap{
		elements:      make([]*vertex, capacity),
		eleIdx:        make(map[*vertex]int, capacity),
		shortDistance: make(map[*vertex]int, capacity),
	}
}

// AddOrUpdateOrIgnore 新增 or 更新 or 什么也不做(已经锁定的顶点)
func (h *EnhanceMinHeap) AddOrUpdateOrIgnore(ele *vertex, distance int) {
	// 得判断堆里是否已经有这个元素了
	if !h.isEntered(ele) {
		// 说明第一次入堆
		h.shortDistance[ele] = distance
		h.eleIdx[ele] = h.size
		h.elements[h.size] = ele
		// 上滤
		h.siftUp(h.size)
		h.size++
	} else if !h.isLocked(ele) {
		// 说明没有锁定这个点，需要更新其值
		h.shortDistance[ele] = min(distance, h.shortDistance[ele])
		// 但是需要对当前元素进行上滤操作
		h.siftUp(h.eleIdx[ele])
	}
}

// Pop 弹出最小的顶点，及其最小的路径距离
func (h *EnhanceMinHeap) Pop() (*vertex, int) {
	// 先取出原先的值
	ele := h.elements[0]
	d := h.shortDistance[ele]

	h.size--
	// 将堆尾换至堆顶
	h.elements[0] = h.elements[h.size]
	h.elements[h.size] = nil
	// 下滤操作
	h.siftDown(0)

	// 标记为锁定
	h.eleIdx[ele] = -1
	// 删除距离
	delete(h.shortDistance, ele)

	return ele, d
}

func (h *EnhanceMinHeap) Size() int {
	return h.size
}

// 是否已经加入过堆了
func (h *EnhanceMinHeap) isEntered(ele *vertex) bool {
	if _, ok := h.eleIdx[ele]; ok {
		return true
	}

	return false
}

// 是否已经锁住了，如果一个顶点使用完后，我们将它在堆上的索引设置为 -1
func (h *EnhanceMinHeap) isLocked(ele *vertex) bool {
	d, ok := h.eleIdx[ele]
	if ok && d == -1 {
		return true
	}
	return false
}

func (h *EnhanceMinHeap) siftUp(idx int) {
	// 取出上滤节点的值
	upEle := h.elements[idx]

	// 有父节点才上滤
	for idx > 0 {
		// 取出父节点
		parentIdx := (idx - 1) >> 1
		parent := h.elements[parentIdx]

		// 如果不比父节点的距离要小，就返回了
		if h.shortDistance[parent] <= h.shortDistance[upEle] {
			break
		}
		// 否则说明需要往上走
		h.elements[idx] = parent
		// 父节点往下走，记得更新索引
		h.eleIdx[parent] = idx
		idx = parentIdx
	}

	h.elements[idx] = upEle
	h.eleIdx[upEle] = idx
}

func (h *EnhanceMinHeap) siftDown(idx int) {
	downEle := h.elements[idx]
	leafSize := h.size >> 1

	// 从第一个非叶子节点遍历
	for idx < leafSize {
		// 取出子节点，默认是左子节点
		childIdx := (idx << 1) + 1
		child := h.elements[childIdx]

		rightIdx := childIdx + 1
		if rightIdx < h.size && h.shortDistance[h.elements[rightIdx]] < h.shortDistance[child] {
			// 说明右边更小
			childIdx = rightIdx
			child = h.elements[rightIdx]
		}

		// 看看能不能往下走
		if h.shortDistance[child] >= h.shortDistance[downEle] {
			// 说明子节点本身就大了，走不了一点
			break
		}

		// 将子节点上走
		h.elements[idx] = child
		h.eleIdx[child] = idx
		idx = childIdx
	}

	// 保存节点
	h.elements[idx] = downEle
	h.eleIdx[downEle] = idx
}
