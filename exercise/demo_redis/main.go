package main

import (
	"context"
	"log"

	str "demo_redis/commands"
	"demo_redis/config"
)

var ctx = context.Background()

func main() {
	// 连接 Redis
	conf := &config.DefaultConfigurator{}
	rdb, err := conf.ConnectRedis(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// 获取 Redis 连接信息
	// conf.ConnectInfoServer2(rdb, ctx)

	// 操作 Redis
	// str.TestSet(rdb, ctx)
	// str.TestGet(rdb, ctx)
	// str.TestIncDec(rdb, ctx)
	// str.TestAppend(rdb, ctx)
	str.TestRateLimiter(rdb, ctx)

	// 关闭连接
	defer rdb.Close()
}
