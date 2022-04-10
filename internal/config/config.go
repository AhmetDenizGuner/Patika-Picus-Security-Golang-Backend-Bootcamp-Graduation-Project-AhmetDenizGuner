package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	defaultEnv = "local"
)

var cfgReader *configReader

type (
	Configuration struct {
		DatabaseSettings
		JwtSettings
		RedisSettings
	}

	DatabaseSettings struct {
		DatabaseURI  string
		DatabaseName string
		Username     string
		Password     string
	}

	JwtSettings struct {
		SecretKey string
	}
	RedisSettings struct {
		AddrUrI string
	}

	configReader struct {
		configFile string
		v          *viper.Viper
	}
)

func GetAllConfigValues(configFile string) (configuration *Configuration, err error) {
	newConfigReader(configFile)
	if err = cfgReader.v.ReadInConfig(); err != nil {
		fmt.Printf("Failed to read config file : %s\n", err)
		return nil, err
	}

	err = cfgReader.v.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Failed to unmarshal yaml file to configuration struct : %s\n", err)
		return nil, err
	}

	return configuration, err
}

func newConfigReader(configFile string) {
	v := viper.GetViper()
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	cfgReader = &configReader{
		configFile: configFile,
		v:          v,
	}
}
