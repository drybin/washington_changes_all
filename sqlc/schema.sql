CREATE TABLE days (
     id   SERIAL PRIMARY KEY, -- id и номер дня
     date TIMESTAMPTZ      NOT NULL, --реальная дата
     account_balance  DOUBLE PRECISION NOT NULL, --баланс usdt на бирже
     overal_amount_usdt  DOUBLE PRECISION NOT NULL, --баланс всех монет в usdt эквиваленте
     overal_coin_count  INTEGER NOT NULL, -- количество уникальных монет в портфеле
     tierName  TEXT NOT NULL, -- корзина из которой покупали
     coin_to_buy  TEXT NOT NULL -- монета которую купили сегодня
);

CREATE TABLE buy_log (
      id  SERIAL PRIMARY KEY, --id
      day_id INTEGER NOT NULL, --id дня из days
      coin TEXT NOT NULL, -- монета которую купили сегодня
      amount  DOUBLE PRECISION NOT NULL, -- количество которое купили
      price  DOUBLE PRECISION NOT NULL -- цена по которой купили
);

CREATE TABLE sell_log (
     id   SERIAL PRIMARY KEY, --id
     day_id INTEGER NOT NULL, --id дня из days
     coin TEXT NOT NULL, -- монета которую продали сегодня
     amount  DOUBLE PRECISION NOT NULL, -- количество которое продали
     price  DOUBLE PRECISION NOT NULL -- цена по которой продали
);

CREATE TABLE coin_avg_prices (
      coin   TEXT PRIMARY KEY, --монета
      price  DOUBLE PRECISION NOT NULL -- средняя цена
);

CREATE TABLE coin_amount (
     coin   TEXT PRIMARY KEY, --монета
     amount  DOUBLE PRECISION NOT NULL -- средняя цена
);

CREATE TABLE market_info_history (
         date             TIMESTAMPTZ      NOT NULL, --реальная дата
         coin_first       TEXT NOT NULL,
         coin_second      TEXT NOT NULL,
         symbol_name      TEXT NOT NULL,
         buy              DOUBLE PRECISION NOT NULL, --buy
         sell              DOUBLE PRECISION NOT NULL, --sell
         change_rate       DOUBLE PRECISION NOT NULL, -- 24h change rate
         change_price      DOUBLE PRECISION NOT NULL, -- 24h change price
         high_price        DOUBLE PRECISION NOT NULL, -- 24h highest price
         low_price         DOUBLE PRECISION NOT NULL, -- 24h lowest price
         vol_btc           DOUBLE PRECISION NOT NULL, -- 24h volume，the aggregated trading volume in BTC
         vol_value         DOUBLE PRECISION NOT NULL, --24h total, the trading volume in quote currency of last 24 hours
         last_price        DOUBLE PRECISION NOT NULL, --last price
         average_price     DOUBLE PRECISION NOT NULL, --24h average transaction price yesterday
         taker_fee_rate     DOUBLE PRECISION NOT NULL, --Basic Taker Fee
         maker_fee_rate     DOUBLE PRECISION NOT NULL, --Basic Maker Fee
         taker_coefficient DOUBLE PRECISION NOT NULL, -- Taker Fee Coefficient
         maker_coefficient DOUBLE PRECISION NOT NULL -- Maker Fee Coefficient
);