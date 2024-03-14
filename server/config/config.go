package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Key  string
	Host string
	Port int32
	Dsn  string
}

func (c *Config) setEnv() {
	c.Host = viper.GetString("HOST")
	c.Port = viper.GetInt32("PORT")
	c.Key = viper.GetString("KEY")
	c.Dsn = viper.GetString("DSN")

}

func (c Config) GetEnvVar() Config {
	if _, err := os.Stat("/.env"); os.IsNotExist(err) {
		viper.SetConfigFile("./.env")
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				panic("[-] File not found!")
			}
		}
	}

	c.setEnv()
	return c
}

func (c Config) GetEnvVarTest() Config {
	viper.SetConfigFile("./../.env")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("[-] File not found!")
		}
	}
	c.setEnv()
	return c

}
