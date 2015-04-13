package authorize

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/eliothedeman/authorize/auth"
)

const (
	BODY_TYPE      = "application/json"
	SANDBOX_URL    = "https://apitest.authorize.net/xml/v1/request.api"
	PRODUCTION_URL = "https://api.authorize.net/xml/v1/request.api"
)

// A base client for the authorize.net api
type Client struct {
	url        string
	production bool
	httpClient *http.Client
	buffer     *bytes.Buffer
	auth       *auth.MerchantAuth
}

func NewClient(name, transactionKey string, production bool) *Client {
	c := &Client{
		production: production,
		buffer:     bytes.NewBuffer([]byte{}),
		httpClient: http.DefaultClient,
		auth: &auth.MerchantAuth{
			Name:           name,
			TransactionKey: transactionKey,
		},
	}

	if production {
		c.url = PRODUCTION_URL
	} else {
		c.url = SANDBOX_URL
	}

	return c
}

func (c *Client) Do(r Request) (resp *Response) {
	resp = &Response{}
	r.SetAuth(c.auth)
	buff, err := json.Marshal(RequestBody(r))
	if err != nil {
		resp.err = err
		return
	}

	c.buffer.Reset()
	c.buffer.Write(buff)

	req, err := http.NewRequest("POST", c.url, c.buffer)
	if err != nil {
		resp.err = err
		return
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	httpResp, err := c.httpClient.Do(req)
	if err != nil {
		resp.err = err
		return
	}

	buff, err = ioutil.ReadAll(httpResp.Body)
	httpResp.Body.Close()
	if err != nil {
		resp.err = err
		return
	}

	resp = ParseResponse(buff)
	log.Println(resp.err)
	return
}
