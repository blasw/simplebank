package util

import (
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource      string `mapstructure:"DB_SOURCE"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	ServerAddress string `mapstructure:"SERVER_ADDR"`
}

func LoadConfig(path string) (config Config, err error) {
	os := runtime.GOOS

	if os == "windows" {
		config.DBSource = "postgresql://root:secret@172.29.46.53:5432/simple_bank?sslmode=disable"
	} else {
		config.DBSource = "postgresql://root:secret@0.0.0.0:5432/simple_bank?sslmode=disable"
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
