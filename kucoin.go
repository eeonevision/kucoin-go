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
func (b *Kucoin) SetDebug(enable bool) {
	b.client.debug = enable
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
// Symbol is required parameter, and side (or type of order in kucoin docs) may be empty.
func (b *Kucoin) ListActiveOrders(symbol string, side string) (activeOrders ActiveOrders, err error) {
	if len(symbol) < 1 {
		return activeOrders, fmt.Errorf("The symbol is required")
	}
	payload := map[string]string{}
	payload["symbol"] = strings.ToUpper(symbol)
	if len(side) < 1 {
		payload["side"] = strings.ToUpper(side)
	}

	r, err := b.client.do("GET", "order/active", payload, true)
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

// AccountHistory is used to get the information about list deposit & withdrawal at Kucoin along with other meta data.
// Coin, Side (type in Kucoin docs.) and Status are required parameters. Limit and page may be zeros.
// Example:
// - Coin = KCS
// - Side = DEPOSIT | WITHDRAW
// - Status = FINISHED | CANCEL | PENDING
func (b *Kucoin) AccountHistory(coin, side, status string, limit, page int) (accountHistory AccountHistory, err error) {
	if len(coin) < 1 || len(side) < 1 || len(status) < 1 {
		return accountHistory, fmt.Errorf("The not all required parameters are presented")
	}
	payload := map[string]string{}
	payload["type"] = side
	payload["status"] = status
	if limit == 0 {
		payload["limit"] = fmt.Sprintf("%v", 1000)
	} else {
		payload["limit"] = fmt.Sprintf("%v", limit)
	}
	if page != 0 {
		payload["page"] = fmt.Sprintf("%v", page)
	}

	r, err := b.client.do("GET", fmt.Sprintf(
		"account/%s/wallet/records", strings.ToUpper(coin)), payload, true)
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
	var rawRes AccountHistory
	err = json.Unmarshal(r, &rawRes)
	accountHistory = rawRes
	return
}

// ListSpecificDealtOrders is used to get the information about dealt orders for specific symbol at Kucoin along with other meta data.
// Symbol, Side (type in Kucoin docs.) are required parameters. Limit and page may be zeros.
// Example:
// - Symbol = KCS-BTC
// - Side = BUY | SELL
func (b *Kucoin) ListSpecificDealtOrders(symbol, side string, limit, page int) (specificDealtOrders SpecificDealtOrder, err error) {
	if len(symbol) < 1 || len(side) < 1 {
		return specificDealtOrders, fmt.Errorf("The not all required parameters are presented")
	}
	payload := map[string]string{}
	payload["symbol"] = symbol
	payload["type"] = side
	if limit == 0 {
		payload["limit"] = fmt.Sprintf("%v", 1000)
	} else {
		payload["limit"] = fmt.Sprintf("%v", limit)
	}
	if page != 0 {
		payload["page"] = fmt.Sprintf("%v", page)
	}

	r, err := b.client.do("GET", "deal-orders", payload, true)
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
	var rawRes SpecificDealtOrder
	err = json.Unmarshal(r, &rawRes)
	specificDealtOrders = rawRes
	return
}

// ListMergedDealtOrders is used to get the information about dealt orders for all symbols at Kucoin along with other meta data.
// All parameters are optional. Timestamp must be in milliseconds from Unix epoch.
func (b *Kucoin) ListMergedDealtOrders(symbol, side string, limit, page int, since, before int64) (mergedDealtOrders MergedDealtOrder, err error) {
	payload := map[string]string{}
	payload["symbol"] = symbol
	payload["type"] = side
	if len(symbol) > 1 {
		payload["symbol"] = symbol
	}
	if len(side) > 1 {
		payload["type"] = side
	}
	if (limit == 0 || limit > 100) && len(symbol) > 1 {
		payload["limit"] = fmt.Sprintf("%v", 100)
	} else if (limit == 0 || limit > 20) && len(symbol) > 1 {
		payload["limit"] = fmt.Sprintf("%v", 20)
	} else {
		payload["limit"] = fmt.Sprintf("%v", limit)
	}
	if page != 0 {
		payload["page"] = fmt.Sprintf("%v", page)
	}
	if since != 0 {
		payload["since"] = fmt.Sprintf("%v", since)
	}
	if before != 0 {
		payload["before"] = fmt.Sprintf("%v", before)
	}

	r, err := b.client.do("GET", "order/dealt", payload, true)
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
	var rawRes MergedDealtOrder
	err = json.Unmarshal(r, &rawRes)
	mergedDealtOrders = rawRes
	return
}
