// @Author: Ciusyan 4/14/24

package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "",               // 密码，如果没有设置则留空
		DB:       0,                // 默认数据库选择
	})

	ctx := context.Background()

	// 间隔性更新 key
	ticker := time.NewTicker(5 * time.Second) // 每 5 秒执行一次
	defer ticker.Stop()

	for range ticker.C {
		// 更新 n9e-k1 的值
		newValueK1 := r.Intn(81) // 生成 0 到 50 的随机数
		if err := rdb.Set(ctx, "n9e-k1", newValueK1, 0).Err(); err != nil {
			fmt.Println("Error setting n9e-k1:", err)
			continue
		}
		fmt.Println("Updated n9e-k1 to", newValueK1)

		// 获取 n9e-k2 当前值
		currentValK2, err := rdb.Get(ctx, "n9e-k2").Int()
		if err != nil {
			fmt.Println("Error getting n9e-k2:", err)
			continue
		}

		// 更新 n9e-k2 的值
		changeK2 := -10 + r.Intn(21) // 生成 -10 到 10 的随机数
		newValueK2 := currentValK2 + changeK2
		if err := rdb.Set(ctx, "n9e-k2", newValueK2, 0).Err(); err != nil {
			fmt.Println("Error setting n9e-k2:", err)
			continue
		}
		fmt.Println("Updated n9e-k2 to", newValueK2)
	}
}
