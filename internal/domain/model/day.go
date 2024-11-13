package model

import "time"

type Day struct {
	ID               int
	Date             time.Time
	AccountBalance   float64
	OveralAmountUsdt float64
	OveralCoinCount  int
	TierName         string
	CointToBuy       string
}
