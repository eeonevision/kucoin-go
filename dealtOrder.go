package kucoin

type SpecificDealtOrder struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	Data    struct {
		Datas []struct {
			Oid       string `json:"oid"`
			DealPrice int    `json:"dealPrice"`
			OrderOid  string `json:"orderOid"`
			Direction string `json:"direction"`
			Amount    int    `json:"amount"`
			DealValue int    `json:"dealValue"`
			CreatedAt int64  `json:"createdAt"`
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
	} `json:"data"`
}

type MergedDealtOrder struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
	Data      struct {
		Total int `json:"total"`
		Datas []struct {
			CreatedAt     int64   `json:"createdAt"`
			Amount        float64 `json:"amount"`
			DealValue     float64 `json:"dealValue"`
			DealPrice     float64 `json:"dealPrice"`
			Fee           float64 `json:"fee"`
			FeeRate       int     `json:"feeRate"`
			Oid           string  `json:"oid"`
			OrderOid      string  `json:"orderOid"`
			CoinType      string  `json:"coinType"`
			CoinTypePair  string  `json:"coinTypePair"`
			Direction     string  `json:"direction"`
			DealDirection string  `json:"dealDirection"`
		} `json:"datas"`
		Limit int `json:"limit"`
		Page  int `json:"page"`
	} `json:"data"`
}
