package repo

import (
    "github.com/Kucoin/kucoin-go-sdk"
    "github.com/drybin/washington_changes_all/internal/domain/model"
)

type ICryptoExchangeRepository interface {
    GetBalance() (float64, error)
    GetMarketsHistory() (*[]model.MarketInfo, error)
    BuyByMarket(coin model.Coin) (*kucoin.OrderModel, error)
    SellByMarket(coin model.Coin, amount string) (*kucoin.OrderModel, error)
    GetPairInfo(pair model.CoinsPair) (*model.PairInfo, error)
}
