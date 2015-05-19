package cim

import "us-west/auth"

type CreateCustomerPaymentProfileRequest struct {
	*auth.MerchantAuth `json:"merchantAuthentication,omitempty"`
	RefId              string          `json:"refId,omitempty"`
	CustomerProfileId  string          `json:"customerProfileId,omitempty"`
	PaymentProfile     *PaymentProfile `json:"paymentProfile,omitempty"`
}

type CreateCustomerPaymentProfileResponse struct {
	CustomerPaymentProfileId string `json:"customerPaymentProfileId"`
}

func (c *CreateCustomerPaymentProfileRequest) ResponseStruct() interface{} {
	return &CreateCustomerPaymentProfileResponse{}
}

func (c *CreateCustomerPaymentProfileRequest) SetAuth(a *auth.MerchantAuth) {
	c.MerchantAuth = a
}

func (c *CreateCustomerPaymentProfileRequest) Method() string {
	return "createCustomerPaymentProfileRequest"
}
