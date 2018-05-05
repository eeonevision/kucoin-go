package kucoin

// CoinBalance struct represents kucoin data model.
type CoinBalance struct {
	CoinType      string  `json:"coinType"`
	Balance       float64 `json:"balance"`
	FreezeBalance float64 `json:"freezeBalance"`
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
