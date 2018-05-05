package kucoin

// OrderDetails structs represents kucoin data model.
type OrderDetails struct {
	CoinType         string  `json:"coinType"`
	DealValueTotal   float64 `json:"dealValueTotal"`
	DealPriceAverage float64 `json:"dealPriceAverage"`
	FeeTotal         float64 `json:"feeTotal"`
	UserOid          string  `json:"userOid"`
	DealAmount       float64 `json:"dealAmount"`
	DealOrders       struct {
		Total     int  `json:"total"`
		FirstPage bool `json:"firstPage"`
		LastPage  bool `json:"lastPage"`
		Datas     []struct {
			Amount    float64 `json:"amount"`
			DealValue float64 `json:"dealValue"`
			Fee       float64 `json:"fee"`
			DealPrice float64 `json:"dealPrice"`
			FeeRate   float64 `json:"feeRate"`
		} `json:"datas"`
		CurrPageNo int `json:"currPageNo"`
		Limit      int `json:"limit"`
		PageNos    int `json:"pageNos"`
	} `json:"dealOrders"`
	CoinTypePair  string  `json:"coinTypePair"`
	OrderPrice    float64 `json:"orderPrice"`
	Type          string  `json:"type"`
	OrderOid      string  `json:"orderOid"`
	PendingAmount float64 `json:"pendingAmount"`
}

type rawOrderDetails struct {
	Success   bool         `json:"success"`
	Code      string       `json:"code"`
	Msg       string       `json:"msg"`
	Timestamp int64        `json:"timestamp"`
	Data      OrderDetails `json:"data"`
}
