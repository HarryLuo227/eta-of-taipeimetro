package configuration

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Conf struct {
	Address         string `mapstucture:"ADDRESS`
	Port            int    `mapstructure:"PORT"`
	MongoDB_Address string `mapstructure:"MONGODB_ADDRESS"`
	MongoDB_Port    int    `mapstructure:"MONGODB_PORT`
}

func LoadConfig(configPath string) {
	viper.SetConfigName("local")    // Set configuration file name.
	viper.SetConfigType("env")      // Set configuration file type.
	viper.AddConfigPath(configPath) // Set configuration file path.

	/**
	 * Bind environment variables
	 */
	// Web service server address
	if err := viper.BindEnv("ADDRESS"); err != nil {
		log.Println(err.Error())
	}
	// Web service port
	if err := viper.BindEnv("PORT"); err != nil {
		log.Println(err.Error())
	}
	// MongoDB server Address
	if err := viper.BindEnv("MONGODB_ADDRESS"); err != nil {
		log.Println(err.Error())
	}
	// MongoDB service Port
	if err := viper.BindEnv("MONGODB_PORT"); err != nil {
		log.Println(err.Error())
	}

	// Find and read configuration file with error handling.
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file, %s", err))
	}

	// Unmarshal the values into the variable `Conf`.
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("Error Unmarshaling values, %s", err))
	}
}
