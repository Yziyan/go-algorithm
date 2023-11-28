// @Author: Ciusyan 11/27/23

package day_25

/**
给定一个数组 arr，arr[i] 代表第 i 号咖啡机泡一杯咖啡的时间
给定一个正数 N，表示 N 个人等着咖啡机泡咖啡，每台咖啡机只能轮流泡咖啡
只有一台咖啡机，一次只能洗一个杯子，时间耗费 a，洗完才能洗下一杯
每个咖啡杯也可以自己挥发干净，时间耗费 b，咖啡杯可以并行挥发
假设所有人拿到咖啡之后立刻喝干净，
返回从开始等到所有咖啡机变干净的最短时间
三个参数：int[] arr、int N，int a、int b
*/

// WashTime 动态规划方法
func WashTime(cookTimes []int, n, washTime, selfTime int) int {
	if cookTimes == nil || len(cookTimes) == 0 || n < 1 {
		return 0
	}

	// 也是先获取所有人喝完咖啡，杯子可洗的时间
	drinkTimes := getDrinkTime(cookTimes, n)

	// 有两个可变参数，cur 和 free，分别代表当前正在洗的杯子、洗杯机可用的时间。
	// 它们的范围是：cur ∈ [0, n]，free ∈ [0, maxFree]，那么 maxFree 代表什么呢？代表最大的洗杯时间。
	maxFree := 0
	for i := 0; i < n; i++ {
		// 看谁最晚能用 + 洗杯时间
		maxFree = max(maxFree, drinkTimes[i]) + washTime
	}

	// 准备缓存
	// dp[cur][free] 的含义是：从第 cur 个杯子开始洗，洗杯机最早可用的时间是 free
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, maxFree+1)
	}

	// 根据递归基可知， dp[n][...] = 0
	// 根据依赖关系可知，cur 依赖 cur+1，那么从下往上求解
	for cur := n - 1; cur >= 0; cur-- {
		// 反正是依赖 cur+1 层的，都填好了
		for free := 0; free <= maxFree; free++ {
			// 1.使用洗杯机洗当前杯子
			curTime := max(drinkTimes[cur], free) + washTime
			if curTime > maxFree {
				// 代表后面的不用填了，因为比全部使用洗杯机洗都慢了，那肯定不会是答案
				break
			}
			// 可求剩余的杯子的最早时间
			remainTime := dp[cur+1][curTime]
			// 洗当前的杯子和剩下的杯子，最大值
			p1 := max(curTime, remainTime)

			// 2.使用自净的方式
			curTime = drinkTimes[cur] + selfTime
			remainTime = dp[cur+1][free]
			// 自净当前杯子和剩下的杯子最大值
			p2 := max(curTime, remainTime)

			// 使用两种清洗策略的最优值
			dp[cur][free] = min(p1, p2)
		}
	}

	// 返回值代表，从第 0 个杯子开始洗，洗杯机最早可用的时间是 0 时刻。，
	return dp[0][0]
}

// WashTime1 暴力递归方法（业务尝试模型） n 个人喝咖啡，所有人都喝完，并且洗完杯子的最短时间
func WashTime1(cookTimes []int, n, washTime, selfTime int) int {
	if cookTimes == nil || len(cookTimes) == 0 || n < 1 {
		return 0
	}

	drinkTimes := getDrinkTime(cookTimes, n)
	// 递归含义：drinkTimes 代表杯子需要洗的情况 drinkTimes[i] -> 第 i 个杯子能洗的时间
	// washTime, selfTime 分别代表洗咖啡机洗杯子的时间和杯子自己挥发干净的时间
	// cur 代表洗到第 cur 个杯子了，free 代表咖啡机空闲的时间点
	var process func(drinkTimes []int, washTime, selfTime int, cur, free int) int
	process = func(drinkTimes []int, washTime, selfTime int, cur, free int) int {
		if len(drinkTimes) == cur {
			// 说明杯子洗完了，不需要耗时
			return 0
		}

		// 要么使用咖啡机洗，要么自净
		// 1.使用咖啡机洗
		// 洗 cur 这个杯子，什么时候才能洗呢？要求杯子能洗时，洗杯机也要能使用
		curTime := max(drinkTimes[cur], free) + washTime
		// 接下来的被子是 cur+1，要 curTime 后才能使用洗杯机
		nextBestTime := process(drinkTimes, washTime, selfTime, cur+1, curTime)

		// 当前杯子洗干净的时间，和之后最快洗净的时间最大的那个
		p1 := max(curTime, nextBestTime)

		// 2.自净
		curTime = drinkTimes[cur] + selfTime
		// 接下来的被子是 cur+1，要 free 后才能使用洗杯机，因为没有使用洗杯机
		nextBestTime = process(drinkTimes, washTime, selfTime, cur+1, free)
		// 当前杯子洗干净的时间，和之后最快洗净的时间最大的那个
		p2 := max(curTime, nextBestTime)

		// 两种策略种，耗时最少的一个
		return min(p1, p2)
	}

	// 从第 0 个杯子开始洗，咖啡机从 0 时可用
	return process(drinkTimes, washTime, selfTime, 0, 0)
}

// 获取 n 个人，最快喝完咖啡的时间
func getDrinkTime(cookTimes []int, n int) []int {
	res := make([]int, n)
	// 准备一个堆，用于映射排队信息
	minHeap := NewHeap()
	for _, cm := range cookTimes {
		// 先把所有咖啡机的情况入堆
		minHeap.Add(&cookMachine{timeConsumer: cm})
	}

	// 搞出排队结果出来
	for i := 0; i < n; i++ {
		// 当前这个人，需要使用的机器弹出来
		machine := minHeap.Remove()
		// 当前这个人的耗时，就是当前可用的时间 + 做咖啡需要的时间
		res[i] = machine.nextCookTime + machine.timeConsumer

		// 然后捏，将这台咖啡机的下一次可用时间，更新一下，放入堆里面
		// 下一次可用的时间就是，当前时间 + 做完这一杯的耗时
		machine.nextCookTime += machine.timeConsumer
		// 然后再到堆里去排队
		minHeap.Add(machine)
	}

	return res
}

// 煮咖啡机
type cookMachine struct {
	nextCookTime int // 下一次什么时候做咖啡
	timeConsumer int // 煮咖啡耗时
}

func (c *cookMachine) compare(x *cookMachine) int {
	// 需要共同决定效率
	return (c.timeConsumer + c.nextCookTime) - (x.timeConsumer + x.nextCookTime)
}

type Heap struct {
	elements []*cookMachine
	size     int
}

func NewHeap() *Heap {
	return &Heap{
		elements: make([]*cookMachine, 10),
	}
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Add(ele *cookMachine) {
	h.ensureCapacity(h.size + 1)

	// 将其放置在最后
	h.elements[h.size] = ele
	// 再对其最后一个元素进行上滤操作
	h.siftUp(h.size)
	// 然后数量+1
	h.size++
}

func (h *Heap) Remove() *cookMachine {
	ele := h.elements[0]

	h.size--
	// 将堆尾放置堆顶
	h.elements[0] = h.elements[h.size]
	// 将堆尾清空
	h.elements[h.size] = nil
	// 然后对堆顶进行下滤操作
	h.siftDown(0)

	return ele
}

// 确保容量足够
func (h *Heap) ensureCapacity(ensureCapacity int) {
	oldCapacity := len(h.elements)
	if ensureCapacity < oldCapacity {
		// 说明当前容量够，没必要扩容
		return
	}

	// 否则说明需要扩容
	capacity := oldCapacity + oldCapacity>>1
	newElements := make([]*cookMachine, capacity)

	for i := 0; i < h.size; i++ {
		newElements[i] = h.elements[i]
	}

	h.elements = newElements
}

// 上滤操作
func (h *Heap) siftUp(idx int) {
	// 先取出上滤节点
	upEle := h.elements[idx]
	// 要有父节点才上滤
	for idx > 0 {
		// 先算出父节点索引
		parentIdx := (idx - 1) >> 1
		parent := h.elements[parentIdx]

		if parent.compare(upEle) <= 0 {
			// 说明父节点本身就较小
			break
		}

		// 说明父节点要大，往下走
		h.elements[idx] = parent
		// 父节点也进行上滤
		idx = parentIdx
	}
	// 最后将上滤元素放置合适位置
	h.elements[idx] = upEle
}

// 下滤操作
func (h *Heap) siftDown(idx int) {
	downEle := h.elements[idx]

	halfSize := h.size >> 1
	// 要有儿子才下滤
	for idx < halfSize {
		// 至少有左孩子
		childIdx := (idx << 1) + 1
		child := h.elements[childIdx]

		// 看看有没有右孩子，并且比左孩子还小
		rightIdx := childIdx + 1
		if rightIdx < h.size && child.compare(h.elements[rightIdx]) < 0 {
			childIdx = rightIdx
			child = h.elements[rightIdx]
		}

		// 来到这里，肯定是最小的子节点了
		if child.compare(downEle) >= 0 {
			// 说明最小的子节点都比自己还大了，就没必要往下走了
			break
		}
		// 否则需要往下走
		h.elements[idx] = child
		idx = childIdx
	}
	// 将下滤元素放在合适位置
	h.elements[idx] = downEle
}
