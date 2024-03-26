// @Author: Ciusyan 3/26/24

package phase_2

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		preMax = nums[0]
		maxRes = nums[0]
	)

	for i := 1; i < len(nums); i++ {
		p1 := nums[i]
		p2 := preMax + nums[i]

		curMax := max(p1, p2)
		maxRes = max(maxRes, curMax)

		preMax = curMax
	}

	return maxRes
}
