package cim

import "us-west/auth"

type CreateCustomerProfileRequest struct {
	*auth.MerchantAuth `json:"merchantAuthentication,omitempty"`
	RefId              string  `json:"refId,omitempty"`
	Profile            Profile `json:"profile,omitempty"`
}

type CreateCustomerProfileResponse struct {
	CustomerProfileId string `json:"customerProfileId"`
}

func (c *CreateCustomerProfileRequest) ResponseStruct() interface{} {
	return &CreateCustomerProfileResponse{}
}

func (c *CreateCustomerProfileRequest) SetAuth(a *auth.MerchantAuth) {
	c.MerchantAuth = a
}

func (c *CreateCustomerProfileRequest) Method() string {
	return "createCustomerProfileRequest"
}
