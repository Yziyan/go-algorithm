// @Author: Ciusyan 3/25/24

package cycle_3_3_23_3_28

type LRUCache struct {
	cache    map[int]*node // 缓存
	first    *node         // 用于维护淘汰顺序，虚拟头结点
	last     *node         // 用于维护淘汰顺序，虚拟尾节点
	capacity int           // 容量
}

// 准备双向链表，需要维护节点顺序
type node struct {
	key  int
	val  int
	prev *node
	next *node
}

func newNode(key, val int) *node {
	return &node{
		key: key,
		val: val,
	}
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cache:    make(map[int]*node, capacity),
		first:    &node{},
		last:     &node{},
		capacity: capacity,
	}
	cache.first.next = cache.last
	cache.last.prev = cache.first

	return cache
}

func (this *LRUCache) Get(key int) int {
	nd, ok := this.cache[key]
	if !ok {
		// 说明没有这个元素
		return -1
	}
	// 来到这里，说明有这个元素，需要更新对应值的优先级，
	// 删除对应节点
	this.removeNode(nd)
	// 插在头部
	this.addNode2First(nd)
	return nd.val
}

func (this *LRUCache) Put(key int, value int) {
	nd, ok := this.cache[key]
	if ok {
		// 说明以前有，需要更新 val，并且更新位置
		nd.val = value
		// 需要删除 nd 的位置
		this.removeNode(nd)
	} else {
		// 说明之前不存在，需要添加，但是需要看看容量满了吗
		if len(this.cache) == this.capacity {
			// 说明容量满了，需要淘汰末尾的节点
			delete(this.cache, this.last.prev.key)
			this.removeNode(this.last.prev)
		}
		// 现在肯定可以加入 cache 了
		nd = newNode(key, value)
		this.cache[key] = nd
	}
	// 更新 nd 的位置
	this.addNode2First(nd)
}

// 删除 nd 在链表中的位置
func (this *LRUCache) removeNode(nd *node) {
	nd.prev.next = nd.next
	nd.next.prev = nd.prev
}

// 将 nd 添加至链表的头部
func (this *LRUCache) addNode2First(nd *node) {
	// first.next 与 nd 的线接好咯
	nd.next = this.first.next
	this.first.next.prev = nd

	// nd 的线维护好了
	nd.prev = this.first
	this.first.next = nd
}
