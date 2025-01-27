package pg

import (
    "context"
    
    "github.com/drybin/washington_changes_all/internal/domain/model"
    "github.com/drybin/washington_changes_all/pkg/wrap"
    washington_database "github.com/drybin/washington_changes_all/sqlc/gen"
    "github.com/jackc/pgx/v5"
)

type CoinAmountRepository struct {
    Postgree *pgx.Conn
}

func NewCoinAmountRepository(pg *pgx.Conn) CoinAmountRepository {
    return CoinAmountRepository{
        Postgree: pg,
    }
}

func (u CoinAmountRepository) Get(ctx context.Context, coin model.Coin) (float64, error) {
    db := washington_database.New(u.Postgree)
    res, err := db.GetCoinAmount(
        ctx,
        coin.Name.String(),
    )
    
    if err != nil {
        if err == pgx.ErrNoRows {
            return 0.0, nil
        }
        return 0.0, wrap.Errorf("failed to get coin amount: %w", err)
    }
    
    return res.Amount, nil
}

func (u CoinAmountRepository) Save(ctx context.Context, coin model.Coin, amount float64) error {
    db := washington_database.New(u.Postgree)
    _, err := db.SaveCoinAmount(
        ctx,
        washington_database.SaveCoinAmountParams{
            Coin:   coin.Name.String(),
            Amount: amount,
        },
    )
    
    if err != nil {
        return wrap.Errorf("failed to save coin amount: %w", err)
    }
    
    return nil
}
