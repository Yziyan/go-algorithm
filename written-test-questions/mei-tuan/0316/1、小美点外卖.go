// @Author: Ciusyan 3/21/24

package _316

func orderTakeout(num int, price []int, x, y int) int {
	if num != len(price) {
		return -1
	}

	res := 0
	for _, p := range price {
		res += p
	}

	res -= x

	if res > y {
		res -= y
	} else {
		res = 0
	}

	return res
}
