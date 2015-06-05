package authorize

import (
	"errors"
	"log"
)

var (
	INVALID_CONTENT      = errors.New("authorize: request content is invalid")
	INVALID_VALUE_LENGTH = errors.New("authorize: invalid length of field")
	DUPLICATE_RECORD     = errors.New("authorize: duplicate record")
)

func parseError(e *Error) error {
	err, ok := errMap[e.Code]
	if !ok {
		log.Printf("unknown error code %+v", e)
		return errors.New(e.Text)
	}

	return err
}

var (
	// maps error codes to errors
	errMap = map[string]error{
		"E00015": INVALID_VALUE_LENGTH,
		"E00013": INVALID_CONTENT,
		"E00039": DUPLICATE_RECORD,
	}
)
