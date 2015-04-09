package authorize

import "github.com/eliothedeman/authorize/auth"

// An api request for authorize.net
type Request interface {
	Method() string // the method name of the request
	ResponseStruct() interface{}
	SetAuth(a *auth.MerchantAuth)
}

func RequestBody(r Request) map[string]Request {
	return map[string]Request{
		r.Method(): r,
	}
}
