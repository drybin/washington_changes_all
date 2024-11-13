package config

import (
	"errors"
	"github.com/drybin/washington_changes_all/pkg/env"
	"github.com/drybin/washington_changes_all/pkg/wrap"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ztrue/tracerr"
	"time"
)

type Config struct {
	ServiceName  string
	PostgreeDsn  string
	TgConfig     *TgConfig
	KucoinConfig *KucoinConfig
	CoinConfig   CoinConfig
}

type TgConfig struct {
	BotToken string
	ChatId   string
	Timeout  time.Duration
}

type KucoinConfig struct {
	Key        string
	Secret     string
	Passphrase string
}

func (c KucoinConfig) Validate() error {

	return validation.ValidateStruct(&c,
		validation.Field(&c.Key, validation.Required),
		validation.Field(&c.Secret, validation.Required),
		validation.Field(&c.Passphrase, validation.Required),
	)
}

func (c TgConfig) Validate() error {

	return validation.ValidateStruct(&c,
		validation.Field(&c.BotToken, validation.Required),
		validation.Field(&c.ChatId, validation.Required),
	)
}

func (c Config) Validate() error {
	var errs []error

	err := validation.ValidateStruct(&c,
		validation.Field(&c.ServiceName, validation.Required),
		validation.Field(&c.PostgreeDsn, validation.Required),
		validation.Field(&c.TgConfig, validation.Required),
		validation.Field(&c.KucoinConfig, validation.Required),
	)
	if err != nil {
		return wrap.Errorf("failed to validate cli config: %w", err)
	}

	if err := c.TgConfig.Validate(); err != nil {
		errs = append(errs, tracerr.Errorf("failed to validate Telegram config: \n%w", err))
	}

	if err := c.KucoinConfig.Validate(); err != nil {
		errs = append(errs, tracerr.Errorf("failed to validate Kucoin config: \n%w", err))
	}

	return errors.Join(errs...)
}

func initTgConfig() *TgConfig {
	return &TgConfig{
		BotToken: env.GetString("TG_BOT_TOKEN", ""),
		ChatId:   env.GetString("TG_CHAT_ID", ""),
		Timeout:  10 * time.Second,
	}
}

func initKucoinConfig() *KucoinConfig {
	return &KucoinConfig{
		Key:        env.GetString("KUCOIN_KEY", ""),
		Secret:     env.GetString("KUCOIN_SECRET", ""),
		Passphrase: env.GetString("KUCOIN_PASSPHRASE", ""),
	}
}

func InitConfig() (*Config, error) {
	config := Config{
		ServiceName:  env.GetString("APP_NAME", "WashingtonChangesAll"),
		PostgreeDsn:  env.GetString("POSTGREE_DSN", ""),
		TgConfig:     initTgConfig(),
		KucoinConfig: initKucoinConfig(),
		CoinConfig:   InitCoinConfig(),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil
}
