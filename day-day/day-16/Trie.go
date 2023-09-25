// @Author: Ciusyan 9/25/23

package day_16

// 前缀树的实现：https://leetcode.cn/problems/implement-trie-prefix-tree/

type Trie struct {
	root *node
}

type node struct {
	pass  int     // 有多少字符通过过当前节点
	end   int     // 有多少个完整的单词
	paths []*node // 总共有 26 条路
}

func newNode() *node {
	return &node{
		// 默认先把 26 条路建好（26 个字符）
		//	node[0] -> 'a' ... node[25] -> 'z'
		//	node[i] == nil -> 不存在 i 字符
		//	node[i] != nil -> 存在 i 字符
		paths: make([]*node, 26),
	}
}

func Constructor() Trie {
	return Trie{root: newNode()}
}

func (this *Trie) Insert(word string) {
	if word == "" {
		return
	}
	// 先转成字符数组，方便操作
	chars := []byte(word)

	nd := this.root
	nd.pass++
	// 遍历每一个字符
	for _, c := range chars {
		path := c - 'a'
		if nd.paths[path] == nil {
			// 说明当前字符还未添加过，添加到路径中
			nd.paths[path] = newNode()
		}
		// 来到这里，肯定存在当前的路了
		nd = nd.paths[path]
		nd.pass++
	}

	// 来到这里，肯定所有的字符都遍历结束了，设置最后 nd 的状态
	nd.end++
}

func (this *Trie) Search(word string) bool {
	if word == "" {
		return true
	}

	nd := this.searchNode(word)
	if nd == nil {
		// 说明在某个字符中断了
		return false
	}

	// 来到最后，说明所有字符都在树上，但是是不是一个完整的单词呢？
	return nd.end != 0
}

func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}
	nd := this.searchNode(prefix)

	// 来到这里，如果 nd 不是 nil 说明全部字符都在 trie 上面
	return nd != nil
}

// 返回 str 能走到的目的地，如果树上还没有 str 的某个字符，返回 nil
func (this *Trie) searchNode(str string) *node {
	chars := []byte(str)
	nd := this.root
	// 遍历每一个字符
	for _, c := range chars {
		path := c - 'a'
		if nd.paths[path] == nil {
			// 说明当前数上没有此字符
			return nil
		}
		nd = nd.paths[path]
	}

	return nd
}

func (this *Trie) Delete(word string) {
	if !this.Search(word) {
		return
	}
	chars := []byte(word)
	// 来到这里，说明树上肯定有这个单词了，将其删除
	nd := this.root
	nd.pass--
	for _, c := range chars {
		path := c - 'a'
		nd.paths[path].pass--
		if nd.paths[path].pass == 0 {
			// 说明此字符只添加过一次，直接删除后面的字符即可，不用往下遍历了
			nd.paths[path] = nil
			return
		}
		nd = nd.paths[path]
	}
	// 来到这里，说明到达了末尾
	nd.end--
}
