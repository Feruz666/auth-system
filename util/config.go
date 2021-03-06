package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configurations of the application
type Config struct {
	DBDriver               string        `mapstructure:"DB_DRIVER"`
	DBSource               string        `mapstructure:"DB_SOURCE"`
	ServerAddress          string        `mapstructure:"SERVER_ADDRESS"`
	TokenSecretKey         string        `mapstructure:"SECRET_KEY"`
	AccessTokenDuration    time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration   time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	MAPS_SYSTEM_ADDRESS    string        `mapstructure:"MAPS_SYSTEM_ADDRESS"`
	SENSORS_SYSTEM_ADDRESS string        `mapstructure:"SENSORS_SYSTEM_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables
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
