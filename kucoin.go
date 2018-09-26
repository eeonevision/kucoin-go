// Package kucoin is an implementation of the Kucoin API in Golang.
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
	kucoinUrl = "https://api.kucoin.com/v1/"
)

// New returns an instantiated Kucoin struct.
func New(apiKey, apiSecret string) *Kucoin {
	client := newClient(apiKey, apiSecret)
	return &Kucoin{client}
}

// NewCustomClient returns an instantiated Kucoin struct with custom http client.
func NewCustomClient(apiKey, apiSecret string, httpClient http.Client) *Kucoin {
	client := newClient(apiKey, apiSecret)
	client.httpClient = httpClient
	return &Kucoin{client}
}

// NewCustomTimeout returns an instantiated Kucoin struct with custom timeout.
func NewCustomTimeout(apiKey, apiSecret string, timeout time.Duration) *Kucoin {
	client := newClient(apiKey, apiSecret)
	client.httpClient.Timeout = timeout
	return &Kucoin{client}
}

func doArgs(args ...string) map[string]string {
	m := make(map[string]string)
	var lastK = ""
	for i, v := range args {
		if i&1 == 0 {
			lastK = v
		} else {
			m[lastK] = v
		}
	}
	return m
}

// handleErr gets JSON response from livecoin API en deal with error.
func handleErr(r interface{}) error {
	switch v := r.(type) {
	case map[string]interface{}:
		err := r.(map[string]interface{})["error"]
		if err != nil {
			switch v := err.(type) {
			case map[string]interface{}:
				errorMessage := err.(map[string]interface{})["message"]
				return errors.New(errorMessage.(string))
			default:
				return fmt.Errorf("don't recognized type %T", v)
			}
		}
	case []interface{}:
		return nil
	default:
		return fmt.Errorf("don't recognized type %T", v)
	}

	return nil
}

// Kucoin represent a Kucoin client.
type Kucoin struct {
	client *client
}

// SetDebug enables/disables http request/response dump.
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

// GetUserSymbols is used to get the all open and available trading markets at Kucoin along with other meta data.
// The user should be logged to call this method.
// Filter parameter can be whether 'FAVOURITE' or 'STICK',
// market and symbol parameters can be any as presented at exchange.
func (b *Kucoin) GetUserSymbols(market, symbol, filter string) (symbols []Symbol, err error) {
	payload := map[string]string{}
	if len(market) > 1 {
		payload["market"] = market
	}
	if len(symbol) > 1 {
		payload["symbol"] = symbol
	}
	if len(filter) > 1 {
		payload["filter"] = filter
	}
	r, err := b.client.do("GET", "market/symbols", payload, true)
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
	r, err := b.client.do("GET",
		"open/tick", doArgs("symbol", strings.ToUpper(market)), false,
	)
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
	r, err := b.client.do(
		"GET", "market/open/coin-info", doArgs("coin", strings.ToUpper(c)), false,
	)
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

// ListActiveMapOrders is used to get the information about active orders in user-friendly view
// at Kucoin along with other meta data.
// Symbol is required parameter, and side (or type of order in kucoin docs) may be empty.
func (b *Kucoin) ListActiveMapOrders(symbol string, side string) (activeMapOrders ActiveMapOrder, err error) {
	if len(symbol) < 1 {
		return activeMapOrders, fmt.Errorf("Symbol is required")
	}
	payload := make(map[string]string)
	payload["symbol"] = strings.ToUpper(symbol)
	if len(side) > 1 {
		payload["side"] = strings.ToUpper(side)
	}

	r, err := b.client.do("GET", "order/active-map", payload, true)
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
	var rawRes rawActiveMapOrder
	err = json.Unmarshal(r, &rawRes)
	activeMapOrders = rawRes.Data
	return
}

// ListActiveOrders is used to get the information about active orders in array mode
// at Kucoin along with other meta data.
// Symbol is required parameter, and side (or type of order in kucoin docs) may be empty.
func (b *Kucoin) ListActiveOrders(symbol string, side string) (activeOrders ActiveOrder, err error) {
	if len(symbol) < 1 {
		return activeOrders, fmt.Errorf("The symbol is required")
	}
	payload := make(map[string]string)
	payload["symbol"] = strings.ToUpper(symbol)
	if len(side) > 1 {
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
	var rawRes rawActiveOrder
	err = json.Unmarshal(r, &rawRes)
	activeOrders = rawRes.Data
	return
}

// OrdersBook is used to get the information about active orders at Kucoin along with other meta data.
// Symbol is required parameter, geoup and limit may be empty.
func (b *Kucoin) OrdersBook(symbol string, group, limit int) (ordersBook OrdersBook, err error) {
	if len(symbol) < 1 {
		return ordersBook, fmt.Errorf("The symbol is required")
	}
	payload := map[string]string{}
	payload["symbol"] = strings.ToUpper(symbol)
	if group > 0 {
		payload["group"] = fmt.Sprintf("%v", group)
	}
	if limit == 0 {
		payload["limit"] = fmt.Sprintf("%v", 1000)
	} else {
		payload["limit"] = fmt.Sprintf("%v", limit)
	}

	r, err := b.client.do("GET", "open/orders", payload, true)
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
	var rawRes rawOrdersBook
	err = json.Unmarshal(r, &rawRes)
	ordersBook = rawRes.Data
	return
}

// CreateOrder is used to create order at Kucoin along with other meta data.
func (b *Kucoin) CreateOrder(symbol, side string, price, amount float64) (orderOid string, err error) {
	payload := make(map[string]string)
	payload["amount"] = strconv.FormatFloat(amount, 'f', 8, 64)
	payload["price"] = strconv.FormatFloat(price, 'f', 8, 64)
	payload["type"] = strings.ToUpper(side)

	r, err := b.client.do("POST", fmt.Sprintf("%s/order", strings.ToUpper(symbol)), payload, true)
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
	if err != nil {
		return
	}
	if !rawRes.Success {
		err = errors.New(string(r))
		return
	}
	orderOid = rawRes.Data.OrderOid
	return
}

// CreateOrderByString is used to create order at Kucoin along with other meta data. This ByString version is fix precise problem.
func (b *Kucoin) CreateOrderByString(symbol, side string, price, amount string) (orderOid string, err error) {
	payload := make(map[string]string)
	payload["amount"] = amount
	payload["price"] = price
	payload["type"] = strings.ToUpper(side)

	r, err := b.client.do("POST", fmt.Sprintf("%s/order", strings.ToUpper(symbol)), payload, true)
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
	if err != nil {
		return
	}
	if !rawRes.Success {
		err = errors.New(string(r))
		return
	}
	orderOid = rawRes.Data.OrderOid
	return
}

// AccountHistory is used to get the information about list deposit & withdrawal
// at Kucoin along with other meta data. Coin, Side (type in Kucoin docs.)
// and Status are required parameters. Limit and page may be zeros.
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
	var rawRes rawAccountHistory
	err = json.Unmarshal(r, &rawRes)
	accountHistory = rawRes.Data
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
	var rawRes rawSpecificDealtOrder
	err = json.Unmarshal(r, &rawRes)
	specificDealtOrders = rawRes.Data
	return
}

// ListMergedDealtOrders is used to get the information about dealt orders for
// all symbols at Kucoin along with other meta data.
// All parameters are optional. Timestamp must be in milliseconds from Unix epoch.
func (b *Kucoin) ListMergedDealtOrders(symbol, side string, limit, page int, since, before int64) (mergedDealtOrders MergedDealtOrder, err error) {
	payload := map[string]string{}
	if len(symbol) > 1 {
		payload["symbol"] = symbol
	}
	if len(side) > 1 {
		payload["type"] = side
	}
	if (limit == 0 || limit > 100) && len(symbol) > 1 {
		payload["limit"] = fmt.Sprintf("%v", 100)
	} else if (limit == 0 || limit > 20) && len(symbol) < 1 {
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
	var rawRes rawMergedDealtOrder
	err = json.Unmarshal(r, &rawRes)
	mergedDealtOrders = rawRes.Data
	return
}

// OrderDetails is used to get the information about orders for specific symbol at Kucoin along with other meta data.
// Symbol, Side (type in Kucoin docs.) are required parameters.
// Limit may be zero, and not greater than 20. Page may be zero and by default is equal to 1.
// Example:
// - Symbol = KCS-BTC
// - Side = BUY | SELL
func (b *Kucoin) OrderDetails(symbol, side, orderOid string, limit, page int) (orderDetails OrderDetails, err error) {
	if len(symbol) < 1 || len(side) < 1 || len(orderOid) < 1 {
		return orderDetails, fmt.Errorf("The not all required parameters are presented")
	}
	payload := map[string]string{}
	payload["orderOid"] = orderOid
	payload["symbol"] = symbol
	payload["type"] = side
	if limit == 0 {
		payload["limit"] = fmt.Sprintf("%v", 20)
	} else {
		payload["limit"] = fmt.Sprintf("%v", limit)
	}
	if page == 0 {
		payload["page"] = fmt.Sprintf("%v", 1)
	} else {
		payload["page"] = fmt.Sprintf("%v", page)
	}

	r, err := b.client.do("GET", "order/detail", payload, true)
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
	var rawRes rawOrderDetails
	err = json.Unmarshal(r, &rawRes)
	orderDetails = rawRes.Data
	return
}

// CreateWithdrawalApply is used to create withdrawal for specific coin
// at Kucoin along with other meta data.
// coin, address and amount are required parameters.
// Example:
// - coin = KCS
// - address = example_address
// - amount 0.68
// Result:
// - Nothing.
func (b *Kucoin) CreateWithdrawalApply(coin, address string, amount float64) (withdrawalApply Withdrawal, err error) {
	if len(coin) < 1 || len(address) < 1 || amount == 0 {
		return withdrawalApply, fmt.Errorf("The not all required parameters are presented")
	}
	payload := map[string]string{}
	payload["coin"] = coin
	payload["address"] = address
	payload["amount"] = fmt.Sprintf("%v", amount)

	r, err := b.client.do("POST", fmt.Sprintf(
		"account/%s/withdraw/apply", strings.ToUpper(coin)), payload, true)
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
	var rawRes rawWithdrawal
	err = json.Unmarshal(r, &rawRes)
	withdrawalApply = rawRes.Data
	return
}

// CancelWithdrawal used to cancel withdrawal for specific coin
// at Kucoin along with other meta data.
// coin, txOid are required parameters.
// Example:
// - coin = KCS
// - txOid = example_tx
// Result:
// - Nothing.
func (b *Kucoin) CancelWithdrawal(coin, txOid string) (withdrawal Withdrawal, err error) {
	if len(coin) < 1 || len(txOid) < 1 {
		return withdrawal, fmt.Errorf("The not all required parameters are presented")
	}
	payload := map[string]string{}
	payload["txOid"] = txOid

	r, err := b.client.do("POST", fmt.Sprintf(
		"account/%s/withdraw/cancel", strings.ToUpper(coin)), payload, true)
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
	var rawRes rawWithdrawal
	err = json.Unmarshal(r, &rawRes)
	withdrawal = rawRes.Data
	return
}

// CancelOrder is used to cancel execution of current order at Kucoin along with other meta data.
// Side (type in Kucoin docs.) and order ID are required parameters. Symbol is optional.
func (b *Kucoin) CancelOrder(orderOid, side, symbol string) error {
	if len(symbol) < 1 || len(side) < 1 || len(orderOid) < 1 {
		return fmt.Errorf("The not all required parameters are presented")
	}
	payload := map[string]string{}
	payload["orderOid"] = orderOid
	payload["type"] = side

	r, err := b.client.do("POST", fmt.Sprintf("%s/cancel-order", strings.ToUpper(symbol)), payload, true)
	if err != nil {
		return err
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return err
	}
	return handleErr(response)
}

// CancelAllOrders is used to cancel execution of all orders at Kucoin along with other meta data.
// Symbol, Side (type in Kucoin docs.) are optional parameters.
func (b *Kucoin) CancelAllOrders(symbol, side string) error {
	payload := map[string]string{}
	if len(symbol) > 1 {
		payload["symbol"] = strings.ToUpper(symbol)
	}
	if len(side) > 1 {
		payload["type"] = side
	}

	r, err := b.client.do("POST", "order/cancel-all", payload, true)
	if err != nil {
		return err
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return err
	}
	return handleErr(response)
}
