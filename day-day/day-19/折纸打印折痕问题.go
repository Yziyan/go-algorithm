// @Author: Ciusyan 10/12/23

package day_19

import "fmt"

// 折纸问题
// 请把一段纸条竖着放在桌子上，然后从纸条的下边向上方对折1次，压出折痕后展开
// 此时折痕是凹下去的，即折痕突起的方向指向纸条的背面
// 如果从纸条的下边向上方连续对折2次，压出折痕后展开
// 此时有三条折痕，从上到下依次是下折痕、下折痕和上折痕。
// 给定一个输入参数N，代表纸条都从下边向上方连续对折N次
// 请从上到下打印所有折痕的方向。
// N=1时，打印: down
// N=2时，打印: down down up

func PrintAllFolds(n int) {
	if n <= 0 {
		fmt.Println("请传入正确的折叠次数")
		return
	}

	inorderFold(1, n, true)
}

//	down
//
// 中序遍历的方式打印
// num 折叠次数，down 是否是下折痕
func inorderFold(i, n int, down bool) {
	if i > n {
		return
	}

	// 中序遍历打印左边折痕
	inorderFold(i+1, n, false)

	if down {
		fmt.Println("凹")
	} else {
		fmt.Println("凸")
	}

	// 中序遍历打印右边折痕
	inorderFold(i+1, n, true)
}
