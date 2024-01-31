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

// CarModel 结构体假设
type CarModel struct {
	ID    string
	Brand string
}

// BrandScatterModule 结构体假设
type BrandScatterModule struct {
	// ... 可能的其他配置字段 ...
}

// DoAction 实现
func (m *BrandScatterModule) DoAction(ctx context.Context, reqList []CarModel) []CarModel {
	// 最大连续品牌数
	const maxConsecutive = 2

	// 初始化最终列表
	finalList := make([]CarModel, len(reqList))
	copy(finalList, reqList)

	// 遍历 ReqList，寻找超过 maxConsecutive 限制的品牌
	for i := 0; i < len(finalList); {
		// 检查当前品牌是否超出连续出现限制
		start := i
		end := i
		for end < len(finalList) && finalList[start].Brand == finalList[end].Brand {
			end++
			if end-start > maxConsecutive {
				// 寻找一个不同品牌的元素来交换
				swapIndex := findSwapIndex(finalList, end, finalList[start].Brand)
				if swapIndex != -1 {
					finalList[end-1], finalList[swapIndex] = finalList[swapIndex], finalList[end-1]
					break
				}
			}
		}
		i = end
	}

	return finalList
}

// findSwapIndex 查找可用于交换的不同品牌元素的索引
func findSwapIndex(list []CarModel, start int, brand string) int {
	for i := start; i < len(list); i++ {
		if list[i].Brand != brand {
			return i
		}
	}
	return -1 // 没有找到可交换的元素
}

// 测试用例结构
type testCase struct {
	name    string
	reqList []CarModel
	want    []CarModel
}

// 测试 DoAction 函数
func TestBrandScatterModule_DoAction(t *testing.T) {
	tests := []testCase{
		{
			name: "No consecutive brands",
			reqList: []CarModel{
				{"1", "Toyota"},
				{"2", "Ford"},
				{"3", "BMW"},
			},
			want: []CarModel{
				{"1", "Toyota"},
				{"2", "Ford"},
				{"3", "BMW"},
			},
		},
		{
			name: "Consecutive brands, simple case",
			reqList: []CarModel{
				{"1", "Volkswagen"},
				{"2", "Volkswagen"},
				{"3", "Volkswagen"},
				{"4", "Toyota"},
			},
			want: []CarModel{
				{"1", "Volkswagen"},
				{"2", "Volkswagen"},
				{"4", "Toyota"},
				{"3", "Volkswagen"},
			},
		},
		{
			name: "All same brand",
			reqList: []CarModel{
				{"1", "Audi"},
				{"2", "Audi"},
				{"3", "Audi"},
			},
			want: []CarModel{
				{"1", "Audi"},
				{"2", "Audi"},
				{"3", "Audi"},
			},
		},
		{
			name: "Multiple brands with consecutive cases",
			reqList: []CarModel{
				{"1", "BMW"},
				{"2", "BMW"},
				{"3", "BMW"},
				{"4", "Ford"},
				{"5", "BMW"},
				{"6", "Toyota"},
				{"7", "Toyota"},
				{"8", "Toyota"},
				{"9", "BMW"},
			},
			want: []CarModel{
				{"1", "BMW"},
				{"2", "BMW"},
				{"4", "Ford"},
				{"3", "BMW"},
				{"5", "BMW"},
				{"6", "Toyota"},
				{"7", "Toyota"},
				{"9", "BMW"},
				{"8", "Toyota"},
			},
		},
	}

	// 初始化 BrandScatterModule
	module := BrandScatterModule{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := module.DoAction(context.Background(), tt.reqList)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoAction() got = %v, want %v", got, tt.want)
			}
		})
	}

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
