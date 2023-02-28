package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ChannelSecret string `yaml:"channel_secret"`
	ChannelToken  string `yaml:"channel_token"`
	FireStore     struct {
		ServiceAccountFile string `yaml:"service_account_file"`
	} `yaml:"firestore"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{
		ChannelSecret: viper.GetString("channel_secret"),
		ChannelToken:  viper.GetString("channel_token"),
	}

	return config, nil
}
