package config

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	type fields struct {
		ServiceName  string
		PostgreeDsn  string
		TgConfig     *TgConfig
		KucoinConfig *KucoinConfig
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "All config is setted, no error",
			fields: fields{
				ServiceName: "name",
				PostgreeDsn: "dsn",
				TgConfig: &TgConfig{
					BotToken: "some_token",
					ChatId:   "some_chat_id",
					Timeout:  10 * time.Second,
				},
				KucoinConfig: &KucoinConfig{
					Key:        "key",
					Secret:     "secret",
					Passphrase: "passpharse",
				},
			},
			wantErr: false,
		},
		{
			name: "TgConfig not setted, error",
			fields: fields{
				ServiceName:  "name",
				TgConfig:     nil,
				KucoinConfig: nil,
			},
			wantErr: true,
		},
		{
			name: "Kukoin not setted, error",
			fields: fields{
				ServiceName: "name",
				TgConfig: &TgConfig{
					BotToken: "some_token",
					ChatId:   "some_chat_id",
					Timeout:  10 * time.Second,
				},
				KucoinConfig: nil,
			},
			wantErr: true,
		},
		{
			name: "Service name is empty, error",
			fields: fields{
				ServiceName:  "",
				TgConfig:     nil,
				KucoinConfig: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := Config{
				ServiceName:  tt.fields.ServiceName,
				PostgreeDsn:  tt.fields.PostgreeDsn,
				TgConfig:     tt.fields.TgConfig,
				KucoinConfig: tt.fields.KucoinConfig,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTgConfig_Validate(t *testing.T) {
	t.Parallel()

	type fields struct {
		BotToken string
		ChatId   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Token,chatId is setted, no error",
			fields: fields{
				BotToken: "some_token",
				ChatId:   "some_chat_id",
			},
			wantErr: false,
		},
		{
			name:    "Token not setted, error",
			fields:  fields{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := TgConfig{
				BotToken: tt.fields.BotToken,
				ChatId:   tt.fields.ChatId,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_initTgConfig(t *testing.T) {
	type envArgs struct {
		name  string
		value string
	}

	tests := []struct {
		name    string
		envArgs []envArgs
		want    *TgConfig
	}{
		{
			name: "test TgConfig",
			envArgs: []envArgs{
				{
					name:  "TG_BOT_TOKEN",
					value: "some_token",
				},
				{
					name:  "TG_CHAT_ID",
					value: "some_chat_id",
				},
			},
			want: &TgConfig{
				BotToken: "some_token",
				ChatId:   "some_chat_id",
				Timeout:  10 * time.Second,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, env := range tt.envArgs {
				t.Setenv(env.name, env.value)
			}

			require.Equal(t, tt.want, initTgConfig())
		})
	}
}

func TestKucoinConfig_Validate(t *testing.T) {
	t.Parallel()

	type fields struct {
		Key        string
		Secret     string
		Passphrase string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "All setted, no error",
			fields: fields{
				Key:        "key",
				Secret:     "secret",
				Passphrase: "passphrase",
			},
			wantErr: false,
		},
		{
			name: "Key not setted, error",
			fields: fields{
				Key:        "",
				Secret:     "secret",
				Passphrase: "passphrase",
			},
			wantErr: true,
		},
		{
			name: "Secret not setted, error",
			fields: fields{
				Key:        "key",
				Secret:     "",
				Passphrase: "passphrase",
			},
			wantErr: true,
		},
		{
			name: "Passphrase not setted, error",
			fields: fields{
				Key:        "key",
				Secret:     "secret",
				Passphrase: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := KucoinConfig{
				Key:        tt.fields.Key,
				Secret:     tt.fields.Secret,
				Passphrase: tt.fields.Passphrase,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_initKucoinConfig(t *testing.T) {
	type envArgs struct {
		name  string
		value string
	}

	tests := []struct {
		name    string
		envArgs []envArgs
		want    *KucoinConfig
	}{
		{
			name: "test TgConfig",
			envArgs: []envArgs{
				{
					name:  "KUCOIN_KEY",
					value: "key",
				},
				{
					name:  "KUCOIN_SECRET",
					value: "secret",
				},
				{
					name:  "KUCOIN_PASSPHRASE",
					value: "passphrase",
				},
			},
			want: &KucoinConfig{
				Key:        "key",
				Secret:     "secret",
				Passphrase: "passphrase",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, env := range tt.envArgs {
				t.Setenv(env.name, env.value)
			}

			require.Equal(t, tt.want, initKucoinConfig())
		})
	}
}
