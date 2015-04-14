package authorize

import (
	"encoding/json"
	"errors"
	"log"
)

type Error struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

func parseError(e *Error) error {
	err, ok := errMap[e.Code]
	if !ok {
		log.Println("unknown error code", e.Code)
		return errors.New(e.Text)
	}
	return err
}

type Message struct {
	Code     string            `json:"resultCode"`
	Messages []json.RawMessage `json:"message"`
}

type Response struct {
	Messages Message `json:"messages"`
	Err      error
}

func ParseResponse(buff []byte) (r *Response) {
	r = &Response{}
	r.Err = json.Unmarshal(buff, r)

	if r.Err != nil {
		return
	}

	// parse out to see if we have an error in the message
	for i := range r.Messages.Messages {
		if r.Messages.Code == "Error" {
			mErr := &Error{}
			r.Err = json.Unmarshal(r.Messages.Messages[i], mErr)
			if r.Err != nil {
				return
			}
			r.Err = parseError(mErr)
			return
		}
	}
	return
}
