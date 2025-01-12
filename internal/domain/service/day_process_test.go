package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/drybin/washington_changes_all/internal/app/cli/config"
	"github.com/drybin/washington_changes_all/internal/domain/model"
	"github.com/drybin/washington_changes_all/internal/domain/types"
	"github.com/drybin/washington_changes_all/internal/domain/types/coin_name"
)

func Test_getTierName(t *testing.T) {
	t.Parallel()

	type args struct {
		day *model.Day
	}
	tests := []struct {
		name string
		args args
		want types.Tier
	}{
		{
			name: "day is nil",
			args: args{
				day: nil,
			},
			want: types.TierOne,
		},
		{
			name: "day with one",
			args: args{
				day: &model.Day{
					TierName: "TierOne",
				},
			},
			want: types.TierTwo,
		},
		{
			name: "day with two",
			args: args{
				day: &model.Day{
					TierName: "TierTwo",
				},
			},
			want: types.TierThree,
		},
		{
			name: "day with three",
			args: args{
				day: &model.Day{
					TierName: "TierThree",
				},
			},
			want: types.TierOne,
		},
		{
			name: "day unknown",
			args: args{
				day: &model.Day{
					TierName: "unknown",
				},
			},
			want: types.TierOne,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := getTierName(tt.args.day); got != tt.want {
				t.Errorf("getTierName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTierCoins(t *testing.T) {
	t.Parallel()

	type args struct {
		tierName types.Tier
		config   config.CoinConfig
	}
	tests := []struct {
		name string
		args args
		want []config.Coin
	}{
		{
			name: "get Tier One",
			args: args{
				tierName: types.TierOne,
				config: config.CoinConfig{
					TierOne: []config.Coin{
						{
							Name: coin_name.CoinName("BTC"),
							ATH:  100.2,
						},
					},
					TierTwo: []config.Coin{
						{
							Name: coin_name.CoinName("ADA"),
							ATH:  2.32,
						},
					},
					TierThree: []config.Coin{
						{
							Name: coin_name.CoinName("PEPE"),
							ATH:  0.006,
						},
					},
				},
			},
			want: []config.Coin{
				{
					Name: coin_name.CoinName("BTC"),
					ATH:  100.2,
				},
			},
		},
		{
			name: "get Tier Two",
			args: args{
				tierName: types.TierTwo,
				config: config.CoinConfig{
					TierOne: []config.Coin{
						{
							Name: coin_name.CoinName("BTC"),
							ATH:  100.2,
						},
					},
					TierTwo: []config.Coin{
						{
							Name: coin_name.CoinName("ADA"),
							ATH:  2.32,
						},
					},
					TierThree: []config.Coin{
						{
							Name: coin_name.CoinName("PEPE"),
							ATH:  0.006,
						},
					},
				},
			},
			want: []config.Coin{
				{
					Name: coin_name.CoinName("ADA"),
					ATH:  2.32,
				},
			},
		},
		{
			name: "get Tier Three",
			args: args{
				tierName: types.TierThree,
				config: config.CoinConfig{
					TierOne: []config.Coin{
						{
							Name: coin_name.CoinName("BTC"),
							ATH:  100.2,
						},
					},
					TierTwo: []config.Coin{
						{
							Name: coin_name.CoinName("ADA"),
							ATH:  2.32,
						},
					},
					TierThree: []config.Coin{
						{
							Name: coin_name.CoinName("PEPE"),
							ATH:  0.006,
						},
					},
				},
			},
			want: []config.Coin{
				{
					Name: coin_name.CoinName("PEPE"),
					ATH:  0.006,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, getTierCoins(tt.args.tierName, tt.args.config))
		})
	}
}
