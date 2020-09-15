package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "anuragcheela"
	password = ""
	dbname   = "users_db"
)

var (
	client *sql.DB
)

func init() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	client, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(); err != nil {
		panic(err)
	}
}

// GetClient function
func GetClient() *sql.DB {
	return client
}
