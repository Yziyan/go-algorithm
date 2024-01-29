// @Author: Ciusyan 1/23/24

package day_32

import (
	"math"
	"testing"
)

func TestName(t *testing.T) {
	s := "*"

	t.Log(len(s))
	t.Log(string(s[0]))
	t.Log(math.MaxInt)

	t.Log(!true)
	t.Log(!false)

	t.Logf("256TB = %dMB", 256*1024*1024)
	t.Logf("2^22 * 64MB = %dMB", int(math.Pow(2, 22)*64))
}
