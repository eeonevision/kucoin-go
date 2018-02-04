# Kucoin REST-API Golang
Unofficial [Kucoin API](https://kucoinapidocs.docs.apiary.io/) implementation written on Golang.

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
| Tick (symbols) | Open    | ✔ |
| Get coin info | Open | ✔ |
| List coins | Open    | ✔ |
| Get coin deposit address | Auth | ✔ |
| Get balance of coin | Auth | ✔ |
| Create an order | Auth | ✔ |
| Get user info | Auth | ✔ |
| List active orders | Auth | ✔ |
| List deposit & withdrawal records | Auth | ✔ |
| List dealt orders (Specific and Merged) | Auth | ✔ |
| Create withdrawal apply |  |
| Cancel withdrawal |  |
| Cancel orders |  |
| Cancel all orders |  |
| Order details |  |
| Order books |  |

## Donate
Your **★Star** will be best donation to my work)
