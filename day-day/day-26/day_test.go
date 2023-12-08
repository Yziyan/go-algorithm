// @Author: Ciusyan 12/4/23

package day_26

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestLivePossibility(t *testing.T) {
	tests := []struct {
		name string
		row  int
		col  int
		k    int
		n    int
		m    int
		want float64
	}{
		{
			name: "Test1",
			row:  1,
			col:  1,
			k:    2,
			n:    3,
			m:    3,
			want: 0.75,
		},
		{
			name: "Test2",
			row:  0,
			col:  0,
			k:    3,
			n:    2,
			m:    2,
			want: 0.125,
		},
		{
			name: "Test3",
			row:  2,
			col:  2,
			k:    4,
			n:    5,
			m:    5,
			want: 0.84375,
		},
		// 可以添加更多测试用例...
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := livePossibility(tc.row, tc.col, tc.k, tc.n, tc.m)
			assert.Equal(t, tc.want, got)

			got1 := livePossibility1(tc.row, tc.col, tc.k, tc.n, tc.m)
			assert.Equal(t, tc.want, got1)
		})
	}
}

func TestKillMonster(t *testing.T) {
	tests := []struct {
		name string
		n    int
		m    int
		k    int
		want float64
	}{
		{
			name: "Test1",
			n:    5,
			m:    2,
			k:    3,
			want: 0.14814814814814814,
		},
		{
			name: "Test2",
			n:    10,
			m:    3,
			k:    5,
			want: 0.216796875,
		},
		{
			name: "Test3",
			n:    20,
			m:    5,
			k:    7,
			want: 0.3321616369455876,
		},
		{
			name: "Test4",
			n:    7,
			m:    4,
			k:    2,
			want: 0.12,
		},
		// 可以添加更多测试用例...
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := KillMonster(tc.n, tc.m, tc.k)
			assert.Equal(t, tc.want, got)

			got1 := KillMonster1(tc.n, tc.m, tc.k)
			assert.Equal(t, tc.want, got1)

			got2 := KillMonster2(tc.n, tc.m, tc.k)
			assert.Equal(t, tc.want, got2)
		})
	}
}

func TestMinCoinsNoLimit(t *testing.T) {
	testCases := []struct {
		name  string
		coins []int
		aim   int
		want  int
	}{
		{"Test1", []int{1, 2, 5}, 11, 3},
		{"Test2", []int{2, 3, 5}, 10, 2},
		{"Test3", []int{2, 3, 7}, 14, 2},
		{"Test4", []int{3, 5, 7}, 17, 3},
		{"Test5", []int{1, 3, 4}, 6, 2},
		{"Test6", []int{2, 4}, 7, math.MaxInt},

		{name: "ComplexTest1", coins: []int{1, 4, 6, 8}, aim: 15, want: 3},
		{name: "ComplexTest2", coins: []int{5, 7, 8, 9}, aim: 22, want: 3},
		{name: "ComplexTest3", coins: []int{1, 3, 7, 10}, aim: 14, want: 2},
		{name: "ComplexTest4", coins: []int{2, 6, 9, 12}, aim: 24, want: 2},
		{name: "ComplexTest5", coins: []int{1, 5, 6, 8}, aim: 20, want: 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := minCoinsNoLimit(tc.coins, tc.aim)
			assert.Equal(t, tc.want, got)

			got1 := minCoinsNoLimit1(tc.coins, tc.aim)
			assert.Equal(t, tc.want, got1)

			got2 := minCoinsNoLimit2(tc.coins, tc.aim)
			assert.Equal(t, tc.want, got2)
		})
	}
}

func TestSplitNumber(t *testing.T) {
	testCases := []struct {
		name string
		n    int

		want int
	}{
		{"Case1", 1, 1},
		{"Case2", 2, 2},
		{"Case3", 3, 3},
		{"Case4", 4, 5},
		{"Case5", 5, 7},
		{"Case6", 6, 11},
		{"Case7", 7, 15},
		{"Case8", 8, 22},
		{"Case9", 9, 30},
		{"Case10", 10, 42},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := splitNumber(tc.n)
			assert.Equal(t, tc.want, got)

			got2 := splitNumber2(tc.n)
			assert.Equal(t, tc.want, got2)

			got1 := splitNumber1(tc.n)
			assert.Equal(t, tc.want, got1)
		})
	}
}

func TestSplitSumClosed(t *testing.T) {
	var tests = []struct {
		name string
		arr  []int
		want int
	}{
		{name: "Large Consecutive", arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, want: 27},
		{name: "Odd Numbers", arr: []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, want: 50},
		{name: "Mixed Positives and Negatives", arr: []int{-5, -4, -3, -2, -1, 1, 2, 3, 4, 5}, want: 0},
		{name: "Uniform Large", arr: []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, want: 50},
		{name: "Very High Values", arr: []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400, 1500, 1600, 1700, 1800, 1900, 2000}, want: 10500},
		{name: "Small Consecutive", arr: []int{1, 2, 3, 4}, want: 5},
		{name: "Uniform Small", arr: []int{2, 2, 2, 2}, want: 4},
		{name: "Three Tens", arr: []int{10, 10, 10}, want: 10},
		{name: "Mixed Values", arr: []int{1, 6, 5, 11}, want: 11},
		{name: "Empty Array", arr: []int{}, want: 0},
		{name: "Nil Array", arr: nil, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitSumClosed(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
