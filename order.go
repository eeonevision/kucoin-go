package kucoin

// Order structs represents kucoin data model.
type Order struct {
	OrderOid string `json:"orderOid"`
}

type rawOrder struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	Data    Order  `json:"data"`
}
