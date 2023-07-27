// @Author: Ciusyan 2023/7/27

package dfs

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {

	s := []int{1, 2, 3}
	res := Permute(s)
	for _, v := range res {
		fmt.Println(v)
	}

}
