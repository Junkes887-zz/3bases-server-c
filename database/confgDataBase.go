package database

import (
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

// CreateConnection Cria conexão com o banco
func CreateConnection() *elastic.Client {
	DATABASE_URL := os.Getenv("DATABASE_URL")

	client, err := elastic.NewClient(elastic.SetURL(DATABASE_URL),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	return client
}
