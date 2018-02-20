package kucoin

type Withdrawal struct {
}

type rawWithdrawal struct {
	Success   bool       `json:"success, omitempty"`
	Code      string     `json:"code, omitempty"`
	Msg       string     `json:"msg, omitempty"`
	Timestamp int64      `json:"timestamp, omitempty"`
	Data      Withdrawal `json:"data, omitempty"`
}
