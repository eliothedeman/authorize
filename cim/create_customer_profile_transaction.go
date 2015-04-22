package cim

import "github.com/eliothedeman/authorize/auth"

type CreateCustomerProfileTransactionRequest struct {
	*auth.MerchantAuth `json:"merchantAuthentication,omitempty"`
	RefId              string `json:"refId,omitempty"`
	Transaction        struct {
		ProfileTransAuthCapture *Transaction `json:"profileTransAuthCapture"`
	} `json:"transaction,omitempty"`
}

type CreateCustomerProfileTransactionResponse struct {
	CustomerPaymentProfileId string `json:"customerPaymentProfileId"`
}

func (c *CreateCustomerProfileTransactionRequest) ResponseStruct() interface{} {
	return &CreateCustomerProfileTransactionResponse{}
}

func (c *CreateCustomerProfileTransactionRequest) SetAuth(a *auth.MerchantAuth) {
	c.MerchantAuth = a
}

func (c *CreateCustomerProfileTransactionRequest) Method() string {
	return "createCustomerProfileTransactionRequest"
}
