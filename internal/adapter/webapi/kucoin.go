package webapi

import (
	"encoding/json"
	"fmt"
	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/drybin/washington_changes_all/internal/domain/model"
	"github.com/drybin/washington_changes_all/internal/domain/types"
	"github.com/drybin/washington_changes_all/internal/domain/types/coin_name"
	"github.com/drybin/washington_changes_all/pkg/wrap"
	"strconv"
	"strings"
	"time"
)

type KucoinWebapi struct {
	client *kucoin.ApiService
}

func NewKucoinWebapi(client *kucoin.ApiService) *KucoinWebapi {
	return &KucoinWebapi{client: client}
}

func (c *KucoinWebapi) GetMarketsHistory() (*[]model.MarketInfo, error) {

	res, err := c.client.Tickers()
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

func (c *KucoinWebapi) GetBalance() (float64, error) {

	balances, err := c.GetBalances()
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

func (c *KucoinWebapi) GetBalances() ([]*model.Balance, error) {
	rsp, err := c.client.Accounts("", "")
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

func (c *KucoinWebapi) BuyByMarket(coin model.Coin) (*kucoin.OrderModel, error) {

	resp, err := c.client.CreateOrder(
		&kucoin.CreateOrderModel{
			ClientOid: kucoin.IntToString(time.Now().UnixNano()),
			Side:      "buy",
			Symbol:    fmt.Sprintf("%s-USDT", coin.Name),
			Type:      "market",
			Funds:     "1",
		},
	)

	if err != nil {
		return nil, wrap.Errorf("failed to place market order to buy: %w %s", err, string(resp.RawData))
	}

	orderInfo := kucoin.CreateOrderResultModel{}
	err = json.Unmarshal(resp.RawData, &orderInfo)
	if err != nil {
		return nil, wrap.Errorf("failed to unmarshal create order result model: %w", err)
	}

	return c.GetOrderInfo(orderInfo.OrderId)
}

func (c *KucoinWebapi) SellByMarket(coin model.Coin) (*kucoin.OrderModel, error) {

	resp, err := c.client.CreateOrder(
		&kucoin.CreateOrderModel{
			ClientOid: kucoin.IntToString(time.Now().UnixNano()),
			Side:      "sell",
			Symbol:    fmt.Sprintf("%s-USDT", coin.Name),
			Type:      "market",
			//Funds:     "0.0000154",
			Size: "0.0000154",
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

	return c.GetOrderInfo(orderInfo.OrderId)
}

func (c *KucoinWebapi) GetOrderInfo(orderId string) (*kucoin.OrderModel, error) {
	resp, err := c.client.Order(orderId)
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

func getCoinNameFirst(s string) string {
	res := strings.Split(s, "-")
	return res[0]
}

func getCoinNameSecond(s string) string {
	res := strings.Split(s, "-")
	return res[1]
}
