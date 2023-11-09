// @Author: Ciusyan 11/9/23

package other

import (
	"bytes"
	"encoding/gob"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	type user struct {
		Name string
		Age  int
		sex  int
	}

	gob.Register([]interface{}{})
	gob.Register([]user{})
	gob.Register(user{})
	gob.Register([3]string{})
	gob.Register(map[string]interface{}{})
	gob.Register(map[string]interface{}{})

	src := map[string]interface{}{
		"user": user{
			Name: "Ciusyan",
			Age:  20,
			sex:  1,
		},
		"int":    23,
		"string": "志颜",
		"float":  3.92,
		"slice":  []int{2, 3, 4},
		"user_slice": []user{{
			Name: "Ciusyan",
			Age:  20,
			sex:  1,
		}, {
			Name: "Zhiyan",
			Age:  22,
			sex:  1,
		}},
		"array": [3]string{"11", "22", "33"},
	}
	dst := make(map[string]interface{})
	err := DeepCopy(&dst, src)
	require.NoError(t, err)
	dst["user_slice"] = nil
	t.Log(dst)
	t.Log(src)
}

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
