package postgres

import (
	"database/sql"
	"fmt"

	"github.com/AnuragCheela/go-utils/config"
	"github.com/AnuragCheela/go-utils/logger"

	_ "github.com/lib/pq"
)

var (
	client *sql.DB
)

// GlobalConfig struct
type GlobalConfig struct {
	DbConfig DBConfig `json:"postgres_config"`
}

// DBConfig struct
type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

func init() {

	logger.Info("init function of postgres started")

	var globalConfig GlobalConfig
	configError := config.GetDecodedConfig().Decode(&globalConfig)
	if configError != nil {
		panic(configError)
	}
	logger.Info(fmt.Sprintf(" config is %+v", globalConfig.DbConfig))

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		globalConfig.DbConfig.Host, globalConfig.DbConfig.Port, globalConfig.DbConfig.User, globalConfig.DbConfig.Password, globalConfig.DbConfig.Dbname)

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
