package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Values Config

type Config struct {
	BatchWorkers int    `mapstructure:"batchworkers"`
	LogFilePath  string `mapstructure:"logfilepath"`
}

func Init(path, fileName, format string) {
	var err error
	Values, err = loadConfig(path, fileName, format)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func loadConfig(path, fileName, format string) (Config, error) {
	viper.SetConfigName(fileName) // name of config file (without extension)
	viper.SetConfigType(format)   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)     // path to look for the config file in
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
