package storage

import "errors"

func setData(chat_id, user_id int64, json_data string) error {
	switch Database {
	case "postgres":
		return postgresSetData(chat_id, user_id, json_data)
	case "redis":
		return redisSetData(chat_id, user_id, json_data)
	}
	return errors.New("undefind database")
}

func getData(chat_id, user_id int64) (string, error) {
	switch Database {
	case "postgres":
		return postgresGetData(chat_id, user_id)
	case "redis":
		return redisGetData(chat_id, user_id)
	}
	return "", errors.New("undefind database")
}

func setStage(chat_id, user_id int64, stage string) error {
	switch Database {
	case "postgres":
		return postgresSetStage(chat_id, user_id, stage)
	case "redis":
		return redisSetStage(chat_id, user_id, stage)
	}
	return errors.New("undefind database")
}

func getStage(chat_id, user_id int64) (string, error) {
	switch Database {
	case "postgres":
		return postgresGetStage(chat_id, user_id)
	case "redis":
		return redisGetStage(chat_id, user_id)
	}
	return "", errors.New("undefind database")

}
