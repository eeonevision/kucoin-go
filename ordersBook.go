package kucoin

// OrdersBook struct represents kucoin data model.
type OrdersBook struct {
	Comment string      `json:"_comment"`
	SELL    [][]float64 `json:"SELL"`
	BUY     [][]float64 `json:"BUY"`
}

type rawOrdersBook struct {
	Success bool       `json:"success"`
	Code    string     `json:"code"`
	Msg     string     `json:"msg"`
	Data    OrdersBook `json:"data"`
}
