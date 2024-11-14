package pg

import (
	"context"
	"errors"
	
	"github.com/drybin/washington_changes_all/internal/domain/model"
	"github.com/drybin/washington_changes_all/pkg/wrap"
	washington_database "github.com/drybin/washington_changes_all/sqlc/gen"
	"github.com/jackc/pgx/v5"
)

type DaysRepository struct {
    Postgree *pgx.Conn
}

func NewDaysRepository(pg *pgx.Conn) DaysRepository {
    return DaysRepository{
        Postgree: pg,
    }
}

func (u DaysRepository) GetLastDayInfo(ctx context.Context) (*model.Day, error) {
    db := washington_database.New(u.Postgree)
    day, err := db.GetLastDayInfo(ctx)
    
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return nil, nil
        }
        return nil, wrap.Errorf("failed to get Last Day Info From Postgree: %w", err)
    }
    return mapToDomainModel(day), nil
}

func (u DaysRepository) Save(ctx context.Context, day model.Day) (*model.Day, error) {
    db := washington_database.New(u.Postgree)
    res, err := db.SaveDayInfo(ctx, mapToSaveDayInfoParams(day))
    
    if err != nil {
        return nil, wrap.Errorf("failed to save Last Day Info From Postgree: %w", err)
    }
    
    return mapToDomainModel(res), nil
}

func mapToDomainModel(m washington_database.Day) *model.Day {
    return &model.Day{
        ID:                m.ID,
        Date:              m.Date,
        AccountBalance:    m.AccountBalance,
        OverallAmountUsdt: m.OveralAmountUsdt,
        OverallCoinCount:  m.OveralCoinCount,
        TierName:          m.Tiername,
        CoinToBuy:         m.CoinToBuy,
    }
}

func mapToSaveDayInfoParams(day model.Day) washington_database.SaveDayInfoParams {
    return washington_database.SaveDayInfoParams{
        Date:             day.Date,
        AccountBalance:   day.AccountBalance,
        OveralAmountUsdt: day.OverallAmountUsdt,
        OveralCoinCount:  day.OverallCoinCount,
        Tiername:         day.TierName,
        CoinToBuy:        day.CoinToBuy,
    }
}
