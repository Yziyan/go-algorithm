// @Author: Ciusyan 1/2/24

package day_30

import (
	"testing"
)

func TestMyAtoi(t *testing.T) {
	atoi := myAtoi("2147483646")
	t.Log(atoi)

	s := "3415"
	chars := []rune(s)
	for _, c := range chars {

		t.Log('0' - c)
	}
}
