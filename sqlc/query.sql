-- name: GetDayInfo :one
SELECT * FROM days
WHERE id = $1 LIMIT 1;

-- name: GetLastDayInfo :one
SELECT * FROM days
ORDER BY id DESC LIMIT 1;

-- name: SaveDayInfo :one
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
    RETURNING *;

-- name: GetCoinAvgPrice :one
SELECT * FROM coin_avg_prices
WHERE coin = $1 LIMIT 1;

-- name: GetCoinAmount :one
SELECT * FROM coin_amount
WHERE coin = $1 LIMIT 1;

-- name: SaveCoinAmount :one
INSERT INTO
    coin_amount (
    coin,
    amount
)
VALUES (
           $1,
           $2
       )
    ON CONFLICT(coin) DO UPDATE
    SET amount = $2
    RETURNING *;

-- name: GetCoinAvgPrices :many
SELECT * FROM coin_avg_prices
WHERE coin = ANY($1::TEXT[]);

-- name: GetCoinAvgPricesAll :many
SELECT * FROM coin_avg_prices;

-- name: SaveCoinAvgPricesAll :one
INSERT INTO coin_avg_prices (coin, price)
    VALUES ($1, $2)
ON CONFLICT(coin) DO UPDATE
    SET price = $2
    RETURNING *;


-- name: SaveMarketInfo :exec
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
);


-- name: SaveBuyLog :one
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
    RETURNING *;

-- name: SaveSellLog :one
INSERT INTO
    sell_log (
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
    RETURNING *;