package buy_strategy

import (
	"github.com/drybin/washington_changes_all/internal/app/cli/config"
	"github.com/drybin/washington_changes_all/internal/domain/model"
)

func MapConfigCoinToModel(c config.Coin) model.Coin {
	return model.Coin{
		Name: c.Name,
	}
}

func MapConfigCoinArrayToModelArray(c []config.Coin) []model.Coin {
	result := make([]model.Coin, 0, len(c))
	for _, i := range c {
		result = append(result, MapConfigCoinToModel(i))
	}
	return result
}
