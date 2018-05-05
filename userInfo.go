package kucoin

// UserInfo struct represents kucoin data model.
type UserInfo struct {
	ReferrerCode             string      `json:"referrer_code"`
	PhotoCredentialValidated bool        `json:"photoCredentialValidated"`
	VideoValidated           bool        `json:"videoValidated"`
	Language                 string      `json:"language"`
	Currency                 string      `json:"currency"`
	Oid                      string      `json:"oid"`
	BaseFeeRate              float64     `json:"baseFeeRate"`
	HasCredential            bool        `json:"hasCredential"`
	CredentialNumber         string      `json:"credentialNumber"`
	PhoneValidated           bool        `json:"phoneValidated"`
	Phone                    string      `json:"phone"`
	CredentialValidated      bool        `json:"credentialValidated"`
	GoogleTwoFaBinding       bool        `json:"googleTwoFaBinding"`
	Nickname                 interface{} `json:"nickname"`
	Name                     string      `json:"name"`
	HasTradePassword         bool        `json:"hasTradePassword"`
	EmailValidated           bool        `json:"emailValidated"`
	Email                    string      `json:"email"`
	LoginRecord              struct {
		Last struct {
			IP      string      `json:"ip"`
			Context interface{} `json:"context"`
			Time    int64       `json:"time"`
		} `json:"last"`
		Current struct {
			IP      string      `json:"ip"`
			Context interface{} `json:"context"`
			Time    int64       `json:"time"`
		} `json:"current"`
	} `json:"loginRecord"`
}

type rawUserInfo struct {
	Data UserInfo `json:"data"`
}
