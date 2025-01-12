package pg

import (
	"context"
	"errors"

	"github.com/drybin/washington_changes_all/internal/domain/model"
	"github.com/drybin/washington_changes_all/internal/domain/types/coin_name"
	"github.com/drybin/washington_changes_all/pkg/wrap"
	washington_database "github.com/drybin/washington_changes_all/sqlc/gen"
	"github.com/jackc/pgx/v5"
)

type CoinAvgPricesRepository struct {
	Postgree *pgx.Conn
}

func NewCoinAvgPricesRepository(pg *pgx.Conn) CoinAvgPricesRepository {
	return CoinAvgPricesRepository{
		Postgree: pg,
	}
}

func (u CoinAvgPricesRepository) Get(ctx context.Context, coin model.Coin) (*float64, error) {
	db := washington_database.New(u.Postgree)
	price, err := db.GetCoinAvgPrice(ctx, string(coin.Name))

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, wrap.Errorf("failed to get Avg Price: %w", err)
	}
	return &price.Price, nil
}

func (u CoinAvgPricesRepository) GetList(ctx context.Context, coins []model.Coin) (*[]model.CoinPrice, error) {
	db := washington_database.New(u.Postgree)
	coinNames := make([]string, 0, len(coins))

	for _, coinName := range coins {
		coinNames = append(coinNames, string(coinName.Name))
	}

	prices, err := db.GetCoinAvgPrices(ctx, coinNames)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, wrap.Errorf("failed to get Avg Price: %w", err)
	}

	result := make([]model.CoinPrice, 0, len(coins))
	for _, price := range prices {
		result = append(result, mapToDomainPriceModel(price))
	}
	return &result, nil
}

func (u CoinAvgPricesRepository) GetAll(ctx context.Context) (*[]model.CoinPrice, error) {
	db := washington_database.New(u.Postgree)

	prices, err := db.GetCoinAvgPricesAll(ctx)

	if err != nil {
		return nil, wrap.Errorf("failed to get Avg Price all: %w", err)
	}

	result := make([]model.CoinPrice, 0, len(prices))
	for _, price := range prices {
		result = append(result, mapToDomainPriceModel(price))
	}
	return &result, nil
}

func (u CoinAvgPricesRepository) Save(ctx context.Context, coinPrice model.CoinPrice) error {
	db := washington_database.New(u.Postgree)

	_, err := db.SaveCoinAvgPricesAll(
		ctx,
		washington_database.SaveCoinAvgPricesAllParams{
			Coin:  coinPrice.Coin.Name.String(),
			Price: coinPrice.Price,
		},
	)

	if err != nil {
		return wrap.Errorf("failed to save coin avg price: %w", err)
	}

	return nil
}

func mapToDomainPriceModel(m washington_database.CoinAvgPrice) model.CoinPrice {
	return model.CoinPrice{
		Coin: model.Coin{
			Name: coin_name.CoinName(m.Coin),
		},
		Price: m.Price,
	}
}
