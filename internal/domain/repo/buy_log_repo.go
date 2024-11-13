package repo

import (
	"context"
	"github.com/drybin/washington_changes_all/internal/domain/model"
)

type IBuyLogRepository interface {
	//Get(ctx context.Context, coin config.Coin) (*float64, error)
	Save(ctx context.Context, dayId int, coin model.Coin, amount float64, price float64) error
}
