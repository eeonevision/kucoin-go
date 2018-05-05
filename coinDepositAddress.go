package kucoin

// CoinDepositAddress struct represents kucoin data model.
type CoinDepositAddress struct {
	Oid            string      `json:"oid"`
	Address        string      `json:"address"`
	Context        interface{} `json:"context"`
	UserOid        string      `json:"userOid"`
	CoinType       string      `json:"coinType"`
	CreatedAt      int64       `json:"createdAt"`
	DeletedAt      interface{} `json:"deletedAt"`
	UpdatedAt      int64       `json:"updatedAt"`
	LastReceivedAt int64       `json:"lastReceivedAt"`
}

type rawCoinDepositAddress struct {
	Data CoinDepositAddress `json:"data"`
}
