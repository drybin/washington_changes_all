package coin_name

type CoinName string

const (
	//Stable
	BTC   CoinName = "BTC"
	ETH   CoinName = "ETH"
	SOL   CoinName = "SOL"
	BNB   CoinName = "BNB"
	XRP   CoinName = "XRP"
	DOGE  CoinName = "DOGE"
	ADA   CoinName = "ADA"
	AVAX  CoinName = "AVAX"
	TON   CoinName = "TON"
	DOT   CoinName = "DOT"
	LINK  CoinName = "LINK"
	MATIC CoinName = "MATIC"
	ATOM  CoinName = "ATOM"
	LTC   CoinName = "LTC"
	TRON  CoinName = "TRX"
	APT   CoinName = "APT"
	OP    CoinName = "OP"
	NEAR  CoinName = "NEAR"
	ARB   CoinName = "ARB"

	//Middle
	BCH   CoinName = "BCH"
	ICP   CoinName = "ICP"
	UNI   CoinName = "UNI"
	FIL   CoinName = "FIL"
	ETC   CoinName = "ETC"
	XLM   CoinName = "XLM"
	INJ   CoinName = "INJ"
	XMR   CoinName = "XMR"
	SUI   CoinName = "SUI"
	ALGO  CoinName = "ALGO"
	FLOW  CoinName = "FLOW"
	AAVE  CoinName = "AAVE"
	EGLD  CoinName = "EGLD"
	AXS   CoinName = "AXS"
	SAND  CoinName = "SAND"
	WLD   CoinName = "WLD"
	MANA  CoinName = "MANA"
	EOS   CoinName = "EOS"
	KCS   CoinName = "KCS"
	NEO   CoinName = "NEO"
	IOTA  CoinName = "IOTA"
	INCH  CoinName = "1INCH"
	COMP  CoinName = "COMP"
	ZIL   CoinName = "ZIL"
	DASH  CoinName = "DASH"
	SUSHI CoinName = "SUSHI"
	HNT   CoinName = "HNT"
	CRO   CoinName = "CRO"
	LDO   CoinName = "LDO"
	RAY   CoinName = "RAY"
	BTT   CoinName = "BTT"
	JUP   CoinName = "JUP"
	FTM   CoinName = "FTM"
	MNT   CoinName = "MNT"
	ORDI  CoinName = "ORDI"
	GRT   CoinName = "GRT"
	RVN   CoinName = "RVN"
	//High
	STX   CoinName = "STX"
	IMX   CoinName = "IMX"
	RNDR  CoinName = "RNDR"
	TAO   CoinName = "TAO"
	HBAR  CoinName = "HBAR"
	MKR   CoinName = "MKR"
	VET   CoinName = "VET"
	KAS   CoinName = "KAS"
	THETA CoinName = "THETA"
	RUNE  CoinName = "RUNE"
	FET   CoinName = "FET"
	AR    CoinName = "AR"
	TIA   CoinName = "TIA"
	SEI   CoinName = "SEI"
	GALA  CoinName = "GALAX"
	BSV   CoinName = "BCHSV"
	QNT   CoinName = "QNT"
	CFX   CoinName = "CFX"
	FLR   CoinName = "FLR"
	STRK  CoinName = "STRK"
	AGIX  CoinName = "AGIX"
	PYTH  CoinName = "PYTH"
	ONDO  CoinName = "ONDO"
	APE   CoinName = "APE"
	KAVA  CoinName = "KAVA"
	KLAY  CoinName = "KLAY"
	OSMO  CoinName = "OSMO"
	LUNC  CoinName = "LUNC"
	CRV   CoinName = "CRV"
	LUNA  CoinName = "LUNA"
	OCEAN CoinName = "OCEAN"
	ENJ   CoinName = "ENJ"
	ANKR  CoinName = "ANKR"
	BAT   CoinName = "BAT"
	CSPR  CoinName = "CSPR"
	ONE   CoinName = "ONE"
	HFT   CoinName = "HFT"
	TFUEL CoinName = "TFUEL"
	VELO  CoinName = "VELO"
	ENA   CoinName = "ENA"
	//Meme
	SHIB     CoinName = "SHIB"
	WIF      CoinName = "WIF"
	PEPE     CoinName = "PEPE"
	FLOKI    CoinName = "FLOKI"
	BONK     CoinName = "BONK"
	BOME     CoinName = "BOME"
	MEME     CoinName = "MEME"
	BABYDOGE CoinName = "BABYDOGE"
	DEGEN    CoinName = "DEGEN"
	MEW      CoinName = "MEW"

	USDT    CoinName = "USDT"
	UNKNOWN CoinName = "UNKNOWN"
)

func (t CoinName) String() string {
	return string(t)
}

func FromString(s string) CoinName {
	switch s {
	case BTC.String():
		return BTC
	case ETH.String():
		return ETH
	case SOL.String():
		return SOL
	case BNB.String():
		return BNB
	case XRP.String():
		return XRP
	case DOGE.String():
		return DOGE
	case ADA.String():
		return ADA
	case AVAX.String():
		return AVAX
	case TON.String():
		return TON
	case DOT.String():
		return DOT
	case LINK.String():
		return LINK
	case MATIC.String():
		return MATIC
	case ATOM.String():
		return ATOM
	case LTC.String():
		return LTC
	case TRON.String():
		return TRON
	case APT.String():
		return APT
	case OP.String():
		return OP
	case NEAR.String():
		return NEAR
	case ARB.String():
		return ARB
	case BCH.String():
		return BCH
	case ICP.String():
		return ICP
	case UNI.String():
		return UNI
	case FIL.String():
		return FIL
	case ETC.String():
		return ETC
	case XLM.String():
		return XLM
	case INJ.String():
		return INJ
	case XMR.String():
		return XMR
	case SUI.String():
		return SUI
	case ALGO.String():
		return ALGO
	case FLOW.String():
		return FLOW
	case AAVE.String():
		return AAVE
	case EGLD.String():
		return EGLD
	case AXS.String():
		return AXS
	case SAND.String():
		return SAND
	case WLD.String():
		return WLD
	case MANA.String():
		return MANA
	case EOS.String():
		return EOS
	case KCS.String():
		return KCS
	case NEO.String():
		return NEO
	case IOTA.String():
		return IOTA
	case INCH.String():
		return INCH
	case COMP.String():
		return COMP
	case ZIL.String():
		return ZIL
	case DASH.String():
		return DASH
	case SUSHI.String():
		return SUSHI
	case HNT.String():
		return HNT
	case CRO.String():
		return CRO
	case LDO.String():
		return LDO
	case RAY.String():
		return RAY
	case BTT.String():
		return BTT
	case JUP.String():
		return JUP
	case FTM.String():
		return FTM
	case MNT.String():
		return MNT
	case ORDI.String():
		return ORDI
	case GRT.String():
		return GRT
	case RVN.String():
		return RVN
	case STX.String():
		return STX
	case IMX.String():
		return IMX
	case RNDR.String():
		return RNDR
	case TAO.String():
		return TAO
	case HBAR.String():
		return HBAR
	case MKR.String():
		return MKR
	case VET.String():
		return VET
	case KAS.String():
		return KAS
	case THETA.String():
		return THETA
	case RUNE.String():
		return RUNE
	case FET.String():
		return FET
	case AR.String():
		return AR
	case TIA.String():
		return TIA
	case SEI.String():
		return SEI
	case GALA.String():
		return GALA
	case BSV.String():
		return BSV
	case QNT.String():
		return QNT
	case CFX.String():
		return CFX
	case FLR.String():
		return FLR
	case STRK.String():
		return STRK
	case AGIX.String():
		return AGIX
	case PYTH.String():
		return PYTH
	case ONDO.String():
		return ONDO
	case APE.String():
		return APE
	case KAVA.String():
		return KAVA
	case KLAY.String():
		return KLAY
	case OSMO.String():
		return OSMO
	case LUNC.String():
		return LUNC
	case CRV.String():
		return CRV
	case LUNA.String():
		return LUNA
	case OCEAN.String():
		return OCEAN
	case ENJ.String():
		return ENJ
	case ANKR.String():
		return ANKR
	case BAT.String():
		return BAT
	case CSPR.String():
		return CSPR
	case ONE.String():
		return ONE
	case HFT.String():
		return HFT
	case TFUEL.String():
		return TFUEL
	case VELO.String():
		return VELO
	case ENA.String():
		return ENA
	case SHIB.String():
		return SHIB
	case WIF.String():
		return WIF
	case PEPE.String():
		return PEPE
	case FLOKI.String():
		return FLOKI
	case BONK.String():
		return BONK
	case BOME.String():
		return BOME
	case MEME.String():
		return MEME
	case BABYDOGE.String():
		return BABYDOGE
	case DEGEN.String():
		return DEGEN
	case MEW.String():
		return MEW
	case USDT.String():
		return USDT
	}

	return UNKNOWN
}
