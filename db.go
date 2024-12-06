package todo

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var pool *sql.DB

func Init() {
	godotenv.Load()
	source := os.Getenv("DB_SOURCE_NAME")
	pool, err := sql.Open("postgres", source)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = pool.Ping()

	if err != nil {
		log.Fatal(err)
		return
	}
}
