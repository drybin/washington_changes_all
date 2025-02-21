package service

import (
    "context"
    
    "github.com/Kucoin/kucoin-go-sdk"
    "github.com/drybin/washington_changes_all/internal/domain/model"
    "github.com/drybin/washington_changes_all/internal/domain/repo"
    "github.com/drybin/washington_changes_all/pkg/wrap"
)

type ICryptoExchangeService interface {
    GetBalance(ctx context.Context) (float64, error)
    GetMarketsHistory(ctx context.Context) (*[]model.MarketInfo, error)
    BuyByMarket(ctx context.Context, coin model.Coin) (*kucoin.OrderModel, error)
    SellByMarket(ctx context.Context, coin model.Coin, amount string) (*kucoin.OrderModel, error)
    GetPairInfo(ctx context.Context, pair model.CoinsPair) (*model.PairInfo, error)
}

type CryptoExchangeServiceService struct {
    Repo repo.ICryptoExchangeRepository
}

func NewCryptoExchangeServiceService(
    repo repo.ICryptoExchangeRepository,
) ICryptoExchangeService {
    return &CryptoExchangeServiceService{repo}
}

func (s *CryptoExchangeServiceService) GetBalance(ctx context.Context) (float64, error) {
    balance, err := s.Repo.GetBalance(ctx)
    if err != nil {
        return 0.0, wrap.Errorf("failed to get last day info: %w", err)
    }
    
    return balance, nil
}

func (s *CryptoExchangeServiceService) GetMarketsHistory(ctx context.Context) (*[]model.MarketInfo, error) {
    t, err := s.Repo.GetMarketsHistory(ctx)
    if err != nil {
        return nil, wrap.Errorf("failed to get markets history: %w", err)
    }
    
    return t, nil
}

func (s *CryptoExchangeServiceService) BuyByMarket(ctx context.Context, coin model.Coin) (*kucoin.OrderModel, error) {
    res, err := s.Repo.BuyByMarket(ctx, coin)
    if err != nil {
        return nil, wrap.Errorf("failed to buy by market: %w", err)
    }
    
    return res, nil
}

func (s *CryptoExchangeServiceService) SellByMarket(ctx context.Context, coin model.Coin, amount string) (*kucoin.OrderModel, error) {
    res, err := s.Repo.SellByMarket(ctx, coin, amount)
    if err != nil {
        return nil, wrap.Errorf("failed to sell by market: %w", err)
    }
    
    return res, nil
}

func (s *CryptoExchangeServiceService) GetPairInfo(ctx context.Context, pair model.CoinsPair) (*model.PairInfo, error) {
    res, err := s.Repo.GetPairInfo(ctx, pair)
    if err != nil {
        return nil, wrap.Errorf("failed to get pair info: %w", err)
    }
    
    return res, nil
}
