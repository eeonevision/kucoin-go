package kucoin

type CoinBalance struct {
	CoinType      string `json:"coinType"`
	Balance       int    `json:"balance"`
	FreezeBalance int    `json:"freezeBalance"`
}

type rawCoinBalances struct {
	Success bool          `json:"success"`
	Code    string        `json:"code"`
	Data    []CoinBalance `json:"data"`
}

type rawCoinBalance struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Data    CoinBalance `json:"data"`
}
