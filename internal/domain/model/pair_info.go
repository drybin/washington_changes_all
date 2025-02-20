package model

type PairInfo struct {
    Pair            CoinsPair
    Symbol          string  `json:"symbol"`
    Name            string  `json:"name"`
    BaseCurrency    string  `json:"baseCurrency"`
    QuoteCurrency   string  `json:"quoteCurrency"`
    Market          string  `json:"market"`
    BaseMinSize     float64 `json:"baseMinSize"`
    QuoteMinSize    float64 `json:"quoteMinSize"`
    BaseMaxSize     float64 `json:"baseMaxSize"`
    QuoteMaxSize    float64 `json:"quoteMaxSize"`
    BaseIncrement   float64 `json:"baseIncrement"`
    QuoteIncrement  float64 `json:"quoteIncrement"`
    PriceIncrement  float64 `json:"priceIncrement"`
    FeeCurrency     string  `json:"feeCurrency"`
    EnableTrading   bool    `json:"enableTrading"`
    IsMarginEnabled bool    `json:"isMarginEnabled"`
    PriceLimitRate  float64 `json:"priceLimitRate"`
    MinFunds        float64 `json:"minFunds"`
}
