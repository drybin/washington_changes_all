package types

type Tier string

const (
	TierOne   Tier = "TierOne"
	TierTwo   Tier = "TierTwo"
	TierThree Tier = "TierThree"
)

func (t Tier) String() string {
	return string(t)
}
