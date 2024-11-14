package pg

import (
	"context"
	
	"github.com/drybin/washington_changes_all/internal/domain/model"
	"github.com/drybin/washington_changes_all/pkg/wrap"
	washington_database "github.com/drybin/washington_changes_all/sqlc/gen"
	"github.com/jackc/pgx/v5"
)

type BuyLogRepository struct {
	Postgree *pgx.Conn
}

func NewBuyLogRepository(pg *pgx.Conn) BuyLogRepository {
	return BuyLogRepository{
		Postgree: pg,
	}
}

func (u BuyLogRepository) Save(ctx context.Context, dayId int, coin model.Coin, amount float64, price float64) error {
	db := washington_database.New(u.Postgree)
	_, err := db.SaveBuyLog(
		ctx,
		washington_database.SaveBuyLogParams{
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
