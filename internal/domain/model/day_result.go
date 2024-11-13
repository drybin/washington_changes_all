package model

import (
	"github.com/drybin/washington_changes_all/internal/domain/types/coin_name"
)

type DayResult struct {
	Balance float64

	Name             coin_name.CoinName
	Amount           float64
	Price            float64
	CoinCount        int
	PrevCoinAvgPrice float64
	CoinAvgPrice     float64
	DayNumber        int
	CoinAmount       float64
}
