package storage

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func redisSetData(chat_id, user_id int64, json_data string) error {
	conn := redisConnection()
	defer conn.Close()
	key := fmt.Sprintf("%s-%d-%d-data", RedisName, chat_id, user_id)
	err := conn.Set(key, json_data, 0).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func redisGetData(chat_id, user_id int64) (string, error) {
	conn := redisConnection()
	defer conn.Close()
	key := fmt.Sprintf("%s-%d-%d-data", RedisName, chat_id, user_id)
	val, err := conn.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		log.Println(err)
		return "", err
	}
	return val, nil
}

func redisSetStage(chat_id, user_id int64, stage string) error {
	conn := redisConnection()
	defer conn.Close()
	key := fmt.Sprintf("%s-%d-%d-stage", RedisName, chat_id, user_id)
	err := conn.Set(key, stage, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func redisGetStage(chat_id, user_id int64) (string, error) {
	conn := redisConnection()
	defer conn.Close()
	key := fmt.Sprintf("%s-%d-%d-stage", RedisName, chat_id, user_id)
	val, err := conn.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		log.Println(err)
		return "", err
	}
	return val, nil
}

func redisConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: RedisPassword, // no password set
		DB:       RedisDatabase, // use default DB
	})
	return rdb
}
