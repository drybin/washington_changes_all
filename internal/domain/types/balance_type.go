package types

type BalanceType string

const (
	Main  BalanceType = "main"
	Trade BalanceType = "trade"
)

func (t BalanceType) String() string {
	return string(t)
}
