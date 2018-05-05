package kucoin

// SpecificDealtOrder struct represents kucoin data model.
type SpecificDealtOrder struct {
	Datas []struct {
		Oid       string  `json:"oid"`
		DealPrice float64 `json:"dealPrice"`
		OrderOid  string  `json:"orderOid"`
		Direction string  `json:"direction"`
		Amount    float64 `json:"amount"`
		DealValue float64 `json:"dealValue"`
		CreatedAt int64   `json:"createdAt"`
	} `json:"datas"`
	Total           int         `json:"total"`
	Limit           int         `json:"limit"`
	PageNos         int         `json:"pageNos"`
	CurrPageNo      int         `json:"currPageNo"`
	NavigatePageNos []int       `json:"navigatePageNos"`
	UserOid         string      `json:"userOid"`
	Direction       interface{} `json:"direction"`
	StartRow        int         `json:"startRow"`
	FirstPage       bool        `json:"firstPage"`
	LastPage        bool        `json:"lastPage"`
}

type rawSpecificDealtOrder struct {
	Success bool               `json:"success"`
	Code    string             `json:"code"`
	Msg     string             `json:"msg"`
	Data    SpecificDealtOrder `json:"data"`
}

// MergedDealtOrder struct represents kucoin data model.
type MergedDealtOrder struct {
	Total int `json:"total"`
	Datas []struct {
		CreatedAt     int64   `json:"createdAt"`
		Amount        float64 `json:"amount"`
		DealValue     float64 `json:"dealValue"`
		DealPrice     float64 `json:"dealPrice"`
		Fee           float64 `json:"fee"`
		FeeRate       float64 `json:"feeRate"`
		Oid           string  `json:"oid"`
		OrderOid      string  `json:"orderOid"`
		CoinType      string  `json:"coinType"`
		CoinTypePair  string  `json:"coinTypePair"`
		Direction     string  `json:"direction"`
		DealDirection string  `json:"dealDirection"`
	} `json:"datas"`
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type rawMergedDealtOrder struct {
	Success   bool             `json:"success"`
	Code      string           `json:"code"`
	Msg       string           `json:"msg"`
	Timestamp int64            `json:"timestamp"`
	Data      MergedDealtOrder `json:"data"`
}
