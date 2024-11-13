package pg

import (
	"context"
	"github.com/drybin/washington_changes_all/internal/domain/model"
	"github.com/drybin/washington_changes_all/pkg/wrap"
	washington_database "github.com/drybin/washington_changes_all/sqlc/gen"
	"github.com/jackc/pgx/v5"
	"time"
)

type MarketsHistoryRepository struct {
	Postgree *pgx.Conn
}

func NewMarketsHistoryRepository(pg *pgx.Conn) MarketsHistoryRepository {
	return MarketsHistoryRepository{
		Postgree: pg,
	}
}

func (u MarketsHistoryRepository) SaveMany(ctx context.Context, marketsInfo *[]model.MarketInfo) error {
	for _, item := range *marketsInfo {
		err := u.Save(ctx, item)
		if err != nil {
			return wrap.Errorf("failed to save markets history: %w", err)
		}
	}
	return nil
}

func (u MarketsHistoryRepository) Save(ctx context.Context, marketInfo model.MarketInfo) error {
	db := washington_database.New(u.Postgree)

	err := db.SaveMarketInfo(
		ctx,
		mapToRepoModel(marketInfo),
	)

	if err != nil {
		return wrap.Errorf("repo: failed to save market history: %w", err)
	}

	return nil
}

func mapToRepoModel(m model.MarketInfo) washington_database.SaveMarketInfoParams {
	return washington_database.SaveMarketInfoParams{
		Date:             time.Now(),
		CoinFirst:        m.Pair.CoinFirst.Name.String(),
		CoinSecond:       m.Pair.CoinSecond.Name.String(),
		SymbolName:       m.SymbolName,
		Buy:              m.Buy,
		Sell:             m.Sell,
		ChangeRate:       m.ChangeRate,
		ChangePrice:      m.ChangePrice,
		HighPrice:        m.HighPrice,
		LowPrice:         m.LowPrice,
		VolBtc:           m.VolBTC,
		VolValue:         m.VolValue,
		LastPrice:        m.LastPrice,
		AveragePrice:     m.AveragePrice,
		TakerFeeRate:     m.TakerFeeRate,
		MakerFeeRate:     m.MakerFeeRate,
		TakerCoefficient: m.TakerCoefficient,
		MakerCoefficient: m.MakerCoefficient,
	}
}
