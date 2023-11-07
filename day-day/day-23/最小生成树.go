// @Author: Ciusyan 11/7/23

package day_23

// KruskalBST 最小生成树，K 算法，返回选择的每一条边的 Weight
// 利用贪心策略，每一次选择一条最小权边，只要选择的边不会形成环，就无脑选择，如果会形成环，就不要那条边即可
func (g *Graph) KruskalBST() []int {

	// 准备一个最小堆，用于快速获取最小权边
	minHeap := NewMinHeap()
	for i := range g.edges {
		// 将所有边加入最小堆中
		minHeap.Add(g.edges[i])
	}

	l := len(g.vertexes) - 1
	// 收集结果
	res := make([]int, 0, l)

	// 准备一个并查集，用于判断待选边是否会形成环路
	uf := NewUnionFind()
	for k := range g.vertexes {
		uf.MakeSets(g.vertexes[k])
	}

	// 只要堆还有元素，除非提前收集完所有结果了，即 BST 的边数 = 顶点数 - 1
	for minHeap.Size() != 0 && len(res) < l {
		// 弹出一条最小的边
		minEd := minHeap.Remove()
		// 判断是否需要选择这条边
		if uf.IsSame(minEd.from, minEd.to) {
			// 说明这条边的起点和终点本身就属于一个集合了，会形成环
			continue
		}

		// 否则可以选择这两条边，
		res = append(res, minEd.weight)
		// 但是别忘了要将这条边的起点和终点放在一个集合
		uf.Union(minEd.from, minEd.to)
	}

	return res
}
