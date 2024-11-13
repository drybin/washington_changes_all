// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package washington_database

import (
	"context"
	"time"
)

const getCoinAmount = `-- name: GetCoinAmount :one
SELECT coin, amount FROM coin_amount
WHERE coin = $1 LIMIT 1
`

func (q *Queries) GetCoinAmount(ctx context.Context, coin string) (CoinAmount, error) {
	row := q.db.QueryRow(ctx, getCoinAmount, coin)
	var i CoinAmount
	err := row.Scan(&i.Coin, &i.Amount)
	return i, err
}

const getCoinAvgPrice = `-- name: GetCoinAvgPrice :one
SELECT coin, price FROM coin_avg_prices
WHERE coin = $1 LIMIT 1
`

func (q *Queries) GetCoinAvgPrice(ctx context.Context, coin string) (CoinAvgPrice, error) {
	row := q.db.QueryRow(ctx, getCoinAvgPrice, coin)
	var i CoinAvgPrice
	err := row.Scan(&i.Coin, &i.Price)
	return i, err
}

const getCoinAvgPrices = `-- name: GetCoinAvgPrices :many
SELECT coin, price FROM coin_avg_prices
WHERE coin = ANY($1::TEXT[])
`

func (q *Queries) GetCoinAvgPrices(ctx context.Context, dollar_1 []string) ([]CoinAvgPrice, error) {
	rows, err := q.db.Query(ctx, getCoinAvgPrices, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CoinAvgPrice
	for rows.Next() {
		var i CoinAvgPrice
		if err := rows.Scan(&i.Coin, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCoinAvgPricesAll = `-- name: GetCoinAvgPricesAll :many
SELECT coin, price FROM coin_avg_prices
`

func (q *Queries) GetCoinAvgPricesAll(ctx context.Context) ([]CoinAvgPrice, error) {
	rows, err := q.db.Query(ctx, getCoinAvgPricesAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CoinAvgPrice
	for rows.Next() {
		var i CoinAvgPrice
		if err := rows.Scan(&i.Coin, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDayInfo = `-- name: GetDayInfo :one
SELECT id, date, account_balance, overal_amount_usdt, overal_coin_count, tiername, coin_to_buy FROM days
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetDayInfo(ctx context.Context, id int) (Day, error) {
	row := q.db.QueryRow(ctx, getDayInfo, id)
	var i Day
	err := row.Scan(
		&i.ID,
		&i.Date,
		&i.AccountBalance,
		&i.OveralAmountUsdt,
		&i.OveralCoinCount,
		&i.Tiername,
		&i.CoinToBuy,
	)
	return i, err
}

const getLastDayInfo = `-- name: GetLastDayInfo :one
SELECT id, date, account_balance, overal_amount_usdt, overal_coin_count, tiername, coin_to_buy FROM days
ORDER BY id DESC LIMIT 1
`

func (q *Queries) GetLastDayInfo(ctx context.Context) (Day, error) {
	row := q.db.QueryRow(ctx, getLastDayInfo)
	var i Day
	err := row.Scan(
		&i.ID,
		&i.Date,
		&i.AccountBalance,
		&i.OveralAmountUsdt,
		&i.OveralCoinCount,
		&i.Tiername,
		&i.CoinToBuy,
	)
	return i, err
}

const saveBuyLog = `-- name: SaveBuyLog :one
INSERT INTO
    buy_log (
    day_id,
    coin,
    amount,
    price
)
VALUES (
        $1,
        $2,
        $3,
        $4
       )
    RETURNING id, day_id, coin, amount, price
`

type SaveBuyLogParams struct {
	DayID  int
	Coin   string
	Amount float64
	Price  float64
}

func (q *Queries) SaveBuyLog(ctx context.Context, arg SaveBuyLogParams) (BuyLog, error) {
	row := q.db.QueryRow(ctx, saveBuyLog,
		arg.DayID,
		arg.Coin,
		arg.Amount,
		arg.Price,
	)
	var i BuyLog
	err := row.Scan(
		&i.ID,
		&i.DayID,
		&i.Coin,
		&i.Amount,
		&i.Price,
	)
	return i, err
}

const saveCoinAmount = `-- name: SaveCoinAmount :one
INSERT INTO
    coin_amount (
    coin,
    amount
)
VALUES (
           $1,
           $2
       )
    RETURNING coin, amount
`

type SaveCoinAmountParams struct {
	Coin   string
	Amount float64
}

func (q *Queries) SaveCoinAmount(ctx context.Context, arg SaveCoinAmountParams) (CoinAmount, error) {
	row := q.db.QueryRow(ctx, saveCoinAmount, arg.Coin, arg.Amount)
	var i CoinAmount
	err := row.Scan(&i.Coin, &i.Amount)
	return i, err
}

const saveCoinAvgPricesAll = `-- name: SaveCoinAvgPricesAll :one
INSERT INTO coin_avg_prices (coin, price) VALUES ($1, $2)  RETURNING coin, price
`

type SaveCoinAvgPricesAllParams struct {
	Coin  string
	Price float64
}

func (q *Queries) SaveCoinAvgPricesAll(ctx context.Context, arg SaveCoinAvgPricesAllParams) (CoinAvgPrice, error) {
	row := q.db.QueryRow(ctx, saveCoinAvgPricesAll, arg.Coin, arg.Price)
	var i CoinAvgPrice
	err := row.Scan(&i.Coin, &i.Price)
	return i, err
}

const saveDayInfo = `-- name: SaveDayInfo :one
INSERT INTO
    days (
    date,
    account_balance,
    overal_amount_usdt,
    overal_coin_count,
    tierName,
    coin_to_buy
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
       )
    RETURNING id, date, account_balance, overal_amount_usdt, overal_coin_count, tiername, coin_to_buy
`

type SaveDayInfoParams struct {
	Date             time.Time
	AccountBalance   float64
	OveralAmountUsdt float64
	OveralCoinCount  int
	Tiername         string
	CoinToBuy        string
}

func (q *Queries) SaveDayInfo(ctx context.Context, arg SaveDayInfoParams) (Day, error) {
	row := q.db.QueryRow(ctx, saveDayInfo,
		arg.Date,
		arg.AccountBalance,
		arg.OveralAmountUsdt,
		arg.OveralCoinCount,
		arg.Tiername,
		arg.CoinToBuy,
	)
	var i Day
	err := row.Scan(
		&i.ID,
		&i.Date,
		&i.AccountBalance,
		&i.OveralAmountUsdt,
		&i.OveralCoinCount,
		&i.Tiername,
		&i.CoinToBuy,
	)
	return i, err
}

const saveMarketInfo = `-- name: SaveMarketInfo :exec
INSERT INTO
    market_info_history (
    date,
    coin_first,
    coin_second,
    symbol_name,
    buy,
    sell,
    change_rate,
    change_price,
    high_price,
    low_price,
    vol_btc,
    vol_value,
    last_price,
    average_price,
    taker_fee_rate,
    maker_fee_rate,
    taker_coefficient,
    maker_coefficient
)
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18
)
`

type SaveMarketInfoParams struct {
	Date             time.Time
	CoinFirst        string
	CoinSecond       string
	SymbolName       string
	Buy              float64
	Sell             float64
	ChangeRate       float64
	ChangePrice      float64
	HighPrice        float64
	LowPrice         float64
	VolBtc           float64
	VolValue         float64
	LastPrice        float64
	AveragePrice     float64
	TakerFeeRate     float64
	MakerFeeRate     float64
	TakerCoefficient float64
	MakerCoefficient float64
}

func (q *Queries) SaveMarketInfo(ctx context.Context, arg SaveMarketInfoParams) error {
	_, err := q.db.Exec(ctx, saveMarketInfo,
		arg.Date,
		arg.CoinFirst,
		arg.CoinSecond,
		arg.SymbolName,
		arg.Buy,
		arg.Sell,
		arg.ChangeRate,
		arg.ChangePrice,
		arg.HighPrice,
		arg.LowPrice,
		arg.VolBtc,
		arg.VolValue,
		arg.LastPrice,
		arg.AveragePrice,
		arg.TakerFeeRate,
		arg.MakerFeeRate,
		arg.TakerCoefficient,
		arg.MakerCoefficient,
	)
	return err
}