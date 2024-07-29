// @Author: Ciusyan 2024/7/29

package cycle_23_7_26_7_30

// https://leetcode.cn/problems/implement-trie-prefix-tree/description/

/**
Trie 前缀树，是将单词的每一个字母，都加入一颗多叉树中。
这样就可以快速判断，某单词或者某前缀，是否出现在树上了。

首先是节点的构造，我们这里可以实现简单一些。因为之后 26 个字母，所以我们将节点的 nexts 大小，设置为 26。
然后每一个字母对应的索引，是：idx = c - 'a'

对于添加而言，从根节点出发，没遇到一层，就看看当前字符的索引是否已经构建 node 了，构建过就往下一层走，没有构建就构建一个再往下走

对于查找而言，也是从根节点出发，一层层往下走，如果遇到某一层没有节点，就说嘛没有这个单词，直接返回 false。
如果单词找完了，还得看看是不是一个完整的单词，如果不是也不行。
*/

type Trie2 struct {
	root *node2
}

type node2 struct {
	nexts  []*node2
	isWord bool
}

func getRoot2() *node2 {
	return &node2{nexts: make([]*node2, 26)}
}

func NewTrie() *Trie2 {
	return &Trie2{root: getRoot2()}
}

func (t *Trie2) Add(word string) {

	nd := t.root

	for _, c := range word {
		idx := c - 'a'
		if nd.nexts[idx] == nil {
			nd.nexts[idx] = getRoot2()
		}
		nd = nd.nexts[idx]
	}
	nd.isWord = true
}

func (t *Trie2) Exist(word string) bool {
	nd := t.root

	for _, c := range word {
		idx := c - 'a'
		nd = nd.nexts[idx]
		if nd == nil {
			return false
		}
	}
	return nd.isWord

}
func (t *Trie2) Prefix(word string) bool {
	nd := t.root

	for _, c := range word {
		idx := c - 'a'
		nd = nd.nexts[idx]
		if nd == nil {
			return false
		}
	}
	return true
}

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
