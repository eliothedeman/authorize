package authorize

import (
	"log"
	"math/rand"
	"time"
)

const (
	TEST_API_LOGIN_ID        = "7wm8uRmTz7m7"
	TEST_API_TRANSACTION_KEY = "6A5HvU67k6Xqu66L"
)

func init() {
	rand.Seed(time.Now().Unix())
}
func NewTestClient() *Client {
	log.SetFlags(log.Llongfile)
	c := NewClient(TEST_API_LOGIN_ID, TEST_API_TRANSACTION_KEY)
	c.url = SANDBOX_URL
	return c
}

func NewTestClientWithCredentials(name, transactionKey string) *Client {
	log.SetFlags(log.Llongfile)
	c := NewClient(name, transactionKey)
	c.url = SANDBOX_URL
	return c
}

