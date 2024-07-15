// @Author: Ciusyan 2024/7/16

package cycle_21_7_16_7_20

// https://leetcode.cn/problems/word-break/description/

// PrefixNode 前缀树节点
type PrefixNode struct {
	isWord bool // 是否是单词的结尾
	nexts  []*PrefixNode
}

func NewPrefixNode() *PrefixNode {
	// 只能添加小写字母到前缀树上，c - 'a' 既是对应的索引
	return &PrefixNode{nexts: make([]*PrefixNode, 26)}
}

// Add 将 words 的所有单词添加到前缀树上
func (n *PrefixNode) Add(words []string) {
	for _, word := range words {
		// 从根节点开始找
		root := n
		// 挨个字符添加到前缀树上
		for _, c := range word {
			// 算出当前字符的索引
			curIdx := int(c - 'a')
			if root.nexts[curIdx] == nil {
				// 说明 c 不在前缀树上
				root.nexts[curIdx] = NewPrefixNode()
			}
			// 来到这里，说明 c 存在了，继续往下一层查找
			root = root.nexts[curIdx]
		}
		// 到达这里，说明 word 被添加完成了，标记结尾为单词
		root.isWord = true
	}
}

func wordBreak(s string, wordDict []string) bool {
	// 先将单词的字典全部添加到前缀树上
	root := NewPrefixNode()
	root.Add(wordDict)

	n := len(s)
	// dp[cur] 代表，s[cur ...] 这些字符，能否被前缀树构建出来
	dp := make([]bool, n+1)
	dp[n] = true // 代表 "" 肯定能被构建出来

	// 若 dp[end+1] 能够被构建出来，s[cur... end] 又是一个单词的结尾，那么肯定能构建出 dp[cur]
	for cur := n - 1; cur >= 0; cur-- {
		nd := root
		for end := cur; end < n; end++ {
			nd = nd.nexts[s[end]-'a'] // 先看看 s[end] 是否在这个前缀树上
			if nd == nil {
				// 说明不在，直接返回
				break
			}

			// 这里说明在前缀树上，看看是否是结尾
			if nd.isWord && dp[end+1] {
				// 说明 s[cur ... end] 是一个字典中的单词，
				// 并且 dp[end+1] 之前又计算过了，能够构建出来，
				// 那么就说明 s[cur ...] 能够被构建出来
				dp[cur] = true
				break
			}
		}
	}

	// 代表：s[0 ...] 这些字符，能否被构建出来
	return dp[0]
}
