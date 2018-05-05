package kucoin

// Coin struct represents kucoin data model.
type Coin struct {
	WithdrawMinFee    float64     `json:"withdrawMinFee"`
	WithdrawMinAmount float64     `json:"withdrawMinAmount"`
	WithdrawFeeRate   float64     `json:"withdrawFeeRate"`
	ConfirmationCount int         `json:"confirmationCount"`
	WithdrawRemark    string      `json:"withdrawRemark"`
	InfoURL           interface{} `json:"infoUrl"`
	Name              string      `json:"name"`
	TradePrecision    int         `json:"tradePrecision"`
	DepositRemark     interface{} `json:"depositRemark"`
	EnableWithdraw    bool        `json:"enableWithdraw"`
	EnableDeposit     bool        `json:"enableDeposit"`
	Coin              string      `json:"coin"`
}

type rawCoins struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
	Data      []Coin `json:"data"`
}

type rawCoin struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
	Data      Coin   `json:"data"`
}
