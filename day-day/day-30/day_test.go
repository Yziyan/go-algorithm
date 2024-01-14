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

func TestIsValid(t *testing.T) {

	valid := isValid("()")
	t.Log(valid)
}

func TestSearchRange(t *testing.T) {
	ints := searchRange([]int{2, 2}, 2)
	t.Log(ints)
}