package cim

import "github.com/eolexe/authorize/auth"

type CreateTransactionRequest struct {
	*auth.MerchantAuth `json:"merchantAuthentication,omitempty"`
	RefId              string             `json:"refId,omitempty"`
	TransactionRequest TransactionRequest `json:"transactionRequest"`
}

type TransactionRequest struct {
	Type    string `json:"transactionType"`
	Amount  string `json:"amount"`
	Profile struct {
		CustomerProfileId string `json:"customerProfileId"`
		PaymentProfile    struct {
			PaymentProfileId string  `json:"paymentProfileId"`
			CardCode         *string `json:"cardCode,omitempty"`
		} `json:"paymentProfile"`
	} `json:"profile"`
	Tax                        *Charge `json:"tax,omitempty"`
	Duty                       *Charge `json:"duty,omitempty"`
	Shipping                   *Charge `json:"shipping,omitempty"`
	RefundTransactionRequestId *string `json:"refTransId,omitempty"`
}

type CreateTransactionResponse struct {
	AuthCode      string `json:"authCode"`
	TransactionId string `json:"transId"`
	TransHash     string `json:"transHash"`
}

func (c *CreateTransactionRequest) ResponseStruct() interface{} {
	return &CreateTransactionResponse{}
}

func (c *CreateTransactionRequest) SetAuth(a *auth.MerchantAuth) {
	c.MerchantAuth = a
}

func (c *CreateTransactionRequest) Method() string {
	return "createTransactionRequest"
}
