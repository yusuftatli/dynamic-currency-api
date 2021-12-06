package redis

import (
	"encoding/json"

	"github.com/go-redis/redis/v7"
	common "github.com/yusuftatli/hepsiburada/config"
)

var (
	client = &RedisClient{}
)

type RedisClient struct {
	c *redis.Client
}

//GetClient get the redis client
func Initialize() *RedisClient {
	c := redis.NewClient(&redis.Options{
		Addr: common.GetEnvironment().RedisAddress,
		Password: common.GetEnvironment().RedisPassword,
		DB: 0,
	})
	if err := c.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}
	return &RedisClient{c}
}



//GetKey get key
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

//SetKey set key
func (client *RedisClient) SetKey(key string, value interface{}) error {
	 cacheEntry, err := json.Marshal(value)
	 if err != nil {
	 	return err
	 }
	err = client.c.Set(key, cacheEntry, 0).Err()
	if err != nil {
		return err
	}
	return nil
}