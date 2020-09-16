package config

import (
	"encoding/json"
	"os"

	"github.com/AnuragCheela/go-utils/logger"
)

const (
	fileName = "CONFIG_FILE"
)

var (
	config     *json.Decoder
	configFile *os.File
)

func init() {
	logger.Info("initializing config")
	var err error
	configFile, err = os.Open(os.Getenv(fileName))
	if err != nil {
		logger.Info("error reading file")
		panic(err)
	}
	config = json.NewDecoder(configFile)
}

//GetDecodedConfig Function
func GetDecodedConfig() *json.Decoder {
	return config
}

// CloseFilePointer function
func CloseFilePointer() {
	configFile.Close()
}
