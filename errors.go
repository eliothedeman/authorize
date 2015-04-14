package authorize

import "errors"

var (
	INVALID_CARD_NUMBER  = errors.New("authorize: invalid card number")
	INVALID_VALUE_LENGTH = errors.New("authorize: invalid length of field")
	DUPLICATE_RECORD     = errors.New("authorize: duplicate record")
)

var (
	// maps error codes to errors
	errMap = map[string]error{
		"E00015": INVALID_VALUE_LENGTH,
		"E00013": INVALID_CARD_NUMBER,
		"E00039": DUPLICATE_RECORD,
	}
)
