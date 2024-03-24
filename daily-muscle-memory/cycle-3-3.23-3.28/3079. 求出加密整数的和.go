// @Author: Ciusyan 3/23/24

package cycle_3_3_23_3_28

// https://leetcode.cn/problems/find-the-sum-of-encrypted-integers/

/*
*
思路重复：
这个题最直接的一个想法就是，将每个数字，转换成字符串。
然后遍历找出其中的最大位，并且记录下有多少位，到时候就能构成多少个最到位。
然后再将其构建出来的最大位，转换成数字，加到最终结果中去。

但是这样需要经过字符串中转，能否直接对数字进行操作，得到对应的结果呢？
也不是不行，可以将其每个数字的每一位，挨次取出最后一位，然后得到最大位是多少。
并且在这个过程中，去构建位数个 1，那么最终此次的最大值，就是最大位*构建的位数个1了。
比如：num = 781，其中最大位 mx = 8，这个数有 3 位，base = 111
那么此次最大值就是：8*111 = 888，将其加到结果中即可
*/
func sumOfEncryptedInt2(nums []int) int {
	res := 0

	for _, num := range nums {
		mx, base := 0, 0
		for num > 0 {
			mx = max(mx, num%10)
			base = base*10 + 1

			num /= 10
		}
		res += base * mx
	}

	return res
}

func sumOfEncryptedInt(nums []int) int {

	res := 0
	for _, num := range nums {
		// num 的最大位是 mx，有多少位1
		mx, base := 0, 0
		for num > 0 {
			// 有几位，到时候就有多少位
			mx = max(num%10, mx)
			// 每进来一次，就要在末尾多增加一个 1，有多少位，就有多少个 1
			base = base*10 + 1
			// 去算下一位，去掉刚计算的最后一位
			num /= 10
		}
		// 此次 num 的产出就是：位数个1*最大位
		res += base * mx
	}
	return res
}
