package util

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func main() {
	config, err := LoadConfig("path/to/config")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// Use the config...
	fmt.Println("DB Driver:", config.DBDriver)
	fmt.Println("DB Source:", config.DBSource)
	fmt.Println("Server Address:", config.ServerAddress)
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}



