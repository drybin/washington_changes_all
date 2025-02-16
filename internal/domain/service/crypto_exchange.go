package service

import (
	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/drybin/washington_changes_all/internal/domain/model"
	"github.com/drybin/washington_changes_all/internal/domain/repo"
	"github.com/drybin/washington_changes_all/pkg/wrap"
)

type ICryptoExchangeService interface {
	GetBalance() (float64, error)
	GetMarketsHistory() (*[]model.MarketInfo, error)
	BuyByMarket(coin model.Coin) (*kucoin.OrderModel, error)
	SellByMarket(coin model.Coin, amount string) (*kucoin.OrderModel, error)
}

type CryptoExchangeServiceService struct {
	Repo repo.ICryptoExchangeRepository
}

func NewCryptoExchangeServiceService(
	repo repo.ICryptoExchangeRepository,
) ICryptoExchangeService {
	return &CryptoExchangeServiceService{repo}
}

func (s *CryptoExchangeServiceService) GetBalance() (float64, error) {
	balance, err := s.Repo.GetBalance()
	if err != nil {
		return 0.0, wrap.Errorf("failed to get last day info: %w", err)
	}

	return balance, nil
}

func (s *CryptoExchangeServiceService) GetMarketsHistory() (*[]model.MarketInfo, error) {
	t, err := s.Repo.GetMarketsHistory()
	if err != nil {
		return nil, wrap.Errorf("failed to get markets history: %w", err)
	}

	return t, nil
}

func (s *CryptoExchangeServiceService) BuyByMarket(coin model.Coin) (*kucoin.OrderModel, error) {
	res, err := s.Repo.BuyByMarket(coin)
	if err != nil {
		return nil, wrap.Errorf("failed to buy by market: %w", err)
	}

	return res, nil
}

func (s *CryptoExchangeServiceService) SellByMarket(coin model.Coin, amount string) (*kucoin.OrderModel, error) {
	res, err := s.Repo.SellByMarket(coin, amount)
	if err != nil {
		return nil, wrap.Errorf("failed to sell by market: %w", err)
	}

	return res, nil
}
