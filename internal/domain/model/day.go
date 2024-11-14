package model

import "time"

type Day struct {
	ID                int
	Date              time.Time
	AccountBalance    float64
	OverallAmountUsdt float64
	OverallCoinCount  int
	TierName          string
	CoinToBuy         string
}
