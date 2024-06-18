package str

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func TestSet(rdb *redis.Client, ctx context.Context) {
	// 设置值
	err := rdb.Set(ctx, "name", "Tom", 0).Err()
	handleErr(err)

	err = rdb.Set(ctx, "age", 18, time.Minute*30).Err()
	handleErr(err)

	// 设置值并设置过期时间
	err = rdb.SetEx(ctx, "address", "Shanghai", time.Minute*30).Err()
	handleErr(err)

	// 如果值不存在设置值并设置过期时间
	err = rdb.SetNX(ctx, "address", "Shanghai", time.Minute*30).Err()
	handleErr(err)

	// GetSet
	val, err := rdb.GetSet(ctx, "name", "Jerry").Result()
	handleErr(err)
	fmt.Println(val)
}

func TestGet(rdb *redis.Client, ctx context.Context) {
	// 获取值
	keys, err := rdb.Keys(ctx, "*").Result()
	handleErr(err)
	fmt.Println(keys)
}
