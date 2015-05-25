package authorize

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/eolexe/authorize/auth"
)

const (
	BODY_TYPE      = "application/json"
	SANDBOX_URL    = "https://apitest.authorize.net/xml/v1/request.api"
	PRODUCTION_URL = "https://api.authorize.net/xml/v1/request.api"
)

// A base client for the authorize.net api
type Client struct {
	URL        string
	httpClient *http.Client
	buffer     *bytes.Buffer
	auth       *auth.MerchantAuth
}

func NewClient(name, transactionKey string) *Client {
	c := &Client{
		buffer:     bytes.NewBuffer([]byte{}),
		httpClient: http.DefaultClient,
		URL:        PRODUCTION_URL,
		auth: &auth.MerchantAuth{
			Name:           name,
			TransactionKey: transactionKey,
		},
	}

	return c
}

func (c *Client) Do(r Request) (resp *Response) {
	resp = &Response{}
	r.SetAuth(c.auth)
	buff, err := json.Marshal(RequestBody(r))
	if err != nil {
		resp.Err = err
		return
	}

	c.buffer.Reset()
	c.buffer.Write(buff)

	req, err := http.NewRequest("POST", c.URL, c.buffer)
	if err != nil {
		resp.Err = err
		return
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	httpResp, err := c.httpClient.Do(req)
	if err != nil {
		resp.Err = err
		return
	}

	buff, err = ioutil.ReadAll(httpResp.Body)

	// this is super gross, but auth.net sends us back responses with garbage at the beginning
	// while using the json api
	index := bytes.Index(buff, []byte("{"))
	buff = buff[index:]
	if index == -1 {
		resp.Err = errors.New("Invalid json")
		return
	}
	httpResp.Body.Close()
	if err != nil {
		resp.Err = err
		return
	}

	resp.ResponseStruct = r.ResponseStruct()
	resp = ParseResponse(resp, buff)
	return
}
