package repo

import (
    "context"
    "github.com/Kucoin/kucoin-go-sdk"
    "github.com/drybin/washington_changes_all/internal/domain/model"
)

type ICryptoExchangeRepository interface {
    GetBalance(ctx context.Context) (float64, error)
    GetMarketsHistory(ctx context.Context) (*[]model.MarketInfo, error)
    BuyByMarket(ctx context.Context, coin model.Coin) (*kucoin.OrderModel, error)
    SellByMarket(ctx context.Context, coin model.Coin, amount string) (*kucoin.OrderModel, error)
    GetPairInfo(ctx context.Context, pair model.CoinsPair) (*model.PairInfo, error)
}
