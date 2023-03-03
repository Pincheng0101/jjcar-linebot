package config

import (
	"github.com/spf13/viper"
)

var configCache *Config = nil

type Firebase struct {
	ServiceAccountFile string `yaml:"service_account_file"`
	StorageBucket      string `yaml:"storage_bucket"`
}

type Config struct {
	ChannelSecret string   `yaml:"channel_secret"`
	ChannelToken  string   `yaml:"channel_token"`
	Firebase      Firebase `yaml:"firebase"`
}

func LoadConfig() (*Config, error) {
	if configCache != nil {
		return configCache, nil
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	config := &Config{
		ChannelSecret: viper.GetString("channel_secret"),
		ChannelToken:  viper.GetString("channel_token"),
		Firebase: Firebase{
			ServiceAccountFile: viper.GetString("firebase.service_account_file"),
			StorageBucket:      viper.GetString("firebase.storage_bucket"),
		},
	}

	return config, nil
}
