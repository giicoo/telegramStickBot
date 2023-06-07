package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Token    string
	PathIn   string `mapstructure:"path_in"`
	PathOut  string `mapstructure:"path_out"`
	Messages Messages
}

type Messages struct {
	Responses `mapstructure:"responses"`
	Errors    `mapstructure:"errors"`
}

type Responses struct {
	Start          string `mapstructure:"start"`
	Default        string `mapstructure:"default"`
	DefaultCommand string `mapstructure:"default_command"`
}

type Errors struct {
	SendNotFile string `mapstructure:"send_not_file"`
}

func InitConfig() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.response", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.error", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	return &cfg, nil
}
