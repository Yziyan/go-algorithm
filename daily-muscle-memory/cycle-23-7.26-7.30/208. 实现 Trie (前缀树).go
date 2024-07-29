// @Author: Ciusyan 2024/7/29

package cycle_23_7_26_7_30

// https://leetcode.cn/problems/implement-trie-prefix-tree/description/

type Trie struct {
	root *node
}

type node struct {
	nexts  []*node
	isWord bool
}

func getRoot() *node {
	// 使用 a-‘a’ 作为索引
	return &node{nexts: make([]*node, 26)}
}

func Constructor() Trie {
	return Trie{root: getRoot()}
}

func (t *Trie) Insert(word string) {
	nd := t.root

	for _, c := range word {
		idx := c - 'a'
		if nd.nexts[idx] == nil {
			nd.nexts[idx] = getRoot()
		}
		nd = nd.nexts[idx]
	}

	nd.isWord = true
}

func (t *Trie) Search(word string) bool {
	nd := t.root

	for _, c := range word {
		nd = nd.nexts[c-'a']
		if nd == nil {
			return false
		}
	}

	return nd.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	nd := t.root

	for _, c := range prefix {
		nd = nd.nexts[c-'a']
		if nd == nil {
			return false
		}

	}

	return true
}
