package model

import "github.com/drybin/washington_changes_all/internal/domain/types"

type Balance struct {
	BalanceType types.BalanceType
	Coin        Coin
	Balance     float64
}
