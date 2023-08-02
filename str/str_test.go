// @Author: Ciusyan 2023/8/2

package str

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	string1 := "189790"
	string2 := "987210"

	resultString := multiply(string1, string2)
	fmt.Println(resultString) // 输出："1219326311370217958215609403036575474090"
}
