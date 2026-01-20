package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
	Logging  LoggingConfig  `mapstructure:"logging"`
	Server   ServerConfig   `mapstructure:"server"`
}

type (
	PostgresConfig struct {
		Conn string `mapstructure:"conn"`
	}
	LoggingConfig struct {
		Level string `mapstructure:"level"`
	}
	ServerConfig struct {
		Port string `mapstructure:"port"`
	}
)

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath("./config/")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return v, nil
}

func GetConfig() (*Config, error) {
	v, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
