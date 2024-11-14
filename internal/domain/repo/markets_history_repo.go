package repo

import (
	"context"

	"github.com/drybin/washington_changes_all/internal/domain/model"
)

type IMarketsHistoryRepository interface {
	Save(ctx context.Context, marketInfo model.MarketInfo) error
	SaveMany(ctx context.Context, marketsInfo *[]model.MarketInfo) error
	//Save(ctx context.Context, day model.Day) error
}
