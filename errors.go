package authorize

import "errors"

var (
	INVALID_CARD_NUMBER = errors.New("authorize: invalid card number")
	DUPLICATE_RECORD    = errors.New("authorize: duplicate record")
)

var (
	// maps error codes to errors
	errMap = map[string]error{
		"E00013": INVALID_CARD_NUMBER,
		"E00039": DUPLICATE_RECORD,
	}
)
