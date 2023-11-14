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

// 傻缓存法
func cardsWin2(cards []int) int {

	// 准备两个缓存
	n := len(cards)
	// preDp[l][r] 代表：cards 从 l~r 先拿牌获取的最大分数
	preDp := make([][]int, n)
	// postDp[l][r] 代表：cards 从 l~r 后拿牌获取的最大分数
	postDp := make([][]int, n)

	// 初始化缓存
	for i := 0; i < n; i++ {
		preDp[i] = make([]int, n)
		postDp[i] = make([]int, n)

		for j := 0; j < n; j++ {
			preDp[i][j] = -1
			postDp[i][j] = -1
		}
	}

	// 准备两个递归函数，并且将缓存一起传递下去
	var pre func(cards []int, l, r int, preDp, postDp [][]int) int
	var post func(cards []int, l, r int, preDp, postDp [][]int) int
	pre = func(cards []int, l, r int, preDp, postDp [][]int) int {
		if preDp[l][r] != -1 {
			// 说明命中了缓存
			return preDp[l][r]
		}

		res := 0
		if l == r {
			res = cards[l]
		} else {
			// 选的是 cards[l]
			p1 := cards[l] + post(cards, l+1, r, preDp, postDp)
			// 选的是 cards[r]
			p2 := cards[r] + post(cards, l, r-1, preDp, postDp)

			// 两种可能的最大值
			res = max(p1, p2)
		}

		// 返回时设置缓存
		preDp[l][r] = res
		return res
	}

	post = func(cards []int, l, r int, preDp, postDp [][]int) int {
		if postDp[l][r] != -1 {
			// 说明命中了缓存
			return postDp[l][r]
		}
		res := 0

		if l != r {
			// 说明对手选了 cards[l]
			p1 := pre(cards, l+1, r, preDp, postDp)
			// 说明对手选了 cards[r]
			p2 := pre(cards, l, r-1, preDp, postDp)

			// 两种最优解的最小值
			res = min(p1, p2)
		}

		// 返回时设置缓存
		postDp[l][r] = res
		return res
	}

	// 还是 先拿 和 后拿 得到的最大值
	return max(pre(cards, 0, n-1, preDp, postDp), post(cards, 0, n-1, preDp, postDp))
}

// 动态规划方法
func cardsWin3(cards []int) int {
	n := len(cards)
	// 准备两个缓存
	// preDp[L][R] 代表从 L~R 上先拿，获取的最大分数
	preDp := make([][]int, n)
	// postDp[L][R] 代表从 L~R 上后拿，获取的最大分数
	postDp := make([][]int, n)

	for i := 0; i < n; i++ {
		preDp[i] = make([]int, n)
		postDp[i] = make([]int, n)
	}

	for L := 0; L < n; L++ {
		// 当 L == R 时，先拿的可以拿到这张牌的分数，后拿的没有分数
		preDp[L][L] = cards[L]
	}

	// 需要从斜线，依次填写下来
	for startCol := 1; startCol < n; startCol++ {
		L := 0
		R := startCol
		for R < n {
			// 两种情况的最大值
			preDp[L][R] = max(cards[L]+postDp[L+1][R], cards[R]+postDp[L][R-1])
			// 两种最优情况的最小值
			postDp[L][R] = min(preDp[L+1][R], preDp[L][R-1])

			L++
			R++
		}
	}

	// 代表从 L~R 上，分别先拿和后拿，获取的最大分数，返回最大值
	return max(preDp[0][n-1], postDp[0][n-1])
}
