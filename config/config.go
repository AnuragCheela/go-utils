package config

import (
	"encoding/json"
	"os"
)

const (
	fileNameStage   = "CONFIG_DEV_FILE"
	fileNameProd    = "CONFIG_DEV_FILE"
	environment     = "SERVICE_ENV"
	prodEnvironment = "prod"
)

var (
	config     *json.Decoder
	configFile *os.File
)

func init() {
	var err error
	if os.Getenv(environment) == prodEnvironment {
		configFile, err = os.Open(os.Getenv(fileNameProd))
	} else {
		configFile, err = os.Open(os.Getenv(fileNameStage))
	}
	if err != nil {
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
