package usecase

import (
    "context"
    "fmt"
    
    "github.com/drybin/washington_changes_all/internal/domain/service"
    "github.com/drybin/washington_changes_all/pkg/wrap"
)

type IKucoinExplore interface {
    Process(ctx context.Context) error
}

type KucoinExploreUsecase struct {
    service service.ICryptoExchangeService
}

func NewKucoinExploreUsecase(service service.ICryptoExchangeService) *KucoinExploreUsecase {
    return &KucoinExploreUsecase{service: service}
}

func (u *KucoinExploreUsecase) Process(
    ctx context.Context,
) error {
    
    balance, err := u.service.GetBalance(ctx)
    if err != nil {
        return wrap.Errorf("failed to get balance: %w", err)
    }
    
    fmt.Printf("balance: %.2f \n", balance)
    
    //tickers, err := u.service.GetMarketsHistory()
    //if err != nil {
    //	return wrap.Errorf("failed to get markets history: %w", err)
    //}
    //
    //fmt.Printf("%v", tickers)
    
    fmt.Println("=====================")
    //res, err := u.service.SellByMarket(model.Coin{Name: coin_name.BTC})
    //if err != nil {
    //	return wrap.Errorf("failed to buy by market: %w", err)
    //}
    
    //fmt.Printf("%v", res)
    fmt.Println("TEST ended")
    return nil
}
