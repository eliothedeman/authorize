package cim

import "us-west/auth"

type CreateShippingAddressRequest struct {
	*auth.MerchantAuth `json:"merchantAuthentication,omitempty"`
	RefId              string   `json:"refId,omitempty"`
	CustomerProfileId  string   `json:"customerProfileId"`
	Address            *Address `json:"address"`
}

type CreateShippingAddressResponse struct {
	CustomerShippingAddressId string `json:"customerAddressId"`
}

func (c *CreateShippingAddressRequest) ResponseStruct() interface{} {
	return &CreateShippingAddressResponse{}
}

func (c *CreateShippingAddressRequest) SetAuth(a *auth.MerchantAuth) {
	c.MerchantAuth = a
}

func (c *CreateShippingAddressRequest) Method() string {
	return "createCustomerShippingAddressRequest"
}
