// @Author: Ciusyan 2/21/24

package other

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseJson2Pairs(t *testing.T) {

	testCases := []struct {
		name string

		jsonData string

		want    []map[string]string
		wantErr error
	}{
		{
			name:     "没配置依赖的键值对",
			jsonData: "",
			wantErr:  errors.New("依赖数据配置有误"),
		},
		{
			name:     "需要配置键值对",
			jsonData: "[]",
			wantErr:  errors.New("依赖数据配置有误"),
		},
		{
			name:     "每一个 Item，KV只能有一个",
			jsonData: `["{\"ClickKey\":\"uid\", \"HotKey\":\"cityId\"}"]`,
			wantErr:  errors.New("依赖数据配置有误"),
		},
		{
			name:     "没有配置依赖的ItemKey",
			jsonData: `["{\"ClickKey\":\"\"}"]`,
			wantErr:  errors.New("依赖数据配置有误"),
		},
		{
			name:     "依赖ClickKey、TestKey和NoDataKey",
			jsonData: `["{\"ClickKey\":\"uid\"}", "{\"TestKey\":\"uid\"}", "{\"NoDataKey\":\"uid\"}"]`,
			want: []map[string]string{
				{"ClickKey": "uid"},
				{"TestKey": "uid"},
				{"NoDataKey": "uid"},
			},
		},
		{
			name:     "只依赖ClickKey",
			jsonData: `["{\"ClickKey\":\"uid\"}"]`,
			want: []map[string]string{
				{"ClickKey": "uid"},
			},
		},
		{
			name:     "依赖ClickKey和HotKey",
			jsonData: `["{\"ClickKey\":\"uid\"}", "{\"HotKey\":\"cityId\"}"]`,
			want: []map[string]string{
				{"ClickKey": "uid"},
				{"HotKey": "cityid"},
			},
		},
		{
			name:     "依赖HotKey和ClickKey",
			jsonData: `["{\"HotKey\":\"cityId\"}", "{\"ClickKey\":\"uid\"}"]`,
			want: []map[string]string{
				{"HotKey": "cityid"},
				{"ClickKey": "uid"},
			},
		},
	}

	equals := func(want, got []map[string]string) (string, bool) {
		if len(want) != len(got) {
			return fmt.Sprintf("want 长度不一致 want: %d, got %d", len(want), len(got)), false
		}

		for i, wantMap := range want {
			if len(want) != len(got) {
				return fmt.Sprintf("wantMap 长度不一致 want: %d, got %d", len(want), len(got)), false
			}
			for wK, wV := range wantMap {
				gV, ok := got[i][wK]
				if !ok {
					return fmt.Sprintf("未发现 Key：%s", wK), false
				}

				if wV != gV {
					return fmt.Sprintf("Value 不匹配，want: %s, got: %s", wV, gV), false
				}
			}
		}
		return "", true
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pairs, err := parseJson2Pairs(tc.jsonData)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			msg, ok := equals(tc.want, pairs)
			assert.True(t, ok, msg)
		})
	}

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
