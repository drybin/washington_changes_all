package config

import (
	"github.com/drybin/washington_changes_all/internal/domain/types/coin_name"
)

type Coin struct {
	Name coin_name.CoinName
	ATH  float64
}

type CoinConfig struct {
	TierOne   []Coin
	TierTwo   []Coin
	TierThree []Coin
}

func InitCoinConfig() CoinConfig {
	config := CoinConfig{
		TierOne: []Coin{
			{
				Name: coin_name.BTC,
				ATH:  109500.0,
			},
			{
				Name: coin_name.ETH,
				ATH:  4891.0,
			},
			{
				Name: coin_name.SOL,
				ATH:  263.0,
			},
			{
				Name: coin_name.BNB,
				ATH:  793.0,
			},
			{
				Name: coin_name.XRP,
				ATH:  3.84,
			},
			{
				Name: coin_name.DOGE,
				ATH:  0.73,
			},
			{
				Name: coin_name.ADA,
				ATH:  3.10,
			},
			{
				Name: coin_name.AVAX,
				ATH:  146.22,
			},
			{
				Name: coin_name.TON,
				ATH:  8.24,
			},
			{
				Name: coin_name.DOT,
				ATH:  55.0,
			},
			{
				Name: coin_name.LINK,
				ATH:  53.88,
			},
			{
				Name: coin_name.MATIC,
				ATH:  2.92,
			},
			{
				Name: coin_name.ATOM,
				ATH:  44.7,
			},
			{
				Name: coin_name.TRON,
				ATH:  0.44,
			},
			{
				Name: coin_name.NEAR,
				ATH:  20.42,
			},
			{
				Name: coin_name.XLM,
				ATH:  0.93,
			},
			{
				Name: coin_name.SUI,
				ATH:  2.18,
			},
		},
		TierTwo: []Coin{
			{
				Name: coin_name.LTC,
				ATH:  412.96,
			},
			{
				Name: coin_name.APT,
				ATH:  19.9,
			},
			{
				Name: coin_name.OP,
				ATH:  4.85,
			},
			{
				Name: coin_name.ARB,
				ATH:  2.4,
			},
			{
				Name: coin_name.BCH,
				ATH:  4355,
			},
			{
				Name: coin_name.ICP,
				ATH:  750.0,
			},
			{
				Name: coin_name.UNI,
				ATH:  44.97,
			},
			{
				Name: coin_name.FIL,
				ATH:  237.24,
			},
			{
				Name: coin_name.ETC,
				ATH:  176.1,
			},
			{
				Name: coin_name.INJ,
				ATH:  52.7,
			},
			{
				Name: coin_name.XMR,
				ATH:  517.6,
			},
			{
				Name: coin_name.ALGO,
				ATH:  3.28,
			},
			{
				Name: coin_name.FLOW,
				ATH:  46.16,
			},
			{
				Name: coin_name.AAVE,
				ATH:  666.8,
			},
			{
				Name: coin_name.EGLD,
				ATH:  542.5,
			},
			{
				Name: coin_name.AXS,
				ATH:  165.3,
			},
			{
				Name: coin_name.SAND,
				ATH:  8.44,
			},
			{
				Name: coin_name.WLD,
				ATH:  11.82,
			},
			{
				Name: coin_name.MANA,
				ATH:  5.9,
			},
			{
				Name: coin_name.EOS,
				ATH:  22.89,
			},
			{
				Name: coin_name.NEO,
				ATH:  196.8,
			},
			{
				Name: coin_name.KCS,
				ATH:  28.8,
			},
			{
				Name: coin_name.IOTA,
				ATH:  5.6,
			},
			{
				Name: coin_name.INCH,
				ATH:  7.87,
			},
			{
				Name: coin_name.COMP,
				ATH:  911.2,
			},
			{
				Name: coin_name.ZIL,
				ATH:  0.25,
			},
			{
				Name: coin_name.DASH,
				ATH:  1642.0,
			},
			{
				Name: coin_name.SUSHI,
				ATH:  23.3,
			},
			{
				Name: coin_name.HNT,
				ATH:  55.22,
			},
			{
				Name: coin_name.CRO,
				ATH:  0.96,
			},
			{
				Name: coin_name.LDO,
				ATH:  11.0,
			},
			{
				Name: coin_name.RAY,
				ATH:  16.93,
			},
			{
				Name: coin_name.BTT,
				ATH:  0.000003054,
			},
			{
				Name: coin_name.JUP,
				ATH:  2.04,
			},
			{
				Name: coin_name.FTM,
				ATH:  3.48,
			},
			{
				Name: coin_name.MNT,
				ATH:  1.51,
			},
			{
				Name: coin_name.ORDI,
				ATH:  96.17,
			},
			{
				Name: coin_name.GRT,
				ATH:  2.88,
			},
			{
				Name: coin_name.HBAR,
				ATH:  0.57,
			},
		},
		TierThree: []Coin{
			{
				Name: coin_name.RVN,
				ATH:  0.28,
			},
			{
				Name: coin_name.STX,
				ATH:  3.84,
			},
			{
				Name: coin_name.IMX,
				ATH:  9.5,
			},
			{
				Name: coin_name.RNDR,
				ATH:  13.6,
			},
			{
				Name: coin_name.TAO,
				ATH:  767.6,
			},
			{
				Name: coin_name.MKR,
				ATH:  6339.02,
			},
			{
				Name: coin_name.VET,
				ATH:  0.27,
			},
			{
				Name: coin_name.KAS,
				ATH:  0.18,
			},
			{
				Name: coin_name.THETA,
				ATH:  15.9,
			},
			{
				Name: coin_name.RUNE,
				ATH:  21.26,
			},
			{
				Name: coin_name.FET,
				ATH:  3.47,
			},
			{
				Name: coin_name.AR,
				ATH:  90.9,
			},
			{
				Name: coin_name.TIA,
				ATH:  20.91,
			},
			{
				Name: coin_name.SEI,
				ATH:  1.14,
			},
			{
				Name: coin_name.GALA,
				ATH:  0.83,
			},
			{
				Name: coin_name.BSV,
				ATH:  491.6,
			},
			{
				Name: coin_name.QNT,
				ATH:  428.3,
			},
			{
				Name: coin_name.CFX,
				ATH:  1.7,
			},
			{
				Name: coin_name.FLR,
				ATH:  0.07,
			},
			{
				Name: coin_name.STRK,
				ATH:  3.66,
			},
			{
				Name: coin_name.AGIX,
				ATH:  1.86,
			},
			{
				Name: coin_name.PYTH,
				ATH:  1.15,
			},
			{
				Name: coin_name.ONDO,
				ATH:  1.05,
			},
			{
				Name: coin_name.LUNC,
				ATH:  119.18,
			},
			{
				Name: coin_name.CRV,
				ATH:  60.5,
			},
			{
				Name: coin_name.LUNA,
				ATH:  19.54,
			},
			{
				Name: coin_name.OCEAN,
				ATH:  1.94,
			},
			{
				Name: coin_name.ENJ,
				ATH:  4.85,
			},
			{
				Name: coin_name.ANKR,
				ATH:  0.22,
			},
			{
				Name: coin_name.BAT,
				ATH:  1.92,
			},
			{
				Name: coin_name.CSPR,
				ATH:  1.36,
			},
			{
				Name: coin_name.ONE,
				ATH:  0.37,
			},
			{
				Name: coin_name.HFT,
				ATH:  2.58,
			},
			{
				Name: coin_name.TFUEL,
				ATH:  0.68,
			},
			{
				Name: coin_name.VELO,
				ATH:  2.07,
			},
			{
				Name: coin_name.ENA,
				ATH:  1.52,
			},
			{
				Name: coin_name.TAO,
				ATH:  767.8,
			},
			{
				Name: coin_name.SHIB,
				ATH:  0.00008845,
			},
			{
				Name: coin_name.WIF,
				ATH:  4.85,
			},
			{
				Name: coin_name.PEPE,
				ATH:  0.00001074,
			},
			{
				Name: coin_name.FLOKI,
				ATH:  0.0003437,
			},
			{
				Name: coin_name.BONK,
				ATH:  0.00004704,
			},
			{
				Name: coin_name.BOME,
				ATH:  0.02805,
			},
			{
				Name: coin_name.MEME,
				ATH:  0.05706,
			},
			{
				Name: coin_name.BABYDOGE,
				ATH:  0.000000006356,
			},
			{
				Name: coin_name.DEGEN,
				ATH:  0.09459,
			},
			{
				Name: coin_name.MEW,
				ATH:  0.01034,
			},
			{
				Name: coin_name.VINU,
				ATH:  0.00000002703,
			},
		},
	}

	return config
}
