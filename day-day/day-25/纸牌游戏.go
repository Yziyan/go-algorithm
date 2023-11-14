// @Author: Ciusyan 11/14/23

package day_25

/**
给定一个整型数组 arr，代表数值不同的纸牌排成一条线
玩家A 和 玩家B 依次拿走每张纸牌
规定 玩家A 先拿，玩家B 后拿
但是每个玩家每次只能拿走最左或最右的纸牌
玩家A 和 玩家B 都绝顶聪明
请返回最后获胜者的分数
*/

// 暴力递归方法
func cardsWin1(cards []int) int {

	// 从 cards[l ... r] 先拿的牌获得的最大分数
	var pre func(cards []int, l, r int) int
	// 从 cards[l ... r] 后拿牌获得的最大分数
	var post func(cards []int, l, r int) int

	pre = func(cards []int, l, r int) int {
		if l == r {
			// 说明只有一张牌了，又是先手，那么肯定能拿到这张牌
			return cards[l]
		}

		// 拿左边的牌，总共得到的分数就是 左边 + 从 l+1 ~ r 后拿牌得到的分数
		p1 := cards[l] + post(cards, l+1, r)
		// 拿右边的牌，总共得到的分数就是 右边 + 从 l ~ r-1 后拿牌得到的分数
		p2 := cards[r] + post(cards, l, r-1)

		// 因为是先拿，结果就是两种方案种的最好方案
		return max(p1, p2)
	}

	post = func(cards []int, l, r int) int {
		if l == r {
			// 说明只有一张牌了，又是后手，那么肯定拿不到这张牌
			return 0
		}

		// 对手先拿走了 cards[l] 的牌，该我从 l+1 ~ r 先拿牌了
		p1 := pre(cards, l+1, r)
		// 对手拿走了 cards[r] 的牌，该我从 l ~ r-1 先拿牌了
		p2 := pre(cards, l, r-1)

		// 因为作为后拿牌的人，没得选，只能拿到两种方案的最小值
		return min(p1, p2)
	}

	last := len(cards) - 1

	// 那么最终返回的最大分数，只能是 先拿 和 后拿 中的最大值
	return max(pre(cards, 0, last), post(cards, 0, last))
}
