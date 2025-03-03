package webapi

import (
    "context"
    "encoding/json"
    "fmt"
    "strconv"
    "strings"
    "time"
    
    "github.com/Kucoin/kucoin-go-sdk"
    "github.com/drybin/washington_changes_all/internal/domain/model"
    "github.com/drybin/washington_changes_all/internal/domain/types"
    "github.com/drybin/washington_changes_all/internal/domain/types/coin_name"
    "github.com/drybin/washington_changes_all/pkg/wrap"
)

type KucoinWebapi struct {
    client *kucoin.ApiService
}

func NewKucoinWebapi(client *kucoin.ApiService) *KucoinWebapi {
    return &KucoinWebapi{client: client}
}

func (c *KucoinWebapi) GetMarketsHistory(ctx context.Context) (*[]model.MarketInfo, error) {
    
    res, err := c.client.Tickers(ctx)
    if err != nil {
        return nil, wrap.Errorf("failed to get kucoin tickers: %w", err)
    }
    
    tickers := kucoin.TickersResponseModel{}
    if err := res.ReadData(&tickers); err != nil {
        return nil, wrap.Errorf("failed to read kucoin tickers: %w", err)
    }
    
    result := make([]model.MarketInfo, 0, len(tickers.Tickers))
    
    for _, ticker := range tickers.Tickers {
        if coin_name.FromString(getCoinNameFirst(ticker.Symbol)) == coin_name.UNKNOWN {
            continue
        }
        if coin_name.FromString(getCoinNameSecond(ticker.Symbol)) == coin_name.UNKNOWN {
            continue
        }
        
        result = append(result, mapTickersResponseToDomainModel(ticker))
    }
    
    return &result, nil
}

func (c *KucoinWebapi) GetBalance(ctx context.Context) (float64, error) {
    
    balances, err := c.GetBalances(ctx)
    if err != nil {
        return 0.0, wrap.Errorf("failed to get kucoin balances: %w", err)
    }
    
    for _, balance := range balances {
        if balance.BalanceType == types.Trade && balance.Coin.Name == coin_name.USDT {
            return balance.Balance, nil
        }
    }
    
    return 0.0, wrap.Errorf("trade balance in usdt not found: %w", err)
}

func (c *KucoinWebapi) GetBalances(ctx context.Context) ([]*model.Balance, error) {
    rsp, err := c.client.Accounts(ctx, "", "")
    if err != nil {
        return nil, wrap.Errorf("failed to get kucoin balances: %w", err)
    }
    
    data := kucoin.AccountsModel{}
    if err := rsp.ReadData(&data); err != nil {
        return nil, wrap.Errorf("failed to read kucoin balances: %w", err)
    }
    
    result := make([]*model.Balance, 0, len(data))
    for _, item := range data {
        
        balanceFloat, err := strconv.ParseFloat(item.Balance, 32)
        if err != nil {
            return nil, wrap.Errorf("failed to parse float in balances: %w", err)
        }
        
        result = append(result, &model.Balance{
            BalanceType: types.BalanceType(item.Type),
            Coin: model.Coin{
                Name: coin_name.CoinName(item.Currency),
            },
            Balance: balanceFloat,
        })
    }
    
    return result, nil
}

func (c *KucoinWebapi) GetPairInfo(ctx context.Context, pair model.CoinsPair) (*model.PairInfo, error) {
    symbol := fmt.Sprintf("%s-%s", pair.CoinFirst.Name, pair.CoinSecond.Name)
    resp, err := c.client.SymbolsDetail(ctx, symbol)
    if err != nil {
        return nil, wrap.Errorf("failed to symbol info: %w", err)
    }
    
    symbolInfo := kucoin.SymbolModelV2{}
    err = json.Unmarshal(resp.RawData, &symbolInfo)
    if err != nil {
        return nil, wrap.Errorf("failed to unmarshal symbolV2 model: %w", err)
    }
    
    result, err := mapSymbolV2ResponseToDomainModel(symbolInfo, pair)
    if err != nil {
        return nil, wrap.Errorf("failed to map symbolV2 to model: %w", err)
    }
    
    return result, nil
}

func (c *KucoinWebapi) BuyByMarket(ctx context.Context, coin model.Coin) (*kucoin.OrderModel, error) {
    
    resp, err := c.client.CreateOrder(
        ctx,
        &kucoin.CreateOrderModel{
            ClientOid: kucoin.IntToString(time.Now().UnixNano()),
            Side:      "buy",
            Symbol:    fmt.Sprintf("%s-USDT", coin.Name),
            Type:      "market",
            Funds:     "1",
        },
    )
    
    if err != nil {
        respAsString := "response is empty"
        if resp != nil && resp.RawData != nil {
            respAsString = string(resp.RawData)
        }
        return nil, wrap.Errorf("failed to place market order to buy: %w %s", err, respAsString)
    }
    
    orderInfo := kucoin.CreateOrderResultModel{}
    err = json.Unmarshal(resp.RawData, &orderInfo)
    if err != nil {
        return nil, wrap.Errorf("failed to unmarshal create order result model: %w", err)
    }
    
    return c.GetOrderInfo(ctx, orderInfo.OrderId)
}

func (c *KucoinWebapi) SellByMarket(ctx context.Context, coin model.Coin, amount string) (*kucoin.OrderModel, error) {
    
    resp, err := c.client.CreateOrder(
        ctx,
        &kucoin.CreateOrderModel{
            ClientOid: kucoin.IntToString(time.Now().UnixNano()),
            Side:      "sell",
            Symbol:    fmt.Sprintf("%s-USDT", coin.Name),
            Type:      "market",
            //Funds:     "0.0000154",
            Size: amount,
        },
    )
    
    if err != nil {
        return nil, wrap.Errorf("failed to place market order to sell: %w", err)
    }
    
    orderInfo := kucoin.CreateOrderResultModel{}
    err = json.Unmarshal(resp.RawData, &orderInfo)
    if err != nil {
        return nil, wrap.Errorf("failed to unmarshal create order result model: %w %s", err, string(resp.RawData))
    }
    
    return c.GetOrderInfo(ctx, orderInfo.OrderId)
}

func (c *KucoinWebapi) GetOrderInfo(ctx context.Context, orderId string) (*kucoin.OrderModel, error) {
    resp, err := c.client.Order(ctx, orderId)
    if err != nil {
        return nil, wrap.Errorf("failed to get order info: %w", err)
    }
    
    order := kucoin.OrderModel{}
    err = json.Unmarshal(resp.RawData, &order)
    if err != nil {
        return nil, wrap.Errorf("failed to unmarshal order info: %w", err)
    }
    
    return &order, nil
}

func mapTickersResponseToDomainModel(m *kucoin.TickerModel) model.MarketInfo {
    buy, _ := strconv.ParseFloat(m.Buy, 64)
    sell, _ := strconv.ParseFloat(m.Sell, 64)
    changeRate, _ := strconv.ParseFloat(m.ChangeRate, 64)
    changePrice, _ := strconv.ParseFloat(m.ChangePrice, 64)
    highPrice, _ := strconv.ParseFloat(m.High, 64)
    lowPrice, _ := strconv.ParseFloat(m.Low, 64)
    volBtc, _ := strconv.ParseFloat(m.Vol, 64)
    volValue, _ := strconv.ParseFloat(m.VolValue, 64)
    lastPrice, _ := strconv.ParseFloat(m.Last, 64)
    averagePrice, _ := strconv.ParseFloat(m.AveragePrice, 64)
    takerFeeRate, _ := strconv.ParseFloat(m.TakerFeeRate, 64)
    makerFeeRate, _ := strconv.ParseFloat(m.MakerFeeRate, 64)
    takerCoefficient, _ := strconv.ParseFloat(m.TakerCoefficient, 64)
    makerCoefficient, _ := strconv.ParseFloat(m.MakerCoefficient, 64)
    
    return model.MarketInfo{
        Pair: model.CoinsPair{
            CoinFirst: model.Coin{
                Name: coin_name.FromString(getCoinNameFirst(m.Symbol)),
            },
            CoinSecond: model.Coin{
                Name: coin_name.FromString(getCoinNameSecond(m.Symbol)),
            },
        },
        Buy:              buy,
        Sell:             sell,
        ChangeRate:       changeRate,
        ChangePrice:      changePrice,
        HighPrice:        highPrice,
        LowPrice:         lowPrice,
        VolBTC:           volBtc,
        VolValue:         volValue,
        LastPrice:        lastPrice,
        AveragePrice:     averagePrice,
        TakerFeeRate:     takerFeeRate,
        MakerFeeRate:     makerFeeRate,
        TakerCoefficient: takerCoefficient,
        MakerCoefficient: makerCoefficient,
    }
}

func mapSymbolV2ResponseToDomainModel(symbolInfo kucoin.SymbolModelV2, pair model.CoinsPair) (*model.PairInfo, error) {
    
    QuoteMinSize, err := strconv.ParseFloat(symbolInfo.QuoteMinSize, 64)
    if err != nil {
        fmt.Println("Error parse QuoteMinSize:", err)
        return nil, err
    }
    
    BaseMinSize, err := strconv.ParseFloat(symbolInfo.BaseMinSize, 64)
    if err != nil {
        fmt.Println("Error parse BaseMinSize:", err)
        return nil, err
    }
    
    BaseMaxSize, err := strconv.ParseFloat(symbolInfo.BaseMaxSize, 64)
    if err != nil {
        fmt.Println("Error parse BaseMaxSize:", err)
        return nil, err
    }
    
    BaseIncrement, err := strconv.ParseFloat(symbolInfo.BaseIncrement, 64)
    if err != nil {
        fmt.Println("Error parse BaseIncrement:", err)
        return nil, err
    }
    
    QuoteIncrement, err := strconv.ParseFloat(symbolInfo.QuoteIncrement, 64)
    if err != nil {
        fmt.Println("Error parse QuoteIncrement:", err)
        return nil, err
    }
    
    QuoteMaxSize, err := strconv.ParseFloat(symbolInfo.QuoteMaxSize, 64)
    if err != nil {
        fmt.Println("Error parse QuoteMaxSize:", err)
        return nil, err
    }
    
    PriceIncrement, err := strconv.ParseFloat(symbolInfo.PriceIncrement, 64)
    if err != nil {
        fmt.Println("Error parse PriceIncrement:", err)
        return nil, err
    }
    
    MinFunds, err := strconv.ParseFloat(symbolInfo.MinFunds, 64)
    if err != nil {
        fmt.Println("Error parse MinFunds:", err)
        return nil, err
    }
    
    PriceLimitRate, err := strconv.ParseFloat(symbolInfo.PriceLimitRate, 64)
    if err != nil {
        fmt.Println("Error parse PriceLimitRate:", err)
        return nil, err
    }
    
    return &model.PairInfo{
        Pair:            pair,
        Symbol:          symbolInfo.Symbol,
        Name:            symbolInfo.Name,
        BaseCurrency:    symbolInfo.BaseCurrency,
        QuoteCurrency:   symbolInfo.QuoteCurrency,
        Market:          symbolInfo.Market,
        BaseMinSize:     BaseMinSize,
        QuoteMinSize:    QuoteMinSize,
        BaseMaxSize:     BaseMaxSize,
        QuoteMaxSize:    QuoteMaxSize,
        BaseIncrement:   BaseIncrement,
        QuoteIncrement:  QuoteIncrement,
        PriceIncrement:  PriceIncrement,
        FeeCurrency:     symbolInfo.FeeCurrency,
        EnableTrading:   symbolInfo.EnableTrading,
        IsMarginEnabled: symbolInfo.IsMarginEnabled,
        PriceLimitRate:  PriceLimitRate,
        MinFunds:        MinFunds,
    }, nil
}

func getCoinNameFirst(s string) string {
    res := strings.Split(s, "-")
    return res[0]
}

func getCoinNameSecond(s string) string {
    res := strings.Split(s, "-")
    return res[1]
}
