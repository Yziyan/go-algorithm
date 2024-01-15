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

func TestSolveSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
}

func TestCountAndSay(t *testing.T) {
	say := countAndSay(4)
	t.Log(say)
}
