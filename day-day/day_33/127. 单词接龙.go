// @Author: Ciusyan 1/31/24

package day_33

// https://leetcode.cn/problems/word-ladder/

type Set map[string]struct{}

func NewDefaultSet() Set {
	return make(map[string]struct{}, 1)
}

func NewSet(keys []string) Set {
	res := make(map[string]struct{}, len(keys))

	for _, key := range keys {
		res[key] = struct{}{}
	}

	return res
}

func (s *Set) Put(str string) {
	(*s)[str] = struct{}{}
}

func (s *Set) Contains(key string) bool {
	_, ok := (*s)[key]
	return ok
}

func (s *Set) Size() int {
	return len(*s)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 先将所有 wordList 放入 Set 中，稍后好查找
	wordSet := NewSet(wordList)
	if !wordSet.Contains(endWord) {
		return 0
	}

	var (
		// 准备几个用于 BFS 的 Set
		startSet = NewDefaultSet() // 从头部开始接龙用
		endSet   = NewDefaultSet() // 从尾部开始接龙用
		visitSet = NewDefaultSet() // 记录已接过的单词

		// 长度得算上首尾单词，所以从 2 开始
		resL = 2
	)

	// 将起始和结束单词放置在对应的集合中
	startSet.Put(beginWord)
	endSet.Put(endWord)

	// 默认每次都从 startSet 开始（单词少的）
	for startSet.Size() != 0 {
		// 每次都从 start 开始
		nextSet := NewDefaultSet() // 用于过渡下一次从哪边开始遍历

		// 从 startSet 开始，挨个单词尝试接龙
		for word := range startSet {

			// 将 word 能用的邻居，挨个尝试看看
			for cur := 0; cur < len(word); cur++ {
				curWord := []rune(word)
				for c := 'a'; c <= 'z'; c++ {
					if curWord[cur] == c {
						// 说明和当前字符相等，就是当前单词
						continue
					}

					// 来到这里，说明可以作为邻居
					curWord[cur] = c
					nextWord := string(curWord)

					if endSet.Contains(nextWord) {
						// 说明接龙用的单词，已经存在于 endSet 中了，说明找到了一条最短的路径
						return resL
					}
					// 否则将这个单词作为一个备选接龙项。
					if wordSet.Contains(nextWord) && !visitSet.Contains(nextWord) {
						// 说明单词表中拥有这个单词，并且这个单词没有使用过
						nextSet.Put(nextWord)
						visitSet.Put(nextWord)
					}
				}
			}

		}

		// 让 startSet 变成单词少的
		if nextSet.Size() > endSet.Size() {
			// 说明 endSet 单词少
			startSet = endSet
			endSet = nextSet
		} else {
			startSet = nextSet
		}

		resL++
	}

	return 0
}
