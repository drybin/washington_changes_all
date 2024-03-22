package config

import (
	"errors"
	"github.com/drybin/washington_changes_all/pkg/wrap"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/joho/godotenv/cmd/godotenv/pkg/env"
	"github.com/ztrue/tracerr"
)

type Config struct {
	ServiceName string
	TgConfig    *TgConfig
}

type TgConfig struct {
	BotToken string
}

func (c TgConfig) Validate() error {

	return validation.ValidateStruct(&c,
		validation.Field(&c.BotToken, validation.Required),
	)
}

func (c Config) Validate() error {
	var errs []error

	err := validation.ValidateStruct(&c,
		validation.Field(&c.ServiceName, validation.Required),
		validation.Field(&c.TgConfig, validation.Required),
	)
	if err != nil {
		return wrap.Errorf("failed to validate cli config: %w", err)
	}

	if err := c.TgConfig.Validate(); err != nil {
		errs = append(errs, tracerr.Errorf("failed to validate Server config: \n%w", err))
	}

	return errors.Join(errs...)
}

func initTgConfig() *TgConfig {
	return &TgConfig{
		BotToken: env.GetString("TG_BOT_TOKEN", ""),
	}
}

func InitConfig() (*Config, error) {
	config := Config{
		ServiceName: env.GetString("APP_NAME", "WashingtonChangesAll"),
		TgConfig:    initTgConfig(),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil
}
