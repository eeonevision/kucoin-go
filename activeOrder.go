package kucoin

// ActiveMapOrder struct represents kucoin data model.
type ActiveMapOrder struct {
	SELL []struct {
		Oid           string      `json:"oid"`
		Type          string      `json:"type"`
		UserOid       interface{} `json:"userOid"`
		CoinType      string      `json:"coinType"`
		CoinTypePair  string      `json:"coinTypePair"`
		Direction     string      `json:"direction"`
		Price         float64     `json:"price"`
		DealAmount    float64     `json:"dealAmount"`
		PendingAmount float64     `json:"pendingAmount"`
		CreatedAt     int64       `json:"createdAt"`
		UpdatedAt     int64       `json:"updatedAt"`
	} `json:"SELL"`
	BUY []struct {
		Oid           string      `json:"oid"`
		Type          string      `json:"type"`
		UserOid       interface{} `json:"userOid"`
		CoinType      string      `json:"coinType"`
		CoinTypePair  string      `json:"coinTypePair"`
		Direction     string      `json:"direction"`
		Price         float64     `json:"price"`
		DealAmount    float64     `json:"dealAmount"`
		PendingAmount float64     `json:"pendingAmount"`
		CreatedAt     int64       `json:"createdAt"`
		UpdatedAt     int64       `json:"updatedAt"`
	} `json:"BUY"`
}

type rawActiveMapOrder struct {
	Success   bool           `json:"success"`
	Code      string         `json:"code"`
	Msg       string         `json:"msg,omitempty"`
	Timestamp int64          `json:"timestamp,omitempty"`
	Data      ActiveMapOrder `json:"data"`
}

// ActiveOrder struct represents kucoin data model.
type ActiveOrder struct {
	SELL [][]interface{} `json:"SELL"`
	BUY  [][]interface{} `json:"BUY"`
}

type rawActiveOrder struct {
	Comment   string      `json:"_comment"`
	Success   bool        `json:"success"`
	Code      string      `json:"code"`
	Msg       string      `json:"msg,omitempty"`
	Timestamp int64       `json:"timestamp,omitempty"`
	Data      ActiveOrder `json:"data"`
}
