package service

import (
	"testing"

	"github.com/drybin/washington_changes_all/internal/domain/model"
	"github.com/drybin/washington_changes_all/internal/domain/types"
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
