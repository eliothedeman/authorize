package cim

import "us-west/auth"

type RefundTransactionRequest struct {
	*auth.MerchantAuth `json:"merchantAuthentication,omitempty"`
	RefId              string             `json:"refId,omitempty"`
	TransactionReqeust TransactionRequest `json:"transactionRequest"`
}

type RefundTransactionResponse struct {
}

func (c *RefundTransactionRequest) ResponseStruct() interface{} {
	return &RefundTransactionResponse{}
}

func (c *RefundTransactionRequest) SetAuth(a *auth.MerchantAuth) {
	c.MerchantAuth = a
}

func (c *RefundTransactionRequest) Method() string {
	return "createTransactionRequest"
}
