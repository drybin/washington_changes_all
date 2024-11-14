package repo

import (
	"context"

	"github.com/drybin/washington_changes_all/internal/domain/model"
)

type ICoinAmountRepository interface {
	Get(ctx context.Context, coin model.Coin) (float64, error)
	Save(ctx context.Context, coin model.Coin, amount float64) error
}
