package authorize

import (
	"log"
	"testing"

	"github.com/eliothedeman/authorize/transaction"
)

const (
	TEST_API_LOGIN_ID        = "7wm8uRmTz7m7"
	TEST_API_TRANSACTION_KEY = "6A5HvU67k6Xqu66L"
)

func newTestClient() *Client {
	return NewClient(TEST_API_LOGIN_ID, TEST_API_TRANSACTION_KEY, false)
}

func TestAuthenticateDo(t *testing.T) {
	c := newTestClient()
	g := transaction.NewGetTransactionList("1")
	r, err := c.Do(g)
	if err != nil {
		log.Println(err, r)
	}
}
