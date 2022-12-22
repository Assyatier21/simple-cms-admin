package database

import (
	"cms-admin/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Init() *sql.DB {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&search_path=%s", config.User, config.Password, config.Host, config.Database, config.Schema)
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
