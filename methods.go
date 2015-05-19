package authorize

import "github.com/eolexe/authorize/cim"

const (
	AUTH_CAPTURE_REQUEST = "authCaptureTransaction"
	REFUND_REQUEST       = "refundTransaction"
)

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

func (c *Client) CreateCustomerSippingAddress(profile_id string, a *cim.Address) (string, error) {

	req := &cim.CreateShippingAddressRequest{}
	req.CustomerProfileId = profile_id
	req.Address = a
	r := c.Do(req)
	resp := r.ResponseStruct.(*cim.CreateShippingAddressResponse)
	return resp.CustomerShippingAddressId, r.Err
}

func (c *Client) CreateCustomerPaymentProfileTransaction(t *cim.Transaction) (string, error) {
	req := &cim.CreateTransactionRequest{}
	req.TransactionRequest.Amount = t.Amount
	req.TransactionRequest.Profile.CustomerProfileId = t.CustomerProfileId
	req.TransactionRequest.Profile.PaymentProfile.PaymentProfileId = t.PaymentProfileId
	if t.CardCode != "" {
		req.TransactionRequest.Profile.PaymentProfile.CardCode = &t.CardCode
	}
	req.TransactionRequest.Tax = t.Tax
	req.TransactionRequest.Duty = t.Duty
	req.TransactionRequest.Shipping = t.Shipping
	req.TransactionRequest.Type = AUTH_CAPTURE_REQUEST
	r := c.Do(req)
	resp := r.ResponseStruct.(*cim.CreateTransactionResponse)

	return resp.TransactionId, r.Err
}

func (c *Client) RefundTransaction(t *cim.Transaction, transactionId string) error {
	req := &cim.RefundTransactionRequest{}
	req.TransactionReqeust.Type = REFUND_REQUEST
	req.TransactionReqeust.RefundTransactionRequestId = &transactionId
	req.TransactionReqeust.Amount = t.Amount
	req.TransactionReqeust.Profile.CustomerProfileId = t.CustomerProfileId
	req.TransactionReqeust.Profile.PaymentProfile.PaymentProfileId = t.PaymentProfileId

	r := c.Do(req)
	return r.Err
}
