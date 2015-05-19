package cim

import "us-west/auth"

type GetCustomerProfileRequest struct {
	*auth.MerchantAuth `json:"merchantAuthentication,omitempty"`
	RefId              string `json:"refId,omitempty"`
	CustomerProfileId  string `json:"customerProfileId"`
}

type GetCustomerProfileResponse struct {
	Profile *Profile `json:"profile"`
}

func (c *GetCustomerProfileRequest) ResponseStruct() interface{} {
	return &GetCustomerProfileResponse{}
}

func (c *GetCustomerProfileRequest) SetAuth(a *auth.MerchantAuth) {
	c.MerchantAuth = a
}

func (c *GetCustomerProfileRequest) Method() string {
	return "getCustomerProfileRequest"
}
