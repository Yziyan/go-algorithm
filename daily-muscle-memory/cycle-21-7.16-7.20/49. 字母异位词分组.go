// @Author: Ciusyan 2024/7/17

package cycle_21_7_16_7_20

import "sort"

// https://leetcode.cn/problems/group-anagrams/description/

/**
思路重复：
要想找出所有的异位词，其实就是给他分桶。
当然最简单的方式就是直接使用 map，然后挨个 strs 进行分桶。
将每个单词，排序后，作为 key，异位词排序后肯定相同，那么就能够被放到一个桶中，
所以也就能够直接分组了。
*/

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
