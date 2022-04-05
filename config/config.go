package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName         string `mapstructure:"APP_NAME"`
	Environment     string `mapstructure:"ENVIRONMENT"`
	ServerPort      string `mapstructure:"SERVER_PORT"`
	DbAutoInit      bool   `mapstructure:"DB_AUTOINIT"`
	DbDriver        string `mapstructure:"DB_DRIVER"`
	DbHost          string `mapstructure:"DB_HOST"`
	DbPort          int    `mapstructure:"DB_PORT"`
	DbUser          string `mapstructure:"DB_USER"`
	DbName          string `mapstructure:"DB_NAME"`
	DbPassword      string `mapstructure:"DB_PASSWORD"`
	DbSingulartable bool   `mapstructure:"DB_SINGULARTABLE"`
	RedisAddress    string `mapstructure:"REDIS_ADDRESS"`
	RedisPassword   string `mapstructure:"REDIS_PASSWORD"`
	RedisDb         int    `mapstructure:"REDIS_DB"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
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
