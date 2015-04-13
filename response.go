package authorize

import (
	"encoding/json"
	"errors"
	"log"
)

type Error struct {
	Code string `json:"error"`
	Text string `json:"text"`
}

var (
	// maps error codes to errors
	errMap = map[string]error{}
)

func parseError(e *Error) error {
	err, ok := errMap[e.Code]
	if !ok {
		return errors.New(e.Text)
	}
	return err
}

type Message struct {
	Code    string          `json:"resultCode"`
	Message json.RawMessage `json:"message"`
}

type Response struct {
	Messages []Message `json:"messages"`
	err      error
}

func (r *Response) Error() string {
	return r.err.Error()
}

func ParseResponse(buff []byte) (r *Response) {
	r = &Response{}
	err := json.Unmarshal(buff, r)

	r.err = err
	if err != nil {
		log.Println(string(buff))
		return
	}

	// parse out to see if we have an error in the message
	for i := range r.Messages {
		if r.Messages[i].Code == "Error" {
			mErr := &Error{}
			err = json.Unmarshal(r.Messages[i].Message, mErr)
			if err != nil {
				r.err = err
				return
			}
			err = parseError(mErr)
			r.err = err
			return
		}
	}
	return
}
