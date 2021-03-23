package configuration

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagForRedisHost = "redis_host"
	flagForRedisPort = "redis_port"
)

type RedisConfig struct {
	Host string
	Port int
}

func LoadRedisConfig() RedisConfig {
	c := RedisConfig{
		Host: "redis",
		Port: 6379,
	}

	flag.String(flagForRedisHost, c.Host, "Redis host")
	flag.Int(flagForRedisPort, c.Port, "Redis port")

	flag.Parse()

	viper.BindPFlag(flagForRedisHost, flag.Lookup(flagForRedisHost))
	viper.BindPFlag(flagForRedisPort, flag.Lookup(flagForRedisPort))

	c.Host = viper.GetString(flagForRedisHost)
	c.Port = viper.GetInt(flagForRedisPort)

	return c
}

func (c RedisConfig) GetRingOptions() *redis.RingOptions {
	addr1 := fmt.Sprintf("%s:%d", c.Host, c.Port)
	return &redis.RingOptions{
		Addrs: map[string]string{
			"server1": addr1,
			//"server2": ":6380",
		},
	}
}

func (c RedisConfig) GetOptions() *redis.Options {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	return &redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	}
}
