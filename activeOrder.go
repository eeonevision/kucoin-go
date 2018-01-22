package kucoin

type ActiveOrders struct {
	SELL []interface{}   `json:"SELL"`
	BUY  [][]interface{} `json:"BUY"`
}

type rawActiveOrders struct {
	Comment string       `json:"_comment"`
	Success bool         `json:"success"`
	Code    string       `json:"code"`
	Data    ActiveOrders `json:"data"`
}
