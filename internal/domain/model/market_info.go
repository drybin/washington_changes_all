package model

type MarketInfo struct {
	Pair             CoinsPair
	SymbolName       string
	Buy              float64 //buy
	Sell             float64 //sell
	ChangeRate       float64 // 24h change rate
	ChangePrice      float64 // 24h change price
	HighPrice        float64 // 24h highest price
	LowPrice         float64 // 24h lowest price
	VolBTC           float64 // 24h volume，the aggregated trading volume in BTC
	VolValue         float64 // 24h total, the trading volume in quote currency of last 24 hours
	LastPrice        float64 // last price
	AveragePrice     float64 // 24h average transaction price yesterday
	TakerFeeRate     float64 // Basic Taker Fee
	MakerFeeRate     float64 // Basic Maker Fee
	TakerCoefficient float64 // Taker Fee Coefficient
	MakerCoefficient float64 // Maker Fee Coefficient
}
