package cim

import "github.com/eliothedeman/authorize/auth"

type CreateCustomerProfileRequest struct {
	Auth    *auth.MerchantAuth `json:"merchantAuthentication,omitempty"`
	RefId   string             `json:"refId,omitempty"`
	Profile struct {
		MerchantCustomerId string `json:"merchantCustomerId,omitempty"`
		Description        string `json:"description,omitempty"`
		Email              string `json:"email,omitempty"`
		PaymentProfiles    struct {
			CustomerType string `json:"customerType,omitempty"`
			BillTo       struct {
				FirstName   string `json:"firstName,omitempty"`
				LastName    string `json:"lastName,omitempty"`
				Company     string `json:"company,omitempty"`
				Address     string `json:"address,omitempty"`
				City        string `json:"city,omitempty"`
				State       string `json:"state,omitempty"`
				Zip         string `json:"zip,omitempty"`
				Country     string `json:"country,omitempty"`
				PhoneNumber string `json:"phoneNumber,omitempty"`
				FaxNumber   string `json:"faxNumber,omitempty"`
			} `json:"billTo,omitempty"`
			Payment struct {
				CreditCard *struct {
					CardNumber     string `json:"cardNumber,omitempty"`
					ExpirationDate string `json:"expirationDate,omitempty"`
					CardCode       string `json:"cardCode,omitempty"`
				} `json:"creditCard,omitempty"`
				BankAccount *struct {
					AccountType    string `json:"accountType,omitempty"`
					RouteingNumber string `json:"routingNumber,omitempty"`
					NameOnAccount  string `json:"nameOnAccount,omitempty"`
					ECheckType     string `json:"echeckType,omitempty"`
					BankName       string `json:"bankName,omitempty"`
				}
			} `json:"payment,omitempty"`
		} `json:"paymentProfiles,omitempty"`
	} `json:"profile,omitempty"`

	Response *CreateCustomerProfileResponse `json:"-,omitempty"`
}

type CreateCustomerProfileResponse struct {
}

func (c *CreateCustomerProfileRequest) ResponseStruct() interface{} {
	return c.Response
}

func (c *CreateCustomerProfileRequest) SetAuth(a *auth.MerchantAuth) {
	c.Auth = a
}

func (c *CreateCustomerProfileRequest) Method() string {
	return "createCustomerProfileRequest"
}
