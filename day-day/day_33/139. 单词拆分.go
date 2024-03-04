// @Author: Ciusyan 2/23/24

package day_33

// https://leetcode.cn/problems/word-break/description/

// PrefixNode 前缀树节点
type PrefixNode struct {
	isWord bool
	nexts  []*PrefixNode
}

func NewPrefixNode() *PrefixNode {
	return &PrefixNode{nexts: make([]*PrefixNode, 26)}
}

func (n *PrefixNode) AddWithList(wordDict []string) {
	for _, wordStr := range wordDict {
		// 从根节点开始找
		node := n
		word := []byte(wordStr)
		curIdx := 0

		// 挨个字母加入
		for cur := 0; cur < len(word); cur++ {
			// 先看看当前 node 的 nexts 中是否已经存在了
			curIdx = int(word[cur] - 'a')
			if node.nexts[curIdx] == nil {
				// 说明以前不存在 cur 这个前缀
				node.nexts[curIdx] = NewPrefixNode()
			}
			// 然后将下一个字母加入
			node = node.nexts[curIdx]
		}
		// 一个单词加完了，标记结尾是单词
		node.isWord = true
	}
}

func wordBreak1(s string, wordDict []string) bool {
	// 先准备一个前缀树的根节点，然后将 wordDict 的单词全部加上去
	root := NewPrefixNode()
	root.AddWithList(wordDict)
	// 当所有单词都添加到前缀树后，动态规划来查找是否可以切分出单词

	chars := []byte(s)
	n := len(chars)
	// 准备缓存，dp[cur] 代表：利用 s[cur ...] 是否能被拼出
	dp := make([]bool, n+1)
	dp[n] = true // 代表 "" 肯定能被拼出

	// dp[cur] 依赖 dp[cur+x]
	for cur := n - 1; cur >= 0; cur-- {
		// 从 root 开始查找是否有 chars[cur] 这个前缀
		curNode := root
		for end := cur; end < n; end++ {
			curNode = curNode.nexts[chars[end]-'a']
			if curNode == nil {
				// 说明不存在这样的前缀，不用看了
				break
			}
			// 说明有这样的前缀，看看是不是某个单词的结尾
			if curNode.isWord && dp[end+1] {
				// 说明是某个单词的结尾，即 [cur ... end] 能被拼出了，
				// 如果 [end+1 ... n] 也能被拼出，就代表可以拼出 [cur ... n]
				dp[cur] = true
				break
			}
		}
	}

	// 代表：s[0 ...] 这个单词，能否被拼接出来
	return dp[0]
}
