// @Author: Ciusyan 3/4/24

package day_33

import "strings"

// https://leetcode.cn/problems/word-break-ii/

func wordBreak(s string, wordDict []string) []string {
	// 构建前缀树
	prefixRoot := NewPrefixNode2()
	prefixRoot.AddWords(wordDict)
	// 获取 dp 表，知道 s 的构建情况
	dp := getWordBreakDp(s, prefixRoot)
	res := make([]string, 0, 1)
	track := make([]string, 0, len(s))
	// 从第 0 层开始搜索，传入 prefixRoot dp 用于快速剪枝
	wordBreakDfs(0, []byte(s), prefixRoot, dp, &track, &res)
	return res
}

// 从第 level 层开始搜索，搜索出 word[level ...] 为止
func wordBreakDfs(level int, word []byte, root *PrefixNode2, dp []bool, track *[]string, res *[]string) {
	if level == len(word) {
		// 收集路径，并且将其使用 " " 拼接起来
		*res = append(*res, strings.Join(*track, " "))
		return
	}

	curNode := root
	// 否则查看所有可能
	for end := level; end < len(word); end++ {
		// 先看看能不能找出一条路
		curNode = curNode.nexts[word[end]-'a']
		if curNode == nil {
			break
		}
		// 说明有路走
		if curNode.paths != "" && dp[end+1] {
			// 说明能记录一个轨迹
			*track = append(*track, curNode.paths)
			// 说明找到一个单词，去这个单词的后面搜索
			wordBreakDfs(end+1, word, root, dp, track, res)
			// 搜索完毕，记得还原现场
			*track = (*track)[:len(*track)-1]
		}
	}
}

type PrefixNode2 struct {
	paths string
	nexts []*PrefixNode2
}

func NewPrefixNode2() *PrefixNode2 {
	return &PrefixNode2{nexts: make([]*PrefixNode2, 26)}
}

func (n *PrefixNode2) AddWords(words []string) {

	for _, wordStr := range words {
		// 对于每一个单词，都从根节点开始查找
		curNode := n
		word := []byte(wordStr)

		// 挨个字符添加到前缀树上
		for cur := 0; cur < len(word); cur++ {
			curIdx := word[cur] - 'a'

			if curNode.nexts[curIdx] == nil {
				// 说明需要添加当前字符
				curNode.nexts[curIdx] = NewPrefixNode2()
			}
			// 然后去下一层
			curNode = curNode.nexts[curIdx]
		}
		// 标记当前单词是一个完整的单词
		curNode.paths = wordStr
	}
}

// 构建出 dp，dp[cur] 代表能否构建出 s[cur...] 这个单词
func getWordBreakDp(s string, root *PrefixNode2) []bool {

	word := []byte(s)
	n := len(word)

	// dp[cur] 代表能否构建出 word[cur ...] 这个单词
	dp := make([]bool, n+1)
	// word[n...] = ""，肯定能构建出 "" 这个单词
	dp[n] = true

	// 因为 dp[cur] 依赖 dp[cur+1]，所以从后往前求解
	for cur := n - 1; cur >= 0; cur-- {
		curNode := root
		// 看看 [cur... n] 能否被求解出来
		for end := cur; end < n; end++ {
			curNode = curNode.nexts[word[end]-'a']
			if curNode == nil {
				break
			}

			// 来到这里，说明有路
			if curNode.paths != "" && dp[end+1] {
				// 来到这里，说明是一个单词的结尾，并且之后的单词能被构建出来。
				dp[cur] = true
				// 知道有路就行了
				break
			}
		}
	}

	return dp
}

type Node struct {
	path  string
	end   bool
	nexts []*Node
}

func newNode() *Node {
	return &Node{
		path:  "",
		end:   false,
		nexts: make([]*Node, 26),
	}
}

func wordBreak22(s string, wordDict []string) []string {
	root := getTrie(wordDict)
	dp := getDp(s, root)
	var path []string
	var ans []string
	process([]rune(s), 0, root, dp, &path, &ans)
	return ans
}

func process(str []rune, index int, root *Node, dp []bool, path *[]string, ans *[]string) {
	if index == len(str) {
		*ans = append(*ans, strings.Join(*path, " "))
	} else {
		cur := root
		for end := index; end < len(str); end++ {
			road := str[end] - 'a'
			if cur.nexts[road] == nil {
				break
			}
			cur = cur.nexts[road]
			if cur.end && dp[end+1] {
				*path = append(*path, cur.path)
				process(str, end+1, root, dp, path, ans)
				*path = (*path)[:len(*path)-1]
			}
		}
	}
}

func getTrie(wordDict []string) *Node {
	root := newNode()
	for _, str := range wordDict {
		node := root
		for _, ch := range str {
			index := ch - 'a'
			if node.nexts[index] == nil {
				node.nexts[index] = newNode()
			}
			node = node.nexts[index]
		}
		node.path = str
		node.end = true
	}
	return root
}

func getDp(s string, root *Node) []bool {
	str := []rune(s)
	N := len(str)
	dp := make([]bool, N+1)
	dp[N] = true
	for i := N - 1; i >= 0; i-- {
		cur := root
		for end := i; end < N; end++ {
			path := str[end] - 'a'
			if cur.nexts[path] == nil {
				break
			}
			cur = cur.nexts[path]
			if cur.end && dp[end+1] {
				dp[i] = true
				break
			}
		}
	}
	return dp
}
