// @Author: Ciusyan 10/31/23

package day_21

/*
给定一个字符串str，只由'X'和'.'两种字符构成
'X'表示墙，不能放灯，也不需要点亮；'.'表示居民点，可以放灯，需要点亮
如果灯放在i位置，可以让i-1，i和i+1三个位置被点亮
返回如果点亮str中所有需要点亮的位置，至少需要几盏灯
*/

// MinLight 道路的最少点灯数
//
// road 只会包含 X . 的字符串，返回最少的点灯数
func MinLight(road string) int {
	if road == "" {
		return 0
	}

	chars := []byte(road)
	idx := 0
	light := 0
	l := len(chars)

	// 遍历所有的字符
	for idx < l {
		if chars[idx] == 'X' {
			// 遇到墙，直接往下走
			idx++
			continue
		}

		// 来到这里说明 idx 是街道，至少需要一盏灯了
		light++

		// 但是看这盏灯放在哪里，才能贪最多
		if idx+1 == l {
			// 说明到最后了，不需要放灯了
			break
		}

		if chars[idx+1] == 'X' {
			// 说明只有一个街道，去墙后面继续看
			idx += 2
		} else {
			// 说明至少有 2 个街道，但是我们们根本不关心有没有第三个街道，
			// 如果有：那灯就放 idx+1，否则随便，所以我们可以直接跳到 idx+3 的位置去判断了
			idx += 3
		}
	}

	return light
}
