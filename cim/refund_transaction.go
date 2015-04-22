package cim

import "github.com/eliothedeman/authorize/auth"

type RefundTransactionRequest struct {
	*auth.MerchantAuth `json:"merchantAuthentication,omitempty"`
	RefId              string   `json:"refId,omitempty"`
	CustomerProfileId  string   `json:"customerProfileId"`
	Address            *Address `json:"address"`
}

type RefundTransactionResponse struct {
	CustomerShippingAddressId string `json:"customerAddressId"`
}

func (c *RefundTransactionRequest) ResponseStruct() interface{} {
	return &RefundTransactionResponse{}
}

func (c *RefundTransactionRequest) SetAuth(a *auth.MerchantAuth) {
	c.MerchantAuth = a
}

func (c *RefundTransactionRequest) Method() string {
	return "createCustomerShippingAddressRequest"
}
