package authorize

import "errors"

var (
	INVALID_CARD_NUMBER = errors.New("authorize: invalid card number")
)

var (
	// maps error codes to errors
	errMap = map[string]error{
		"E00013": INVALID_CARD_NUMBER,
	}
)
