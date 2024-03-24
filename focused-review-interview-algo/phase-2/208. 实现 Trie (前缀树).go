// @Author: Ciusyan 3/24/24

package phase_2

type node struct {
	pass int  // 有多少单词穿过此节点
	word bool // 是否是一个完整的单词结尾

	nexts []*node // 有 26 条路径
}

func newNode() *node {
	return &node{nexts: make([]*node, 26)}
}

type Trie struct {
	root *node
}

func Constructor() Trie {
	return Trie{root: newNode()}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	// 从根节点开始，往下找
	nd := this.root
	nd.pass++
	for _, c := range word {
		cIdx := c - 'a' // 映射到索引里面
		// 看看 root 上有没有挂这个字符
		next := nd.nexts[cIdx]
		if next == nil {
			// 说明没有挂载过，直接新建一个
			nd.nexts[cIdx] = newNode()
		}

		// 到达这里，说明 next 一定有值了
		next.pass++
		// 往下传
		nd = nd.nexts[cIdx]
	}

	// word 遍历完了，nd 是一恶搞单词的结尾了
	nd.word = true
}

func (this *Trie) Search(word string) bool {
	if word == "" {
		// 空串一定在
		return true
	}

	nd := this.searchNode(word)
	return nd != nil && nd.word
}

func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		// 空串一定在
		return true
	}

	nd := this.searchNode(prefix)
	return nd != nil
}

// 搜索对应的 node
func (this *Trie) searchNode(word string) *node {
	if len(word) == 0 {
		return nil
	}

	// 从根节点开始搜索
	nd := this.root
	for _, c := range word {
		cIdx := c - 'a'
		next := nd.nexts[cIdx]
		if next == nil {
			// 说明没有下一个节点了，找不到对应的节点
			return nil
		}
		nd = next
	}

	return nd
}
