// @Author: Ciusyan 11/23/23

package day_25

/**
可搜索或者想象一个象棋的棋盘，大小为：10*9
然后把整个棋盘放入第一象限，棋盘的最左下角是 (0,0) 位置
那么整个棋盘就是横坐标上 9 条线、纵坐标上 10 条线的区域
给你三个 参数 x，y，k
返回“马”从 (0,0) 位置出发，必须走 k 步，注意：马只能走日
最后落在 (x,y) 上的方法数有多少种?
*/

// 动态规划方法
func jump(x, y int, k int) int {
	if x < 0 || x > 9 || y < 0 || y > 8 || k < 0 {
		// 说明是在棋盘外的点，永远走不到
		return 0
	}

	// 可以发现，这个题的可变参数是三个，可能需要三维的动态规划
	// curX, curY, remain，他们的范围分别是：curX ∈ [0, 9], curY ∈ [0, 8], remain ∈ [0, k]
	// dp[curX][curY][remain] 代表，当前处于 (curX, curY)， 还剩余 remain 步要走，到达 (x, y) 位置的方法数
	dp := make([][][]int, 10)
	// 先准备好 dp
	for xi := range dp {
		dp[xi] = make([][]int, 9)
		for yi := range dp[xi] {
			dp[xi][yi] = make([]int, k+1)
		}
	}
	// 填充 dp
	// 根据递归基可知，remain == 0 时，(curX, curY) 是 (x, y) 时有一种方案，其余的在这第一层，其实都为 0
	dp[x][y][0] = 1

	// pick 下，因为可能会越界
	picker := func(dp [][][]int, curX, curY, remain int) int {
		if curX < 0 || curX > 9 || curY < 0 || curY > 8 {
			return 0
		}

		return dp[curX][curY][remain]
	}

	// 要不然对于一般情况，根据依赖情况，应该是上层依赖下层，所以应该从下往上填
	for remain := 1; remain <= k; remain++ {
		for curX := 0; curX < 10; curX++ {
			for curY := 0; curY < 9; curY++ {
				// 有 8 个方向可以走，那么总方法数就是八个方向的总和
				ways := picker(dp, curX+2, curY+1, remain-1)
				ways += picker(dp, curX+1, curY+2, remain-1)
				ways += picker(dp, curX-1, curY+2, remain-1)
				ways += picker(dp, curX-2, curY+1, remain-1)
				ways += picker(dp, curX-2, curY-1, remain-1)
				ways += picker(dp, curX-1, curY-2, remain-1)
				ways += picker(dp, curX+1, curY-2, remain-1)
				ways += picker(dp, curX+2, curY-1, remain-1)

				dp[curX][curY][remain] = ways
			}
		}
	}

	// 当前处于 (0, 0) 位置，剩余 k 步需要走，走到 (x, y) 的方法数
	return dp[0][0][k]
}

// 暴力递归方法：目标位置：(x, y)，剩余 k 步
func jump1(x, y int, k int) int {
	if x < 0 || x > 9 || y < 0 || y > 8 || k < 0 {
		// 说明是在棋盘外的点，永远走不到
		return 0
	}

	// 憋一个暴力递归，
	// 当前处于 (curX, curY) 位置，还剩余 remain 步可以走，能够到达 (x, y) 位置的方法数
	var process func(curX, curY int, x, y int, remain int) int
	process = func(curX, curY int, x, y int, remain int) int {
		if curX < 0 || curX > 9 || curY < 0 || curY > 8 {
			// 说明跳到棋盘外了
			return 0
		}

		if remain == 0 {
			// 没有步数了，看看是否已经出现在 (x, y) 位置了
			if curX == x && curY == y {
				// 说明当前刚好停在 (c, y)
				return 1
			}

			// 否则说明根本没有停在这
			return 0
		}

		// 一般情况的话，应该有八个方向可以走，能走到目标的方法数，就应该是八个方向相加
		ways := process(curX+2, curY+1, x, y, remain-1)
		ways += process(curX+1, curY+2, x, y, remain-1)
		ways += process(curX-1, curY+2, x, y, remain-1)
		ways += process(curX-2, curY+1, x, y, remain-1)
		ways += process(curX-2, curY-1, x, y, remain-1)
		ways += process(curX-1, curY-2, x, y, remain-1)
		ways += process(curX+1, curY-2, x, y, remain-1)
		ways += process(curX+2, curY-1, x, y, remain-1)

		return ways
	}

	// 当前处于 (0, 0) 位置，还剩余 k 步可以走，能够到达 (x, y) 位置的方法数
	return process(0, 0, x, y, k)
}
