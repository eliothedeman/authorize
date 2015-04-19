package authorize

import "github.com/eliothedeman/authorize/cim"

// create a new customer profile in authorize.net
func (c *Client) CreateCustomerProfile(p cim.Profile) (string, err error) {
	req := &cim.CreateCustomerProfileRequest{}
	req.Profile = p
	r := c.Do(req)
	err = r.Err
	resp := r.ResponseStruct.(*cim.CreateCustomerProfileResponse)
	return resp.CustomerProfileId, err
}

func (c *Client) GetCustomerProfile(profileId string) (*cim.Profile, err error) {
	req := &cim.GetCustomerProfileRequest{}
	req.CustomerProfileId = profileId
	r := c.Do(req)
	err = r.Err
	resp := r.ResponseStruct.(*cim.GetCustomerProfileResponse)
	return resp.Profile, err
}
