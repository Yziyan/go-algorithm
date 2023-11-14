// @Author: Ciusyan 11/12/23

package day_25

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWays(t *testing.T) {

	type args struct{ n, start, aim, k int }
	testCases := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{n: 4, start: 2, aim: 3, k: 3},
			want: 3,
		},
		{
			name: "case2",
			args: args{n: 10, start: 5, aim: 7, k: 4},
			want: 4,
			// 1 2 3 4 5 6 7 8 9 10
		},
		{
			name: "case3",
			args: args{n: 10, start: 5, aim: 7, k: 5},
			want: 0,
			// 1 2 3 4 5 6 7 8 9 10
		},
		{
			name: "case4",
			args: args{n: 8, start: 2, aim: 4, k: 6},
			want: 14,
		},
		{
			name: "case5",
			args: args{n: 6, start: 2, aim: 5, k: 5},
			want: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Now()
			got1 := ways1(tc.args.n, tc.args.start, tc.args.aim, tc.args.k)
			assert.Equal(t, tc.want, got1)
			t.Logf("暴力递归耗时：%s", time.Since(start))

			start = time.Now()
			got2 := ways2(tc.args.n, tc.args.start, tc.args.aim, tc.args.k)
			assert.Equal(t, tc.want, got2)
			t.Logf("傻缓存耗时：%s", time.Since(start))

			start = time.Now()
			got3 := ways3(tc.args.n, tc.args.start, tc.args.aim, tc.args.k)
			assert.Equal(t, tc.want, got3)
			t.Logf("动态规划耗时：%s", time.Since(start))
		})
	}

}

func TestCardsWin(t *testing.T) {

	testCases := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "case1",
			args: []int{5, 7, 4, 5, 8, 1, 6, 0, 3, 4, 6, 1, 7},
			want: 32,
		},
		{
			name: "case2",
			args: []int{1, 2, 3, 4},
			want: 6,
		},
		{
			name: "case3",
			args: []int{1, 100, 2, 50},
			want: 150,
		},
		{
			name: "case4",
			args: []int{10, 20, 30, 40, 50, 60},
			want: 120,
		},
		{
			name: "case5",
			args: []int{1, 100, 2},
			want: 100,
		},
		{
			name: "case6",
			args: []int{3, 5, 1, 2},
			want: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Now()
			got1 := cardsWin1(tc.args)
			assert.Equal(t, tc.want, got1)
			t.Logf("暴力递归耗时: %s", time.Since(start))
			start = time.Now()

			got2 := cardsWin2(tc.args)
			assert.Equal(t, tc.want, got2)
			t.Logf("傻缓存法耗时: %s", time.Since(start))
			start = time.Now()

			got3 := cardsWin3(tc.args)
			assert.Equal(t, tc.want, got3)
			t.Logf("动态规划耗时: %s", time.Since(start))
		})
	}
}
