// @Author: Ciusyan 10/31/23

package day_22

// UnionFind 通用版并查集
//
// Quick Union + rank + 路径减半
type UnionFind[V comparable] struct {
	// 标识所有的集合
	sets map[V]*node[V]
}

// 内部节点类
type node[V comparable] struct {
	val    V
	parent *node[V]
	rank   int
}

func newNode[V comparable](val V) *node[V] {
	n := &node[V]{
		val:  val,
		rank: 1,
	}
	// 默认自己就是自己的 parent
	n.parent = n
	return n
}

func NewUnionFind[V comparable]() *UnionFind[V] {
	return &UnionFind[V]{
		sets: make(map[V]*node[V]),
	}
}

// MakeSets 添加一个元素（自成一个集合）
func (uf *UnionFind[V]) MakeSets(vals ...V) {
	// 将其 val 建立集合
	for i := range vals {
		val := vals[i]
		uf.sets[val] = newNode(val)
	}
}

// IsSame 查看 v1 v2 是否处于同一个集合
func (uf *UnionFind[V]) IsSame(v1, v2 V) bool {
	p1 := uf.Find(v1)
	p2 := uf.Find(v2)

	return p1 == p2
}

// Find 返回 val 所在集合的根节点
func (uf *UnionFind[V]) Find(val V) V {
	root := uf.findRoot(val)
	if root == nil {
		return uf.zero()
	}

	return root.val
}

// Union 将 v1 和 v2 所在集合合并，基于 rank
func (uf *UnionFind[V]) Union(v1, v2 V) {
	root1 := uf.findRoot(v1)
	root2 := uf.findRoot(v2)
	if root1 == nil || root2 == nil {
		// 说明至少有一个都还不在集合中
		return
	}
	if root1.val == root2.val {
		// 说明本身就处于一个集合了
		return
	}

	// 来到这里，说明需要合并，根据 rank 合并（矮的挂高的）
	if root1.rank < root2.rank {
		// 说明 v1 所在集合矮，将 v1 挂载到 v2 的集合
		root1.parent = root2
	} else if root1.rank > root2.rank {
		// 说明 v2 所在集合矮，将 v2 挂载到 v1 的集合
		root2.parent = root1
	} else {
		// 一样高，谁挂谁都可以，但是被挂的那边 rank 需要增加
		root1.parent = root2
		root2.rank++
	}
}

// 找到 val 的根节点
//
// 这里采用基于路径减半的方式优化
func (uf *UnionFind[V]) findRoot(val V) *node[V] {
	// 先获取 val 的节点
	nd, ok := uf.sets[val]
	if !ok {
		return nil
	}

	// 一路向上，直到 nd.parent = nd
	for nd.parent != nd {
		// 将自己挂载到祖父节点上面
		nd.parent = nd.parent.parent

		// 然后让祖父节点也做这样的操作
		nd = nd.parent
	}

	return nd
}

// 返回泛型的零值
func (uf *UnionFind[V]) zero() (v V) {
	return
}
