// @Author: Ciusyan 10/27/23

package other

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/require"
	"math/rand"
	"os"
	"sync"
	"testing"
)

var (
	once   sync.Once
	client *redis.Client

	sha string
)

func init() {
	// 连接 Redis
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // 如果没有密码，就写空字符串
		DB:       0,  // 使用默认 DB
	})
}

func getSha() {
	once.Do(func() {
		scriptContent, err := os.ReadFile("random-pop.lua")
		if err != nil {
			panic(err)
		}
		sha, err = client.ScriptLoad(string(scriptContent)).Result()

		if err != nil {
			panic(err)
		}
	})

	fmt.Println(sha)
}

func TestRedis_Lua(t *testing.T) {

	key := "mylist"

	// 延迟关闭连接
	defer client.Close()

	// 向 list 中添加一些元素
	//client.RPush(key, "a", "b", "c", "d", "e")
	//
	//getSha()
	//
	//// 加载 Lua 脚本，并获取 SHA1 校验码（假设你要操作的列表键名是 mylist）
	//
	//// 执行已加载的 Lua 脚本，并传入 key 和 count 参数（假设你要弹出 3 个元素）
	//result, err := client.EvalSha(sha, []string{key}, 2).Result()
	//require.NoError(t, err)
	//
	//// 打印结果（可能会有不同的顺序）
	//t.Log("EvalSha result:", result)
	//
	//// 删除 list 中的所有元素
	//_, err = client.Del(key).Result()
	//require.NoError(t, err)

	// 向 list 中添加一些元素

	randVals := func(size int, max int64) []interface{} {
		res := make([]interface{}, size)
		for i := 0; i < size; i++ {
			res[i] = rand.Int63n(max) + 1
		}
		return res
	}

	// 执行前，先刷一些数据
	client.RPush(key, randVals(100, 3000)...)

	// 加载 Lua 脚本，并获取 SHA1 校验码（假设你要操作的列表键名是 mylist）

	// 执行已加载的 Lua 脚本，并传入 key 和 count 参数（假设你要弹出 3 个元素）
	result, err := client.EvalSha("3d707a8b804a8de0a29621d07d4cbecd0059d40b", []string{key}, 3).Result()
	require.NoError(t, err)

	// 打印结果（可能会有不同的顺序）
	t.Log("2EvalSha result:", result)

	// 删除 list 中的所有元素
	_, err = client.Del(key).Result()
	require.NoError(t, err)
}
