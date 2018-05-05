# Kucoin REST-API Golang
Unofficial [Kucoin API](https://kucoinapidocs.docs.apiary.io/) implementation written on Golang.

[![GoDoc](https://godoc.org/github.com/eeonevision/kucoin-go?status.svg)](https://godoc.org/github.com/eeonevision/kucoin-go)[![Go Report Card](https://goreportcard.com/badge/github.com/eeonevision/kucoin-go)](https://goreportcard.com/report/github.com/eeonevision/kucoin-go)

## Features
- Ready to go solution. Just import the package
- The most needed methods are implemented
- Simple authorization handling
- Pure and stable code
- Built-in Golang performance

## How to use
```bash
go get -u github.com/eeonevision/kucoin-go
```
```golang
package main

import (
	"github.com/eeonevision/kucoin-go"
)

func main() {
	// Set your own API key and secret
	k := kucoin.New("API_KEY", "API_SECRET")
	// Call resource
	k.GetCoinBalance("BTC")
}
```
## Checklist
| API Resource | Type | Done  |
| -------------| ----- | ----- |
| Tick (symbols) | Open | ✔ |
| Get coin info | Open | ✔ |
| List coins | Open | ✔ |
| Tick (symbols) for logged user | Auth | ✔ |
| Get coin deposit address | Auth | ✔ |
| Get balance of coin | Auth | ✔ |
| Create an order | Auth | ✔ |
| Get user info | Auth | ✔ |
| List active orders (Both map and array) | Auth | ✔ |
| List deposit & withdrawal records | Auth | ✔ |
| List dealt orders (Both Specific and Merged) | Auth | ✔ |
| Order details | Auth | ✔ |
| Create withdrawal apply | Auth | ✔ |
| Cancel withdrawal | Auth | ✔ |
| Cancel orders | Auth | ✔ |
| Cancel all orders | Auth | ✔ |
| Order books | Auth | ✔ |

## Donate
Your **★Star** will be best donation to my work)
