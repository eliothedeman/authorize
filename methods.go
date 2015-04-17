package authorize

import "github.com/eliothedeman/authorize/cim"

// create a new customer profile in authorize.net
func (c *Client) CreateCustomerProfile(p cim.Profile) (resp *cim.CreateCustomerProfileResponse, err error) {
	req := &cim.CreateCustomerProfileRequest{}
	req.Profile = p
	r := c.Do(req)
	err = r.Err
	return r.ResponseStruct.(*cim.CreateCustomerProfileResponse), err
}

func (c *Client) GetCustomerProfile(profileId string) (resp *cim.GetCustomerProfileResponse, err error) {
	req := &cim.GetCustomerProfileRequest{}
	req.CustomerProfileId = profileId
	r := c.Do(req)
	err = r.Err
	return r.ResponseStruct.(*cim.GetCustomerProfileResponse), err
}
