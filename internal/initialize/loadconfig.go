package initialize

import (
	"log"

	"github.com/csc13010-student-management/config"
	"github.com/spf13/viper"
)

func LoadConfig() *config.Config {
	cfg := config.NewConfig()

	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Error reading config file")
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Println("Error unmarshalling config")
		panic(err)
	}

	log.Println("Config loaded successfully")
	return cfg
}
