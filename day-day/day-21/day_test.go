// @Author: Ciusyan 10/25/23

package day_21

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"slices"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	idx := 3
	strs := []int{44, 22, 133, 765, 43}
	strs = append(strs[0:idx], strs[idx+1:]...)
	fmt.Println(strs)
}

func TestLowestString(t *testing.T) {

	randString := func(strLen int) string {
		const charset = "abcdefghijklmnopqrstuvwxyz"
		b := make([]byte, strLen)
		for i := range b {
			b[i] = charset[rand.Intn(len(charset))]
		}
		return string(b)
	}

	randStringSlice := func(size, strLen int) []string {
		rand.New(rand.NewSource(time.Now().UnixNano()))

		res := make([]string, size)
		for i := 0; i < size; i++ {
			res[i] = randString(strLen)
		}
		return res
	}

	loopCount := 50
	arrLen := 6
	strLen := 4
	for i := 0; i < loopCount; i++ {
		strs := randStringSlice(arrLen, strLen)
		t.Log(strs)
		str2 := slices.Clone(strs)
		got1 := LowestString(strs)
		got2 := LowestString2(str2)
		assert.Equal(t, got1, got2)
		t.Log(got1)
		t.Log(got2)
		fmt.Println()
	}

	strs := []string{"b", "ba"}
	got1 := LowestString(strs)

	strs = []string{"b", "ba"}
	got2 := LowestString2(strs)

	t.Log(got1)
	t.Log(got2)
	assert.Equal(t, got1, got2)

}
