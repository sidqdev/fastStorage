package storage

func setData(chat_id, user_id int64, json_data string) error {
	switch Database {
	case "postgres":
		return postgresSetData(chat_id, user_id, json_data)
	case "redis":
		return redisSetData(chat_id, user_id, json_data)
	}
	return nil
}

func getData(chat_id, user_id int64) (string, error) {
	switch Database {
	case "postgres":
		return postgresGetData(chat_id, user_id)
	case "redis":
		return redisGetData(chat_id, user_id)
	}
	return "", nil
}

func setStage(chat_id, user_id int64, stage string) error {
	switch Database {
	case "postgres":
		return postgresSetStage(chat_id, user_id, stage)
	case "redis":
		return redisSetStage(chat_id, user_id, stage)
	}
	return nil
}

func getStage(chat_id, user_id int64) (string, error) {
	switch Database {
	case "postgres":
		return postgresGetStage(chat_id, user_id)
	case "redis":
		return redisGetStage(chat_id, user_id)
	}
	return "", nil

}
