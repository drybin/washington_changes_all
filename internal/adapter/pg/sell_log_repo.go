package pg

import (
	"context"

	"github.com/drybin/washington_changes_all/internal/domain/model"
	"github.com/drybin/washington_changes_all/pkg/wrap"
	washington_database "github.com/drybin/washington_changes_all/sqlc/gen"
	"github.com/jackc/pgx/v5"
)

type SellLogRepository struct {
	Postgree *pgx.Conn
}

func NewSellLogRepository(pg *pgx.Conn) SellLogRepository {
	return SellLogRepository{
		Postgree: pg,
	}
}

func (u SellLogRepository) Save(ctx context.Context, dayId int, coin model.Coin, amount float64, price float64) error {
	db := washington_database.New(u.Postgree)
	_, err := db.SaveSellLog(
		ctx,
		washington_database.SaveSellLogParams{
			DayID:  dayId,
			Coin:   coin.Name.String(),
			Amount: amount,
			Price:  price,
		},
	)

	if err != nil {
		return wrap.Errorf("failed to save buy log: %w", err)
	}

	return nil
}
