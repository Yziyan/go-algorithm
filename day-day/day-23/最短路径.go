// @Author: Ciusyan 11/8/23

package day_23

import "math"

// DijkstraShortPath 迪杰克斯拉算法，求解单源最短路径
// 从起点开始，将能够直接到达的顶点，加入一个记录中，记录最短路径，
// 记录完成后将起点锁住，再从刚刚的记录中，选择一个最短路径到达的顶点，往返做上面的事情，
// 但是在途中，需要尝试更新已存在顶点的最短路径，还要将新能够到达的顶点，加入其中。
// @startVal: 求解哪个点的最短路径 @return: 到所有点的最短路径
// ep: a <a, 0> <c, 10> <b, 5> 代表 a -> a 的最短路径为 0，a -> c 的最短路径为 10，a -> b 的最短路径为 5

func (g *Graph) DijkstraShortPath(startVal int) map[int]int {
	// 查看起点是否存在
	start, ok := g.vertexes[startVal]
	if !ok {
		return nil
	}

	vtL := len(g.vertexes)
	// 准备一个加强堆(小根堆)，用于获取 spEndVt，并且方便更新存在于堆上的内容
	heap := NewEnhanceMinHeap(vtL)
	// 先将起点加入堆中
	heap.AddOrUpdateOrIgnore(start, 0)

	result := make(map[int]int, vtL)
	// 只要堆里还有元素，就说明还没求解完成
	for heap.Size() != 0 {
		// 将堆顶弹出
		spEndVt, spDistance := heap.Pop()

		// 然后遍历所有的出边
		for _, ed := range spEndVt.edges {
			// 将对应的终点加入堆中
			heap.AddOrUpdateOrIgnore(ed.to, spDistance+ed.weight)
		}

		// 保存结果
		result[spEndVt.val] = spDistance
	}

	return result
}

func (g *Graph) DijkstraShortPath1(startVal int) map[int]int {
	// 查看起点是否存在
	if _, ok := g.vertexes[startVal]; !ok {
		return nil
	}

	vtL := len(g.vertexes)
	// 建立等待求解的结果集
	waitResolve := make(map[int]int, vtL)
	// 将起点加入
	waitResolve[startVal] = 0
	// 准备一个已经锁定的点集
	lockedSet := make(map[int]struct{}, vtL)
	spEndVt := g.getSPEndVt(waitResolve, lockedSet)

	// 只要 spEndVt 还有值，就别停
	for spEndVt != nil {
		// 取出当前的最短路径是多少
		curPath := waitResolve[spEndVt.val]

		// 遍历 end 的所有出边
		for _, eg := range spEndVt.edges {
			newDistance := curPath + eg.weight
			oldDistance, ok := waitResolve[eg.to.val]
			if !ok {
				// 说明以前没有路径，第一次找到
				waitResolve[eg.to.val] = newDistance
			} else {
				// 说明以前有路径，但是取老路和新路权值一个最小的
				waitResolve[eg.to.val] = min(oldDistance, newDistance)
			}
		}

		// 然后将当前 spEndVt 标记为已锁定
		lockedSet[spEndVt.val] = struct{}{}
		// 再获取一个新的
		spEndVt = g.getSPEndVt(waitResolve, lockedSet)
	}

	return waitResolve
}

// 从 waitResolve 中获取最短路径的终点，但是不能选择已经存在于 lockedSet 中的顶点
func (g *Graph) getSPEndVt(waitResolve map[int]int, lockedSet map[int]struct{}) *vertex {

	minDistance := math.MaxInt
	var spEndVt *vertex
	// 遍历 waitResolve
	for vtVal, distance := range waitResolve {
		if _, ok := lockedSet[vtVal]; ok {
			// 说明已经求解完成被锁定了
			continue
		}
		// 说明可选，挑一个目前距离最小的
		if distance < minDistance {
			// 说明比之前的还小
			spEndVt = g.vertexes[vtVal]
			minDistance = distance
		}
	}

	return spEndVt
}
