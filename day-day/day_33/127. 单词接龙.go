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
	// 先将单词表放在集合中，方便后面查找
	wordSet := NewSet(wordList)

	if !wordSet.Contains(endWord) {
		return 0
	}

	var (
		beginSet   = NewDefaultSet() // 从开始端 BFS 的集合
		endSet     = NewDefaultSet() // 从结束端 BFS 的集合
		visitedSet = NewDefaultSet() // 记录哪些单词已经被使用过了
	)

	// 将开始单词和结尾单词放在对应的集合中
	beginSet.Put(beginWord)
	endSet.Put(endWord)

	// 长度默认包含首尾两个单词
	resL := 2

	for beginSet.Size() != 0 {
		nextSet := NewDefaultSet() // 用于过渡下一次从那端开始的集合

		// 默认 begin 端单词少，从这开始 BFS 要快一些
		for curWord := range beginSet {

			// 暴力枚举出能用于当前单词接龙的所有可能
			for cur := 0; cur < len(curWord); cur++ {
				// 每次尝试的时候，都置为原始单词
				chars := []rune(curWord)
				// 每个字符尝试换一下
				for c := 'a'; c <= 'z'; c++ {
					if c == chars[cur] {
						// 和当前单词相同，就别尝试了
						continue
					}

					// 说明可以尝试接龙
					chars[cur] = c
					nextWord := string(chars)

					if endSet.Contains(nextWord) {
						// 说明接龙的单词已经在另一端的集合中了，可以链接成龙了
						return resL
					}

					// 但是用于接龙的单词，必须要保证：1.在可选单词集合里 2.没有使用过
					if wordSet.Contains(nextWord) && !visitedSet.Contains(nextWord) {
						// 说明可以接龙，将其放置在 Next 中，去下一层接龙
						nextSet.Put(nextWord)
						visitedSet.Put(nextWord)
					}
				}
			}
		}

		// 保持 begin 端单词最少
		if nextSet.Size() > endSet.Size() {
			// 说明 end 端单词要少一点，从这边 BFS 快一点
			beginSet = endSet
			endSet = nextSet
		} else {
			// 说明从开始端要快一点
			beginSet = nextSet
		}

		// 到达这里说明选了一个单词，去下一层
		resL++
	}

	// 能来到这里，说明全部尝试完了，也么有成功接轨，说明连不成龙
	return 0
}
