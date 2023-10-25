// @Author: Ciusyan 10/25/23

package day_21

import (
	"sort"
	"strings"
)

// CompareStr 要自定义排序，所以定义一个类型
type CompareStr []string

func (c CompareStr) Len() int {
	return len(c)
}

func (c CompareStr) Less(i, j int) bool {
	// 采用贪心策略，每次都取 strs[i] + strs[j] 字典序最小的，拼接在 res 的后面
	// return c[i] < c[j]
	return c[i]+c[j] < c[j]+c[i]
}

func (c CompareStr) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// LowestString 给定一个由字符串组成的数组strs，必须把所有的字符串拼接起来，返回所有可能的拼接结果中字典序最小的结果。
func LowestString(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	// 采用贪心策略，每次都取 strs[i] + strs[j] 字典序最小的，拼接在 res 的后面
	sort.Sort(CompareStr(strs))

	// 排序过后，得到的结果，从前往后拼接成一个字符串返回即可
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}

	return sb.String()
}

// LowestString2 对数器写法
func LowestString2(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	// 删除 idx 位置的元素，返回新的切片
	removeIdx := func(strs []string, idx int) []string {
		rmCurStrs := append([]string(nil), strs[0:idx]...)
		return append(rmCurStrs, strs[idx+1:]...)
	}

	// 先收集所有结果
	var getAllRes func(strs1 []string) []string
	getAllRes = func(strs1 []string) []string {
		res := make([]string, 0, 1)
		if strs1 == nil || len(strs1) == 0 {
			res = append(res, "")
			return res
		}

		// 否则说明可以去收集结果，挨个收集
		for i, str := range strs1 {
			// str 作为第一个字符串，除去 str 后，去收集之后的
			rmCurStrs := removeIdx(strs1, i)
			// 收集之后的结果
			nexts := getAllRes(rmCurStrs)
			// 收集到结果后，将后面收集到的所有结果保存到当前结果中
			for _, next := range nexts {
				// 但是别忘了，都要加上第一个字符串
				res = append(res, str+next)
			}
		}

		return res
	}

	allRes := getAllRes(strs)
	res := allRes[0]
	// 选一个字典序最小的结果
	for i := 1; i < len(allRes); i++ {
		if allRes[i] < res {
			res = allRes[i]
		}
	}

	return res
}
