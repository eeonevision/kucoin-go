package kucoin

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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

// doTimeoutRequest do a HTTP request with timeout.
func (c *client) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	// Do the request in the background so we can check the timeout
	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		if c.debug {
			c.dumpRequest(req)
		}
		resp, err := c.httpClient.Do(req)
		if c.debug {
			c.dumpResponse(resp)
		}
		done <- result{resp, err}
	}()
	// Wait for the read or the timeout
	select {
	case r := <-done:
		return r.resp, r.err
	case <-timer.C:
		return nil, errors.New("timeout on reading data from Kucoin API")
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
func (c *client) do(method string, resource string, payload map[string]string, authNeeded bool) (response []byte, err error) {
	var rawurl string
	if strings.HasPrefix(resource, "http") {
		rawurl = resource
	} else {
		rawurl = fmt.Sprintf("%s%s/%s", APIBase, APIPrefix, resource)
	}
	var formData string
	if method == "GET" {
		var URL *url.URL
		URL, err = url.Parse(rawurl)
		if err != nil {
			return
		}
		q := URL.Query()
		for key, value := range payload {
			q.Set(key, value)
		}
		formData = q.Encode()
		URL.RawQuery = formData
		rawurl = URL.String()
	} else {
		formValues := url.Values{}
		for key, value := range payload {
			formValues.Set(key, value)
		}
		formData = formValues.Encode()
	}

	req, err := http.NewRequest(method, rawurl, strings.NewReader(formData))
	if err != nil {
		return
	}
	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	}
	req.Header.Add("Accept", "application/json")

	// Auth
	if authNeeded {
		if len(c.apiKey) == 0 || len(c.apiSecret) == 0 {
			err = errors.New("You need to set API Key and API Secret to call this method")
			return
		}

		nonce := time.Now().UnixNano() / int64(time.Millisecond)
		req.Header.Add("KC-API-KEY", c.apiKey)
		req.Header.Add("KC-API-NONCE", fmt.Sprintf("%v", nonce))
		req.Header.Add("KC-API-SIGNATURE", c.sign(fmt.Sprintf("%s/%s", APIPrefix, resource), nonce, formData))
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}
	return response, err
}

func computeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *client) sign(path string, nonce int64, queryString string) (signature string) {
	strForSign := fmt.Sprintf("%s/%v/%s", path, nonce, queryString)
	signatureStr := b64.StdEncoding.EncodeToString([]byte(strForSign))
	signature = computeHmac256(signatureStr, c.apiSecret)
	return
}
