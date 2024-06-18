package str

import (
	"context"
	"Explore_Go/exercise/demo_redis/utils"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestSet(rdb *redis.Client, ctx context.Context) {
	// 设置值
	err := rdb.Set(ctx, "name", "Tom", 0).Err()
	utils.HandleErr(err)

	err = rdb.Set(ctx, "age", 18, time.Minute*30).Err()
	utils.HandleErr(err)

	// 设置值并设置过期时间
	err = rdb.SetEx(ctx, "address", "Shanghai", time.Minute*30).Err()
	utils.HandleErr(err)

	// 如果值不存在设置值并设置过期时间
	err = rdb.SetNX(ctx, "address", "Shanghai", time.Minute*30).Err()
	utils.HandleErr(err)

	// GetSet
	val, err := rdb.GetSet(ctx, "name", "Jerry").Result()
	utils.HandleErr(err)
	fmt.Println(val)

	// MSet
	err = rdb.MSet(ctx, "name", "Tom", "age", 18).Err()
	utils.HandleErr(err)

	// MSetNX
	err = rdb.MSetNX(ctx, "name", "Tom", "age", 20).Err()
	utils.HandleErr(err)
}

func TestGet(rdb *redis.Client, ctx context.Context) {
	// 获取值
	keys, err := rdb.Keys(ctx, "*").Result()
	utils.HandleErr(err)
	fmt.Println(keys)
}

func TestIncDec(rdb *redis.Client, ctx context.Context) {
	// 自增
	err := rdb.Incr(ctx, "age").Err()
	utils.HandleErr(err)

	// 自减
	err = rdb.Decr(ctx, "age").Err()
	utils.HandleErr(err)

	// 自增指定值
	err = rdb.IncrBy(ctx, "age", 10).Err()
	utils.HandleErr(err)

	// 自减指定值
	err = rdb.DecrBy(ctx, "age", 10).Err()
	utils.HandleErr(err)

	// 自增浮点数
	err = rdb.IncrByFloat(ctx, "age", 0.5).Err()
	utils.HandleErr(err)

	// 自减浮点数
	err = rdb.IncrByFloat(ctx, "age", -1.5).Err()
	utils.HandleErr(err)

	// 设置键过期时间
	err = rdb.Expire(ctx, "age", time.Minute*30).Err()
	utils.HandleErr(err)
}

func TestAppend(rdb *redis.Client, ctx context.Context) {
	// 追加值
	err := rdb.Append(ctx, "name", " is a good student").Err()
	utils.HandleErr(err)

	// 获取字符串长度
	len, err := rdb.StrLen(ctx, "name").Result()
	utils.HandleErr(err)
	fmt.Println(len)
}
