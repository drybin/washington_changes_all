package repo

import (
	"context"

	"github.com/drybin/washington_changes_all/internal/domain/model"
)

type ISellLogRepository interface {
	Save(ctx context.Context, dayId int, coin model.Coin, amount float64, price float64) error
}
