// @Author: Ciusyan 3/19/24

package cycle_2_3_18_3_22

func reverseBits(num uint32) uint32 {
	// 采用分治的思想，依次将：num 的 高16和低16 交换，每一组的高八和第八交换，... 每一组的高一和低一交换

	// 将 num 的 高16和低16 交换
	num = (num >> 16) | (num << 16)
	// 将 num 每组的 高8和低8 交换
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	// 将 num 每组的 高4和低4 交换
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	// 将 num 每组的 高2和低2 交换
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	// 将 num 每组的 高1和低1 交换
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)

	return num
}
