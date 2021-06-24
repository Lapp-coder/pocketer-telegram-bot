package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TelegramBotToken  string
	PocketConsumerKey string
	DBPath            string `mapstructure:"db_path"`
	RedirectURL       string `mapstructure:"redirect_url"`
	Messages          Messages
}

type Messages struct {
	Responses
	Errors
}

type Responses struct {
	Start             string `mapstructure:"start"`
	AlreadyAuthorized string `mapstructure:"already_authorized"`
	SavedSuccessfully string `mapstructure:"saved_successfully"`
	UnknownCommand    string `mapstructure:"unknown_command"`
}

type Errors struct {
	Default            string `mapstructure:"default"`
	InvalidURL         string `mapstructure:"invalid_url"`
	Unauthorized       string `mapstructure:"unauthorized"`
	FailedToSave       string `mapstructure:"failed_to_save"`
	FailedToAuthorized string `mapstructure:"failed_to_authorized"`
	FailedToGenerate   string `mapstructure:"failed_to_generate"`
}

func NewConfig() (*Config, error) {
	viper.AddConfigPath("configs/")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	if err := LoadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func LoadEnv(cfg *Config) error {
	if err := viper.BindEnv("TELEGRAM_BOT_TOKEN"); err != nil {
		return err
	}

	if err := viper.BindEnv("CONSUMER_KEY"); err != nil {
		return err
	}

	cfg.TelegramBotToken = viper.GetString("TELEGRAM_BOT_TOKEN")
	cfg.PocketConsumerKey = viper.GetString("CONSUMER_KEY")

	return nil
}
