package database

import (
	"cms-admin/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Init() *sql.DB {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s", config.POSTGRE_USER, config.POSTGRE_PASSWORD, config.POSTGRE_HOST, config.POSTGRE_PORT, config.POSTGRE_DATABASE, config.POSTGRE_SSLMODE, config.POSTGRE_SCHEMA)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Println("[Database] can't connect to database: ", err.Error())
		return nil
	}

	log.Println("[Database] successfully connected")
	return db
}
