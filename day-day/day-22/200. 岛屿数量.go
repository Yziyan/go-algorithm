// @Author: Ciusyan 11/1/23

package day_22

// https://leetcode.cn/problems/number-of-islands/description/

// 并查集做法（数组实现）
func numIslands(grid [][]byte) int {
	rowLen := len(grid)
	colLen := len(grid[0])

	uf := newUnionFindNd(rowLen, colLen)
	// 将二维数组中的 1，全部建立集合
	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if grid[row][col] == '1' {
				uf.makeSet(row, col)
			}
		}
	}

	// 挨个合并左上
	for col := 1; col < colLen; col++ {
		// 先合并第一行，就不用看上方了
		if grid[0][col-1] == '1' && grid[0][col] == '1' {
			// 说明左边可以合并成一个集合
			uf.union(0, col-1, 0, col)
		}
	}

	for row := 1; row < rowLen; row++ {
		// 再合并第一列，就不用看左边了
		if grid[row-1][0] == '1' && grid[row][0] == '1' {
			// 说明上边可以合并成一个集合
			uf.union(row-1, 0, row, 0)
		}
	}

	for row := 1; row < rowLen; row++ {
		for col := 1; col < colLen; col++ {
			if grid[row][col] != '1' {
				// 如果当前格子不是 1，就没必要看看上左了
				continue
			}
			// 来到这里说明需要看看上和左能否合并
			// 查看左边
			if grid[row][col-1] == '1' {
				uf.union(row, col, row, col-1)
			}
			// 查看上面
			if grid[row-1][col] == '1' {
				uf.union(row, col, row-1, col)
			}
		}
	}

	return uf.sets
}

// row * col 这么多的集合
func newUnionFindNd(row, col int) *unionFindNd {
	// 二维转一维后，至少需要的容量
	l := row * col
	return &unionFindNd{
		parents: make([]int, l),
		rank:    make([]int, l),
		row:     row,
		col:     col,
	}
}

type unionFindNd struct {
	parents []int // 父节点数组
	rank    []int // 集合高度
	sets    int   // 集合数量

	row, col int // 矩阵的行和列
}

// 根据行号列号，转换出 (row, col) 在一维数组中的索引
func (u *unionFindNd) index(row, col int) int {
	return row*u.col + col
}

func (u *unionFindNd) makeSet(row, col int) {
	// 这个位置需要新建一个集合
	idx := u.index(row, col)
	u.parents[idx] = idx
	u.rank[idx] = 1
	u.sets++
}

func (u *unionFindNd) union(row1, col1 int, row2, col2 int) {
	// 先计算出两个的索引
	idx1 := u.index(row1, col1)
	idx2 := u.index(row2, col2)

	root1 := u.findRoot(idx1)
	root2 := u.findRoot(idx2)

	if root1 == root2 {
		// 说明本来就在一个集合
		return
	}

	// 来到这里说明需要合并，将矮的合并到高的上面
	if u.rank[root1] < u.rank[root2] {
		u.parents[root1] = root2
	} else if u.rank[root1] > u.rank[root2] {
		u.parents[root2] = root1
	} else {
		// 合并谁都可以，但是需要长高树
		u.parents[root1] = root2
		u.rank[root2]++
	}

	// 合并了，sets 肯定得减少
	u.sets--
}

func (u *unionFindNd) findRoot(idx int) int {
	// 当 idx 的父亲是自己时，就说明到顶端了
	for idx != u.parents[idx] {
		// 将自己挂载到祖父身上，路径减半
		u.parents[idx] = u.parents[u.parents[idx]]

		// 让祖父也去左这个操作
		idx = u.parents[idx]
	}

	return idx
}

// 并查集做法（通用版）
func numIslands1(grid [][]byte) int {

	rowLen := len(grid)
	colLen := len(grid[0])

	uf := NewUnionFind[*byte]()
	// 将所有岛屿都加入并查集中
	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if grid[row][col] == '1' {
				// 这里使用 map，键不一样才行，所以用 grid[row][col] 的地址
				uf.MakeSets(&grid[row][col])
			}
		}
	}

	// 现在总共有那么多的岛屿，但是挨个查看，是不是相邻的，如果是相邻的，那么我们将其合并
	for col := 1; col < colLen; col++ {
		// 先合并第一行，只需要看左边即可
		if grid[0][col-1] == '1' && grid[0][col] == '1' {
			uf.Union(&grid[0][col-1], &grid[0][col])
		}
	}
	for row := 1; row < rowLen; row++ {
		// 再合并第一列，只需要看上边即可
		if grid[row-1][0] == '1' && grid[row][0] == '1' {
			uf.Union(&grid[row-1][0], &grid[row][0])
		}
	}

	for row := 1; row < rowLen; row++ {
		for col := 1; col < colLen; col++ {
			// 最后再来看其他的，上面和左边
			if grid[row][col] == '1' {
				if grid[row][col-1] == '1' {
					// 上面
					uf.Union(&grid[row][col], &grid[row][col-1])
				}

				if grid[row-1][col] == '1' {
					// 左边
					uf.Union(&grid[row][col], &grid[row-1][col])
				}
			}
		}
	}

	return uf.GetSize()
}

// 染色做法
func numIslands2(grid [][]byte) int {

	rowLen := len(grid)
	colLen := len(grid[0])
	// 定义一个染色方法，从 grid[row][col] 出发，将所有相邻的 1 都染成其他颜色
	var infact func(row, col int)
	infact = func(row, col int) {
		if row >= rowLen || row < 0 || col >= colLen || col < 0 || grid[row][col] != '1' {
			// 1.越界的情况 2.[row][col] 不是岛屿，统统不染色
			return
		}

		// 来到这里，将 [row][col] 的上下左右全染色了，但是需要将当前字符染色，要不然退不出循环
		grid[row][col] = '2'

		infact(row-1, col) // 上
		infact(row+1, col) // 下
		infact(row, col-1) // 左
		infact(row, col+1) // 右
	}

	res := 0
	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if grid[row][col] == '1' {
				// 只要触发一次染色，那么就会
				res++
				infact(row, col)
			}
		}
	}

	return res
}
