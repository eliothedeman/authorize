package cim

import "github.com/eliothedeman/authorize/auth"

type CreateCustomerProfileRequest struct {
	*auth.MerchantAuth `json:"merchantAuthentication,omitempty"`
	RefId              string                         `json:"refId,omitempty"`
	Profile            Profile                        `json:"profile,omitempty"`
	Response           *CreateCustomerProfileResponse `json:"-,omitempty"`
}

type CreateCustomerProfileResponse struct {
}

func (c *CreateCustomerProfileRequest) ResponseStruct() interface{} {
	return c.Response
}

func (c *CreateCustomerProfileRequest) SetAuth(a *auth.MerchantAuth) {
	c.MerchantAuth = a
}

func (c *CreateCustomerProfileRequest) Method() string {
	return "createCustomerProfileRequest"
}
