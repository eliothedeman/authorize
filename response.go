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
	err      error
}

func (r *Response) Error() string {
	if r.err == nil {
		return ""
	}
	return r.err.Error()
}

func ParseResponse(buff []byte) (r *Response) {
	log.Println(string(buff))
	r = &Response{}
	err := json.Unmarshal(buff, r)

	r.err = err
	if err != nil {
		log.Println(string(buff))
		return
	}

	// parse out to see if we have an error in the message
	for i := range r.Messages.Messages {
		if r.Messages.Code == "Error" {
			mErr := &Error{}
			err = json.Unmarshal(r.Messages.Messages[i], mErr)
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
