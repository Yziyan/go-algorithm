// @Author: Ciusyan 11/2/23

package day_22

// 一开始给定一个 m*n 的矩阵。还有一系列的坐标。
// 遍历一系列坐标，将每一坐标挨个变为岛屿，求变化过程中，相邻岛屿的数量。

func numIslands22(m, n int, positions [][]int) []int {
	// 建立一个能容纳 m*n 长度的并查集
	uf := newUnionFindNd2(m, n)

	landsNums := make([]int, len(positions))
	// 挨个坐标便利，并且收集每一次建岛后，相邻岛屿的数量
	for i, position := range positions {
		landsNums[i] = uf.contact(position[0], position[1])
	}

	return landsNums
}

func newUnionFindNd2(m, n int) *unionFindNd2 {
	l := m * n
	return &unionFindNd2{
		parents: make([]int, l),
		ranks:   make([]int, l),
		m:       m, // 行
		n:       n, // 列
	}
}

type unionFindNd2 struct {
	parents []int
	ranks   []int
	sets    int
	m, n    int
}

// 新建一个集合后，返回剩余的集合数量
func (u *unionFindNd2) contact(row, col int) int {
	idx := u.index(row, col)
	if u.ranks[idx] != 0 {
		// 说明这个坐标之前已经建立过集合了
		return u.sets
	}
	// 没建立过，建立，并且看看相邻集合的数量
	u.makeSet(idx)

	// 然后查看上下左右，是否能够合并
	u.union(row-1, col, idx) // 上
	u.union(row+1, col, idx) // 下
	u.union(row, col-1, idx) // 左
	u.union(row, col+1, idx) // 右

	return u.sets
}

func (u *unionFindNd2) index(row, col int) int {
	return row*u.n + col
}

func (u *unionFindNd2) makeSet(idx int) {
	u.parents[idx] = idx
	u.ranks[idx] = 1
	u.sets++
}

// 看看 (row, col) 能否和 idx 合并
func (u *unionFindNd2) union(row, col int, idx1 int) {
	if row >= u.m || row < 0 || col >= u.n || col < 0 {
		// 说明索引不合法
		return
	}
	// 先计算出索引
	idx2 := u.index(row, col)
	if u.ranks[idx2] == 0 {
		// 说明 idx2 位置还不是岛
		return
	}

	root1 := u.findRoot(idx1)
	root2 := u.findRoot(idx2)
	if root1 == root2 {
		// 说明本身就在一个集合中了
		return
	}

	// 需要合并
	if u.ranks[root1] < u.ranks[root2] {
		u.parents[root1] = root2
	} else if u.ranks[root1] > u.ranks[root2] {
		u.parents[root2] = root1
	} else {
		u.parents[root1] = root2
		u.ranks[root2]++
	}

	u.sets--
}

func (u *unionFindNd2) findRoot(idx int) int {
	for idx != u.parents[idx] {
		// 将自己挂载到祖父节点上面，路径减半
		u.parents[idx] = u.parents[u.parents[idx]]
		idx = u.parents[idx]
	}

	return idx
}
