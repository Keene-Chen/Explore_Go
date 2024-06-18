package str

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimiter struct {
	client     *redis.Client
	rate       int
	windowSize time.Duration
	prefix     string
}

func NewRateLimiter(client *redis.Client, rate int, windowSize time.Duration, prefix string) *RateLimiter {
	return &RateLimiter{
		client:     client,
		rate:       rate,
		windowSize: windowSize,
		prefix:     prefix,
	}
}

func (rl *RateLimiter) Allow(ctx context.Context, userID string) (bool, error) {
	key := fmt.Sprintf("%s:%s", rl.prefix, userID)

	// 使用Lua脚本原子性地执行INCR和EXPIRE
	script := `
	local current = redis.call("INCR", KEYS[1])
	if current == 1 then
		redis.call("EXPIRE", KEYS[1], ARGV[1])
	end
	if current > tonumber(ARGV[2]) then
		return 0
	end
	return 1
	`

	allowed, err := rl.client.Eval(ctx, script, []string{key}, int(rl.windowSize.Seconds()), rl.rate).Int()
	if err != nil {
		return false, err
	}

	return allowed == 1, nil
}

func TestRateLimiter(rdb *redis.Client, ctx context.Context) {
	limiter := NewRateLimiter(rdb, 10, time.Minute, "rate_limiter")
	userID := "user123"

	for i := 0; i < 12; i++ {
		allowed, err := limiter.Allow(ctx, userID)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		if allowed {
			fmt.Printf("Request %d: Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: Denied\n", i+1)
		}

		time.Sleep(1 * time.Second)
	}
}
