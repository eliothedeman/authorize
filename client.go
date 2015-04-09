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

func (c *Client) Do(r Request) error {
	r.SetAuth(c.auth)
	buff, err := json.Marshal(RequestBody(r))
	if err != nil {
		return err
	}

	c.buffer.Reset()
	c.buffer.Write(buff)

	req, err := http.NewRequest("POST", c.url, c.buffer)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	buff, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	response := r.ResponseStruct()
	err = json.Unmarshal(buff, response)

	log.Println(string(buff))
	return err
}
