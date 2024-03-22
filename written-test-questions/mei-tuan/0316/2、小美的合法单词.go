// @Author: Ciusyan 3/21/24

package _316

func legalWord(word string) int {

	// 统计 word 的大小写
	// 大写有多少个，小写有多少个
	chars := []byte(word)

	daxieNum := 0
	xiaoxieNum := 0

	for _, c := range chars {
		if c >= 'a' && c <= 'z' {
			xiaoxieNum++
		} else {
			daxieNum++
		}
	}

	if daxieNum == 1 && chars[0] >= 'A' && chars[0] <= 'Z' {
		return 0
	} else if xiaoxieNum == 0 || daxieNum == 0 {
		return 0
	}

	// 上面的都是合法情况，到达这里，说明都是不合法情况，尽量少，就是谁少，返回谁
	if xiaoxieNum > daxieNum {
		if chars[0] >= 'A' && chars[0] <= 'Z' {
			return daxieNum - 1
		}
		return daxieNum
	} else {
		return xiaoxieNum
	}
}
