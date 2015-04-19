package authorize

import "github.com/eliothedeman/authorize/cim"

// create a new customer profile in authorize.net
func (c *Client) CreateCustomerProfile(p cim.Profile) (string, error) {
	req := &cim.CreateCustomerProfileRequest{}
	req.Profile = p
	r := c.Do(req)
	resp := r.ResponseStruct.(*cim.CreateCustomerProfileResponse)
	return resp.CustomerProfileId, r.Err
}

func (c *Client) GetCustomerProfile(profileId string) (*cim.Profile, error) {
	req := &cim.GetCustomerProfileRequest{}
	req.CustomerProfileId = profileId
	r := c.Do(req)
	resp := r.ResponseStruct.(*cim.GetCustomerProfileResponse)
	return resp.Profile, r.Err
}

func (c *Client) CreateCustomerPaymentProfile(customerId string, pp *cim.PaymentProfile) (customerProfileId string, err error) {
	req := &cim.CreateCustomerPaymentProfileRequest{}
	req.PaymentProfile = pp
	req.CustomerProfileId = customerId
	r := c.Do(req)
	resp := r.ResponseStruct.(*cim.CreateCustomerPaymentProfileResponse)
	return resp.CustomerPaymentProfileId, r.Err
}
