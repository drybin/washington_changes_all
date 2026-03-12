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
                Name: coin_name.AVAX,
                ATH:  146.22,
            },
            {
                Name: coin_name.TON,
                ATH:  8.24,
            },
            {
                Name: coin_name.LINK,
                ATH:  53.88,
            },
            {
                Name: coin_name.TRON,
                ATH:  0.44,
            },
            {
                Name: coin_name.SUI,
                ATH:  5.35,
            },
        },
        TierTwo: []Coin{
            {
                Name: coin_name.ADA,
                ATH:  3.10,
            },
            {
                Name: coin_name.NEAR,
                ATH:  20.42,
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
                Name: coin_name.UNI,
                ATH:  44.97,
            },
            {
                Name: coin_name.INJ,
                ATH:  52.7,
            },
            {
                Name: coin_name.AAVE,
                ATH:  666.8,
            },
            {
                Name: coin_name.LDO,
                ATH:  11.0,
            },
            {
                Name: coin_name.JUP,
                ATH:  2.04,
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
                Name: coin_name.HBAR,
                ATH:  0.57,
            },
            {
                Name: coin_name.DASH,
                ATH:  1642.0,
            },
            {
                Name: coin_name.HYPE,
                ATH:  35.0,
            },
            {
                Name: coin_name.PENDLE,
                ATH:  7.5,
            },
            {
                Name: coin_name.JTO,
                ATH:  4.5,
            },
        },
        TierThree: []Coin{
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
                Name: coin_name.KAS,
                ATH:  0.18,
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
                Name: coin_name.STRK,
                ATH:  3.66,
            },
            {
                Name: coin_name.ONDO,
                ATH:  1.05,
            },
            {
                Name: coin_name.ENA,
                ATH:  1.52,
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
                Name: coin_name.BONK,
                ATH:  0.00004704,
            },
            {
                Name: coin_name.VIRTUAL,
                ATH:  5.0,
            },
            {
                Name: coin_name.IO,
                ATH:  6.0,
            },
            {
                Name: coin_name.AKT,
                ATH:  6.5,
            },
        },
    }

    return config
}
