package model

import (
	"github.com/drybin/washington_changes_all/internal/domain/types"
	"github.com/drybin/washington_changes_all/internal/domain/types/coin_name"
)

type DayResult struct {
	Balance float64
	PrevDay Day

	Tier                         types.Tier
	AmountEmptyStrategy          bool
	AmountEmptyStrategyCoinCount int
	CoinName                     coin_name.CoinName
	CoinPriceChange              float64
	CoinAth                      float64
	Amount                       float64
	Price                        float64
	CoinCount                    int
	PrevCoinAvgPrice             float64
	CoinAvgPrice                 float64
	DayNumber                    int
	CoinAmount                   float64
}
