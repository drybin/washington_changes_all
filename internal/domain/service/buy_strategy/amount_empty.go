package buy_strategy

import (
	"context"
	"sort"

	"github.com/drybin/washington_changes_all/internal/app/cli/config"
	"github.com/drybin/washington_changes_all/internal/domain/model"
	"github.com/drybin/washington_changes_all/internal/domain/repo"
	"github.com/drybin/washington_changes_all/internal/domain/types/coin_name"
	"github.com/drybin/washington_changes_all/pkg/wrap"
)

type AmountEmptyStrategy struct {
	Repo       repo.ICoinAvgPricesRepository
	CoinConfig config.CoinConfig
}

type IAmountEmptyStrategy interface {
	Process(ctx context.Context, coins []config.Coin, marketsHistory *[]model.MarketInfo) ([]model.CoinPriceChange, error)
}

func NewAmountEmptyStrategy(
	repo repo.ICoinAvgPricesRepository,
	CoinConfig config.CoinConfig,
) IAmountEmptyStrategy {
	return &AmountEmptyStrategy{Repo: repo, CoinConfig: CoinConfig}
}

func (s *AmountEmptyStrategy) Process(
	ctx context.Context,
	coins []config.Coin,
	marketsHistory *[]model.MarketInfo,
) ([]model.CoinPriceChange, error) {

	coinAvgPrices, err := s.Repo.GetList(ctx, MapConfigCoinArrayToModelArray(coins))
	if err != nil {
		return nil, wrap.Errorf("failed to get coins avg prices: %w", err)
	}

	coinResult := getCoinsWithEmptyAvgPrice(coinAvgPrices, coins)

	athResult := calculateGlobalAthChange(coinResult, marketsHistory)

	athResult = sortAthResultAsc(athResult)

	return athResult, nil
}

func getCoinsWithEmptyAvgPrice(coinAvgPrices *[]model.CoinPrice, coins []config.Coin) []config.Coin {
	coinResult := make([]config.Coin, 0, len(coins))

	for _, coin := range coins {
		found := false
		for _, coinAvgPrice := range *coinAvgPrices {
			if coin.Name == coinAvgPrice.Coin.Name {
				found = true
			}
		}
		if !found {
			coinResult = append(coinResult, coin)
		}
	}

	return coinResult
}

func calculateGlobalAthChange(coinResult []config.Coin, marketsHistory *[]model.MarketInfo) []model.CoinPriceChange {
	athResult := make([]model.CoinPriceChange, 0, len(coinResult))

	for _, coin := range coinResult {
		for _, market := range *marketsHistory {
			if market.Pair.CoinFirst.Name == coin.Name && market.Pair.CoinSecond.Name == coin_name.USDT {
				percent := coin.ATH / 100.0
				athChange := market.LastPrice / percent
				athResult = append(
					athResult,
					model.CoinPriceChange{
						Coin:        model.Coin{Name: coin.Name},
						PriceChange: athChange,
					},
				)
			}
		}
	}
	return athResult
}

func sortAthResultAsc(athResult []model.CoinPriceChange) []model.CoinPriceChange {
	sort.Slice(athResult, func(i, j int) bool {
		return athResult[i].PriceChange < athResult[j].PriceChange
	})

	return athResult
}
