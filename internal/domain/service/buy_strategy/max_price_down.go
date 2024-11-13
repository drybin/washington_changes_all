package buy_strategy

import (
	"context"
	"github.com/drybin/washington_changes_all/internal/app/cli/config"
	"github.com/drybin/washington_changes_all/internal/domain/model"
	"github.com/drybin/washington_changes_all/internal/domain/repo"
	"github.com/drybin/washington_changes_all/internal/domain/types/coin_name"
	"github.com/drybin/washington_changes_all/pkg/wrap"
	"sort"
)

type MaxPriceDownStrategy struct {
	Repo       repo.ICoinAvgPricesRepository
	CoinConfig config.CoinConfig
}

type IMaxPriceDownStrategy interface {
	Process(ctx context.Context, coins []config.Coin, marketsHistory *[]model.MarketInfo) ([]model.CoinPriceChange, error)
}

func NewMaxPriceDownStrategy(
	repo repo.ICoinAvgPricesRepository,
	CoinConfig config.CoinConfig,
) IMaxPriceDownStrategy {
	return &MaxPriceDownStrategy{Repo: repo, CoinConfig: CoinConfig}
}

func (s *MaxPriceDownStrategy) Process(
	ctx context.Context,
	coins []config.Coin,
	marketsHistory *[]model.MarketInfo,
) ([]model.CoinPriceChange, error) {

	coinAvgPrices, err := s.Repo.GetList(ctx, MapConfigCoinArrayToModelArray(coins))
	if err != nil {
		return nil, wrap.Errorf("failed to get coins avg prices: %w", err)
	}

	result := calculateAvgPriceChange(coinAvgPrices, marketsHistory)

	result = sortAthResultDesc(result)

	return result, nil
}

func calculateAvgPriceChange(coinResult *[]model.CoinPrice, marketsHistory *[]model.MarketInfo) []model.CoinPriceChange {
	priceResult := make([]model.CoinPriceChange, 0, len(*coinResult))

	for _, coin := range *coinResult {
		for _, market := range *marketsHistory {
			if market.Pair.CoinFirst.Name == coin.Coin.Name && market.Pair.CoinSecond.Name == coin_name.USDT {
				percent := coin.Price / 100.0
				priceChange := market.LastPrice / percent
				priceResult = append(
					priceResult,
					model.CoinPriceChange{
						Coin:        model.Coin{Name: coin.Coin.Name},
						PriceChange: 100.0 - priceChange,
					},
				)
			}
		}
	}
	return priceResult
}

func sortAthResultDesc(athResult []model.CoinPriceChange) []model.CoinPriceChange {
	sort.Slice(athResult, func(i, j int) bool {
		return athResult[i].PriceChange > athResult[j].PriceChange
	})

	return athResult
}
