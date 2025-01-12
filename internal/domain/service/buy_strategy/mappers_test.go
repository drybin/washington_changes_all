package buy_strategy

import (
	"reflect"
	"testing"

	"github.com/drybin/washington_changes_all/internal/app/cli/config"
	"github.com/drybin/washington_changes_all/internal/domain/model"
)

func TestMapConfigCoinToModel(t *testing.T) {
	type args struct {
		c config.Coin
	}
	tests := []struct {
		name string
		args args
		want model.Coin
	}{
		{
			name: "test Map",
			args: args{
				c: config.Coin{
					Name: "BTC",
					ATH:  100.2,
				},
			},
			want: model.Coin{
				Name: "BTC",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapConfigCoinToModel(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapConfigCoinToModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapConfigCoinArrayToModelArray(t *testing.T) {
	type args struct {
		c []config.Coin
	}
	tests := []struct {
		name string
		args args
		want []model.Coin
	}{
		{
			name: "test Map empty",
			args: args{
				c: []config.Coin{},
			},
			want: []model.Coin{},
		},
		{
			name: "test Map",
			args: args{
				c: []config.Coin{
					{
						Name: "BTC",
						ATH:  100.2,
					},
				},
			},
			want: []model.Coin{
				{
					Name: "BTC",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapConfigCoinArrayToModelArray(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapConfigCoinArrayToModelArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
