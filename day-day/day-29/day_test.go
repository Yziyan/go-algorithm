// @Author: Ciusyan 12/27/23

package day_29

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrixMul(t *testing.T) {

	mul := matrixMul([][]int{{1, 1}}, [][]int{{1, 1}, {1, 0}})
	mul2 := matrixMul([][]int{{1, 1}, {1, 0}}, [][]int{{1}, {1}})
	t.Log(mul)
	t.Log(mul2)
}

func TestCowProblem(t *testing.T) {
	for i := 0; i < 20; i++ {
		got := cowProblem(i)
		got1 := cowProblem1(i)
		got2 := cowProblem2(i)
		assert.Equal(t, got, got1)
		assert.Equal(t, got, got2)
	}

	n := 40
	got := cowProblem(n)
	got1 := cowProblem1(n)
	got2 := cowProblem2(n)
	t.Log(got)
	t.Log(got1)
	t.Log(got2)
}

func TestStr01Num(t *testing.T) {
	for i := 0; i < 20; i++ {
		got1 := str01Num1(i)
		got2 := str01Num2(i)
		assert.Equal(t, got1, got2)
	}

	n := 40
	got1 := str01Num1(n)
	got2 := str01Num2(n)
	t.Log(got1)
	t.Log(got2)
}
