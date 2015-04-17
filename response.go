package authorize

import "encoding/json"

type Error struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type Message struct {
	Code     string            `json:"resultCode"`
	Messages []json.RawMessage `json:"message"`
}

type Response struct {
	Messages       Message `json:"messages"`
	Raw            json.RawMessage
	ResponseStruct interface{}
	Err            error
}

func ParseResponse(r *Response, buff []byte) *Response {
	r.Raw = buff
	r.Err = json.Unmarshal(buff, r)

	if r.Err != nil {
		return r
	}

	// parse out to see if we have an error in the message
	for i := range r.Messages.Messages {
		if r.Messages.Code == "Error" {
			mErr := &Error{}
			r.Err = json.Unmarshal(r.Messages.Messages[i], mErr)
			if r.Err != nil {
				return r
			}
			r.Err = parseError(mErr)
			return r
		}
	}

	r.Err = json.Unmarshal(buff, r.ResponseStruct)
	return r
}
