package authorize

const (
	TEST_API_LOGIN_ID        = "7wm8uRmTz7m7"
	TEST_API_TRANSACTION_KEY = "6A5HvU67k6Xqu66L"
)

func NewTestClient() *Client {
	return NewClient(TEST_API_LOGIN_ID, TEST_API_TRANSACTION_KEY, false)
}
