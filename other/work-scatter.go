// @Author: Ciusyan 2/1/24

package other

import "context"

type scatter struct {
	maxConsecutive int // 同商标允许的最大连续数 C
	stepLength     int // 需要插入的步长（需要选几个插上来）
	topScatterNum  int // 需要保证打散的头部数量，如果是 -1 代表全部

	compare ScatterCompare // 比较器
}

// ScatterCompare 比较器如何比较两个元素
type ScatterCompare func(item1, item2 string) int

// ScatterOption Options 模式选项
type ScatterOption func(*scatter)

// NewScatter 要使用打散器，必须要传入比较器
func NewScatter(compare ScatterCompare, opts ...ScatterOption) *scatter {
	res := &scatter{
		compare:        compare,
		maxConsecutive: 1, // 最大连续数默认为 1
		stepLength:     1, // 步长默认为 1
		topScatterNum:  1, // 对序列头部打散数默认为 1
	}

	// 将所有传入的选项，全部加载好
	for _, opt := range opts {
		opt(res)
	}

	return res
}

// WithMaxConsecutive 可传入最大连续数
func WithMaxConsecutive(maxConsecutive int) ScatterOption {
	return func(s *scatter) {
		s.maxConsecutive = maxConsecutive
	}
}

// WithStepLength 可传入调整步长
func WithStepLength(stepLength int) ScatterOption {
	return func(s *scatter) {
		s.stepLength = stepLength
	}
}

// WithTopScatterNum 可传入调整步长
func WithTopScatterNum(topScatterNum int) ScatterOption {
	return func(s *scatter) {
		s.topScatterNum = topScatterNum
	}
}

// DoScatter 进行打散
// reqList: 需要打散的列表
// compare: 元素间如何比较
// return: 打散后的列表
func (s *scatter) DoScatter(ctx context.Context, reqList []string) []string {
	if reqList == nil || len(reqList) == 0 {
		return reqList
	}

	// 先 copy 一份原序列，不要操作原序列
	finalList := make([]string, len(reqList))
	copy(finalList, reqList)

	size := len(reqList)
	if s.topScatterNum != -1 && s.topScatterNum < size {
		size = s.topScatterNum
	}

	// 最多打散 topScatterNum 个，但是可能总长度都没有那么长
	for cur := 0; cur < size; {
		if s.isConsecutiveTooMuch(finalList, cur) {
			// 说明超过了最大连续数，先找到第一个可交换的索引
			swapIndex := s.findFirstSwapIndex(finalList, cur, cur+s.maxConsecutive)
			if swapIndex == -1 {
				// 说明后面全是这个车牌，打散不了了
				break
			}

			// 挪动对应步长的数上来
			s.swapWithStepLength(finalList, cur+s.maxConsecutive, swapIndex)

			// 去到交换后的位置
			cur += s.maxConsecutive
		} else {
			cur++
		}
	}

	return finalList
}

// isConsecutiveTooMuch 是否超过了最大连续数
// list 待检查的列表
// start 从哪来开始检查
func (s *scatter) isConsecutiveTooMuch(list []string, start int) bool {
	count := 0
	for cur := start; cur < len(list) && s.compare(list[start], list[cur]) == 0; cur++ {
		// 说明 cur 和 start 位置的元素是一样的
		count++
		if count > s.maxConsecutive {
			return true
		}
	}
	return false
}

// findFirstSwapIndex 找到能交换的第一个索引（和当前 compare 不一样的第一个位置）
// list 待检查的列表
// standard 待交换的位置
// start 从哪来开始检查
func (s *scatter) findFirstSwapIndex(list []string, standard, start int) int {
	for i := start; i < len(list); i++ {
		if s.compare(list[standard], list[i]) != 0 {
			// 说明和当前位置的是不一样的
			return i
		}
	}
	return -1
}

// swapWithStepLength 交换 stepLength 个元素上来
// list 待交换的列表
// waitSwapIdx 待交换的位置
// firstSwapIdx 从哪来开始交换
func (s *scatter) swapWithStepLength(list []string, waitSwapIdx, firstSwapIdx int) {
	// 要挑这么多个上来
	for i := 0; i < s.stepLength; i++ {
		if firstSwapIdx >= len(list) {
			// 说明没有那么多可以交换了，直接返回
			return
		}
		// 先交换两个位置
		list[waitSwapIdx], list[firstSwapIdx] = list[firstSwapIdx], list[waitSwapIdx]
		// 然后都往后挪一位
		waitSwapIdx++
		firstSwapIdx++
	}
}
