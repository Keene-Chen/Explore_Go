package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// RedisConfig 代表 Redis 连接的配置
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`     // 地址
	Username string `mapstructure:"username"` // 用户名
	Password string `mapstructure:"password"` // 密码
	DB       int    `mapstructure:"db"`       // 数据库索引
}

// Configurator 定义用于检索和连接 Redis 的接口
type Configurator interface {
	// GetRedisConfig 从环境中获取 Redis 配置
	GetRedisConfig() ([]RedisConfig, error)

	// ConnectRedis 使用提供的配置连接到 Redis 服务器
	ConnectRedis(ctx context.Context) (*redis.Client, error)

	// ConnectInfo 连接信息
	ConnectInfoServer(rdb *redis.Client, ctx context.Context)
}

// DefaultConfigurator 实现 Configurator 接口
type DefaultConfigurator struct{}

// GetRedisConfig 从环境中获取 Redis 配置
func (c *DefaultConfigurator) GetRedisConfig() ([]RedisConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config/config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("致命错误配置文件: %w", err)
	}

	var redisConfigs []RedisConfig
	if err := viper.UnmarshalKey("redis", &redisConfigs); err != nil {
		return nil, fmt.Errorf("无法解码为结构体: %w", err)
	}

	return redisConfigs, nil
}

// ConnectRedis 使用提供的配置连接到 Redis 服务器
func (c *DefaultConfigurator) ConnectRedis(ctx context.Context) (*redis.Client, error) {
	redisConfigs, err := c.GetRedisConfig()
	if err != nil {
		return nil, err
	}

	config := redisConfigs[0]
	opt := &redis.Options{
		Addr:     config.Addr,
		Username: config.Username,
		Password: config.Password,
		DB:       config.DB,
	}

	rdb := redis.NewClient(opt)
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("无法连接到 Redis: %w", err)
	}

	fmt.Printf("已连接到 Redis: %s!\n", pong)
	fmt.Printf("客户端版本: %s\n", redis.Version())

	return rdb, nil
}
