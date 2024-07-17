// @Author: Ciusyan 2024/7/17

package cycle_21_7_16_7_20

import "sort"

// https://leetcode.cn/problems/group-anagrams/description/

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	// 准备 n 个桶
	n := len(strs)
	buckets := make(map[string][]string, n)
	// 挨个单词分桶存放
	for _, str := range strs {
		// 先对 str 的切片进行排序，要保留源 str
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		key := string(chars)
		bucket, ok := buckets[key]
		if !ok {
			// 说明是第一次，初始化这个桶
			bucket = make([]string, 0, 10)
		}
		buckets[key] = append(bucket, str)
	}

	res := make([][]string, 0, n)
	for _, bucket := range buckets {
		res = append(res, bucket)
	}
	return res
}
