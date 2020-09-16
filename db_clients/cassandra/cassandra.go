package cassandra

import (
	"github.com/AnuragCheela/go-utils/config"
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

// GlobalConfig struct
type GlobalConfig struct {
	DbConfig DBConfig `json:"cassandra_config"`
}

// DBConfig struct
type DBConfig struct {
	Host     string `json:"host"`
	Keyspace string `json:"keyspace"`
}

func init() {
	// Connect to Cassandra cluster:

	var globalConfig GlobalConfig
	configError := config.GetDecodedConfig().Decode(&globalConfig)
	if configError != nil {
		panic(configError)
	}
	cluster := gocql.NewCluster(globalConfig.DbConfig.Host)
	cluster.Keyspace = globalConfig.DbConfig.Keyspace
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

// GetSession function to be called from elsewhere
func GetSession() *gocql.Session {
	return session
}
