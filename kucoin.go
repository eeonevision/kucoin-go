// Package Kucoin is an implementation of the Kucoin API in Golang.
package kucoin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	// Kucoin API endpoint
	API_BASE   = "https://api.kucoin.com"
	API_PREFIX = "/v1"
)

// New returns an instantiated Kucoin struct
func New(apiKey, apiSecret string) *Kucoin {
	client := NewClient(apiKey, apiSecret)
	return &Kucoin{client}
}

// NewWithCustomHttpClient returns an instantiated Kucoin struct with custom http client
func NewWithCustomHttpClient(apiKey, apiSecret string, httpClient *http.Client) *Kucoin {
	client := NewClientWithCustomHttpConfig(apiKey, apiSecret, httpClient)
	return &Kucoin{client}
}

// NewWithCustomTimeout returns an instantiated Kucoin struct with custom timeout
func NewWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) *Kucoin {
	client := NewClientWithCustomTimeout(apiKey, apiSecret, timeout)
	return &Kucoin{client}
}

// handleErr gets JSON response from livecoin API en deal with error
func handleErr(r interface{}) error {
	switch v := r.(type) {
	case map[string]interface{}:
		error := r.(map[string]interface{})["error"]
		if error != nil {
			switch v := error.(type) {
			case map[string]interface{}:
				errorMessage := error.(map[string]interface{})["message"]
				return errors.New(errorMessage.(string))
			default:
				return fmt.Errorf("I don't know about type %T!\n", v)
			}
		}
	case []interface{}:
		return nil
	default:
		return fmt.Errorf("I don't know about type %T!\n", v)
	}

	return nil
}

// Kucoin represent a Kucoin client
type Kucoin struct {
	client *client
}

// set enable/disable http request/response dump
func (c *Kucoin) SetDebug(enable bool) {
	c.client.debug = enable
}

// GetUserInfo is used to get the user information at Kucoin along with other meta data.
func (b *Kucoin) GetUserInfo() (userInfo UserInfo, err error) {
	r, err := b.client.do("GET", "user/info", nil, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var rawRes rawUserInfo
	err = json.Unmarshal(r, &rawRes)
	userInfo = rawRes.Data
	return
}

// GetSymbols is used to get the all open and available trading markets at Kucoin along with other meta data.
func (b *Kucoin) GetSymbols() (symbols []Symbol, err error) {
	r, err := b.client.do("GET", "market/open/symbols", nil, false)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var rawRes rawSymbols
	err = json.Unmarshal(r, &rawRes)
	symbols = rawRes.Data
	return
}

// GetSymbol is used to get the open and available trading market at Kucoin along with other meta data.
// Trading symbol e.g. KCS-BTC. If not specified then you will get data of all symbols.
func (b *Kucoin) GetSymbol(market string) (symbol Symbol, err error) {
	r, err := b.client.do("GET", "open/tick?symbol="+strings.ToUpper(market), nil, false)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var rawRes rawSymbol
	err = json.Unmarshal(r, &rawRes)
	symbol = rawRes.Data
	return
}

// GetCoins is used to get the all open and available trading coins at Kucoin along with other meta data.
func (b *Kucoin) GetCoins() (coins []Coin, err error) {
	r, err := b.client.do("GET", "market/open/coins", nil, false)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var rawRes rawCoins
	err = json.Unmarshal(r, &rawRes)
	coins = rawRes.Data
	return
}

// GetCoin is used to get the open and available trading coin at Kucoin along with other meta data.
func (b *Kucoin) GetCoin(c string) (coin Coin, err error) {
	r, err := b.client.do("GET", "market/open/coin-info?coin="+strings.ToUpper(c), nil, false)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var rawRes rawCoin
	err = json.Unmarshal(r, &rawRes)
	coin = rawRes.Data
	return
}

// GetCoinBalance is used to get the balance at chosen coin at Kucoin along with other meta data.
func (b *Kucoin) GetCoinBalance(c string) (coinBalance CoinBalance, err error) {
	r, err := b.client.do("GET", fmt.Sprintf("account/%s/balance", strings.ToUpper(c)), nil, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var rawRes rawCoinBalance
	err = json.Unmarshal(r, &rawRes)
	coinBalance = rawRes.Data
	return
}

// GetCoinDepositAddress is used to get the address at chosen coin at Kucoin along with other meta data.
func (b *Kucoin) GetCoinDepositAddress(c string) (coinDepositAddress CoinDepositAddress, err error) {
	r, err := b.client.do("GET", fmt.Sprintf("account/%s/wallet/address", strings.ToUpper(c)), nil, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var rawRes rawCoinDepositAddress
	err = json.Unmarshal(r, &rawRes)
	coinDepositAddress = rawRes.Data
	return
}

// ListActiveOrders is used to get the information about active orders at Kucoin along with other meta data.
// Symbol is required parameter, and side (or type of order) may be empty.
func (b *Kucoin) ListActiveOrders(symbol string, side string) (activeOrders ActiveOrders, err error) {
	r, err := b.client.do("GET", fmt.Sprintf("order/active?symbol=%s&type=%s", strings.ToUpper(symbol), strings.ToUpper(side)), nil, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var rawRes rawActiveOrders
	err = json.Unmarshal(r, &rawRes)
	activeOrders = rawRes.Data
	return
}

// Create is used to create order at Kucoin along with other meta data.
func (b *Kucoin) CreateOrder(symbol, side string, price, amount float64) (orderOid string, err error) {
	payload := make(map[string]string)
	payload["amount"] = strconv.FormatFloat(amount, 'f', 8, 64)
	payload["price"] = strconv.FormatFloat(price, 'f', 8, 64)
	payload["type"] = strings.ToUpper(side)

	r, err := b.client.do("POST", fmt.Sprintf("order?symbol=%s", strings.ToUpper(symbol)), payload, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var rawRes rawOrder
	err = json.Unmarshal(r, &rawRes)
	orderOid = rawRes.Data.OrderOid
	return
}
