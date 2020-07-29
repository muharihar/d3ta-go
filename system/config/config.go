package config

import (
	"fmt"
	"path/filepath"

	"github.com/muharihar/d3ta-go/system/utils"
	"github.com/spf13/viper"
)

// NewConfig is a function to Load Configuration
func NewConfig(path string) (*Config, *viper.Viper, error) {

	defaultConfigFile, err := GetConfigFilePath(path)
	if err != nil {
		return nil, nil, err
	}

	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err.Error()))
	}
	v.WatchConfig()

	c := new(Config)
	err = v.Unmarshal(&c)
	if err != nil {
		return nil, nil, err
	}

	return c, v, err
}

// GetConfigFilePath get filepath location
func GetConfigFilePath(baseDir string) (string, error) {
	filename := "config.yaml"
	filePath := baseDir

	// first location
	file := filepath.Join(filePath, filename)
	exist, _ := utils.FileIsExist(file)

	if exist == false {
		// check in default location [./] -> binary dir
		filePath = "./"
		file = filepath.Join(filePath, filename)
		exist, _ := utils.FileIsExist(file)

		if exist == false {
			// check in configuration directory [./conf]
			filePath = "./conf"
			file = filepath.Join(filePath, filename)
			_, err := utils.FileIsExist(file)

			// give up, return error
			if err != nil {
				return "", err
			}
			// return file, nil
		}
		// return file, nil
	}

	return file, nil
}
