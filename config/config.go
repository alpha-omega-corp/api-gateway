package config

import "github.com/spf13/viper"

type Config struct {
	HOST   string `mapstruct:"HOST"`
	AUTH   string `mapstruct:"AUTH"`
	DOCKER string `mapstruct:"DOCKER"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
