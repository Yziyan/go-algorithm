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
		got := str01Num(i)
		got1 := str01Num1(i)
		got2 := str01Num2(i)
		assert.Equal(t, got, got1)
		assert.Equal(t, got1, got2)
	}

	n := 40
	got := str01Num(n)
	got1 := str01Num1(n)
	got2 := str01Num2(n)
	t.Log(got)
	t.Log(got1)
	t.Log(got2)
}

func TestFullBlock(t *testing.T) {
	for i := 0; i < 20; i++ {
		got := fullBlock(i)
		got1 := fullBlock1(i)
		assert.Equal(t, got, got1)
	}

	n := 40
	got := fullBlock(n)
	got1 := fullBlock1(n)
	t.Log(got)
	t.Log(got1)
}

func TestMinMoney(t *testing.T) {
	tests := []struct {
		name      string
		abilities []int
		coins     []int
		want      int
	}{
		{
			name:      "case1",
			abilities: []int{1, 2, 3},
			coins:     []int{10, 20, 30},
			want:      30,
		},
		{
			name:      "case2",
			abilities: []int{4, 5, 6},
			coins:     []int{40, 50, 60},
			want:      90,
		},
		{
			name:      "case3",
			abilities: []int{2, 3, 5, 8},
			coins:     []int{5, 10, 20, 40},
			want:      35,
		},
		{
			name:      "case4",
			abilities: []int{10, 20, 30},
			coins:     []int{15, 25, 35},
			want:      40,
		},
		{
			name:      "case5",
			abilities: []int{1, 3, 4, 6},
			coins:     []int{5, 15, 20, 30},
			want:      40,
		},
		{
			name:      "case6",
			abilities: []int{5, 10, 15, 20},
			coins:     []int{10, 20, 30, 40},
			want:      60,
		},
		{
			name:      "case7",
			abilities: []int{3, 6, 9, 12, 15},
			coins:     []int{10, 20, 30, 40, 50},
			want:      60,
		},
		// Add more complex cases if needed
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := minMoney(tc.abilities, tc.coins)
			got2 := minMoneyDp(tc.abilities, tc.coins)
			got3 := minMoney1(tc.abilities, tc.coins)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.want, got2)
			assert.Equal(t, tc.want, got3)
		})
	}
}

func TestSubArrModM(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		m    int
		want int
	}{
		{
			name: "Test 1",
			arr:  []int{3, 1, 4, 2},
			m:    6,
			want: 5,
		},
		{
			name: "Test 2",
			arr:  []int{1, 2, 3},
			m:    4,
			want: 3,
		},
		{
			name: "Test 3",
			arr:  []int{1, 2, 3, 4},
			m:    5,
			want: 4,
		},
		{
			name: "Test 4",
			arr:  []int{10, 12, 7},
			m:    8,
			want: 7,
		},
		{
			name: "Test 5",
			arr:  []int{6, 7, 9},
			m:    5,
			want: 4,
		},
		{
			name: "Large Array Test 1",
			arr:  []int{12, 3, 14, 56, 77, 13, 4, 25, 33, 45, 67, 89, 2, 6, 8},
			m:    10,
			want: 9,
		},
		{
			name: "Large Array Test 2",
			arr:  []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89},
			m:    15,
			want: 14,
		},
		{
			name: "Large Array with Small m",
			arr:  []int{7, 11, 5, 2, 8, 13, 21, 1, 3},
			m:    3,
			want: 2,
		},
		{
			name: "Large Values",
			arr:  []int{102, 304, 506, 708, 910, 111, 313, 515},
			m:    50,
			want: 49,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := subArrModM(tc.arr, tc.m)
			assert.Equal(t, tc.want, got)
		})
	}

}
