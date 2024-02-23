// @Author: Ciusyan 1/30/24

package other

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"reflect"
	"strings"
	"testing"
	// 导入其他需要的包
)

func TestOther(t *testing.T) {
	t.Run("小写", func(t *testing.T) {
		t.Log(strings.Replace(strings.ToLower("userId"), "_", "", -1))
		t.Log(strings.Replace(strings.ToLower("userID"), "_", "", -1))
		t.Log(strings.Replace(strings.ToLower("user_id"), "_", "", -1))
		t.Log(strings.Replace(strings.ToLower("user_ID"), "_", "", -1))
		t.Log(strings.Replace(strings.ToLower("user_id_"), "_", "", -1))
		t.Log(strings.Replace(strings.ToLower("USERID"), "_", "", -1))
		t.Log(strings.Replace(strings.ToLower("USER_ID"), "_", "", -1))
	})

	t.Run("TestUnmarshal", func(t *testing.T) {
		// JSON 数据
		jsonData := `[
        "{ \"clickKey\": \"uid\"}",
        "{ \"topKey\": \"cityId\"}"
    ]`

		// 将 JSON 数组解析为字符串切片
		var jsonStrs []string
		err := json.Unmarshal([]byte(jsonData), &jsonStrs)
		require.NoError(t, err)

		var topKeys []map[string]string
		for _, jsonStr := range jsonStrs {
			var keyMap map[string]string
			err = json.Unmarshal([]byte(jsonStr), &keyMap)
			require.NoError(t, err)
			topKeys = append(topKeys, keyMap)
		}
		// 遍历字符串切片，解析每个 JSON 对象
		for _, topPair := range topKeys {
			for key, value := range topPair {
				t.Logf("Key: %s, Value: %s\n", key, value)
			}
		}
	})
}

// 测试用例结构
type scatterTestCase struct {
	name    string
	module  *scatter
	reqList []string
	want    []string
}

const (
	Separator = "_"
)

// extractFirstElement 从给定的字符串中切割出第一个元素
func extractFirstElement(item string, separator string) string {
	parts := strings.Split(item, separator)
	if len(parts) > 0 {
		return parts[0] // 返回第一个元素
	}
	return "" // 如果没有元素，返回空字符串
}

// 测试 DoScatter 函数
func TestScatter_DoScatter(t *testing.T) {

	cache := make(map[string]string, 10)
	var brandScatterCompare ScatterCompare = func(item1, item2 string) int {
		brand1, ok := cache[item1]
		if !ok {
			// 说明 item1 还没被切割过
			brand1 = extractFirstElement(item1, Separator)
			cache[item1] = brand1
		}

		brand2, ok := cache[item2]
		if !ok {
			// 说明 item2 还没被切割过
			brand2 = extractFirstElement(item2, Separator)
			cache[item2] = brand2
		}

		if brand1 == brand2 {
			// 说明商标相同
			return 0
		}
		return -1
	}

	tests := []scatterTestCase{
		{
			name: "scatter with limited topScatterNum",
			module: NewScatter(
				brandScatterCompare,
				WithMaxConsecutive(2),
				WithStepLength(1),
				WithTopScatterNum(3),
			),
			reqList: []string{
				"Audi_1",
				"Audi_2",
				"Audi_3",
				"BMW_4",
				"Audi_5",
			},
			want: []string{
				"Audi_1",
				"Audi_2",
				"BMW_4",
				"Audi_3",
				"Audi_5",
			},
		},
		{
			name: "scatter with full list",
			module: NewScatter(
				brandScatterCompare,
				WithMaxConsecutive(2),
				WithStepLength(1),
				WithTopScatterNum(5),
			),
			reqList: []string{
				"Audi_1",
				"Audi_2",
				"Audi_3",
				"BMW_4",
				"BMW_5",
				"Audi_6",
			},
			want: []string{
				"Audi_1",
				"Audi_2",
				"BMW_4",
				"Audi_3",
				"BMW_5",
				"Audi_6",
			},
		},
		{
			name: "scatter with full list",
			module: NewScatter(
				brandScatterCompare,
				WithMaxConsecutive(2),
				WithStepLength(2),
				WithTopScatterNum(-1),
			),
			reqList: []string{
				"Audi_1",
				"Audi_2",
				"Audi_3",
				"BMW_4",
				"BMW_5",
				"Ciu_6",
				"BMW_7",
				"BMW_8",
				"BMW_9",
				"Audi_10",
				"Audi_11",
				"Ciu_12",
				"Ciu_13",
				"Audi_14",
				"Audi_15",
				"Ciu_16",
				"Ciu_17",
			},
			want: []string{
				"Audi_1",
				"Audi_2",
				"BMW_4",
				"BMW_5",
				"Audi_3",
				"Ciu_6",
				"BMW_7",
				"BMW_8",
				"Audi_10",
				"Audi_11",
				"BMW_9",
				"Ciu_12",
				"Ciu_13",
				"Audi_14",
				"Audi_15",
				"Ciu_16",
				"Ciu_17",
			},
		},
		{
			name: "scatter with full list",
			module: NewScatter(
				brandScatterCompare,
				WithMaxConsecutive(2),
				WithStepLength(2),
				WithTopScatterNum(-1),
			),
			reqList: []string{
				"Audi_1",
				"Audi_2",
				"Audi_3",
				"Audi_4",
				"Audi_5",
				"BMW_6",
				"BMW_7",
				"Ciu_8",
				"Ciu_9",
				"Ciu_10",
				"BMW_11",
				"BMW_12",
				"BMW_13",
				"BMW_14",
				"Audi_15",
				"Audi_16",
				"Ciu_17",
				"Ciu_18",
				"Audi_19",
				"Audi_20",
				"Ciu_21",
				"Ciu_22",
			},
			want: []string{
				"Audi_1",
				"Audi_2",
				"BMW_6",
				"BMW_7",
				"Audi_5",
				"Audi_3",
				"Ciu_8",
				"Ciu_9",
				"Audi_4",
				"Ciu_10",
				"BMW_11",
				"BMW_12",
				"Audi_15",
				"Audi_16",
				"BMW_13",
				"BMW_14",
				"Ciu_17",
				"Ciu_18",
				"Audi_19",
				"Audi_20",
				"Ciu_21",
				"Ciu_22",
			},
		},
		// 添加更多测试用例以覆盖不同的场景
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got := tt.module.DoScatter(ctx, tt.reqList)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoScatter() got = %v, want %v", got, tt.want)
			}

			// 还原 Cache 的现场
			cache = make(map[string]string, 10)
		})
	}
}
