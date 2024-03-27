package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "docker"
	password = "docker"
	dbname   = "twitter_db"
)

var DB *sql.DB

func OpenDatabase() error {
	var error error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, error = sql.Open("postgres", psqlInfo)

	if error != nil {
		return error
	}
	return nil
}

func CloseDatabase() error {
	return DB.Close()
}
