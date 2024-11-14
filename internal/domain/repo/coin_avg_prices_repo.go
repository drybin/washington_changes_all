package repo

import (
	"context"
	
	"github.com/drybin/washington_changes_all/internal/domain/model"
)

type ICoinAvgPricesRepository interface {
	Get(ctx context.Context, coin model.Coin) (*float64, error)
	GetList(ctx context.Context, coin []model.Coin) (*[]model.CoinPrice, error)
	GetAll(ctx context.Context) (*[]model.CoinPrice, error)
	Save(ctx context.Context, coinPrice model.CoinPrice) error
}
