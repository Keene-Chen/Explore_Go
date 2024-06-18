package main

import (
	"context"
	"log"

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
	conf.ConnectInfoServer2(rdb, ctx)

	// 操作 Redis
	// str.TestStrSet(rdb, ctx)
	// str.TestStrGet(rdb, ctx)

	// 关闭连接
	defer rdb.Close()
}
