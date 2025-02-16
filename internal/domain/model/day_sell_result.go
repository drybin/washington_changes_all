package model

import (
	"github.com/drybin/washington_changes_all/internal/domain/types/coin_name"
)

type DaySellResult struct {
	Balance float64
	DayInfo Day

	Coins []DaySellCoinInfo
}

type DaySellCoinInfo struct {
	CoinName         coin_name.CoinName
	CoinCurrentPrice float64
	CoinAvgPrice     float64
	CoinPrevAmount   float64
	CoinNewAmount    float64
	SellUsdtAmount   float64
}
