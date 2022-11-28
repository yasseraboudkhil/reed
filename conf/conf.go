package conf

import (
	"strings"

	"github.com/spf13/viper"
)

type AWSConfig struct {
	Region    string
	AccessKey string
	SecretKey string
}

type KMSConfig struct {
	KeyId string
}

type ServerConfig struct {
	Host string
	Port int
}

type Config struct {
	Server ServerConfig
	Aws    AWSConfig
	Kms    KMSConfig
}

var C Config

func LoadConfig(configFile string) error {

	// Viper setup for ENV config
	viper.SetEnvPrefix("REED")
	/*
	 * Nested key configuration structures in viper are delimited
	 * with a "."
	 * In order to match this against our environment vars which
	 * are delimited with a "_", we use the replacer below
	 */
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// Ensure viper reads from ENV vars as well
	viper.AutomaticEnv()

	// Viper setup for file config
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("conf")
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&C); err != nil {
		return err
	}

	return nil
}
