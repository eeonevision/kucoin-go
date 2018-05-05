package kucoin

// AccountHistory struct represents kucoin data model.
type AccountHistory struct {
	Datas []struct {
		Fee             float64     `json:"fee"`
		Oid             string      `json:"oid"`
		Type            string      `json:"type"`
		Amount          float64     `json:"amount"`
		Remark          string      `json:"remark"`
		Status          string      `json:"status"`
		Address         string      `json:"address"`
		Context         string      `json:"context"`
		UserOid         string      `json:"userOid"`
		CoinType        string      `json:"coinType"`
		CreatedAt       int64       `json:"createdAt"`
		DeletedAt       interface{} `json:"deletedAt"`
		UpdatedAt       int64       `json:"updatedAt"`
		OuterWalletTxid interface{} `json:"outerWalletTxid"`
	} `json:"datas"`
	Total           int         `json:"total"`
	Limit           int         `json:"limit"`
	PageNos         int         `json:"pageNos"`
	CurrPageNo      int         `json:"currPageNo"`
	NavigatePageNos []int       `json:"navigatePageNos"`
	CoinType        string      `json:"coinType"`
	Type            interface{} `json:"type"`
	UserOid         string      `json:"userOid"`
	Status          interface{} `json:"status"`
	FirstPage       bool        `json:"firstPage"`
	LastPage        bool        `json:"lastPage"`
	StartRow        int         `json:"startRow"`
}

type rawAccountHistory struct {
	Success bool           `json:"success"`
	Code    string         `json:"code"`
	Data    AccountHistory `json:"data"`
}
