package config

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

// RedisInfo Redis服务器信息结构体
type RedisInfo struct {
	RedisVersion      string `json:"redis_version"`
	RedisGitSha1      string `json:"redis_git_sha1"`
	RedisGitDirty     string `json:"redis_git_dirty"`
	RedisBuildId      string `json:"redis_build_id"`
	RedisMode         string `json:"redis_mode"`
	Os                string `json:"os"`
	ArchBits          int    `json:"arch_bits"`
	MonotonicClock    string `json:"monotonic_clock"`
	MultiplexingApi   string `json:"multiplexing_api"`
	AtomicvarApi      string `json:"atomicvar_api"`
	GccVersion        string `json:"gcc_version"`
	ProcessId         int    `json:"process_id"`
	ProcessSupervised string `json:"process_supervised"`
	RunId             string `json:"run_id"`
	TcpPort           int    `json:"tcp_port"`
	ServerTimeUsec    int64  `json:"server_time_usec"`
	UptimeInSeconds   int64  `json:"uptime_in_seconds"`
	UptimeInDays      int    `json:"uptime_in_days"`
	HZ                int    `json:"hz"`
	ConfiguredHZ      int    `json:"configured_hz"`
	LruClock          int    `json:"lru_clock"`
	Executable        string `json:"executable"`
	ConfigFile        string `json:"config_file"`
	IoThreadsActive   int    `json:"io_threads_active"`
}

// ConnectInfoServer 连接信息
func (c *DefaultConfigurator) ConnectInfoServer(rdb *redis.Client, ctx context.Context) {
	// 获取 Redis 服务器信息
	val, err := rdb.Do(ctx, "info", "server").Result()
	if err != nil {
		panic(err)
	}
	// fmt.Println(val)

	// 将val转换为字符串并去除首行的#Server，根据换行符分割字符串为多行
	lines := strings.Split(strings.TrimPrefix(val.(string), "# Server\n"), "\n")

	// 创建RedisInfo实例
	info := &RedisInfo{}

	// 创建一个映射，将每个键映射到相应的设置函数
	fieldSetters := map[string]func(string){
		"redis_version":      func(value string) { info.RedisVersion = value },
		"redis_git_sha1":     func(value string) { info.RedisGitSha1 = value },
		"redis_git_dirty":    func(value string) { info.RedisGitDirty = value },
		"redis_build_id":     func(value string) { info.RedisBuildId = value },
		"redis_mode":         func(value string) { info.RedisMode = value },
		"os":                 func(value string) { info.Os = value },
		"arch_bits":          func(value string) { info.ArchBits, _ = strconv.Atoi(value) },
		"monotonic_clock":    func(value string) { info.MonotonicClock = value },
		"multiplexing_api":   func(value string) { info.MultiplexingApi = value },
		"atomicvar_api":      func(value string) { info.AtomicvarApi = value },
		"gcc_version":        func(value string) { info.GccVersion = value },
		"process_id":         func(value string) { info.ProcessId, _ = strconv.Atoi(value) },
		"process_supervised": func(value string) { info.ProcessSupervised = value },
		"run_id":             func(value string) { info.RunId = value },
		"tcp_port":           func(value string) { info.TcpPort, _ = strconv.Atoi(value) },
		"server_time_usec":   func(value string) { info.ServerTimeUsec, _ = strconv.ParseInt(value, 10, 64) },
		"uptime_in_seconds":  func(value string) { info.UptimeInSeconds, _ = strconv.ParseInt(value, 10, 64) },
		"uptime_in_days":     func(value string) { info.UptimeInDays, _ = strconv.Atoi(value) },
		"hz":                 func(value string) { info.HZ, _ = strconv.Atoi(value) },
		"configured_hz":      func(value string) { info.ConfiguredHZ, _ = strconv.Atoi(value) },
		"lru_clock":          func(value string) { info.LruClock, _ = strconv.Atoi(value) },
		"executable":         func(value string) { info.Executable = value },
		"config_file":        func(value string) { info.ConfigFile = value },
		"io_threads_active":  func(value string) { info.IoThreadsActive, _ = strconv.Atoi(value) },
	}

	// 解析每一行并填充到结构体
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := parts[0]
			// 去除空格
			value := strings.TrimSpace(parts[1])
			if setter, ok := fieldSetters[key]; ok {
				setter(strings.TrimSpace(value))
			} else {
				fmt.Printf("未找到匹配的键：%s\n", key)
			}
		}
	}

	// 打印解析后的结构体
	fmt.Printf("%+v\n", info)
}

// ConnectInfoServer2 连接信息,使用反射优化
func (c *DefaultConfigurator) ConnectInfoServer2(rdb *redis.Client, ctx context.Context) {
	// 获取 Redis 服务器信息
	val, err := rdb.Do(ctx, "info", "server").Result()
	if err != nil {
		panic(err)
	}
	// fmt.Println(val)

	// 将val转换为字符串并去除首行的#Server，根据换行符分割字符串为多行
	lines := strings.Split(strings.TrimPrefix(val.(string), "# Server\n"), "\n")

	// 创建RedisInfo实例
	info := &RedisInfo{}

	// 使用反射填充结构体
	infoValue := reflect.ValueOf(info).Elem()
	infoType := infoValue.Type()

	// 缓存结构体字段的标签和对应的字段值
	fieldMap := make(map[string]reflect.Value)
	for i := 0; i < infoType.NumField(); i++ {
		field := infoType.Field(i)
		tag := field.Tag.Get("json")
		if tag != "" {
			fieldMap[tag] = infoValue.Field(i)
		}
	}

	// 解析每一行并填充到结构体
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := strings.TrimSpace(parts[1])

			if fieldValue, ok := fieldMap[key]; ok {
				switch fieldValue.Kind() {
				case reflect.String:
					fieldValue.SetString(value)
				case reflect.Int, reflect.Int64:
					intValue, _ := strconv.ParseInt(value, 10, fieldValue.Type().Bits())
					fieldValue.SetInt(intValue)
				}
			}
		}
	}

	// 打印解析后的结构体
	fmt.Printf("%+v\n", info)
}
