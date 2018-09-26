package kucoin

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"time"
)

type client struct {
	apiKey     string
	apiSecret  string
	httpClient http.Client
	debug      bool
}

func newClient(apiKey, apiSecret string) (c *client) {
	c = &client{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
	c.httpClient.Timeout = time.Second * 30
	return
}

func (c client) dumpRequest(r *http.Request) {
	if r == nil {
		log.Println("dumpReq ok: <nil>")
	} else {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Printf("dumpReq err: %s\n", err)
		} else {
			log.Printf("dumpReq ok: %s\n", dump)
		}
	}
}

func (c client) dumpResponse(r *http.Response) {
	if r == nil {
		log.Println("dumpResponse ok: <nil>")
	} else {
		dump, err := httputil.DumpResponse(r, true)
		if err != nil {
			log.Printf("dumpResponse err: %s\n", err)
		} else {
			log.Printf("dumpResponse ok: %s\n", dump)
		}
	}
}

// do prepare and process HTTP request to Kucoin API.
/*
	 *  Example
	 *  POST parameters：
	 *    type: BUY
	 *    amount: 10
	 *    price: 1.1
	 *    Arrange the parameters in ascending alphabetical order (lower cases first),
		  then combine them with & (don't urlencode them, don't add ?, don't add extra &),
		  e.g. amount=10&price=1.1&type=BUY
*/
func (c *client) do(method, resource string, payload map[string]string, authNeeded bool) ([]byte, error) {
	var req *http.Request

	Url, err := url.Parse(kucoinUrl)
	if err != nil {
		return nil, err
	}
	Url.Path = path.Join(Url.Path, resource)
	queryString := ""
	if method == "GET" {
		q := Url.Query()
		for key, value := range payload {
			q.Set(key, value)
		}
		Url.RawQuery = q.Encode()
		req, err = http.NewRequest("GET", Url.String(), nil)
		queryString = Url.Query().Encode()
	} else {
		postValues := url.Values{}
		for key, value := range payload {
			postValues.Set(key, value)
		}
		queryString = postValues.Encode()
		req, err = http.NewRequest(
			method, Url.String(), strings.NewReader(
				queryString,
			),
		)
	}
	if err != nil {
		return nil, err
	}
	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	}
	req.Header.Add("Accept", "application/json")

	// Auth
	if authNeeded {
		if len(c.apiKey) == 0 || len(c.apiSecret) == 0 {
			return nil, errors.New("API Key and API Secret must be set")
		}

		nonce := time.Now().UnixNano() / int64(time.Millisecond)
		req.Header.Add("KC-API-KEY", c.apiKey)
		req.Header.Add("KC-API-NONCE", fmt.Sprintf("%v", nonce))
		req.Header.Add(
			"KC-API-SIGNATURE", c.sign(
				Url.Path, queryString, nonce,
			),
		)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}
	return data, err
}

func computeHmac256(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	io.WriteString(h, message)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (c *client) sign(path, queryString string, nonce int64) (signature string) {
	strForSign := fmt.Sprintf("%s/%v/%s", path, nonce, queryString)
	signatureStr := b64.StdEncoding.EncodeToString([]byte(strForSign))
	signature = computeHmac256(signatureStr, c.apiSecret)
	return
}
