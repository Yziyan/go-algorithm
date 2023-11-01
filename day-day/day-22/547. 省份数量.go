// @Author: Ciusyan 11/1/23

package day_22

// https://leetcode.cn/problems/number-of-provinces/description/

func findCircleNum(isConnected [][]int) int {
	capacity := len(isConnected)
	uf := newUnionFindCN(capacity)

	// 遍历对角线上方即可，因为是对称的
	for i := 0; i < capacity; i++ {
		for j := i + 1; j < capacity; j++ {
			// 如果是相邻的，就把他们放到一个集合
			if isConnected[i][j] == 1 {
				// 代表接壤了，将 i 和 j 合并
				uf.union(i, j)
			}
		}
	}

	return uf.sets
}

func newUnionFindCN(capacity int) *unionFindCN {
	parents := make([]int, capacity)
	ranks := make([]int, capacity)

	for i := 0; i < capacity; i++ {
		// 自成为一个集合
		parents[i] = i
		ranks[i] = 1
	}

	return &unionFindCN{
		parents: parents,
		ranks:   ranks,
		sets:    capacity,
	}
}

type unionFindCN struct {
	parents []int
	ranks   []int
	sets    int // 有多少个集合
}

func (u *unionFindCN) union(v1, v2 int) {
	root1 := u.findRoot(v1)
	root2 := u.findRoot(v2)

	if root1 == root2 {
		// 说明已经在一个集合了
		return
	}

	// 来到这里，说明不属于一个集合，将矮的合并到高的上方
	if u.ranks[root1] < u.ranks[root2] {
		u.parents[root1] = root2
	} else if u.ranks[root1] > u.ranks[root2] {
		u.parents[root2] = root1
	} else {
		// 一样高，合并谁都可以，但是需要长高
		u.parents[root1] = root2
		u.ranks[root2]++
	}

	// 合并了，就少一个集合了
	u.sets--
}

func (u *unionFindCN) findRoot(v int) int {
	// 一路向上寻找，直至遇到父节点就是自己，说明到达了根节点
	for v != u.parents[v] {
		// 路径减半，将自己挂载到祖父身上
		u.parents[v] = u.parents[u.parents[v]]
		// 再利用祖父向上查找
		v = u.parents[v]
	}

	return v
}
