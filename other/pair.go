// @Author: Ciusyan 2/21/24

package other

import (
	"encoding/json"
	"errors"
	"strings"
)

/*
parseJson2Pairs 解析成键值对数组，解析形如这样的 Json

		`[
	        "{ \"clickKey\": \"uid\"}",
	        "{ \"topKey\": \"cityId\"}"
	    ]`
*/
func parseJson2Pairs(jsonData string) ([]map[string]string, error) {
	if jsonData == "" {
		return nil, errors.New("依赖数据配置有误")
	}

	var (
		err      error
		jsonStrs []string
	)
	// 先解析成数组
	if err = json.Unmarshal([]byte(jsonData), &jsonStrs); err != nil {
		return nil, err
	}

	// 再解析出动态的字符串
	var pairs []map[string]string
	for _, jsonStr := range jsonStrs {
		var pair map[string]string
		// 再挨个解析成键值对
		if err = json.Unmarshal([]byte(jsonStr), &pair); err != nil {
			return nil, err
		}

		limit := 0
		for k, v := range pair {
			limit++
			if v == "" || limit != 1 {
				return nil, errors.New("依赖数据配置有误")
			}

			// 统一转换成小写，并且去掉蛇形（防止格式配置错误）
			pair[k] = strings.Replace(strings.ToLower(v), "_", "", -1)
		}

		pairs = append(pairs, pair)
	}

	if len(pairs) == 0 {
		return nil, errors.New("依赖数据配置有误")
	}

	return pairs, nil
}
