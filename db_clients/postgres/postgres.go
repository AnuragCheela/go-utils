package postgres

import (
	"database/sql"
	"fmt"

	"github.com/AnuragCheela/go-utils/config"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "anuragcheela"
// 	password = ""
// 	dbname   = "users_db"
// )

var (
	client *sql.DB
)

// DBConfig struct
type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

func init() {

	var dbConfig DBConfig
	configError := config.GetDecodedConfig().Decode(&dbConfig)
	if configError != nil {
		panic(configError)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname)

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
