package redis

import (
	"encoding/json"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/config"
	"github.com/go-redis/redis"
	"time"
)

type RedisClient struct {
	c *redis.Client
}

func NewRedisClient(appConfig *config.Configuration) *RedisClient {
	c := redis.NewClient(&redis.Options{
		Addr: "",
	})

	if err := c.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}

	client := &RedisClient{
		c: c,
	}

	return client
}

//GetKey get redis key
func (client *RedisClient) GetKey(key string, src interface{}) error {
	val, err := client.c.Get(key).Result()
	if err == redis.Nil || err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		return err
	}
	return nil
}

//SetKey set redis key
func (client *RedisClient) SetKey(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = client.c.Set(key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

//DeleteKey delete redis key
func (client *RedisClient) DeleteKey(key string) error {
	_, err := client.c.Del(key).Result()

	if err != nil {
		return err
	}
	return nil
}
