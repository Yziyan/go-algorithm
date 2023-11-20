// @Author: Ciusyan 11/12/23

package day_25

import (
	"github.com/stretchr/testify/assert"
	"sort"
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

func TestMaxValue(t *testing.T) {
	type args struct {
		weights []int
		values  []int
		bag     int
	}
	testCases := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				weights: []int{3, 2, 4, 7},
				values:  []int{5, 6, 3, 19},
				bag:     11,
			},
			want: 25,
		},
		{
			name: "case2",
			args: args{
				weights: []int{1, 2, 3},
				values:  []int{1, 2, 3},
				bag:     4,
			},
			want: 4,
		},
		{
			name: "case3",
			args: args{
				weights: []int{1, 2, 3},
				values:  []int{6, 10, 12},
				bag:     5,
			},
			want: 22,
		},
		{
			name: "case4",
			args: args{
				weights: []int{2, 2, 6, 5, 4},
				values:  []int{6, 3, 5, 4, 6},
				bag:     10,
			},
			want: 15,
		},
		{
			name: "case5",
			args: args{
				weights: []int{3, 5, 1, 2},
				values:  []int{4, 2, 6, 8},
				bag:     7,
			},
			want: 18,
		},
		{
			name: "case5",
			args: args{
				weights: []int{4, 2, 3},
				values:  []int{10, 4, 7},
				bag:     5,
			},
			want: 11,
		},
		{
			name: "Complex Case Large Dataset",
			args: args{
				weights: []int{10, 20, 30, 40, 50, 5, 15, 25, 35, 45, 55, 65, 75, 85, 95, 12, 22, 32, 42, 52, 62, 72, 82, 92, 2, 14, 24, 34, 44, 54, 64, 74, 84, 94, 3, 13, 23, 33, 43, 53, 63, 73, 83, 93},
				values:  []int{35, 25, 55, 45, 20, 50, 65, 75, 85, 95, 15, 25, 35, 45, 55, 65, 75, 85, 95, 5, 15, 25, 35, 45, 55, 65, 75, 85, 95, 10, 30, 50, 70, 90, 11, 22, 33, 44, 55, 66, 77, 88, 99, 100},
				bag:     200,
			},
			want: 741, // The expected value must be calculated based on the algorithm's logic
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Now()
			got1 := maxValue1(tc.args.weights, tc.args.values, tc.args.bag)
			t.Logf("暴力递归耗时:%s", time.Since(start))
			assert.Equal(t, tc.want, got1)

			start = time.Now()
			got2 := maxValue2(tc.args.weights, tc.args.values, tc.args.bag)
			t.Logf("动态规划耗时:%s", time.Since(start))
			assert.Equal(t, tc.want, got2)
		})
	}
}

func TestConvertStrLetter(t *testing.T) {

	testCases := []struct {
		name string
		args string
		want int
	}{
		{name: "case1", args: "111", want: 3},
		{name: "case2", args: "1302", want: 0},
		{name: "case3", args: "1111", want: 5},
		{name: "case4", args: "12345", want: 3},
		{name: "case5", args: "1234567890", want: 0},
		{name: "case6", args: "261812", want: 8},
		{name: "case7", args: "1111111111", want: 89},
		{name: "case8", args: "9999999999", want: 1},
		{name: "case9", args: "12121212", want: 34},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Now()
			got1 := convertStrLetter1(tc.args)
			t.Logf("暴力递归耗时：%s", time.Since(start))
			assert.Equal(t, tc.want, got1)

			start = time.Now()
			got2 := convertStrLetter2(tc.args)
			t.Logf("动态规划耗时：%s", time.Since(start))
			assert.Equal(t, tc.want, got2)
		})
	}

}

func TestMinStickers(t *testing.T) {

	testCases := []struct {
		name     string
		stickers []string
		target   string
		want     int
	}{
		{
			name:     "case1",
			stickers: []string{"with", "example", "science"},
			target:   "thehat",
			want:     3,
		},
		{
			name:     "case2",
			stickers: []string{"notice", "possible"},
			target:   "basicbasic",
			want:     -1,
		},
		{
			name:     "case3",
			stickers: []string{"these", "guess", "about", "garden", "him"},
			target:   "atomher",
			want:     3,
		},
		{
			name:     "case4",
			stickers: []string{"control", "heart", "interest", "stream", "sentence", "soil", "wonder", "them", "month", "slip", "table", "miss", "boat", "speak", "figure", "no", "perhaps", "twenty", "throw", "rich", "capital", "save", "method", "store", "meant", "life", "oil", "string", "song", "food", "am", "who", "fat", "if", "put", "path", "come", "grow", "box", "great", "word", "object", "stead", "common", "fresh", "the", "operate", "where", "road", "mean"},
			target:   "stoodcrease",
			want:     3,
		},
		{
			name:     "case5",
			stickers: []string{"heavy", "claim", "seven", "set", "had", "it", "dead", "jump", "design", "question", "sugar", "dress", "any", "special", "ground", "huge", "use", "busy", "prove", "there", "lone", "window", "trip", "also", "hot", "choose", "tie", "several", "be", "that", "corn", "after", "excite", "insect", "cat", "cook", "glad", "like", "wont", "gray", "especially", "level", "when", "cover", "ocean", "try", "clean", "property", "root", "wing"},
			target:   "travelbell",
			want:     4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := minStickers(tc.stickers, tc.target)
			assert.Equal(t, tc.want, got)

			//got1 := minStickers1(tc.stickers, tc.target)
			//assert.Equal(t, tc.want, got1)
		})
	}

}

func TestOther(t *testing.T) {

	str := "adxasadcb"
	chars := []byte(str)

	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	t.Log(str)
	str = string(chars)
	t.Log(str)
}
