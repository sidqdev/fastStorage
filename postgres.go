package storage

import (
	"errors"
	"log"

	"github.com/jackc/pgx"
)

func postgresSetData(chat_id, user_id int64, json_data string) error {
	conn, err := postgresConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	var record int
	err = conn.QueryRow("SELECT COUNT(*) FROM storage_data WHERE chat_id = $1 AND user_id = $2", chat_id, user_id).Scan(&record)
	if err != nil && err != pgx.ErrNoRows {
		log.Println("Can not select row from database")
		return errors.New("can not select row from database")
	}
	if record == 0 {
		_, err = conn.Exec("INSERT INTO storage_data(chat_id, user_id, json_data, stage) VALUES($1, $2, $3, NULL)", chat_id, user_id, json_data)
	} else {
		_, err = conn.Exec("UPDATE storage_data SET json_data = $3 WHERE chat_id = $1 AND user_id = $2", chat_id, user_id, json_data)
	}
	if err != nil {
		log.Println("Can not select row from database")
		return errors.New("can not select row from database")
	}

	return nil
}

func postgresGetData(chat_id, user_id int64) (string, error) {
	conn, err := postgresConnection()
	if err != nil {
		return "", err
	}
	defer conn.Close()
	var json_data string
	err = conn.QueryRow("SELECT json_data FROM storage_data WHERE chat_id = $1 AND user_id = $2", chat_id, user_id).Scan(&json_data)
	if err != nil {
		log.Println("Can not select row from database")
		return "", errors.New("can not select row from database")
	}
	return json_data, nil
}

func postgresGetStage(chat_id, user_id int64) (string, error) {
	conn, err := postgresConnection()
	if err != nil {
		return "", err
	}
	defer conn.Close()
	var stage string
	err = conn.QueryRow("SELECT stage FROM storage_data WHERE chat_id = $1 AND user_id = $2", chat_id, user_id).Scan(&stage)
	if err != nil && err != pgx.ErrNoRows {
		log.Println(err)
		return "", errors.New("can not select row from database")
	}
	return stage, nil

}

func postgresSetStage(chat_id, user_id int64, stage string) error {
	conn, err := postgresConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	var record int
	err = conn.QueryRow("SELECT COUNT(*) FROM storage_data WHERE chat_id = $1 AND user_id = $2", chat_id, user_id).Scan(&record)
	if err != nil {
		log.Println("Can not select row from database")
		return errors.New("can not select row from database")
	}
	if record == 0 {
		_, err = conn.Exec("INSERT INTO storage_data(chat_id, user_id, json_data, stage) VALUES($1, $2, NULL, $3)", chat_id, user_id, stage)
	} else {
		_, err = conn.Exec("UPDATE storage_data SET stage = $3 WHERE chat_id = $1 AND user_id = $2", chat_id, user_id, stage)
	}
	if err != nil {
		log.Println("Can not select row from database")
		return errors.New("can not select row from database")
	}

	return nil
}

func postgresInitDatabase() error {
	conn, err := postgresConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Exec(
		"CREATE TABLE IF NOT EXISTS storage_data " +
			"(chat_id INT " +
			"chat_id INT " +
			"json_data TEXT " +
			"stage TEXT)")
	return err
}

func postgresConnection() (*pgx.Conn, error) {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		User:     PostgresUsername,
		Password: PostgresPassword,
		Database: PostgresDatabase})
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil, errors.New("no connection to database")
	}

	return conn, nil
}
