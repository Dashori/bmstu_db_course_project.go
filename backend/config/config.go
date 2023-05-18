package config

import (
	"backend/cmd/modes/flags"
	"github.com/spf13/viper"
)

type Config struct {
	Postgres flags.PostgresFlags `mapstructure:"postgres"`
	Address  string              `mapstructure:"address"`
	Port     string              `mapstructure:"port"`
	LogLevel string              `mapstructure:"loglevel"`
	LogFile  string              `mapstructure:"logfile"`
	Mode     string              `mapstructure:"mode"`
}

func (c *Config) ParseConfig(configFileName, pathToConfig string) error {
	v := viper.New()
	v.SetConfigName(configFileName)
	v.SetConfigType("json")
	v.AddConfigPath(pathToConfig)

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(c) //  в  json
	if err != nil {
		return err
	}

	return nil
}
