package configIP

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port           string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connection_string"`
} // Config struct for the application configurations to be loaded from config.json file using Viper package

var AppConfig *Config // AppConfig is the global variable for the application configurations

// LoadAppConfig loads the application configurations from config.json file using Viper package
func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server Configurations Loaded...")
}

