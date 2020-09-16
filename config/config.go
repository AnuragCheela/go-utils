package config

import (
	"encoding/json"
	"os"

	"github.com/AnuragCheela/go-utils/logger"
)

const (
	fileName = "CONFIG_FILE"
)

var config *json.Decoder

func init() {
	logger.Info("initializing config")
	configFile, err := os.Open(os.Getenv(fileName))
	if err != nil {
		logger.Info("error reading file")
	}
	defer configFile.Close()
	config = json.NewDecoder(configFile)
}

//GetDecodedConfig Function
func GetDecodedConfig() *json.Decoder {
	return config
}
