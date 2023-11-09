// @Author: Ciusyan 11/9/23

package day_24

import "fmt"

// 打印 N 层汉诺塔，从最左移动到最优的全部过程，要求
// 1.每次只能挪动一个盘子
// 2.大盘子只能放在小盘子下面

func Hanoi(n int) {
	// 我们最终就是想把所有盘子都从左 -> 右
	left2right(n)
}

// 将 n 个盘子，左 -> 右
func left2right(n int) {
	if n == 1 {
		// 如果只有一个盘子，直接挪动
		fmt.Printf("Move %d from left to right\n", n)
		return
	}

	// 否则先将 n-1 个盘子挪到中间
	left2mid(n - 1)
	// 然后将第 n 个盘子挪动到右边
	fmt.Printf("Move %d from left to right\n", n)
	// 再将中间的 n-1 个盘子挪到右边
	mid2right(n - 1)
}

// 将 n 个盘子，左 -> 中
func left2mid(n int) {
	if n == 1 {
		// 如果只有一个盘子，直接挪动
		fmt.Printf("Move %d from left to mid\n", n)
		return
	}

	// 否则先将 n-1 个盘子挪到右边
	left2right(n - 1)
	// 再将第 n 个盘子挪动到中间
	fmt.Printf("Move %d from left to mid\n", n)
	// 最后将 n-1 个盘子挪到中间
	right2mid(n - 1)
}

// 将 n 个盘子，中 -> 右
func mid2right(n int) {
	if n == 1 {
		// 如果只有一个盘子，直接挪动
		fmt.Printf("Move %d from mid to right\n", n)
		return
	}

	mid2left(n - 1)
	fmt.Printf("Move %d from mid to right\n", n)
	left2right(n - 1)
}

// 将 n 个盘子，右 -> 中
func right2mid(n int) {
	if n == 1 {
		// 如果只有一个盘子，直接挪动
		fmt.Printf("Move %d from right to mid\n", n)
		return
	}

	right2left(n - 1)
	fmt.Printf("Move %d from right to mid\n", n)
	left2mid(n - 1)
}

// 将 n 个盘子，中 -> 左
func mid2left(n int) {
	if n == 1 {
		// 如果只有一个盘子，直接挪动
		fmt.Printf("Move %d from mid to left\n", n)
		return
	}

	mid2right(n - 1)
	fmt.Printf("Move %d from mid to left\n", n)
	right2left(n - 1)
}

// 将 n 个盘子，右 -> 左
func right2left(n int) {
	if n == 1 {
		// 如果只有一个盘子，直接挪动
		fmt.Printf("Move %d from right to left\n", n)
		return
	}

	right2mid(n - 1)
	fmt.Printf("Move %d from right to left\n", n)
	mid2left(n - 1)
}
