// @Author: Ciusyan 1/16/24

package day_30

import "sort"

// https://leetcode.cn/problems/group-anagrams/

func groupAnagrams(strs []string) [][]string {
	if strs == nil || len(strs) == 0 {
		return nil
	}
	l := len(strs)
	// 准备一个哈希表，用于分类异位词
	bucketMap := make(map[string][]string, l)

	for _, str := range strs {
		// 将字符串按照从小到大排序
		chars := []rune(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		// 然后加入 Map 中，进行分组
		values, ok := bucketMap[string(chars)]
		if !ok {
			// 说明是第一次加入 Map
			values = make([]string, 0, 1)
		}
		// 那么所有排序后的字符串，都会被加入一个桶中
		bucketMap[string(chars)] = append(values, str)
	}

	res := make([][]string, 0, len(bucketMap))
	// 然后再遍历 Map，取出所有桶中的值
	for _, values := range bucketMap {
		res = append(res, values)
	}

	return res
}
