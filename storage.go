package storage

import (
	"encoding/json"
	"log"
)

var PostgresUsername, PostgresPassword, PostgresDatabase string
var RedisDatabase int
var RedisPassword, RedisName string

var Database string

func SetPostgresConfig(username, password, database string) {
	PostgresUsername = username
	PostgresPassword = password
	PostgresDatabase = database
	Database = "postgres"
}

func SetRedisConfig(database int, password, name string) {
	RedisDatabase = database
	RedisPassword = password
	RedisName = name
	Database = "redis"
}

func GetData(chat_id, user_id int64, obj interface{}) error {
	raw_json, err := getData(chat_id, user_id)
	if err != nil {
		log.Println(err)
		return err
	}

	json_data := []byte(raw_json)

	err = json.Unmarshal(json_data, &obj)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SetData(chat_id, user_id int64, obj interface{}) error {
	json_data, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
		return err
	}

	err = setData(chat_id, user_id, string(json_data))

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetStage(chat_id, user_id int64) (string, error) {
	stage, err := getStage(chat_id, user_id)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return stage, nil
}

func SetStage(chat_id, user_id int64, stage string) error {
	err := setStage(chat_id, user_id, stage)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
