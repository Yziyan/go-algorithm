// @Author: Ciusyan 3/11/24

package day_34

func titleToNumber(columnTitle string) int {
	if columnTitle == "" {
		return 0
	}

	res := 0
	for _, c := range columnTitle {
		res = res*26 + int((c-'A')+1)
	}

	return res
}
