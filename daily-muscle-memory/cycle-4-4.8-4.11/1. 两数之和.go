// @Author: Ciusyan 4/10/24

package cycle_4_4_8_4_11

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
