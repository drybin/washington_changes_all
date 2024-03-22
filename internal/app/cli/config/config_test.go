package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	type fields struct {
		ServiceName string
		TgConfig    *TgConfig
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
				TgConfig: &TgConfig{
					BotToken: "some_token",
				},
			},
			wantErr: false,
		},
		{
			name: "TgConfig not setted, error",
			fields: fields{
				ServiceName: "name",
				TgConfig:    nil,
			},
			wantErr: true,
		},
		{
			name: "Service name is empty, error",
			fields: fields{
				ServiceName: "",
				TgConfig:    nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := Config{
				ServiceName: tt.fields.ServiceName,
				TgConfig:    tt.fields.TgConfig,
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
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Token is setted, no error",
			fields: fields{
				BotToken: "some_token",
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
			},
			want: &TgConfig{
				BotToken: "some_token",
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
