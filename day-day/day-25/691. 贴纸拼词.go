// @Author: Ciusyan 11/17/23

package day_25

import (
	"math"
	"strings"
)

// https://leetcode.cn/problems/stickers-to-spell-word/

func minStickers1(stickers []string, target string) int {
	if stickers == nil || len(stickers) == 0 || target == "" {
		return 0
	}

	// 使用 src 的字符，能够抵消多少 target 中的字符，返回抵消后的字符
	trim := func(src, target string) string {
		if src == "" || target == "" {
			// 说明不能抵消掉任何字符，或者不需要抵消任何字符
			return target
		}

		// 因为出现的都是英文小写字符
		targetCount := make([]int, 26)
		for _, c := range []byte(target) {
			// 先统计 target 的字符频率
			targetCount[c-'a']++
		}

		// 再使用 src trim
		for _, c := range []byte(src) {
			// 在 targetCount 的基础上 trim
			targetCount[c-'a']--
		}

		// 然后，将所有字符收集起来
		var sb strings.Builder
		for i, count := range targetCount {
			if count <= 0 {
				// 说明不能加这个字符了
				continue
			}
			// 否则写入当前字符，count 次。
			for j := 0; j < count; j++ {
				sb.WriteByte(byte(i) + 'a')
			}
		}

		return sb.String()
	}

	// stickers 是所有贴纸，想要拼出 remain 的最少贴纸数，如果无法拼出，那就返回系统最大值
	var process func(stickers []string, remain string) int
	process = func(stickers []string, remain string) int {
		if len(remain) == 0 {
			// 说明所有字符都拼好了，需要 0 张贴纸
			return 0
		}

		// 默认是系统最大值
		res := math.MaxInt
		for _, first := range stickers {
			// 对每一张贴纸，都当做第一张贴纸，去看看拼出多少字符
			nextRemain := trim(first, remain)
			if len(remain) == len(nextRemain) {
				// 说明这张贴纸根本没有减掉任何字符，直接跳下一张贴纸尝试
				continue
			}

			// 否则说明这张贴纸减掉了一些字符，再看看，使用这些贴纸，要剪出 nextRemain 的贴纸需要的最少张数
			nextRes := process(stickers, nextRemain)
			if nextRes != math.MaxInt {
				// 说明拼出 remain，需要  nextRes + 1，看看是不是比 res 更小
				// 需要加上使用的 first 这张贴纸
				res = min(res, nextRes+1)
			}
		}

		return res
	}

	// 那么主函数该如何调用呢？stickers 的贴纸，要拼出 target 的最少贴纸数
	res := process(stickers, target)
	if res == math.MaxInt {
		return -1
	}

	return res
}
