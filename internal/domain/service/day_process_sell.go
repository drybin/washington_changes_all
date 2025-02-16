package service

import (
    "context"
    "fmt"
    "strconv"
    
    "github.com/Kucoin/kucoin-go-sdk"
    "github.com/drybin/washington_changes_all/internal/app/cli/config"
    "github.com/drybin/washington_changes_all/internal/domain/model"
    "github.com/drybin/washington_changes_all/internal/domain/repo"
    "github.com/drybin/washington_changes_all/internal/domain/service/buy_strategy"
    "github.com/drybin/washington_changes_all/internal/domain/types/coin_name"
    "github.com/drybin/washington_changes_all/pkg/wrap"
)

type DayProcessSellService struct {
    DaysRepo              repo.IDaysRepository
    MarketHistoryRepo     repo.IMarketsHistoryRepository
    CoinAvgPricesRepo     repo.ICoinAvgPricesRepository
    SellLogRepo           repo.ISellLogRepository
    CoinAmountRepo        repo.ICoinAmountRepository
    CryptoExchangeService ICryptoExchangeService
    AmountEmptyStrategy   buy_strategy.IAmountEmptyStrategy
    MaxPriceDownStrategy  buy_strategy.IMaxPriceDownStrategy
    CoinConfig            config.CoinConfig
}

type IDayProcessSellService interface {
    Process(ctx context.Context) (*model.DayResult, error)
}

func NewDayProcessSellService(
    repo repo.IDaysRepository,
    marketHistoryRepo repo.IMarketsHistoryRepository,
    coinAvgPricesRepo repo.ICoinAvgPricesRepository,
    sellLogRepo repo.ISellLogRepository,
    coinAmountRepo repo.ICoinAmountRepository,
    cryptoExchangeService ICryptoExchangeService,
    amountEmptyStrategy buy_strategy.IAmountEmptyStrategy,
    maxPriceDownStrategy buy_strategy.IMaxPriceDownStrategy,
    coinConfig config.CoinConfig,
) IDayProcessSellService {
    return &DayProcessSellService{
        DaysRepo:              repo,
        MarketHistoryRepo:     marketHistoryRepo,
        CoinAvgPricesRepo:     coinAvgPricesRepo,
        SellLogRepo:           sellLogRepo,
        CoinAmountRepo:        coinAmountRepo,
        CryptoExchangeService: cryptoExchangeService,
        AmountEmptyStrategy:   amountEmptyStrategy,
        MaxPriceDownStrategy:  maxPriceDownStrategy,
        CoinConfig:            coinConfig,
    }
}

func (s *DayProcessSellService) Process(
    ctx context.Context,
) (*model.DayResult, error) {
    
    fmt.Printf("Ищем монеты для продажи\n")
    
    marketsHistory, err := s.CryptoExchangeService.GetMarketsHistory()
    if err != nil {
        return nil, wrap.Errorf("failed to get Markets History: %w", err)
    }
    
    allCoin := getAllTierCoins(s.CoinConfig)
    
    coinAvgPrices, err := s.CoinAvgPricesRepo.GetList(ctx, buy_strategy.MapConfigCoinArrayToModelArray(allCoin))
    if err != nil {
        return nil, wrap.Errorf("failed to get coins avg prices: %w", err)
    }
    
    for _, coinAvgPrice := range *coinAvgPrices {
        for _, market := range *marketsHistory {
            if market.Pair.CoinFirst.Name == coinAvgPrice.Coin.Name && market.Pair.CoinSecond.Name == coin_name.USDT {
                if market.AveragePrice >= coinAvgPrice.Price*2 {
                    fmt.Printf("Монета найдена %s avgPrice %f currentPrice %f\n",
                        coinAvgPrice.Coin.Name,
                        coinAvgPrice.Price,
                        market.AveragePrice,
                    )
                    
                    //  sell
                    amountFromRepo, err := s.CoinAmountRepo.Get(ctx, model.Coin{
                        Name: coinAvgPrice.Coin.Name,
                    })
                    if err != nil {
                        return nil, wrap.Errorf("failed to get coins amount: %w", err)
                    }
                    
                    amountToSell := amountFromRepo / 2.0
                    amountToSellString := fmt.Sprintf("%f", amountToSell)
                    
                    fmt.Printf("текущее количество монеты %f, будем продавать %s\n", amountFromRepo, amountToSell)
                    
                    orderInfo, err := s.sell(ctx, coinAvgPrice.Coin.Name, amountToSellString)
                    if err != nil {
                        return nil, wrap.Errorf("failed to sell by market: %w", err)
                    }
                    
                    if orderInfo == nil {
                        return nil, wrap.Errorf("sell orderInfo empty")
                    }
                    
                    amount, err := strconv.ParseFloat(orderInfo.DealSize, 32)
                    if err != nil {
                        return nil, wrap.Errorf("failed to parse amount float in order info: %w", err)
                    }
                    
                    fmt.Println("Логируем продажу")
                    
                    fmt.Println("Логируем новое количество")
                    err = s.CoinAmountRepo.Save(ctx, coinAvgPrice.Coin, amountToSell)
                    if err != nil {
                        return nil, wrap.Errorf("failed to save coin amount: %w", err)
                    }
                    
                    dealSize, err := strconv.ParseFloat(orderInfo.DealSize, 32)
                    if err != nil {
                        return nil, wrap.Errorf("failed to parse deal size float in order info: %w", err)
                    }
                    
                    dealFunds, err := strconv.ParseFloat(orderInfo.DealFunds, 32)
                    if err != nil {
                        return nil, wrap.Errorf("failed to parse deal fund float in order info: %w", err)
                    }
                    
                    price := dealFunds / dealSize
                }
            }
        }
    }
    
    return &model.DayResult{}, nil
}

func (s *DayProcessSellService) sell(
    ctx context.Context,
    coin coin_name.CoinName,
    amount string,
) (*kucoin.OrderModel, error) {
    
    //byu on kukoin
    orderInfo, err := s.CryptoExchangeService.SellByMarket(model.Coin{
        Name: coin,
    },
        amount)
    
    if err != nil {
        return nil, wrap.Errorf("failed to sell by market: %w", err)
    }
    
    return orderInfo, nil
}

func getAllTierCoins(cfg config.CoinConfig) []config.Coin {
    res := []config.Coin{}
    
    for _, item := range cfg.TierOne {
        res = append(res, item)
    }
    
    for _, item := range cfg.TierTwo {
        res = append(res, item)
    }
    
    for _, item := range cfg.TierThree {
        res = append(res, item)
    }
    
    return res
}
