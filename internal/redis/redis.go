package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mabta/clpc/internal/cfg"
)

var client *redis.Client
var ctx = context.Background()

// Init 初始化Redis连接
func Init(addr, password string, db int) error {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return client.Ping(ctx).Err()
}

// InitFrom 从配置中初始化Redis连接
func InitFrom(c *cfg.RedisConfig) error {
	return Init(c.Addr, c.Password, c.DB)
}

// Set 设置值
func Set(key string, value interface{}) error {
	return SetExpiredSecond(key, value, 0)
}

// SetExpiredSecond 设置值，并指定以秒为单位的过期时间
func SetExpiredSecond(key string, value interface{}, expSecs int) error {
	return SetExpiredDuration(key, value, time.Duration(expSecs)*time.Second)
}

// SetExpiredDuration 设置值，并指定具体的duration
func SetExpiredDuration(key string, value interface{}, exp time.Duration) error {
	return client.Set(ctx, key, value, exp).Err()
}

// Get 获取值
func Get(key string) (value string, err error) {
	return GetOr(key, "")
}

// GetRaw 获取值，如果有错误，返回原始的错误信息
func GetRaw(key string) (value string, err error) {
	return client.Get(ctx, key).Result()
}

// GetOr 获取值，如果键不存在则使用指定的默认值
func GetOr(key string, val string) (value string, err error) {
	v, err := GetRaw(key)
	if err == redis.Nil {
		return val, nil
	}
	if err != nil {
		return "", err
	}
	return v, nil
}

// Exists 判断是否存在
func Exists(key string) (bool, error) {
	_, err := GetRaw(key)
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
