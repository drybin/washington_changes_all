package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/drybin/washington_changes_all/internal/app/cli/config"
	"github.com/drybin/washington_changes_all/internal/domain/model"
	"github.com/drybin/washington_changes_all/internal/domain/repo"
	"github.com/drybin/washington_changes_all/internal/domain/service/buy_strategy"
	"github.com/drybin/washington_changes_all/internal/domain/types"
	"github.com/drybin/washington_changes_all/pkg/wrap"
	"github.com/samber/lo"
)

type DayProcessService struct {
	DaysRepo              repo.IDaysRepository
	MarketHistoryRepo     repo.IMarketsHistoryRepository
	CoinAvgPricesRepo     repo.ICoinAvgPricesRepository
	BuyLogRepo            repo.IBuyLogRepository
	CoinAmountRepo        repo.ICoinAmountRepository
	CryptoExchangeService ICryptoExchangeService
	AmountEmptyStrategy   buy_strategy.IAmountEmptyStrategy
	MaxPriceDownStrategy  buy_strategy.IMaxPriceDownStrategy
	CoinConfig            config.CoinConfig
}

type IDayProcessService interface {
	Process(ctx context.Context) (bool, error)
}

func NewDayProcessService(
	repo repo.IDaysRepository,
	marketHistoryRepo repo.IMarketsHistoryRepository,
	coinAvgPricesRepo repo.ICoinAvgPricesRepository,
	buyLogRepo repo.IBuyLogRepository,
	coinAmountRepo repo.ICoinAmountRepository,
	cryptoExchangeService ICryptoExchangeService,
	amountEmptyStrategy buy_strategy.IAmountEmptyStrategy,
	maxPriceDownStrategy buy_strategy.IMaxPriceDownStrategy,
	coinConfig config.CoinConfig,
) IDayProcessService {
	return &DayProcessService{
		DaysRepo:              repo,
		MarketHistoryRepo:     marketHistoryRepo,
		CoinAvgPricesRepo:     coinAvgPricesRepo,
		BuyLogRepo:            buyLogRepo,
		CoinAmountRepo:        coinAmountRepo,
		CryptoExchangeService: cryptoExchangeService,
		AmountEmptyStrategy:   amountEmptyStrategy,
		MaxPriceDownStrategy:  maxPriceDownStrategy,
		CoinConfig:            coinConfig,
	}
}

func (s *DayProcessService) Process(
	ctx context.Context,
) (bool, error) {

	balance, err := s.CryptoExchangeService.GetBalance()
	if err != nil {
		return false, wrap.Errorf("failed to get balance info: %w", err)
	}
	fmt.Printf("Баланс: %.2f\n", balance)
	if balance < 0.0 {
		fmt.Printf("баланс меньше 1$ завершаем работу\n")
		return false, nil
	}

	day, err := s.DaysRepo.GetLastDayInfo(ctx)
	if err != nil {
		return false, wrap.Errorf("failed to get last day info: %w", err)
	}

	tier := getTierName(day)
	//tierCoins := mapConfigCoinArrayToModelCoinArray(getTierCoins(tier, s.CoinConfig))
	tierCoins := getTierCoins(tier, s.CoinConfig)
	fmt.Printf("Корзина из которой будем покупать: %s\n", tier)
	fmt.Printf("Монеты: %v\n", getTierCoins(tier, s.CoinConfig))
	fmt.Printf("Монеты преобразованные: %v\n", tierCoins)
	fmt.Printf("%v", day)

	marketsHistory, err := s.CryptoExchangeService.GetMarketsHistory()
	if err != nil {
		return false, wrap.Errorf("failed to get Markets History: %w", err)
	}

	err = s.MarketHistoryRepo.SaveMany(ctx, marketsHistory)
	if err != nil {
		return false, wrap.Errorf("failed to save Markets History: %w", err)
	}

	coinsToBuy, err := s.AmountEmptyStrategy.Process(ctx, tierCoins, marketsHistory)
	if err != nil {
		return false, wrap.Errorf("failed to apply AmountEmptyStrategy: %w", err)
	}

	if len(coinsToBuy) > 0 {
		fmt.Printf("Монеты которые раньше не покупали, отсортированные по макс падению: %v\n", coinsToBuy)
		//TODO BUY
	}

	fmt.Printf("Монеты которыx раньше не покупали нет\n")
	coinsToBuy, err = s.MaxPriceDownStrategy.Process(ctx, tierCoins, marketsHistory)
	if err != nil {
		return false, wrap.Errorf("failed to apply MaxPriceDownStrategy: %w", err)
	}

	fmt.Printf("Монеты отсортированные по макс падению от средней цены: %v\n", coinsToBuy)
	fmt.Printf("%v", lo.Samples(coinsToBuy, 1))

	return true, nil
}

func (s *DayProcessService) buy(
	ctx context.Context,
	tierName types.Tier,
	balance float64,
	prevDay model.Day,
	coinToBuy model.CoinPriceChange,
) error {
	fmt.Printf("Покупаем монету: %s\n", coinToBuy.Coin.Name)

	//byu on kukoin
	orderInfo, err := s.CryptoExchangeService.BuyByMarket(model.Coin{
		Name: coinToBuy.Coin.Name,
	})

	amount, err := strconv.ParseFloat(orderInfo.DealFunds, 32)
	if err == nil {
		return wrap.Errorf("failed to parse amount float in order info: %w", err)
	}

	dealSize, err := strconv.ParseFloat(orderInfo.DealSize, 32)
	if err == nil {
		return wrap.Errorf("failed to parse deal size float in order info: %w", err)
	}

	price := amount / dealSize

	fmt.Println("Логируем информацию о дне")
	coinWithAvgPrices, err := s.CoinAvgPricesRepo.GetAll(ctx)
	if err != nil {
		return wrap.Errorf("failed to get all coin avg prices: %w", err)
	}

	coinCount, prevCoinAvgPrice := calcCoinsCountAndPrevAvgPrice(coinWithAvgPrices, coinToBuy.Coin)
	fmt.Printf("Количество монет: %d\n", coinCount)
	fmt.Printf("Предыдущая средняя цена: %f\n", prevCoinAvgPrice)

	day, err := s.DaysRepo.Save(
		ctx,
		model.Day{
			Date:              time.Now(),
			AccountBalance:    balance,
			OverallAmountUsdt: prevDay.OverallAmountUsdt + 1.0,
			OverallCoinCount:  coinCount,
			TierName:          tierName.String(),
			CoinToBuy:         coinToBuy.Coin.Name.String(),
		},
	)
	if err != nil {
		return wrap.Errorf("failed to save day info: %w", err)
	}

	fmt.Println("Логируем покупку в таблицу buy_log")
	err = s.BuyLogRepo.Save(
		ctx,
		day.ID,
		coinToBuy.Coin,
		amount,
		price,
	)
	if err != nil {
		return wrap.Errorf("failed to save log buy: %w", err)
	}

	prevCoinAmount, err := s.CoinAmountRepo.Get(ctx, coinToBuy.Coin)
	if err != nil {
		return wrap.Errorf("failed to get coin amount: %w", err)
	}

	avgPrice := price
	if prevCoinAvgPrice != 0.0 {
		avgPrice = (prevCoinAvgPrice*prevCoinAmount + price*amount) / (prevCoinAmount + amount)
	}
	fmt.Printf("Новая посчитаная средняя цена: %f\n", avgPrice)

	fmt.Println("Логируем новую среднюю цену")
	err = s.CoinAvgPricesRepo.Save(ctx, model.CoinPrice{Coin: coinToBuy.Coin, Price: avgPrice})
	if err != nil {
		return wrap.Errorf("failed to save coin avg price: %w", err)
	}

	coinAmount := prevCoinAmount + amount
	fmt.Printf("Новое посчитаное количество монеты: %f\n", avgPrice)

	fmt.Println("Логируем новое количество")
	err = s.CoinAmountRepo.Save(ctx, coinToBuy.Coin, coinAmount)
	if err != nil {
		return wrap.Errorf("failed to save coin amount: %w", err)
	}

	//droptabs log
	//coingecko log
	return nil
}

func calcCoinsCountAndPrevAvgPrice(coins *[]model.CoinPrice, coinToBuy model.Coin) (int, float64) {
	coinCount := len(*coins)
	inc := 0
	prevCoinAvgPrice := 0.0
	for _, item := range *coins {
		if item.Coin.Name == coinToBuy.Name {
			inc = 1
			prevCoinAvgPrice = item.Price
		}
	}

	coinCount = coinCount + inc

	return coinCount, prevCoinAvgPrice
}

func getTierName(day *model.Day) types.Tier {
	if day == nil {
		return types.TierOne
	}

	switch day.TierName {
	case string(types.TierOne):
		return types.TierTwo
	case string(types.TierTwo):
		return types.TierThree
	case string(types.TierThree):
		return types.TierOne
	default:
		return types.TierOne
	}
}

func getTierCoins(tierName types.Tier, config config.CoinConfig) []config.Coin {
	switch tierName {
	case types.TierOne:
		return config.TierOne
	case types.TierTwo:
		return config.TierTwo
	case types.TierThree:
		return config.TierThree
	default:
		return config.TierOne
	}
}
