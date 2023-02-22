package database

import (
	"cms-admin/config"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/olivere/elastic/v7"
)

func Init() *sql.DB {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&search_path=%s", config.User, config.Password, config.Host, config.Database, config.Schema)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("[Database] initialized...")

	err = db.Ping()
	if err != nil {
		log.Println("[Database] failed to connect to database: ", err)
		return nil
	}

	log.Println("[Database] successfully connected")
	return db
}

func InitElasticClient() *elastic.Client {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(config.ESADDRESS),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	log.Println("[Elasticsearch] initialized...")
	if err != nil {
		log.Println("[Elasticsearch] failed to connect to elasticsearch: ", err)
		return nil
	}

	info, _, err := client.Ping(config.ESADDRESS).Do(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("[Elasticsearch] successfully connected. running version %s\n", info.Version.Number)
	return client
}
