// @Author: Ciusyan 4/10/24

package cycle_4_4_8_4_11

/**
思路重复：
经典的两数之和，
核心就是使用 HashTable 存储的数据是：<num, index>
从前往后遍历所有的 num，都使用 target - num 得到 remain
如果 remain 在 HashTable 里面有值，说明找到了能和 num 加起来，得到 target 的数
返回对应的索引即可。
*/

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	numIdx := make(map[int]int, len(nums))
	for idx, num := range nums {
		remain := target - num
		preIdx, ok := numIdx[remain]
		if ok {
			// 说明找到了结果
			return []int{preIdx, idx}
		}
		// 说明没有结果，将次值存起来
		numIdx[num] = idx
	}

	return nil
}
