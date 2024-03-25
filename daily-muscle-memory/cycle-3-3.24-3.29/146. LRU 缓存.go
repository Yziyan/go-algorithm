// @Author: Ciusyan 3/25/24

package cycle_3_3_24_3_29

/**
思路重复：
缓存很简单，使用 Map 即可。
但是要维护缓存中 key 的顺序，这个时候其实就可以使用双向链表来维护了。
当获取元素的时候，获取对应的元素，并且删除在双向链表中对应的节点，然后将其节点加入到链表的头部
当更新元素的时候，将其对应节点的值修改了，然后删除在双向链表中对应的节点，然后将其节点加入链表的头部。
当新增元素的时候，先看看是否达到了最大容量，如果容量满了，就删除在链表尾部的节点，并且删除对应 key 对应的缓存。
	然后再加入新的缓存，并且将其放置在链表的头部。
*/

//type LRUCache struct {
//	cache    map[int]*node // 缓存
//	first    *node         // 用于维护淘汰顺序，虚拟头结点
//	last     *node         // 用于维护淘汰顺序，虚拟尾节点
//	capacity int           // 容量
//}
//
//// 准备双向链表，需要维护节点顺序
//type node struct {
//	key  int
//	val  int
//	prev *node
//	next *node
//}
//
//func newNode(key, val int) *node {
//	return &node{
//		key: key,
//		val: val,
//	}
//}
//
//func Constructor(capacity int) LRUCache {
//	cache := LRUCache{
//		cache:    make(map[int]*node, capacity),
//		first:    &node{},
//		last:     &node{},
//		capacity: capacity,
//	}
//	cache.first.next = cache.last
//	cache.last.prev = cache.first
//
//	return cache
//}
//
//func (this *LRUCache) Get(key int) int {
//	nd, ok := this.cache[key]
//	if !ok {
//		// 说明没有这个元素
//		return -1
//	}
//	// 来到这里，说明有这个元素，需要更新对应值的优先级，
//	// 删除对应节点
//	this.removeNode(nd)
//	// 插在头部
//	this.addNode2First(nd)
//	return nd.val
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	nd, ok := this.cache[key]
//	if ok {
//		// 说明以前有，需要更新 val，并且更新位置
//		nd.val = value
//		// 需要删除 nd 的位置
//		this.removeNode(nd)
//	} else {
//		// 说明之前不存在，需要添加，但是需要看看容量满了吗
//		if len(this.cache) == this.capacity {
//			// 说明容量满了，需要淘汰末尾的节点
//			delete(this.cache, this.last.prev.key)
//			this.removeNode(this.last.prev)
//		}
//		// 现在肯定可以加入 cache 了
//		nd = newNode(key, value)
//		this.cache[key] = nd
//	}
//	// 更新 nd 的位置
//	this.addNode2First(nd)
//}
//
//// 删除 nd 在链表中的位置
//func (this *LRUCache) removeNode(nd *node) {
//	nd.prev.next = nd.next
//	nd.next.prev = nd.prev
//}
//
//// 将 nd 添加至链表的头部
//func (this *LRUCache) addNode2First(nd *node) {
//	// first.next 与 nd 的线接好咯
//	nd.next = this.first.next
//	this.first.next.prev = nd
//
//	// nd 的线维护好了
//	nd.prev = this.first
//	this.first.next = nd
//}

// 双向链表节点
type node struct {
	key  int
	val  int
	prev *node
	next *node
}

func newNode(key, val int) *node {
	return &node{key: key, val: val}
}

type LRUCache struct {
	cache    map[int]*node
	first    *node
	last     *node
	capacity int
}

func Construct(capacity int) LRUCache {
	// 准备两个虚拟节点
	first := &node{}
	last := &node{}
	first.next = last
	last.prev = first

	return LRUCache{
		cache:    make(map[int]*node, capacity),
		first:    first,
		last:     last,
		capacity: capacity,
	}
}

func (c *LRUCache) Get(key int) int {
	nd, ok := c.cache[key]
	if !ok {
		return -1
	}
	// 否则需要更新位置
	c.remove(nd)
	c.addNode2First(nd)

	return nd.val
}

func (c *LRUCache) Put(key, val int) {
	nd, ok := c.cache[key]
	if ok {
		// 说明以前存在，修改值，放入头部
		nd.val = val
		c.remove(nd)
	} else {
		// 说明以前不存在，插入前看看容量满了吗
		if len(c.cache) == c.capacity {
			// 说明需要删除对应节点
			delete(c.cache, c.last.prev.key)
			// 还有链表中
			c.remove(c.last.prev)
		}
		// 到这里，肯定可以添加了
		nd = newNode(key, val)
		// 放入缓存
		c.cache[key] = nd
	}

	// 然后将对应节点加入链表头
	c.addNode2First(nd)
}

// 删除 nd 节点
func (c *LRUCache) remove(nd *node) {
	nd.prev.next = nd.next
	nd.next.prev = nd.prev
}

// 添加 nd 至链表头部
func (c *LRUCache) addNode2First(nd *node) {
	c.first.next.prev = nd
	nd.next = c.first.next

	nd.prev = c.first
	c.first.next = nd
}
